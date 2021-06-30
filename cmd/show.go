/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	//    "fmt"

	"github.com/spf13/cobra"
	"gitlab.com/zwischenloesung/natem/kb"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show View",
	Long: `Show a visualization of the information stored in the knowledge base.
The focus here lies on the content and behaviour of the things.`,
	Run: func(cmd *cobra.Command, args []string) {

		project, _ := rootCmd.PersistentFlags().GetString("project")

		thing, _ := cmd.PersistentFlags().GetString("thing")

		//default: param, _ := cmd.PersistentFlags().GetBool("parameters")
		act_beh, _ := cmd.PersistentFlags().GetString("actions")
		beh, _ := cmd.PersistentFlags().GetBool("behavior")
		cat, _ := cmd.PersistentFlags().GetBool("categories")
		rel_beh, _ := cmd.PersistentFlags().GetString("relations")

		if beh {
			kb.ShowBehavior(project, thing)
		}
		if cat {
			kb.ShowCategories(project, thing)
		}
		if act_beh != "" {
			kb.ShowActions(project, thing, act_beh)
		}
		if rel_beh != "" {
			kb.ShowRelations(project, thing, rel_beh)
		}
		if !beh && !cat && act_beh == "" && rel_beh == "" {
			kb.ShowParameters(project, thing)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	showCmd.PersistentFlags().StringP("thing", "t", "", "summarize the info for this thing")
	showCmd.MarkPersistentFlagRequired("thing")
	showCmd.PersistentFlags().BoolP("parameters", "P", true, "display the parameters set in 'content' (default)")
	showCmd.PersistentFlags().StringP("actions", "A", "", "display the action defined in 'behaviour:<string>:action'")
	showCmd.PersistentFlags().BoolP("behavior", "B", false, "display the capabilities set in 'behaviour:'")
	showCmd.PersistentFlags().BoolP("categories", "C", false, "display the category hierarchies set in 'behaviour:is'")
	showCmd.PersistentFlags().StringP("relations", "R", "", "display the relations set in 'behaviour:<string>:relations'")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
