/*
This is Free Software; feel free to redistribute and/or modify it
under the terms of the GNU General Public License as published by
the Free Software Foundation; version 3 of the License.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

Copyright © 2021 Michael Lustenberger <mic@inofix.ch>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/zwischenloesung/natem/util"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show View",
	Long: `Show a visualization of the information stored in the knowledge base.
The focus here lies on the content and behaviour of the things.`,
	Run: func(cmd *cobra.Command, args []string) {

		viper.BindPFlag("project", rootCmd.PersistentFlags().Lookup("project"))
		project := viper.GetString("project")

		viper.BindPFlag("thing", cmd.PersistentFlags().Lookup("thing"))
		thing := viper.GetString("thing")

		viper.BindPFlag("actions", cmd.PersistentFlags().Lookup("actions"))
		act_beh := viper.GetString("actions")

		viper.BindPFlag("behavior", cmd.PersistentFlags().Lookup("behavior"))
		beh := viper.GetBool("behavior")

		viper.BindPFlag("categories", cmd.PersistentFlags().Lookup("categories"))
		cat := viper.GetBool("categories")

		viper.BindPFlag("relations", cmd.PersistentFlags().Lookup("relations"))
		rel_beh := viper.GetString("relations")

		if beh {
			ShowBehavior(project, thing)
		}
		if cat {
			ShowCategories(project, thing)
		}
		if act_beh != "" {
			ShowActions(project, thing, act_beh)
		}
		if rel_beh != "" {
			ShowRelations(project, thing, rel_beh)
		}
		if !beh && !cat && act_beh == "" && rel_beh == "" {
			ShowParameters(project, thing)
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

func ShowActions(project string, thing string, behavior string) {
	fmt.Println("util.ShowActions(", project, ",", thing, ",", behavior, ") called")
}

func ShowBehavior(project string, thing string) {
	fmt.Println("util.ShowBehavior(", project, ",", thing, ") called")
}

func ShowCategories(project string, thing string) {
	fmt.Println("util.ShowCategories(", project, ",", thing, ") called")
}

func ShowParameters(project string, thing string) {
	fmt.Println("util.ShowParameters(", project, ",", thing, ") called")
	theThing := util.ParseThingFromFile(thing)
	fmt.Println(theThing)
}

func ShowRelations(project string, thing string, behavior string) {
	fmt.Println("util.ShowRelations(", project, ",", thing, ",", behavior, ") called")
}
