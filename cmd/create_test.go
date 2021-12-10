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
func TestExecuteCreateHelp(t *testing.T) {
	a := bytes.NewBufferString("")
	b := bytes.NewBufferString("")
	rootCmd.SetOut(a)
	rootCmd.SetArgs([]string{"help", "create"})
	rootCmd.Execute()
	aOut, err := ioutil.ReadAll(a)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"create", "--help"})
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
func TestExecuteCreateOptions(t *testing.T) {
	rootCmd.SetArgs([]string{"create"})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatal("The create command does not exist.\n", err)
	}
	a := bytes.NewBufferString("")
	b := bytes.NewBufferString("")
	rootCmd.SetOut(a)
	rootCmd.SetArgs([]string{"create"})
	rootCmd.Execute()
	aOut, err := ioutil.ReadAll(a)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"create", "-t", "foo"})
	rootCmd.Execute()
	bOut, err := ioutil.ReadAll(a)
	if err != nil {
		t.Fatal(err)
	}
	if string(aOut) == string(bOut) {
		t.Fatal("The create command should have failed as the required -t was missing in one call.")
	}
}
