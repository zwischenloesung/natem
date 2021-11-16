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
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/zwischenloesung/natem/util"
)

func EditThing(editor string, thing string, context string, isContextless bool) {

	filePath, err := util.GetThingPathURL(thing, context, !isContextless)

	if err == util.UrlThingOutsideContextError {
		fmt.Println("Use the --context-less switch to force editing Things outside any context:", err)
		return
	} else if err != nil {
		fmt.Println("There was an error.\n", err)
		return
	}

	if editor == "" {
		editor, _ = os.LookupEnv("EDITOR")
	}

	cmd := exec.Command(editor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("An error occurred.\n", err)
	}
}

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open a Thing in an editor",
	Long: `Open an editor and edit the information stored in a Thing of the
knowledge base. The focus here lies on the content and behaviour of the
Things.`,
	Run: func(cmd *cobra.Command, args []string) {

		viper.BindPFlag("context", rootCmd.PersistentFlags().Lookup("context"))
		context := viper.GetString("context")

		viper.BindPFlag("thing", cmd.PersistentFlags().Lookup("thing"))
		thing := viper.GetString("thing")

		viper.BindPFlag("editor", cmd.PersistentFlags().Lookup("editor"))
		editor := viper.GetString("editor")

		viper.BindPFlag("context-less", cmd.PersistentFlags().Lookup("context-less"))
		isContextless := viper.GetBool("context-less")

		EditThing(editor, thing, context, isContextless)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	editCmd.PersistentFlags().StringP("thing", "t", "", "summarize the info for this thing")
	editCmd.MarkPersistentFlagRequired("thing")
	editCmd.PersistentFlags().String("editor", "", "specify the editor of choice (default: Environment Variable $EDITOR)")
	editCmd.PersistentFlags().BoolP("context-less", "C", false, "edit a thing outside of any context")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
