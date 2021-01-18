/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/vgcrld/all-things-go/cmd/funcs"
	"github.com/vgcrld/all-things-go/cmd/messages"

	"github.com/spf13/cobra"
)

// logurusCmd represents the logurus command
var logurusCmd = &cobra.Command{
	Use:   "logurus",
	Short: "A brief description of your command",
	Long:  messages.LogurusCommandLong,
	Run: func(cmd *cobra.Command, args []string) {
		// x := messages.NoConfigurationFound
		funcs.RunLogurus()
	},
}

func init() {
	rootCmd.AddCommand(logurusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logurusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logurusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
