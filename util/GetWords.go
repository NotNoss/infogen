package util

import (
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Word string `json:"word"`
}

func GetWords() string {
	response, err := http.Get("https://random-word-api.herokuapp.com/word")
	if err != nil {
		fmt.Println(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	w := string(responseData)

	var str string
	for range w {
		str = fmt.Sprintf(`%s`, w)
		if str[0] == '[' {
			str = str[1 : len(str)-1]
		}

		if str[0] == '"' {
			str = str[1 : len(str)-1]
		}
	}

	return str
}
