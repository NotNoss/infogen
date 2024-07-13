package util

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Domain string
	Key    string
}

func (c Config) GetDomain() string {
	username := GetUser()

	filePath := fmt.Sprintf("/home/%s/.config/infogen/config.toml", username)

	config, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	err = toml.Unmarshal([]byte(config), &c)
	if err != nil {
		fmt.Println(err)
	}

	return c.Domain
}

func (c Config) GetKey() string {
	username := GetUser()

	filePath := fmt.Sprintf("/home/%s/.config/infogen/config.toml", username)

	config, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	err = toml.Unmarshal([]byte(config), &c)
	if err != nil {
		fmt.Println(err)
	}

	return c.Key
}
