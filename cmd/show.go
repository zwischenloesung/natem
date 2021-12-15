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
	"log"

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

		viper.BindPFlag("context", rootCmd.PersistentFlags().Lookup("context"))
		context := viper.GetString("context")

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
			ShowBehavior(context, thing)
		}
		if cat {
			ShowCategories(context, thing)
		}
		if act_beh != "" {
			ShowActions(context, thing, act_beh)
		}
		if rel_beh != "" {
			ShowRelations(context, thing, rel_beh)
		}
		if !beh && !cat && act_beh == "" && rel_beh == "" {
			ShowParameters(context, thing)
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

func ShowActions(context string, thing string, behavior string) {
	fmt.Println("util.ShowActions(", context, ",", thing, ",", behavior, ") called")
}

func ShowBehavior(context string, thing string) {
	fmt.Println("util.ShowBehavior(", context, ",", thing, ") called")
}

func ShowCategories(context string, thing string) {
	fmt.Println("util.ShowCategories(", context, ",", thing, ") called")
}

func ShowParameters(context string, thing string) {
	fmt.Println("util.ShowParameters(", context, ",", thing, ") called")
	theThing, e := util.ParseThingFromFile(thing)
	if e != nil {
		log.Fatalf("Could not parse Thing from file: %s.\n", e)
	}
	fmt.Println(theThing)
}

func ShowRelations(context string, thing string, behavior string) {
	fmt.Println("util.ShowRelations(", context, ",", thing, ",", behavior, ") called")
}
