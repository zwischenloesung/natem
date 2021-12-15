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
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/zwischenloesung/natem/util"
)

func CreateThing(url string, context string, hasContext bool) {

	_, err := util.CreateNewThingFile(url, context, hasContext)
	if err != nil {
		log.Fatal("Could not create and serialize a new Thing to a file.\n", err)
	}
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Thing",
	Long: `Create a new Thing in the knowledge base. A Thing contains
all the information about and points to an abstract or concrete thing.`,
	Run: func(cmd *cobra.Command, args []string) {

		viper.BindPFlag("context", rootCmd.PersistentFlags().Lookup("context"))
		context := viper.GetString("context")

		viper.BindPFlag("thing", cmd.PersistentFlags().Lookup("thing"))
		thing := viper.GetString("thing")

		viper.BindPFlag("context-less", cmd.PersistentFlags().Lookup("context-less"))
		isContextless := !viper.GetBool("context-less")

		CreateThing(thing, context, isContextless)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	createCmd.PersistentFlags().StringP("thing", "t", "", "the thing")
	createCmd.MarkPersistentFlagRequired("thing")
	createCmd.PersistentFlags().BoolP("context-less", "C", false, "create a thing outside of any context")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
