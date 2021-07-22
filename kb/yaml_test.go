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

package kb

import (
	"testing"
)

// Make sure the basic structure is available
func TestThing(t *testing.T) {
	var thing Thing
	thing.Name = "foo"
	if thing.Name != "foo" {
		t.Fatal("the struct Thing does not work as expected")
	}
}

func TestParseThing(t *testing.T) {

	a := "---\n_name: 'example'\n"
	b := []byte(a)
	c := ParseThing(b)
	if c.Name != "example" {
		t.Fatal("the YAML parser does not work as expected")
	}
}

func TestParseThingFromtFile(t *testing.T) {

	a := ParseThingFromFile("testing/example.yml")
	if a.Name != "example" {
		t.Fatal("parsing YAML from file does not work as expected")
	}
}
