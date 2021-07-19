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
func TestExecuteHelp(t *testing.T) {
	a := bytes.NewBufferString("")
	b := bytes.NewBufferString("")
	rootCmd.SetOut(a)
	rootCmd.SetArgs([]string{"help"})
	rootCmd.Execute()
	aOut, err := ioutil.ReadAll(a)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"--help"})
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
func TestExecuteOptions(t *testing.T) {
	// TODO
	rootCmd.SetArgs([]string{"--config", "test.conf"})
	cobra.CheckErr(rootCmd.Execute())
	rootCmd.SetArgs([]string{"--project", "/tmp/foo"})
	cobra.CheckErr(rootCmd.Execute())
}
