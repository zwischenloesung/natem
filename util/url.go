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
	"errors"
	"net/url"
)

type ThingURL struct {
	*url.URL
	Context string
}

var SupportedThingURLSchemes = []string{"file", "http", "https"}

func ParseThingURL(u string, context string) (*ThingURL, error) {
	uri, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	// Only consider URLs pointing to the local file system, allow for
	// absolute or relative paths too.
	if uri.Scheme == "file" || uri.Scheme == "" {
		uri.Scheme = "file"
		if uri.Path[0] != byte('/') {
			uri.Path = context + "/" + uri.Path
		}
		return &ThingURL{uri, context}, nil
	} else {
		return &ThingURL{uri, context}, errors.New("This URL was not of scheme 'file:///' as expected.\n")
	}
}

// Just get the path of the file back
func GetPathFromThingURL(u string, context string) (string, error) {
	uri, err := ParseThingURL(u, context)
	return uri.Path, err
}
