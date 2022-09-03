package request

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)
func Post(url string, json_data []byte)(string, string, error){
	hello := bytes.NewBuffer(json_data)
	req, err := http.NewRequest(http.MethodPost, url, hello)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		os.Exit(-1)}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(errors.New("error when sending request to the server"))
		os.Exit(-1)
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	return resp.Status, string(responseBody), err
}
