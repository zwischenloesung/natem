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
	"strings"
)

type ThingURL struct {
	*url.URL
	ContextPath string
	RW          bool
}

var SupportedThingURLSchemesRO = []string{"file", "http", "https"}
var SupportedThingURLSchemesRW = []string{"file"}

func isSupportedThingURLScheme(schemeString string) (bool, bool) {
	isRO := false
	isRW := false
	for _, i := range SupportedThingURLSchemesRO {
		if i == schemeString {
			isRO = true
		}
	}
	for _, i := range SupportedThingURLSchemesRW {
		if i == schemeString {
			isRW = true
		}
	}
	return isRO, isRW
}

func ParseThingURL(thingURI string, contextURI string) (*ThingURL, error) {
	tu, err := url.Parse(thingURI)
	if err != nil {
		return nil, err
	}
	cu, err := url.Parse(contextURI)
	if err != nil {
		return nil, err
	}

	if tu.Path == "" {
		return nil, errors.New("This thing might not have an empty path.\n")
	}
	if cu.Path[0] != byte('/') {
		return nil, errors.New("The context path of this thing must be absolute.\n")
	}

	if tu.Path[0] != byte('/') {
		tu.Path = cu.Path + "/" + tu.Path
	} else if !strings.HasPrefix(tu.Path, cu.Path) {
		return &ThingURL{tu, cu.Path, false}, errors.New("Thing and context URLs do not match.\n")
	}

	if tu.Scheme == "" {
		if cu.Scheme == "" {
			tu.Scheme = "file"
		} else {
			tu.Scheme = cu.Scheme
		}
	} else {
		if cu.Scheme != tu.Scheme {
			return &ThingURL{tu, cu.Path, false}, errors.New("Thing and context URLs do not match.\n")
		}
	}

	r, w := isSupportedThingURLScheme(tu.Scheme)
	if r {
		return &ThingURL{tu, cu.Path, w}, nil
	} else {
		return &ThingURL{tu, cu.Path, w}, errors.New("This URL does not have a compatible scheme.\n")
	}
}

// Just get the path of the file back
func GetPathFromThingURL(u string, context string) (string, error) {
	uri, err := ParseThingURL(u, context)
	if uri.Scheme != "file" {
		err = errors.New("This path is not local, scheme must be 'file'.\n")
	}
	return uri.Path, err
}
