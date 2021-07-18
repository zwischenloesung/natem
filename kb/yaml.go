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
	"io/ioutil"
	"log"
	"net/url"

	"gopkg.in/yaml.v2"
)

type Thing struct {
	Targets      []string `yaml:"targets"`
	Behavior     struct{} `yaml:"behaviour"`
	Content      struct{} `yaml:"content"`
	Id           string   `yaml:"_id"`
	Alias        []string `yaml:"_alias"`
	Name         string   `yaml:"_name"`
	Urls         []string `yaml:"_urls"`
	Version      string   `yaml:"_version"`
	Dependencies []string `yaml:"_dependencies"`
	Type         struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
		//		Schema  URI `yaml:"schema"`
	} `yaml:"_type"`
	Authors []struct {
		Name string `yaml:"name"`
		Uri  string `yaml:"uri"`
	} `yaml:"_authors"`
	References []struct {
		Name string `yaml:"name"`
		Uri  string `yaml:"uri"`
	} `yaml:"_references"`
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

//func ParseThing(thing string) Thing {

//	err := yaml.Unmarshal(
//}

func ParseThingFromFile(file string) Thing {

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Error reading JSON/YAML file.\n", err)
	}

	var t Thing
	err = yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		log.Fatal("Error parsing JSON/YAML from file.\n", err)
	}
	return t
}

func ShowActions(project string, thing string, behavior string) {
	fmt.Println("kb.ShowActions(", project, ",", thing, ",", behavior, ") called")
}

func ShowBehavior(project string, thing string) {
	fmt.Println("kb.ShowBehavior(", project, ",", thing, ") called")
}

func ShowCategories(project string, thing string) {
	fmt.Println("kb.ShowCategories(", project, ",", thing, ") called")
}

func ShowParameters(project string, thing string) {
	fmt.Println("kb.ShowParameters(", project, ",", thing, ") called")
	theThing := ParseThingFromFile(thing)
	fmt.Println(theThing)
}

func ShowRelations(project string, thing string, behavior string) {
	fmt.Println("kb.ShowRelations(", project, ",", thing, ",", behavior, ") called")
}
