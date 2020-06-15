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
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "f5ToTf",
	Short: "Export Big-IP resources to Terraform code",
	Long: `Exports Big-IP resources to corresponding Terraform code. This code
uses the BigIp Terraform provider. The intention of this tool is to simplify the
process of using Terraform's import command. Currently, Terraform import does
not create a corresponding code for the state it imports. This tool fills that
gap
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.f5ToTf.yaml)")
	rootCmd.PersistentFlags().StringP("username", "u", "", "User to access the Big-IP system as")
	rootCmd.PersistentFlags().StringP("address", "a", "", "Url of the Big-IP instance")

	// Is this wonky? Yes it is. See this GitHub comment as to why
	// https://github.com/spf13/cobra/issues/206#issuecomment-471959800
	pf := rootCmd.PersistentFlags()
	cobra.MarkFlagRequired(pf, "username")
	cobra.MarkFlagRequired(pf, "address")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".f5ToTf" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".f5ToTf")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
