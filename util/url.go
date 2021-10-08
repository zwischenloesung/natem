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

func ParseFileURL(u string, project string) (*Link, error) {
	uri, err := ParseURI(u)
	// Only consider URLs pointing to the local file system, allow for
	// absolute or relative paths too.
	if uri.Scheme == "file" || uri.Scheme == "" {
		uri.Scheme = "file"
		if uri.Path[0] != byte('/') {
			uri.Path = project + "/" + uri.Path
		}
		return &Link{*uri}, nil
	} else {
		log.Fatal("This URL was not of scheme 'file:///' as expected.\n", err)
		return &Link{*uri}, err
	}
}

// Just get the path of the file back
func GetFilePathFromURL(u string, project string) (string, error) {
	uri, err := ParseFileURL(u, project)
	return uri.Path, err
}
