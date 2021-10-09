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

func EditThing(project string, thing string, editor string) {
	fmt.Println("EditThing(", project, ",", thing, ",", editor, "", ") called")

	file, _ := util.GetPathFromThingURL(thing, project)

	if editor == "" {
		editor, _ = os.LookupEnv("EDITOR")
	}

	cmd := exec.Command(editor, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println("an error occurred.\n", err)
	}
}

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit View",
	Long: `Edit a visualization of the information stored in the knowledge base.
The focus here lies on the content and behaviour of the things.`,
	Run: func(cmd *cobra.Command, args []string) {

		viper.BindPFlag("project", rootCmd.PersistentFlags().Lookup("project"))
		project := viper.GetString("project")

		viper.BindPFlag("thing", cmd.PersistentFlags().Lookup("thing"))
		thing := viper.GetString("thing")

		viper.BindPFlag("editor", cmd.PersistentFlags().Lookup("editor"))
		editor := viper.GetString("editor")

		EditThing(project, thing, editor)
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
	editCmd.Flags().String("editor", "", "specify the editor of choice (default: Environment Variable $EDITOR)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
