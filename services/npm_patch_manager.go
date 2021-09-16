package services

import (
	"config-pilot-agent/model"
	"config-pilot-agent/services/json_parser"
	"fmt"
	"log"

	"github.com/tidwall/sjson"
)

type NpmPatchManager struct {
	patchManager  *PatchManager
	Name          string
	packages      model.NpmPackage
	Dependency    []*model.NpmDependencies
	DevDependency []*model.NpmDependencies
	RequireUpdate bool
}

func (patch *NpmPatchManager) LoadPatchData() {
	log.Println("loading package dependencies")
	err := json_parser.JsonToModel(model.GetPackagePath(patch.Name), &patch.packages)
	if err != nil {
		log.Fatalln("could not find package.json file")
		return
	}
	patch.packages.OriginalContent, err = json_parser.LoadFile(model.GetPackagePath(patch.Name))
	if err != nil {
		log.Fatalln("parsing package.json failed")
	}
	for p, v := range patch.packages.Dependencies {
		dependecy := model.NpmDependencies{Name: p, Value: v}
		patch.Dependency = append(patch.Dependency, &dependecy)
	}
	for p, v := range patch.packages.DevDependencies {
		dependecy := model.NpmDependencies{Name: p, Value: v}
		patch.DevDependency = append(patch.DevDependency, &dependecy)
	}
}

func (patch *NpmPatchManager) VerifyAndUpgradePatches() {
	log.Println("verifying package dependencies")
	for _, original := range patch.patchManager.patches {
		for _, current := range patch.Dependency {
			if current.Name == original.PackageName {
				log.Println("found matching package:", current.Name)
				if model.VerifyVersion(current.Value, original.MinVersion) {
					log.Printf("found upgrade for %s with version:%s", current.Name, original.MinVersion)
					if original.ForceUpgrade {
						current.Value = original.MinVersion
						log.Printf("force upgrading -> done")
						patch.RequireUpdate = true
					} else {
						log.Println("force upgrade set to false for this package:", current.Name, "version:", current.Value)
					}
				} else {
					log.Println("package is upto date:", current.Name, "version:", current.Value)
				}
			}
		}
		for _, current := range patch.DevDependency {
			if current.Name == original.PackageName {
				log.Println("found matching package:", current.Name)
				if model.VerifyVersion(current.Value, original.MinVersion) {
					log.Printf("found upgrade for %s with version:%s", current.Name, original.MinVersion)
					if original.ForceUpgrade {
						current.Value = original.MinVersion
						log.Printf("force upgrading -> done")
					}
				}
			}
		}
	}
}
func (patch *NpmPatchManager) SaveChanges() {
	var content string = patch.packages.OriginalContent
	for _, current := range patch.Dependency {
		content, _ = sjson.Set(content, fmt.Sprintf("dependencies.%s", current.Name), current.Value)
	}
	for _, current := range patch.DevDependency {
		content, _ = sjson.Set(content, fmt.Sprintf("dependencies.%s", current.Name), current.Value)
	}
	json_parser.JsonToFile(content, model.GetPackagePath(patch.Name))
}
