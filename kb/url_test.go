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

func TestParseURL(t *testing.T) {
	a := "urn:uuid:00000000-0000-0000-0000-000000000000"
	b, e := ParseURL(a)
	if e != nil {
		t.Fatal("parsing the URI (urn) did no work as expected")
	}
	if b.Scheme != "urn" {
		t.Fatal("setting the 'Scheme' value did no work as expected while parsing the URI (urn)")
	}
	a = "https://example.org"
	b, e = ParseURL(a)
	if e != nil {
		t.Fatal("parsing the URI (https) did no work as expected")
	}
	if b.Host != "example.org" {
		t.Fatal("setting the 'Host' value did no work as expected while parsing the URI (https)")
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
	a, e := GetFilePathFromURL("testing/example.yml")
	if e != nil || a != "testing/example.yml" {
		t.Fatalf("parsing the URI (file) did no work as expected, result is: %s", a)
	}
}
