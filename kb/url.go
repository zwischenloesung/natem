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
	"fmt"
	"log"
	"net/url"
)

type URL struct {
	Whole string
	Items url.URL
}

func (u *URL) UnmarshalYAML(unmarshal func(interface{}) error) error {

	fmt.Println("hi")
	return nil
}

func ParseURL(u string) (*url.URL, error) {
	URL, err := url.Parse(u)
	if err != nil {
		log.Fatal("An error occurs while parsing the URL.\n", err)
	}
	return URL, err
}

func ParseFileURL(u string) (*url.URL, error) {
	URL, err := ParseURL(u)
	// Only consider URLs pointing to the local file system, allow for
	// absolute or relative paths too.
	if URL.Scheme == "file" || URL.Scheme == "" {
		return URL, nil
	} else {
		log.Fatal("This URL was not of scheme 'file:///' as expected.\n", err)
		return URL, err
	}
}

// Just get the path of the file back
func GetFilePathFromURL(u string) (string, error) {
	URL, err := ParseFileURL(u)
	return URL.Path, err
}
