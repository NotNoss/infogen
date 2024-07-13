package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/NotNoss/infogen/util"
)

type ID struct {
	Content Person `json:"contents"`
}

type Person struct {
	PersonName []string `json:"names"`
}

func Name() {
	url := "https://fake-identity-generation.p.rapidapi.com/identity/person/name"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	config := util.Config{}
	key := config.GetKey()

	request.Header.Add("accept", "application/json")
	request.Header.Add("x-rapidapi-host", "fake-identity-generation.p.rapidapi.com")
	request.Header.Add("x-rapidapi-key", key)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var responseObject ID
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Content.PersonName[0])
	os.Exit(0)
}
