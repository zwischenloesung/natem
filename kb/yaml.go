/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
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
