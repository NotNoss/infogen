package util

import (
	"fmt"
	"os/user"
)

func GetUser() string {
	user, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
	}

	return user.Username
}
