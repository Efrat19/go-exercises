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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "add a task",
	Long: `
add a task`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, name := range args {
			add(name)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func add(name string) {
	task, err := db.AddTask(name)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("task %d added", task.Id))
}
