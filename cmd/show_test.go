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
	"io/ioutil"
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
	aOut, err := ioutil.ReadAll(a)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"show", "--help"})
	rootCmd.Execute()
	bOut, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(aOut) != string(bOut) {
		t.Fatalf("expected the same output for `help` and `--help`, but got ...\n\"%s\"\n ... and ... \n\"%s\"", string(aOut), string(bOut))
	}
}

// Test if all expected options are there
func TestExecuteShowOptions(t *testing.T) {
	// TODO
	rootCmd.SetArgs([]string{"show"})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatal("the show command is expected to fail as the required -t is missing")
	}
	a := bytes.NewBufferString("")
	rootCmd.SetOut(a)
	rootCmd.SetArgs([]string{"show", "-t", "foo"})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatal("show command is expected to fail as the required -t is missing")
	}
}
