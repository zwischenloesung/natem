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

func TestParseThingURL(t *testing.T) {
	a := "testing/example.yml"
	b := true
	c := "/home/foo"
	d, e := ParseThingURL(a, c, b)
	if e != nil {
		t.Fatalf("parsing the path did not work as expected: got an error: %s.\n", e)
	}
	if d.Path != c+"/"+a {
		t.Fatal("parsing the path did not work as expected: got wrong path")
	}
	if d.Scheme != "file" {
		t.Fatalf("The string should have been read as a file reference %s.\n", d.String())
	}
	a = "testing/example.yml"
	c = "file:///home/foo"
	if d.Scheme+"://"+d.Path != c+"/"+a {
		t.Fatalf("parsing the URI (file) did not work as expected: got wrong path %s://%s.\n", d.Scheme, d.Path)
	}
	a = "/tmp/somefile.suffix"
	c = "https://example.org/home/foo"
	d, e = ParseThingURL(a, c, false)
	if d.Path != a {
		t.Fatal("parsing the URI (file) did not work as expected: did not get the original path back.\n")
	}
	a = "file:///home/foo/somefile.suffix"
	c = "file:///home/foo"
	d, e = ParseThingURL(a, c, b)
	if d.Scheme+"://"+d.Path != a {
		t.Fatal("parsing the URI (file) did not work as expected: both parameters should accept a URL or a local path string.")
	}
	t.Log("Now failing successfully (URL-Scheme):")
	a = "http://example.com/tmp/somefile.suffix"
	c = "file:///home/foo"
	d, e = ParseThingURL(a, c, b)
	if e == nil {
		t.Fatalf("parsing different URLs should produce an error...")
	} else {
		t.Logf("got the expected error: %s.\n", e)
	}
	t.Log("Now failing successfully (URL-Path):")
	a = "file:///tmp/somefile.suffix"
	c = "file:///home/foo"
	d, e = ParseThingURL(a, c, b)
	if e == nil {
		t.Fatalf("parsing different paths should produce an error...")
	} else {
		t.Logf("got the expected error: %s.\n", e)
	}
	a = "file:///tmp/somefile.suffix"
	c = "file:///home/foo"
	d, e = ParseThingURL(a, c, false)
	if e == nil {
		t.Logf("parsing different paths is now ok...")
	} else {
		t.Fatalf("got the unexpected error: %s.\n", e)
	}
	t.Log("Now failing successfully (URL-Path):")
	a = "file:///tmp/../somefile.suffix"
	d, e = ParseThingURL(a, c, b)
	if e != nil {
		t.Logf("climbing up directories is forbidden. %s.\n", e)
	} else {
		t.Fatalf("climbing up directories should have produced an error.\n")
	}
	a = "file:///tmp/../somefile.suffix"
	c = "file:///home/foo"
	d, e = ParseThingURL(a, c, false)
	if e == nil {
		t.Logf("climbing up directories is now ok...")
	} else {
		t.Fatalf("got the unexpected error: %s.\n", e)
	}
}

func TestGetThingURLPath(t *testing.T) {
	a := "testing/example.yml"
	b := true
	p := "/home/foo"
	d, e := GetThingURLPath(a, p, b)
	if e != nil || d != p+"/"+a {
		t.Fatalf("parsing the URI (file) did not work as expected, result is: %s", d)
	}
	t.Log("Now failing successfully (Thing inside Context)")
	a = "/testing/example.yml"
	p = "/home/foo"
	d, e = GetThingURLPath(a, p, b)
	if e != nil {
		t.Logf("Expected error was: %s.\n", e)
	} else {
		t.Fatalf("parsing the URI (file) did not work as expected, should have thrown an error. The result is: %s", d)
	}
	t.Log("Now failing successfully (remote Things are not allowed)")
	a = "testing/example.yml"
	p = "https://example.org/home/foo"
	d, e = GetThingURLPath(a, p, b)
	if e != nil {
		t.Logf("Expected error was: %s.\n", e)
	} else {
		t.Fatalf("parsing the URI (file) did not work as expected, should have thrown an error. The result is: %s", d)
	}
}
