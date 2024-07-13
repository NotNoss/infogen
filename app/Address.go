package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/NotNoss/infogen/util"
)

type Content struct {
	Content Addresses `json:"contents"`
}

type Addresses struct {
	Addresses []AddressInfo `json:"addresses"`
}

type AddressInfo struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	Country    string `json:"country"`
	PostalCode string `json:"postalcode"`
	State      string `json:"state"`
}

func Address() {
	url := "https://fake-identity-generation.p.rapidapi.com/identity/person/address"

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

	var responseObject Content
	json.Unmarshal(responseData, &responseObject)

	fmt.Println("Address: " + responseObject.Content.Addresses[0].Address)
	fmt.Println("City: " + responseObject.Content.Addresses[0].City)
	fmt.Println("State: " + responseObject.Content.Addresses[0].State)
	fmt.Println("Postal Code: " + responseObject.Content.Addresses[0].PostalCode)
	fmt.Println("Country: " + responseObject.Content.Addresses[0].Country)
	os.Exit(0)
}
