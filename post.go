package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Starting the application...")
	jsonData := map[string]string{"moviename1": "TeneT", "moviename2": "Dark", "moviename3": "Fight Club", "moviename": "Joker"}
	jsonValue, _ := json.Marshal(jsonData)

	// ------ POST REQUEST ---
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodPost, "https://api.jsonbin.io/b", bytes.NewReader(jsonValue))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("secret-key", "$2b$10$31DKDWQKwfYeWIkljMsKE.hdW8BJJ3Pq/3YCzLb0i83KTd6r7Yf/.")
	request.Header.Add("private", "false")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	var responsejson map[string]interface{}

	if err := json.Unmarshal(data, &responsejson); err != nil {
		panic(err)
	}

	fmt.Println(responsejson["id"].(string))
	geturl := "https://api.jsonbin.io/b/" + responsejson["id"].(string)
	//fmt.Println(geturl)

	// --- GET REQUEST ---
	getClinet := &http.Client{}
	getRequest, _ := http.NewRequest(http.MethodGet, geturl, nil)
	getResponse, getErr := getClinet.Do(getRequest)
	if getErr != nil {
		panic(getErr)
	}

	defer getResponse.Body.Close()

	getData, getErr := ioutil.ReadAll(getResponse.Body)
	if getErr != nil {
		panic(getErr)
	}
	fmt.Println(string(getData))
	fmt.Println("Terminating the application...")
}
