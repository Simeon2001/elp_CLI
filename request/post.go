package request

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"net/http"
)
func Post(url string, json_data []byte)(string, string){
	hello := bytes.NewBuffer(json_data)
	req, err := http.NewRequest(http.MethodPost, url, hello)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		fmt.Println(err)}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		fmt.Println(err)
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	
	return resp.Status, string(responseBody)
}
