package main

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type GithubRelease struct {
	TagName string `json:"tag_name"`
}

func main() {

	app := gin.Default()

	app.StaticFile("/gapps-config.txt", "./resources/gapps-config.txt")

    app.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    app.GET("/opengapps", func(c *gin.Context) {
    	resp, _ := http.Get("https://api.github.com/repos/opengapps/arm/releases/latest")
    	defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

    	var release GithubRelease
    	json.Unmarshal(body, &release)

    	c.String(200, "https://github.com/opengapps/arm/releases/download/" + release.TagName + "/open_gapps-arm-6.0-pico-" + release.TagName + ".zip")
    })
    
    app.Run("0.0.0.0:8080") // listen and server on 0.0.0.0:8080
}
