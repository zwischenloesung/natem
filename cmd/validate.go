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
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate",
	Long: `
Validate the thing provided against a schema definition and report back
to the user.`,
	Run: func(cmd *cobra.Command, args []string) {

		viper.BindPFlag("context", rootCmd.PersistentFlags().Lookup("context"))
		context := viper.GetString("context")

		viper.BindPFlag("thing", cmd.PersistentFlags().Lookup("thing"))
		thing := viper.GetString("thing")

		viper.BindPFlag("schema", cmd.PersistentFlags().Lookup("schema"))
		schema := viper.GetString("schema")

		ValidateThing(context, thing, schema)
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	validateCmd.PersistentFlags().StringP("thing", "t", "", "the thing to be validated, either in URL or short form")
	validateCmd.MarkPersistentFlagRequired("thing")
	validateCmd.PersistentFlags().StringP("schema", "s", "", "the schema to use for validation against, either in URL or short form")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ValidateThing(context string, thing string, schema string) {
	fmt.Println("util.ValidateThingFromFile(", context, ",", thing, ",", schema, ") called")
}
