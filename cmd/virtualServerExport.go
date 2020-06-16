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
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
	"github.com/smerrell/f5ToTf/cliutil"
	"github.com/spf13/cobra"
)

// virtualServerExportCmd represents the virtualServerExport command
var virtualServerExportCmd = &cobra.Command{
	Use:   "virtualServerExport",
	Short: "Export a Virtual Server to Terraform code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("virtualServerExport called")
		passwd := cliutil.PromptForPassword()
		fmt.Printf(passwd)
		session, err := bigip.NewTokenSession("", "443", "", passwd, "tmos", nil)
		if err != nil {
			log.Fatal(err)
		}

		server, err := session.GetVirtualServer("")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v", server)
	},
}

func init() {
	rootCmd.AddCommand(virtualServerExportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// virtualServerExportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// virtualServerExportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
