package model

import (
	"errors"
	"strings"
)

func trim(v string) string {
	return strings.TrimSpace(v)
}

func (p *PatchConfiguration) Validate() error {
	if trim(p.PackageName) == "" {
		return errors.New("package name is required")
	} else if trim(p.MinVersion) == "" {
		return errors.New("min version is required")
	}
	return nil
}
func (v Configuration) Validate() error {
	if trim(v.CheckoutBranch) == "" {
		return errors.New("checkout branch is required")
	} else if trim(v.PrTitle) == "" {
		return errors.New("pr title is required")
	} else if trim(v.PrMessage) == "" {
		return errors.New("pr message is required")
	}
	if len(v.AzureDevops.Repository) == 0 && len(v.Github.Repository) == 0 {
		return errors.New("repository details are required")
	}
	if len(v.AzureDevops.Repository) > 0 {
		if trim(v.AzureDevops.ProjectName) == "" {
			return errors.New("azure project name is required")
		} else if trim(v.AzureDevops.Organization) == "" {
			return errors.New("azure organization name is required")
		}
		for _, item := range v.AzureDevops.Repository {
			if trim(item.Name) == "" {
				return errors.New("azure repository name is required")
			} else if trim(item.URL) == "" {
				return errors.New("azure repository url is required")
			} else if trim(item.MergeBranch) == "" {
				return errors.New("azure merge branch is required")
			}
		}
	}
	if len(v.Github.Repository) > 0 {
		if trim(v.Github.Organization) == "" {
			return errors.New("github organization name is required")
		}
		for _, item := range v.Github.Repository {
			if trim(item.Name) == "" {
				return errors.New("github repository name is required")
			} else if trim(item.URL) == "" {
				return errors.New("github repository url is required")
			} else if trim(item.MergeBranch) == "" {
				return errors.New("github merge branch is required")
			}
		}
	}
	return nil
}
