package model

import (
	"fmt"
	"strings"
)

type NpmPackage struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	OriginalContent string
}

type NpmDependencies struct {
	Name  string
	Value string
}

func GetPackagePath(root string) string {
	return fmt.Sprintf("%s\\package.json", root)
}
func cleanVersion(v string) string {
	v = strings.ReplaceAll(v, "^", "")
	v = strings.ReplaceAll(v, "~", "")
	v = strings.ReplaceAll(v, ">", "")
	return v
}
func VerifyVersion(src string, target string) bool {
	sourceVersion := strings.Split(cleanVersion(src), ".")
	targetVersion := strings.Split(cleanVersion(target), ".")
	if targetVersion[0] > sourceVersion[0] {
		return true
	} else if targetVersion[1] > sourceVersion[1] {
		return true
	} else if targetVersion[2] > sourceVersion[2] {
		return true
	}
	return false
}
