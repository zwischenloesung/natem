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
	Long:  `Show a representation of the information stored in the knowledge base.`,
	Run: func(cmd *cobra.Command, args []string) {

		viper.BindPFlag("context", rootCmd.PersistentFlags().Lookup("context"))
		context := viper.GetString("context")

		viper.BindPFlag("thing", cmd.PersistentFlags().Lookup("thing"))
		thing := viper.GetString("thing")

		viper.BindPFlag("parameters", cmd.PersistentFlags().Lookup("parameters"))
		par := viper.GetString("parameters")

		viper.BindPFlag("behavior", cmd.PersistentFlags().Lookup("behavior"))
		beh := viper.GetString("behavior")

		viper.BindPFlag("categories", cmd.PersistentFlags().Lookup("categories"))
		cat := viper.GetBool("categories")

		viper.BindPFlag("relations", cmd.PersistentFlags().Lookup("relations"))
		rel := viper.GetString("relations")

		if beh != "" {
			ShowBehavior(context, thing, beh)
		}
		if cat {
			ShowRelation(context, thing, "is")
		}
		if rel != "" {
			ShowRelation(context, thing, rel)
		}
		if par != "" {
			ShowParameter(context, thing, par)
		} else if beh == "" && !cat && rel == "" {
			ShowParameter(context, thing, "*")
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
	showCmd.PersistentFlags().StringP("parameters", "P", "", "display the values set in 'parameters' (default)")
	showCmd.PersistentFlags().StringP("behavior", "B", "", "display the capabilities set in 'behaviour:'")
	showCmd.PersistentFlags().BoolP("categories", "C", false, "display the category hierarchies set in 'relation:is'")
	showCmd.PersistentFlags().StringP("relations", "R", "", "display the relations set in 'relation'")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ShowBehavior(context string, thing string, behavior string) {
	fmt.Println("util.ShowBehavior(", context, ",", thing, ",", behavior, ") called")
}

func ShowParameter(context string, thing string, parameter string) {
	fmt.Println("util.ShowParameter(", context, ",", thing, ",", parameter, ") called")
	theThing, e := util.ParseThingFromFile(thing)
	if e != nil {
		log.Fatalf("Could not parse Thing from file: %s.\n", e)
	}
	fmt.Println(theThing)
}

func ShowRelation(context string, thing string, relation string) {
	fmt.Println("util.ShowRelation(", context, ",", thing, ",", relation, ") called")
}
