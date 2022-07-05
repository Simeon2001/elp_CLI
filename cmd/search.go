/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"reflect"
)

func connect(json_data []byte){
	hello := bytes.NewBuffer(json_data)
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8000/api/query", hello)
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

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _,info := range args {
			data := info
			values := map[string]string{"search": data}
			jsonData, err := json.Marshal(values)
			
			if err != nil {
				fmt.Println(err)}
			
			connect(jsonData)
			fmt.Println(reflect.TypeOf(jsonData))
		}
		
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
