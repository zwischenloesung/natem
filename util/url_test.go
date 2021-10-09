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
	"testing"
)

// TODO: ATM some of the tests here are almost like 'testing upstream'
// (which we will have to clean up later), the idea is to test the basics
// here first to make sure the later tests succeed/fail for the right reasons.
func TestParseURI(t *testing.T) {
	a := "urn:uuid:00000000-0000-0000-0000-000000000000"
	b, e := ParseURI(a)
	if e != nil {
		t.Fatal("parsing the URI (urn) did not work as expected")
	}
	if b.Scheme != "urn" {
		t.Fatal("setting the 'Scheme' value did not work as expected while parsing the URI (urn)")
	}
	a = "https://example.org"
	b, e = ParseURI(a)
	if e != nil {
		t.Fatal("parsing the URI (https) did not work as expected")
	}
	if b.Host != "example.org" {
		t.Fatal("setting the 'Host' value did not work as expected while parsing the URI (https)")
	}
	if e != nil {
		t.Fatal("parsing the URI (https) did not work as expected")
	}
	if b.Path != "" {
		t.Fatal("'Path' value is expected to be empty while parsing this URI (https)")
	}
	b.Path = "foo/bar.html"
	if b.String() != "https://example.org/foo/bar.html" {
		t.Fatal("manually setting the 'Path' value did not work as expected for URI (https)")
	}
	a = "tag:test@example.org,1970:foobar"
	b, e = ParseURI(a)
	if e != nil {
		t.Fatal("parsing the URI (tag) did not work as expected")
	}
	if b.Opaque != "test@example.org,1970:foobar" {
		t.Fatalf("The value of `Opaque` is not as expected: %s.\n", b.Opaque)
	}
}

func TestParseThingURL(t *testing.T) {
	a := "testing/example.yml"
	p := "/home/foo"
	b, e := ParseThingURL(a, p)
	if e != nil {
		t.Fatal("parsing the URI (file) did not work as expected: got an error")
	}
	if b.Path != p+"/"+a {
		t.Fatal("parsing the URI (file) did not work as expected: got wrong path")
	}
	a = "/tmp/somefile.suffix"
	b, e = ParseThingURL(a, p)
	if b.Scheme != "file" {
		t.Fatalf("The string should have been read as a file reference %s.\n", b.String())
	}
}

func TestGetFilePathFromURL(t *testing.T) {
	a := "testing/example.yml"
	p := "/home/foo"
	b, e := GetFilePathFromURL(a, p)
	if e != nil || b != p+"/"+a {
		t.Fatalf("parsing the URI (file) did not work as expected, result is: %s", b)
	}
}
