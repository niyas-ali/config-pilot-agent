package model

type PatchConfiguration struct {
	PackageName  string `json:"packageName"`
	MinVersion   string `json:"minVersion"`
	ForceUpgrade bool   `json:"forceUpgrade"`
}
