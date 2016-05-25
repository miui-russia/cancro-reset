package main

import (
	"os"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.StaticFile("/gapps-config.txt", "./resources/gapps-config.txt")

	app.GET("/opengapps.zip", func(c *gin.Context) {
		resp, _ := http.Get("https://api.github.com/repos/opengapps/arm/releases/latest")
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		var release GithubRelease
		json.Unmarshal(body, &release)

		c.Redirect(http.StatusTemporaryRedirect, release.DownloadUrl())
	})

	app.GET("/cm.zip", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "http://download.cyanogenmod.org/get/cancro-latest.zip")
	})

	port := os.Getenv("CANCRO_RESET_PORT")
	if port == "" {
		port = "8080"
	}

	app.Run("0.0.0.0:" + port) // listen and server on 0.0.0.0:8080
}
