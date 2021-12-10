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

package util

import (
	"strings"
	"testing"
)

// Make sure the basic structure is available
func TestThing(t *testing.T) {
	var thing Thing
	thing.Name = "foo"
	if thing.Name != "foo" {
		t.Fatal("the struct Thing did not work as expected")
	}
}

func TestNewThing(t *testing.T) {

	a := NewThing()
	if a.Id == "" {
		t.Fatal("Failed to set the UUID for the new Thing.")
	}
}

func TestValidateThing(t *testing.T) {

	d := "{ \"_version\": \"0.1\" }"
	data := []byte(d)
	s := "{ \"type\": \"object\", \"properties\": { \"_version\": { \"type\": \"string\" } } }"
	sc := []byte(s)
	if !ValidateJSONThing(sc, data) {
		t.Fatal("this (JSON) should have validated successfully...")
	}
	s = "{ \"type\": \"object\", \"properties\": { \"_version\": { \"type\": \"array\" } } }"
	sc = []byte(s)
	t.Log("Now failing successfully (wrong type):")
	if ValidateJSONThing(sc, data) {
		t.Fatal("this should not have validated...")
	} else {
		t.Log("this document failed to validate (which is good).")
	}
	d = "---\n_version: \"0.1\""
	data = []byte(d)
	s = "---\ntype: \"object\"\nproperties:\n  _version:\n    type: \"string\""
	sc = []byte(s)
	if !ValidateYAMLThing(sc, data) {
		t.Fatal("this (YAML) should have validated successfully...")
	}
}

func TestParseThing(t *testing.T) {

	a := "---\nname: 'example'\n"
	b := []byte(a)
	c := ParseThing(b)
	if c.Name != "example" {
		t.Fatal("the YAML parser did not work as expected")
	}
}

func TestParseThingFromtFile(t *testing.T) {

	a := ParseThingFromFile("testing/example.yml")
	if a.Name != "example" {
		t.Fatal("Parsing YAML from file did not work as expected.")
	}
}

func TestSerializeThing(t *testing.T) {

	a := NewThing()
	a.Name = "example"
	a.Id = ""
	b, e := SerializeThing(a)
	//t.Log(string(b))
	if e != nil {
		t.Fatal("Serializing Thing failed.")
	}
	if strings.Contains(string(b), "id: \"\"") {
		t.Fatal("Incorrectly serialized Thing.")
	}
	if !strings.Contains(string(b), "name: example") {
		t.Fatal("Incorrectly serialized Thing.")
	}
}
