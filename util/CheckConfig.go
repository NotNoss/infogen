package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func CheckConfig() {
	username := GetUser()

	dirPath := fmt.Sprintf("/home/%s/.config/mail-alias/", username)
	filePath := fmt.Sprintf("/home/%s/.config/mail-alias/config.toml", username)

	_, err := os.Stat(dirPath)
	if err != nil {
		os.MkdirAll(dirPath, os.ModePerm)
	}

	_, err = os.Stat(filePath)
	if err != nil {

		response, err := http.Get("https://raw.githubusercontent.com/NotNoss/mail-alias/main/example-config.toml")
		if err != nil {
			fmt.Println(err)
		}

		defer response.Body.Close()

		if response.StatusCode != 200 {
			fmt.Println("unable to get file from github")
		}

		config, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
		}

		defer config.Close()

		_, err = io.Copy(config, response.Body)
		if err != nil {
			fmt.Println(err)
		}
	}
}
