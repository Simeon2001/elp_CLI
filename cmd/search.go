/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"elp-cli/request"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type Output []struct {
	Title string  `json:"title"`
	Command string `json:"command"`
	Framework string `json:"framework"`
}

var data Output

var link string = "http://127.0.0.1:8000/api/query"


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
		info := strings.Join(args, " ")
		values := map[string]string{"search": info}
		jsonData, err := json.Marshal(values)
		if err != nil {
			fmt.Println(err)}
		
		_,out := request.Post(link,jsonData)
		ata := []byte(out)
		json.Unmarshal(ata,&data)
		fmt.Println(data)
		for _, value := range data{
		fmt.Println(value.Command,value.Framework,value.Title)
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

