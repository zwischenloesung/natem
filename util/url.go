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

type ThingURL struct {
	URI
	Context string
}

var SupportedThingURLSchemes = []string{"file", "http", "https"}

func ParseURI(u string) (*URI, error) {
	uri, err := url.Parse(u)
	if err != nil {
		log.Fatal("An error occured while parsing the URI.\n", err)
	}
	return &URI{uri}, err
}

func ParseThingURL(u string, project string) (*ThingURL, error) {
	uri, err := ParseURI(u)
	// Only consider URLs pointing to the local file system, allow for
	// absolute or relative paths too.
	if uri.Scheme == "file" || uri.Scheme == "" {
		uri.Scheme = "file"
		if uri.Path[0] != byte('/') {
			uri.Path = project + "/" + uri.Path
		}
		return &ThingURL{*uri, project}, nil
	} else {
		log.Fatal("This URL was not of scheme 'file:///' as expected.\n", err)
		return &ThingURL{*uri, project}, err
	}
}

// Just get the path of the file back
func GetPathFromThingURL(u string, project string) (string, error) {
	uri, err := ParseThingURL(u, project)
	return uri.Path, err
}
