/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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

	db "github.com/Efrat19/gophercises/todo/database"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:     "check",
	Aliases: []string{"ck"},
	Short:   "mark task as completed",
	Long: `
mark task as completed!!!!!!!!!!!!!!!!!!`,
	Run: func(cmd *cobra.Command, args []string) {
		check(id)
	},
}

var id int

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().IntVar(&id, "id", 0, "task id to check")
}

func check(id int) {
	newid, err := db.CheckTask(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("task %d checked!", newid))
}
