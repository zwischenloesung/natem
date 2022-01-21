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
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

// Make sure the basic structure is available
func TestThing(t *testing.T) {
	var thing Thing
	thing.Id.Name = "foo"
	if thing.Id.Name != "foo" {
		t.Fatal("the struct Thing did not work as expected")
	}
}

func TestNewThing(t *testing.T) {

	a := NewThing()
	if strings.HasPrefix("urn:uuid:", a.Id.Uuid) {
		t.Fatal("Failed to set the UUID for the new Thing.")
	}
}

func TestValidateThing(t *testing.T) {

	d := "{ \"id\": { \"version\": \"0.1\" } }"
	data := []byte(d)
	s := "{ \"type\": \"object\", \"properties\": { \"id\": { \"type\": \"object\" } } }"
	sc := []byte(s)
	r, e := ValidateJSONThing(sc, data)
	if e != nil {
		t.Fatalf("Got an error validating a Thing: %s.\n", e)
	}
	if !r {
		t.Fatal("this (JSON) should have validated successfully...")
	}
	// This time via the YAML parser..
	r, e = ValidateThing(sc, data)
	if e != nil {
		t.Fatalf("Got an error validating a Thing: %s.\n", e)
	}
	if !r {
		t.Fatal("this (JSON) should have validated successfully...")
	}
	s = "{ \"type\": \"object\", \"properties\": { \"id\": { \"type\": \"array\" } } }"
	sc = []byte(s)
	t.Log("Now failing successfully (wrong type):")
	r, e = ValidateJSONThing(sc, data)
	if e != nil {
		t.Fatalf("Got an error validating a Thing: %s.\n", e)
	}
	if r {
		t.Fatal("this should not have validated...")
	} else {
		t.Log("this document failed to validate (which is good).")
	}
	d = "---\nid:\n  version: \"0.1\""
	data = []byte(d)
	s = "---\ntype: \"object\"\nproperties:\n  id:\n    type: \"object\"\n    properties:\n      version:\n        type: \"string\""
	sc = []byte(s)
	r, e = ValidateThing(sc, data)
	if e != nil {
		t.Fatalf("Got an error validating a Thing: %s.\n", e)
	}
	if !r {
		t.Fatal("this (YAML) should have validated successfully...")
	}
}

func TestParseThing(t *testing.T) {

	a := "---\nid:\n  version: \"0.1\""
	b := []byte(a)
	c, e := ParseThing(b)
	if e != nil {
		t.Fatalf("Error parsing the Thing: %s.\n", e)
	}
	if c.Id.Version != "0.1" {
		t.Fatal("the YAML parser did not work as expected")
	}
}

func TestParseThingFromtFile(t *testing.T) {

	a, e := ParseThingFromFile("testing/example.yml")
	if e != nil {
		t.Fatalf("Error parsing the Thing from file: %s.\n", e)
	}
	if a.Id.Name != "example" {
		t.Fatal("Parsing YAML from file did not work as expected.")
	}
}

func TestSerializeThing(t *testing.T) {

	a := NewThing()
	t.Log("a address: ", &a)
	a.Id.Name = "example"
	a.Id.Uuid = ""
	t.Log("a address: ", &a)
	b, e := SerializeThing(a)
	if e != nil {
		t.Fatal("Serializing Thing failed.")
	}
	if !strings.Contains(string(b), "uuid:") {
		t.Fatal("Incorrectly serialized Thing, no UUID.")
	}
	if strings.Contains(string(b), "uuid: \"\"") {
		t.Fatal("Incorrectly serialized Thing, empty UUID.")
	}
	if !strings.Contains(string(b), "name: example") {
		t.Fatal("Incorrectly serialized Thing, Name not set.")
	}
}

func TestSerializeThingToFile(t *testing.T) {

	a, _ := ParseThingFromFile("testing/example.yml")
	f, err := os.CreateTemp("testing", "natem")
	defer os.Remove(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	err = SerializeThingToFile(&a, f.Name())
	if err != nil {
		t.Fatal(err)
	}
	b, err := ParseThingFromFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	if a.Id.Uuid != b.Id.Uuid {
		t.Fatal("The Ids did not match.")
	}
}

func TestWriteThingFile(t *testing.T) {

	a, _ := ParseThingFromFile("testing/example.yml")
	b, e := os.Getwd()
	if e != nil {
		t.Fatalf("Could not get current working directory: %s.\n", e)
	}
	rand.Seed(int64(time.Now().Nanosecond()))
	c := strconv.Itoa(rand.Intn(900000) + 100000)
	d := filepath.Join(b, "testing", "_thing_test_"+c)
	p, f, e := WriteThingFile(&a, d, b, true, true)
	if e != nil {
		t.Fatalf("Failed to create file: %s.\n", e)
	} else {
		t.Logf("Created this file: %s, in: %s.\n", f, p)
	}
	p, f, e = WriteThingFile(&a, d, b, true, false)
	if e == nil {
		t.Fatal("This time the file creation should have failed...\n")
	} else {
		t.Log("Successfully faild to create an already existing file.")
	}
	p, f, e = WriteThingFile(&a, d, b, true, true)
	if e != nil {
		t.Fatal("This time it should have worked though...\n")
	}
	os.Remove(d)
}

func TestCreateNewThingFile(t *testing.T) {

	b, e := os.Getwd()
	if e != nil {
		t.Fatalf("Could not get current working directory: %s.\n", e)
	}
	rand.Seed(int64(time.Now().Nanosecond()))
	c := strconv.Itoa(rand.Intn(900000) + 100000)
	d := filepath.Join(b, "testing", "_thing_test_"+c)
	a, e := CreateNewThingFile(d, b, false)
	if e != nil {
		t.Fatalf("Failed to create a brand new Thing file: %s.\n", e)
	}
	if a.Id.Name != "" {
		t.Fatal("The name was not empty..")
	}
	os.Remove(d)
}
