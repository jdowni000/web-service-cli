/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "web-service-cli",
	Short: "A simple cli tool that hits a local running image/container on port 8080",
	Long: `A simple cli tool designed to help retrieve json data from a local running 
image/container on port 8080. This tool was designed specifically to target 
https://github.com/jdowni000/gameserver. You may hit the root that retrieves condensed
information for all games available, or provide the id with the flag -g with the id number.

For example: 
	web-service-cli webservice
	
	or

	web-service-cli webservice -g 1
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.web-service-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
