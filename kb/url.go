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
	"log"
	"net/url"
)

type URI struct {
	*url.URL
}

type Link struct {
	URI
}

var SupportedLinkSchemes = []string{"file", "http", "https"}

func ParseURI(u string) (*URI, error) {
	uri, err := url.Parse(u)
	if err != nil {
		log.Fatal("An error occured while parsing the URI.\n", err)
	}
	return &URI{uri}, err
}

func ParseFileURL(u string) (*URI, error) {
	uri, err := ParseURI(u)
	// Only consider URLs pointing to the local file system, allow for
	// absolute or relative paths too.
	if uri.Scheme == "file" || uri.Scheme == "" {
		return uri, nil
	} else {
		log.Fatal("This URL was not of scheme 'file:///' as expected.\n", err)
		return uri, err
	}
}

// Just get the path of the file back
func GetFilePathFromURL(u string) (string, error) {
	uri, err := ParseFileURL(u)
	return uri.Path, err
}
