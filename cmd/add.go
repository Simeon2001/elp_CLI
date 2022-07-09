/*
Copyright © 2022 CISCOQUAN <jesusanyasimeon@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"
	"bufio"
	"github.com/spf13/cobra"
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"strings"
)

func conne(json_data []byte){
	hello := bytes.NewBuffer(json_data)
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8000/api/add", hello)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		fmt.Println(err)}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))
}


// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

// where to take input from user

		readfr := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the language or framework name: ")
		frameworks, _ := readfr.ReadString('\n')
		framework := strings.Trim(frameworks, "\r\n")

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter a short description about the command: ")
		titles, _ := reader.ReadString('\n')
		title := strings.Trim(titles, "\r\n")

		readit := bufio.NewReader(os.Stdin)
		fmt.Println("Type Command: ")
		commands, _ := readit.ReadString('\n')
		command := strings.Trim(commands, "\r\n")

		values := map[string]string{"framework": framework,"title": title,"command":command}
		fmt.Print(values)
		jsonData, err := json.Marshal(values)
			
			if err != nil {
				fmt.Println(err)}
			
			conne(jsonData)


	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
