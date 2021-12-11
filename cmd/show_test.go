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
func TestExecuteShowHelp(t *testing.T) {
	a := bytes.NewBufferString("")
	b := bytes.NewBufferString("")
	rootCmd.SetOut(a)
	rootCmd.SetArgs([]string{"help", "show"})
	rootCmd.Execute()
	aOut, err := io.ReadAll(a)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"show", "--help"})
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
func TestExecuteShowOptions(t *testing.T) {
	rootCmd.SetArgs([]string{"show"})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatal("The show command does not exist.\n", err)
	}
	a := bytes.NewBufferString("")
	b := bytes.NewBufferString("")
	rootCmd.SetOut(a)
	rootCmd.SetArgs([]string{"show"})
	rootCmd.Execute()
	aOut, err := io.ReadAll(a)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"show", "-t", "foo"})
	rootCmd.Execute()
	bOut, err := io.ReadAll(a)
	if err != nil {
		t.Fatal(err)
	}
	if string(aOut) == string(bOut) {
		t.Fatal("The show command should have failed as the required -t was missing in one call.")
	}
}
