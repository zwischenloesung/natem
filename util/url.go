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

var UrlThingOutsideContextError = errors.New("Target and context URLs do not match.\n")

var SupportedThingURLSchemesRO = []string{"file", "http", "https"}
var SupportedThingURLSchemesRW = string("file")

func isSupportedThingURLScheme(schemeString string) (bool, bool) {
	isRO := false
	isRW := false
	for _, i := range SupportedThingURLSchemesRO {
		if i == schemeString {
			isRO = true
		}
	}
	if schemeString == SupportedThingURLSchemesRW {
		isRW = true
	}
	return isRO, isRW
}

/*
	ParseThingURL
      args
		thingURL		the string to be parsed
		contextURL		a string with the context information
		hasContext		bool, whether to consider the context mandatory
	  returns
		ThingURL        pointer to a ThingURL containing the locator info
		error			status of the ThingURL
*/
func ParseThingURL(thingURI string, contextURI string, hasContext bool) (*ThingURL, error) {
	tu, err := url.Parse(thingURI)
	if err != nil {
		return nil, err
	}
	cu, err := url.Parse(contextURI)
	if err != nil {
		return nil, err
	}

	if tu.Path == "" {
		return nil, errors.New("This thing must not have an empty path.\n")
	}
	if cu.Path[0] != byte('/') {
		return nil, errors.New("The context path of this thing must be absolute.\n")
	}

	if hasContext && strings.Contains(tu.Path, "..") {
		return &ThingURL{tu, cu.Path, false}, UrlThingOutsideContextError
	}
	if tu.Path[0] != byte('/') {
		tu.Path = cu.Path + "/" + tu.Path
	} else if hasContext && !strings.HasPrefix(tu.Path, cu.Path) {
		return &ThingURL{tu, cu.Path, false}, UrlThingOutsideContextError
	}

	if tu.Scheme == "" {
		if cu.Scheme == "" {
			tu.Scheme = SupportedThingURLSchemesRW
		} else {
			tu.Scheme = cu.Scheme
		}
	} else {
		if cu.Scheme != tu.Scheme {
			return &ThingURL{tu, cu.Path, false}, UrlThingOutsideContextError
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
func GetThingURLPath(u string, context string, hasContext bool) (string, error) {
	uri, err := ParseThingURL(u, context, hasContext)
	if err == nil && !uri.RW {
		err = errors.New("This path is not local, scheme must be 'file'.\n")
	}
	return uri.Path, err
}
