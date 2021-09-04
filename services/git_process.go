package services

import (
	"fmt"
	"log"
	"os/exec"
)

type GitProcess struct {
	cmd             *exec.Cmd
	RepoName        string
	RepoUrl         string
	npmPatchManager *NpmPatchManager
	PatchManager    *PatchManager
	defaultBranch   string
}

func NewGitProcess(name string, url string, patchManager *PatchManager) *GitProcess {
	git := new(GitProcess)
	git.RepoName = name
	git.RepoUrl = url
	git.PatchManager = patchManager
	git.cmd = exec.Command("bash")
	git.defaultBranch = "features/package-upgrade"
	git.npmPatchManager = &NpmPatchManager{Name: git.RepoName, patchManager: git.PatchManager}
	return git
}
func (git *GitProcess) Clone() {
	log.Println("cloning remote repository: ...")
	git.cmd = exec.Command("git", "clone", git.RepoUrl)
	if out, err := git.cmd.Output(); err != nil {
		log.Fatalln("cloning remote repository failed with error:", err.Error())
	} else {
		log.Println("cloning remote repository: -> done")
		log.Println("output:", string(out))
	}
}
func (git *GitProcess) Clean() {
	log.Println("cleaning up tempory files and folders")
	git.cmd = exec.Command("rm", "-rf", git.RepoName)
	if out, err := git.cmd.Output(); err != nil {
		log.Fatalln("cleaning up failed with error:", err.Error(), string(out))
	} else {
		log.Println(fmt.Sprintf("cleaning up : %s -> done", git.RepoName))
	}
}
func (git *GitProcess) Scan() {
	git.npmPatchManager.LoadPatchData()
	git.npmPatchManager.VerifyAndUpgradePatches()
}
func (git *GitProcess) SaveChanges() {
	git.npmPatchManager.SaveChanges()
}
func (git *GitProcess) CheckoutBranch() {
	log.Println("checking out to new branch: ...")
	git.cmd = exec.Command("git", "checkout", "-b", git.defaultBranch)
	git.cmd.Dir = git.RepoName
	if out, err := git.cmd.Output(); err != nil {
		log.Fatalln("creating new branch failed with error:", err.Error(), string(out))
	} else {
		log.Println(fmt.Sprintf("checked out to new branch: %s -> done", git.defaultBranch))
	}
}
func (git *GitProcess) RaisePR() {
	branch := "features/package-upgrade"
	log.Println("staging current changes to the branch...")
	git.cmd = exec.Command("git", "add", ".")
	git.cmd.Dir = git.RepoName
	if out, err := git.cmd.Output(); err != nil {
		log.Fatalln("staging changes failed with error:", err.Error(), string(out))
	} else {
		log.Println("stagin changes  -> done")
	}
	log.Println("commiting staged changes to the branch...")
	git.cmd = exec.Command("git", "commit", "-m", "changes")
	git.cmd.Dir = git.RepoName
	if out, err := git.cmd.Output(); err != nil {
		log.Fatalln("commiting changes failed with error:", err.Error(), string(out))
	} else {
		log.Println("changes commited  -> done")
		log.Println("output:", string(out))
	}
	log.Println("pushing current branch...")
	git.cmd = exec.Command("git", "push", "origin", branch)
	git.cmd.Dir = git.RepoName
	if out, err := git.cmd.Output(); err != nil {
		log.Fatalln("pushing changes failed with error:", err.Error(), string(out))
	} else {
		log.Println("pushing code -> done")
		log.Println("output:", string(out))
	}
}
func (git *GitProcess) Run() {
	// git.Clone()
	// git.Scan()
	// git.SaveChanges()
	// git.CheckoutBranch()
	// git.RaisePR()
	git.Clean()
}
