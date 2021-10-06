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

	"github.com/ghodss/yaml"
	"github.com/xeipuuv/gojsonschema"
)

type Thing struct {
	Targets      []string `json:"targets"` // URI:Link
	Behavior     struct{} `json:"behaviour"`
	Content      struct{} `json:"content"`
	Id           string   `json:"_id"` // URI:UUID
	Alias        []string `json:"_alias"`
	Name         string   `json:"_name"`         // URI:... (locally unique as path is added?)
	Urls         []string `json:"_urls"`         // URI:Link
	Version      string   `json:"_version"`      // URI tag would be nice!
	Dependencies []string `json:"_dependencies"` // URI:Link
	Type         struct {
		Name    string `json:"name"`    // URI:tag?
		Version string `json:"version"` // URI:tag would be nice!
		//		Schema  URI `json:"schema"` (auto-calculated)
	} `json:"_type"`
	Authors []struct {
		Name string `json:"name"`
		Uri  string `json:"uri"` // URI:Address
	} `json:"_authors"`
	References []struct {
		Name string `json:"name"`
		Uri  string `json:"uri"` // URI:Address
	} `json:"_references"`
}

func ValidateJSONThing(schemaBytes []byte, contentBytes []byte) bool {
	schemaLoader := gojsonschema.NewStringLoader(string(schemaBytes))
	contentLoader := gojsonschema.NewStringLoader(string(contentBytes))

	result, err := gojsonschema.Validate(schemaLoader, contentLoader)
	if err != nil {
		log.Fatal("Error validating the document:\n", err)
		return false
	}

	if result.Valid() {
		return true
	} else {
		log.Print("Invalid document:\n")
		for _, e := range result.Errors() {
			//			fmt.Printf("- %s\n", err)
			log.Printf("- %s\n", e)
		}
		return false
	}
}

// We do only accept JSON compatible YAML anyway. TSENTSAK-YAML is defined to
// be an object/map and has only strings as keys.
func ValidateYAMLThing(schemaBytes []byte, contentBytes []byte) bool {

	JSONSchemaBytes, err := yaml.YAMLToJSON(schemaBytes)
	if err != nil {
		log.Fatal("Parsing the YAML schema failed:", err)
	}
	JSONContentBytes, err := yaml.YAMLToJSON(contentBytes)
	if err != nil {
		log.Fatal("Parsing the YAML document failed:", err)
	}

	return ValidateJSONThing(JSONSchemaBytes, JSONContentBytes)
}

func ParseThing(yamlContent []byte) Thing {

	var thing Thing
	err := yaml.Unmarshal(yamlContent, &thing)
	if err != nil {
		log.Fatal("Error parsing JSON/YAML from file.\n", err)
	}
	return thing
}

func ParseThingFromFile(fileName string) Thing {

	yamlContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading JSON/YAML file.\n", err)
	}

	return ParseThing(yamlContent)
}
