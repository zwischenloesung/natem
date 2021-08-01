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
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Thing struct {
	Targets      []string `yaml:"targets"` // URI:Link
	Behavior     struct{} `yaml:"behaviour"`
	Content      struct{} `yaml:"content"`
	Id           string   `yaml:"_id"` // URI:UUID
	Alias        []string `yaml:"_alias"`
	Name         string   `yaml:"_name"`         // URI:... (locally unique as path is added?)
	Urls         []string `yaml:"_urls"`         // URI:Link
	Version      string   `yaml:"_version"`      // URI tag would be nice!
	Dependencies []string `yaml:"_dependencies"` // URI:Link
	Type         struct {
		Name    string `yaml:"name"`    // URI:tag?
		Version string `yaml:"version"` // URI:tag would be nice!
		//		Schema  URI `yaml:"schema"` (auto-calculated)
	} `yaml:"_type"`
	Authors []struct {
		Name string `yaml:"name"`
		Uri  string `yaml:"uri"` // URI:Address
	} `yaml:"_authors"`
	References []struct {
		Name string `yaml:"name"`
		Uri  string `yaml:"uri"` // URI:Address
	} `yaml:"_references"`
}

func ParseThing(yamlContent []byte) Thing {

	var thing Thing
	err := yaml.Unmarshal(yamlContent, &thing)
	if err != nil {
		log.Fatal("Error parsing JSON/YAML from file.\n", err)
	}
	return thing
}

func ParseThingFromFile(file string) Thing {

	yamlContent, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Error reading JSON/YAML file.\n", err)
	}

	return ParseThing(yamlContent)
}
