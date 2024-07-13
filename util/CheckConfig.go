package util

import (
	"fmt"
	"os"
	"os/user"
)

func CheckConfig() {
	user, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
	}

	username := user.Username

	dirPath := fmt.Sprintf("/home/%s/.config/mail-alias/", username)
	filePath := fmt.Sprintf("/home/%s/.config/mail-alias/config.toml", username)

	_, err = os.Stat(dirPath)
	if err != nil {
		os.MkdirAll(dirPath, os.ModePerm)
	}

	_, err = os.Stat(filePath)
	if err != nil {
		config, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
		}

		defer config.Close()
	}
}
