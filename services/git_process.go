package services

import (
	"bytes"
	"config-pilot-job/model"
	"config-pilot-job/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type GitProcess struct {
	cmd             *exec.Cmd
	repository      model.Repository
	configuration   model.Configuration
	npmPatchManager *NpmPatchManager
	patchManager    *PatchManager
}

func NewGitProcess(config model.Configuration, repo model.Repository, patchManager *PatchManager) *GitProcess {
	git := new(GitProcess)
	git.configuration = config
	git.patchManager = patchManager
	git.repository = repo
	git.npmPatchManager = &NpmPatchManager{Name: git.repository.Name, patchManager: git.patchManager}
	return git
}
func (git *GitProcess) Clone() {
	log.Println("cloning remote repository: ...")
	git.cmd = exec.Command("git", "clone", git.repository.URL)
	if out, err := git.cmd.Output(); err != nil {
		log.Println("cloning remote repository failed with error:", err.Error(), string(out))
	} else {
		log.Println("cloning remote repository: -> done")
	}
}
func (git *GitProcess) Clean() {
	log.Println("cleaning up tempory files and folders")
	os.RemoveAll(git.repository.Name)
	log.Println(fmt.Sprintf("cleaning up : %s -> done", git.repository.Name))
}
func (git *GitProcess) Scan() {
	git.npmPatchManager.LoadPatchData()
	git.npmPatchManager.VerifyAndUpgradePatches()
}
func (git *GitProcess) SaveChanges() {
	if git.npmPatchManager.RequireUpdate {
		git.npmPatchManager.SaveChanges()
		if err := git.CheckoutBranch(); err == nil {
			if err = git.PushingCodeChanges(); err != nil {
				git.cleanUpRemoteBranch()
				git.PushingCodeChanges()
			}
			git.CreatePr()
		}
	} else {
		log.Println("==== skipping save changes -> done")
	}

}
func (git *GitProcess) CheckoutBranch() error {
	log.Println("checking out to new branch: ...")
	git.cmd = exec.Command("git", "checkout", "-b", git.configuration.CheckoutBranch)
	git.cmd.Dir = git.repository.Name
	if out, err := git.cmd.Output(); err != nil {
		log.Println("creating new branch failed with error:", err.Error(), string(out))
		return err
	} else {
		log.Println(fmt.Sprintf("checked out to new branch: %s -> done", git.configuration.CheckoutBranch))
	}
	log.Println("staging current changes to the branch...")
	git.cmd = exec.Command("git", "add", ".")
	git.cmd.Dir = git.repository.Name
	if out, err := git.cmd.Output(); err != nil {
		log.Println("staging changes failed with error:", err.Error(), string(out))
		return err
	} else {
		log.Println("staging changes  -> done")
	}
	log.Println("commiting staged changes to the branch...")
	git.cmd = exec.Command("git", "commit", "-m", "changes")
	git.cmd.Dir = git.repository.Name
	if out, err := git.cmd.Output(); err != nil {
		log.Println("commiting changes failed with error:", string(out))
		return err
	} else {
		log.Println("changes commited  -> done")
		log.Println("output:", string(out))
	}
	return nil
}
func (git *GitProcess) PushingCodeChanges() error {
	log.Println("pushing current branch...")
	git.cmd = exec.Command("git", "push", "origin", git.configuration.CheckoutBranch)
	git.cmd.Dir = git.repository.Name
	if out, err := git.cmd.Output(); err != nil {
		log.Println("pushing changes failed with error:", err.Error(), string(out))
		return err
	} else {
		log.Println("pushing code -> done")
	}
	return nil
}
func (git *GitProcess) cleanUpRemoteBranch() error {
	// git.cmd = exec.Command("git", "pull", "origin", git.defaultBranch)
	log.Println("cleaning up/syncing remote feature branch")
	git.cmd = exec.Command("git", "push", "origin", "--delete", git.configuration.CheckoutBranch)
	git.cmd.Dir = git.repository.Name
	if out, err := git.cmd.Output(); err != nil {
		log.Println("cleaning up remote feature branch failed with error:", string(out))
		return err
	} else {
		log.Println("cleanup done  -> done")
	}
	return nil
}
func (git *GitProcess) Run() {
	log.Println("running git process job for ", git.repository.Name)
	git.Clean()
	git.Clone()
	git.Scan()
	git.SaveChanges()
	git.Clean()

}
func (git *GitProcess) CreatePr() {
	log.Println(fmt.Sprintf("creating pull-request for %s branch to %s branch", git.configuration.CheckoutBranch, git.repository.Branch))
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls", git.configuration.Owner, git.repository.Name)
	head := fmt.Sprintf("%s:%s", git.configuration.Owner, git.configuration.CheckoutBranch)
	postBody, _ := json.Marshal(map[string]string{
		"head":  head,
		"base":  git.repository.Branch,
		"body":  git.configuration.PrRequestMessage,
		"title": git.configuration.PrRequestTitle,
	})
	responseBody := *bytes.NewBuffer(postBody)
	client := utils.NewClient()
	client.SendRequest(url, responseBody)
	log.Println("pull-request created -> done")
}
