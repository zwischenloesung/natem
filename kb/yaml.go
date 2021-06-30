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
	"net/url"
	"os"
	"os/exec"
)

type Thing struct {
	targets       []string
	behavior      struct{}
	content       struct{}
	_id           string
	_alias        []string
	_urls         []string
	_version      string
	_dependencies []string
	_type         struct {
		name    string
		version string
		schema  url.URL
	}
	_authors []struct {
		name string
		uri  []url.URL
	}
	_references []struct {
		name string
		uri  []url.URL
	}
}

func ShowActions(project string, thing string, behavior string) {
	fmt.Println("kb.ShowActions(", project, ",", thing, ",", behavior, ") called")

	//TODO test for url or combine with project

	cmd := exec.Command("vim", thing)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println("an error occurred.\n", err)
	}
}

func ShowBehavior(project string, thing string) {
	fmt.Println("kb.ShowBehavior(", project, ",", thing, ") called")
}

func ShowCategories(project string, thing string) {
	fmt.Println("kb.ShowCategories(", project, ",", thing, ") called")
}

func ShowParameters(project string, thing string) {
	fmt.Println("kb.ShowParameters(", project, ",", thing, ") called")
}

func ShowRelations(project string, thing string, behavior string) {
	fmt.Println("kb.ShowRelations(", project, ",", thing, ",", behavior, ") called")
}
