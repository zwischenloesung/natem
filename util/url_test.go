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

func TestGetPathFromThingURL(t *testing.T) {
	a := "testing/example.yml"
	p := "/home/foo"
	b, e := GetPathFromThingURL(a, p)
	if e != nil || b != p+"/"+a {
		t.Fatalf("parsing the URI (file) did not work as expected, result is: %s", b)
	}
}
