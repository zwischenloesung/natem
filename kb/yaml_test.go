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
	"net/url"
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

func TestParseThing(t *testing.T) {

	a := "---\n_name: 'example'\n"
	b := []byte(a)
	c := ParseThing(b)
	if c.Name != "example" {
		t.Fatal("the YAML parser did not work as expected")
	}
}

func TestParseThingFromtFile(t *testing.T) {

	a := ParseThingFromFile("testing/example.yml")
	if a.Name != "example" {
		t.Fatal("parsing YAML from file did not work as expected")
	}
}

func TestParseURL(t *testing.T) {
	a := "urn:uuid:00000000-0000-0000-0000-000000000000"
	_, e := url.Parse(a)
	if e != nil {
		t.Fatal("parsing the URI (urn) did no work as expected")
	}
	a = "https://example.org"
	_, e = url.Parse(a)
	if e != nil {
		t.Fatal("parsing the URI (https) did no work as expected")
	}
}

func TestParseFileURL(t *testing.T) {
	a := "testing/example.yml"
	b, e := ParseFileURL(a)
	if e != nil {
		t.Fatal("parsing the URI (file) did no work as expected: got an error")
	}
	if b.Path != a {
		t.Fatal("parsing the URI (file) did no work as expected: got wrong path")
	}
}

func TestGetFilePathFromURL(t *testing.T) {
	a, e := GetFilePathFromURL("tessting/example.yml")
	if e != nil && a == "testing/example.yml" {
		t.Fatal("parsing the URI (file) did no work as expected")
	}
}
