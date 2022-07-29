/*
Copyright © 2022 CISCOQUAN <jesusanyasimeon@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"
	"bufio"
	"github.com/spf13/cobra"
	"encoding/json"
	"strings"
	"elp-cli/request"
)

type searchit struct {
	Message string 
}

var out searchit
var url string = "http://127.0.0.1:8000/api/add"

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

		if command == "" || title == "" || framework ==""{
			fmt.Println("fields can't be empty")
		}else{

			values := map[string]string{"framework": framework,"title": title,"command":command}
			jsonData, err := json.Marshal(values)
				
				if err != nil {
					fmt.Println(err)}
				
				status,data := request.Post(url,jsonData)
				if string(status) == "201 Created"{
					ata := []byte(data)
					json.Unmarshal(ata,&out)
					fmt.Println(out.Message)
					fmt.Println(status)
					
				}else{
					fmt.Println("check your internet connection or server is down")	
				}
		}

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
