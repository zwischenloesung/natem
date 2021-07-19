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
)

// Test the basics...
func TestExecuteHelp(t *testing.T) {
	cmd := NewRootCmd()
	a := bytes.NewBufferString("")
	b := bytes.NewBufferString("")
	cmd.SetOut(a)
	cmd.SetArgs([]string{"help"})
	cmd.Execute()
	aOut, err := ioutil.ReadAll(a)
	if err != nil {
		t.Fatal(err)
	}
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	cmd.Execute()
	bOut, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(aOut) != string(bOut) {
		t.Fatalf("expected the same output for `help` and `--help`, but got ...\n\"%s\"\n ... and ... \n\"%s\"", string(aOut), string(bOut))
	}
}
