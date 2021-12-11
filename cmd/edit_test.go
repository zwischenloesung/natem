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
	"bytes"
	"io"
	"testing"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize(initConfig)
}

// Test the basics...
func TestExecuteEditHelp(t *testing.T) {
	a := bytes.NewBufferString("")
	b := bytes.NewBufferString("")
	rootCmd.SetOut(a)
	rootCmd.SetArgs([]string{"help", "edit"})
	rootCmd.Execute()
	aOut, err := io.ReadAll(a)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"edit", "--help"})
	rootCmd.Execute()
	bOut, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(aOut) != string(bOut) {
		t.Fatalf("expected the same output for `help` and `--help`, but got ...\n\"%s\"\n ... and ... \n\"%s\"", string(aOut), string(bOut))
	}
}

// Test if all expected options are there
func TestExecuteEditOptions(t *testing.T) {
	rootCmd.SetArgs([]string{"edit"})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatal("The edit command does not exist.\n", err)
	}
	a := bytes.NewBufferString("")
	b := bytes.NewBufferString("")
	rootCmd.SetOut(a)
	rootCmd.SetArgs([]string{"edit"})
	rootCmd.Execute()
	aOut, err := io.ReadAll(a)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"edit", "-t", "foo"})
	rootCmd.Execute()
	bOut, err := io.ReadAll(a)
	if err != nil {
		t.Fatal(err)
	}
	if string(aOut) == string(bOut) {
		t.Fatal("The edit command should have failed as the required -t was missing in one call.")
	}
}
