/*
Copyright © 2021 Muhammad Yarjuna Rohmat <myarjunar@gmail.com>

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
	"errors"
	"fmt"
	"os"

	"github.com/myarjunar/larviscli/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var port string
var baseUrl string
var cfgFile string
var defaultPort string = "8080"
var defaultBaseUrl string = "localhost"

var rootCmd = &cobra.Command{
	Use:   "larviscli <your_name>",
	Short: "Ask larvis to greet you.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("you need to provide your name")
		}

		name := args[0]
		baseUrl = viper.GetString("LARVIS_URL")

		if port == "" {
			port = defaultPort
		}

		if baseUrl == "" {
			baseUrl = defaultBaseUrl
		}

		apiClient := client.NewApiClient(baseUrl, port)

		apiClient.Greet(name)

		return nil
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&port, "port", "", "port")
}

func initConfig() {
	cfgFile = ".larviscli.yaml"

	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
