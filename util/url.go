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
	ContextPath string
}

var SupportedThingURLSchemes = []string{"file", "http", "https"}

func ParseThingURL(thingURI string, contextURI string) (*ThingURL, error) {
	tu, err := url.Parse(thingURI)
	if err != nil {
		return nil, err
	}
	cu, err := url.Parse(contextURI)
	if err != nil {
		return nil, err
	}

	// Only consider URLs pointing to the local file system, allow for
	// absolute or relative paths too.
	if tu.Scheme == "file" || tu.Scheme == "" {
		tu.Scheme = "file"
		if tu.Path[0] != byte('/') {
			tu.Path = cu.Path + "/" + tu.Path
		}
		return &ThingURL{tu, cu.Path}, nil
	} else {
		return &ThingURL{tu, cu.Path}, errors.New("This URL was not of scheme 'file:///' as expected.\n")
	}
}

// Just get the path of the file back
func GetPathFromThingURL(u string, context string) (string, error) {
	uri, err := ParseThingURL(u, context)
	return uri.Path, err
}
