package services

import (
	"config-pilot-agent/controller"
	"config-pilot-agent/interfaces"
	"config-pilot-agent/model"
	"config-pilot-agent/utils/logger"
	"fmt"
	"os"
	"os/exec"
)

type GitProcess struct {
	cmd             *exec.Cmd
	repository      model.Repository
	npmPatchManager *NpmPatchManager
	patchManager    *PatchManager
	controllerApi   interfaces.ControllerApi
}

func NewGitProcess(patchManager *PatchManager, repo model.Repository, controller interfaces.ControllerApi) *GitProcess {
	git := new(GitProcess)
	git.patchManager = patchManager
	git.repository = repo
	git.controllerApi = controller
	git.npmPatchManager = &NpmPatchManager{Name: git.repository.Name, patchManager: git.patchManager}
	return git
}
func (git *GitProcess) Clone() {
	logger.Println("cloning remote repository: ...")
	git.cmd = exec.Command("git", "clone", git.repository.URL)
	if out, err := git.cmd.Output(); err != nil {
		logger.PrintErrorln("cloning remote repository failed with error:", err.Error(), string(out))
	} else {
		logger.PrintSuccessln("cloning remote repository: -> done")
		git.Scan()
		git.SaveChanges()
	}
}
func (git *GitProcess) Clean() {
	logger.Println("cleaning up tempory files and folders")
	os.RemoveAll(git.repository.Name)
	logger.PrintSuccessln(fmt.Sprintf("cleaning up : %s -> done", git.repository.Name))
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
				logger.Println("retry pushing changes")
				git.cleanUpRemoteBranch()
				git.PushingCodeChanges()
			}
			git.CreatePr()
		}
	} else {
		logger.PrintSuccessln("==== skipping save changes -> done")
	}

}
func (git *GitProcess) CheckoutBranch() error {
	logger.Println("checking out to new branch: ...")
	git.cmd = exec.Command("git", "checkout", "-b", getCheckBranch(git.controllerApi))
	git.cmd.Dir = git.repository.Name
	if out, err := git.cmd.Output(); err != nil {
		logger.PrintErrorln("creating new branch failed with error:", err.Error(), string(out))
		return err
	} else {
		logger.PrintSuccessln(fmt.Sprintf("checked out to new branch: %s -> done", getCheckBranch(git.controllerApi)))
	}
	logger.Println("staging current changes to the branch...")
	git.cmd = exec.Command("git", "add", ".")
	git.cmd.Dir = git.repository.Name
	if out, err := git.cmd.Output(); err != nil {
		logger.PrintErrorln("staging changes failed with error:", err.Error(), string(out))
		return err
	} else {
		logger.PrintSuccessln("staging changes  -> done")
	}
	logger.Println("commiting staged changes to the branch...")
	git.cmd = exec.Command("git", "commit", "-m", "changes")
	git.cmd.Dir = git.repository.Name
	if out, err := git.cmd.Output(); err != nil {
		logger.PrintErrorln("commiting changes failed with error:", string(out))
		return err
	} else {
		logger.PrintSuccessln("changes commited  -> done")
		logger.Println("output:", string(out))
	}
	return nil
}
func (git *GitProcess) PushingCodeChanges() error {
	logger.Println("pushing current branch...")
	git.cmd = exec.Command("git", "push", "origin", getCheckBranch(git.controllerApi))
	git.cmd.Dir = git.repository.Name
	if out, err := git.cmd.Output(); err != nil {
		logger.PrintErrorln("pushing changes failed with error:", err.Error(), string(out))
		return err
	} else {
		logger.PrintSuccessln("pushing code to remote branch -> done")
	}
	return nil
}
func (git *GitProcess) cleanUpRemoteBranch() error {
	logger.Println("cleaning up/syncing remote feature branch")
	git.cmd = exec.Command("git", "push", "origin", "--delete", getCheckBranch(git.controllerApi))
	git.cmd.Dir = git.repository.Name
	if out, err := git.cmd.Output(); err != nil {
		logger.PrintErrorln("cleaning up remote feature branch failed with error:", string(out))
		return err
	} else {
		logger.PrintSuccessln("remote branch cleaning up done  -> done")
	}
	return nil
}
func (git *GitProcess) Run() {
	logger.Println("running git process job for ", git.repository.Name)
	git.Clean()
	git.Clone()
	git.Clean()
}
func getCheckBranch(t interfaces.ControllerApi) string {
	switch c := t.(type) {
	case controller.AzureDevopsApi:
		return c.Request.SourceBranch
	case controller.GithubApi:
		return c.Request.SourceBranch
	default:
		return ""
	}
}
func (git *GitProcess) CreatePr() {
	logger.Println(fmt.Sprintf("creating pull-request for %s branch to %s branch", getCheckBranch(git.controllerApi), git.repository.MergeBranch))
	if result, err := git.controllerApi.CreatePr(); err != nil {
		logger.PrintErrorln("pull-request created -> failed")
		logger.Println(err.Error())
	} else {
		logger.Println(result)
		logger.PrintSuccessln("pull-request created -> done")
	}
}
