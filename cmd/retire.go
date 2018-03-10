// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// retireCmd represents the retire command
var retireCmd = &cobra.Command{
	Use:   "retire",
	Short: "Retires a replicant or multiple replicants",
	Long: `

	- deckard retire frontend-85737-23423
	- deckard retire deployment frontend
	- deckard retire daemonset logging`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("retire called")
	},
}

func init() {
	RootCmd.AddCommand(retireCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// retireCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// retireCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
