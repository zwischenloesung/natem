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
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/google/uuid"
	"github.com/xeipuuv/gojsonschema"
)

type Thing struct {
	Id      string   `json:"id"`   // URI:UUID
	Name    string   `json:"name"` // URI:... (locally unique as path is added?)
	Urls    []string `json:"urls"` // URI:Link
	Targets struct {
		Url     string `json:"url"` // URI:Link
		Version string `json:"version"`
	} `json:"targets"`
	Behavior struct {
		Is       []Thing `json:"is"`
		IsLike   []Thing `json:"is_like"`
		IsPartOf []Thing `json:"is_part_of"`
		Seems    []Thing `json:"seems"`
		IsNot    []Thing `json:"is_not"`
		Has      []Thing `json:"has"`
		//		hosts
		//		inhabits
		// 		is_located_at
		BelongsTo []Thing `json:"belongs_to"`
		//		execute
	} `json:"behavior"`
	Property struct {
		Name        []struct{} `json:"name"`  // map of `<lang>: <value>`
		Alias       []struct{} `json:"alias"` // map of `<lang>: <value>`
		Description []struct{} `json:"name"`  // map of `<long/short>:: <lang>: <value>`
		Tags        []string   `json:"tags"`
	} `json:"property"`
	Schema struct {
		Name    string   `json:"name"`    // URI:tag?
		Version string   `json:"version"` // URI:tag would be nice!
		Urls    []string `json:"urls"`    //(auto-calculated?)
		//Compatibility string `json:"compatibility"`
	} `json:"_schema"`
	Version      string   `json:"_version"`      // URI tag would be nice!
	Dependencies []string `json:"_dependencies"` // URI:Link
	Authors      []struct {
		Name string `json:"name"`
		Uri  string `json:"uri"` // URI:Address
	} `json:"_authors"`
	References []struct {
		Name string `json:"name"`
		Uri  string `json:"uri"` // URI:Address
	} `json:"_references"`
}

func NewThing() *Thing {

	t := new(Thing)
	t.Id = uuid.New().String()
	return t
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
func ValidateThing(schemaBytes []byte, contentBytes []byte) bool {

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

	yamlContent, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading JSON/YAML file.\n", err)
	}

	return ParseThing(yamlContent)
}

func SerializeThing(theThing *Thing) ([]byte, error) {

	// Make sure every Thing always has its UUID set
	if theThing.Id == "" {
		theThing.Id = uuid.New().String()
	}

	resultBytes, err := yaml.Marshal(theThing)
	/// Unnused ATM
	//	if err != nil {
	//		return resultBytes, err
	//	}

	return resultBytes, err
}

func SerializeThingToFile(theThing *Thing, fileName string) error {

	thingBytes, err := SerializeThing(theThing)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, thingBytes, 0644)
}

func WriteThingFile(theThing *Thing, url string, context string, hasContext bool, overwrite bool) (string, string, error) {

	path, err := GetThingURLPath(url, context, hasContext)
	if err != nil {
		return path, "", err
	}

	if !overwrite {
		_, err = os.Stat(path)
		if !os.IsNotExist(err) {
			return path, "", fmt.Errorf("Not overwriting: %s.\n%s", path, err)
		}
	}
	dir, file := filepath.Split(path)
	dh, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return dir, file, err
		}
	} else if !dh.IsDir() {
		return dir, file, fmt.Errorf("Existing but not a dir: %s.\n%s", path, err)
	}
	return dir, file, SerializeThingToFile(theThing, path)
}

func CreateNewThingFile(url string, context string, hasContext bool) (*Thing, error) {

	t := NewThing()
	_, _, e := WriteThingFile(t, url, context, hasContext, false)
	return t, e
}
