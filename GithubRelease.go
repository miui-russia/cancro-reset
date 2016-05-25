package main

type GithubRelease struct {
	TagName string `json:"tag_name"`
}

func (gr *GithubRelease) DownloadUrl() string {
	return "https://github.com/opengapps/arm/releases/download/" + gr.TagName + "/open_gapps-arm-6.0-pico-" + gr.TagName + ".zip"
}