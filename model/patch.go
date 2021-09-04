package model

type Patch struct {
	PackageName  string `json:"packageName"`
	PackageUrl   string `json:"packageUrl"`
	MinVersion   string `json:"minVersion"`
	ForceUpgrade bool   `json:"forceUpgrade"`
}
