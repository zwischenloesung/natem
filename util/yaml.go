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
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/google/uuid"
	"github.com/xeipuuv/gojsonschema"
)

type NameUrl struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type NameUrlVersion struct {
	*NameUrl
	Version string `json:"version"`
}

type DateGeo struct {
	Date   string `json:"date"`
	Geotag string `json:"geotag"`
}

type NameUrlVersionDateGeo struct {
	*NameUrlVersion
	*DateGeo
	Target bool `json:"target"`
}

type ThingTarget struct {
	Url      string `json:"url"`
	Checksum string `json:"checksum"`
	Tag      string `json"tag"`
	Uuid     string `json"uuid"`
	*DateGeo
}

type ThingRelation struct {
	ThingUrl string `json:"thing_url"`
	Version  string `json:"version"`
	Priority string `json:"priority"`
	Kind     string `json:"kind"`
}

type ThingAction struct {
	Environment NameUrlVersion `json:"environment"`
	Run         struct {
		Command string   `json:"command"`
		Option  []string `json:"option"`
	} `json:"run"`
	Dependency []NameUrlVersion `json:"dependency"`
	Condition  []string         `json:"condition"`
}

type ThingId struct {
	Uuid    string   `json:"uuid"`
	Name    string   `json:"name"`
	Version string   `json:"version"`
	Url     []string `json:"url"`
}

type ThingLegal struct {
	Author    []NameUrlVersionDateGeo `json:"author"` //TODO or URI actually
	Reference []NameUrlVersionDateGeo `json:"reference"`
	License   []NameUrlVersionDateGeo `json:"license"`
}

type ThingPermission struct {
	// it definitely makes sense to define only one owner, but who are
	// we to decide / enforce it here?
	Owner    string   `json:"owner"`
	Editor   []string `json:"editor"`
	Consumer []string `json:"consumer"`
}

type Thing struct {
	// usually we have either "null" or "one" target, but there might
	// exist exceptions..
	Target   []ThingTarget    `json:"target"`
	Relation []ThingRelation  `json:"relation"`
	Id       ThingId          `json:"id"`
	Schema   []NameUrlVersion `json:"schema"`
	//	Behavior  ThingBehavior    `json:"behavior"`
	Behavior  map[string]interface{} `json:"behavior"`
	Parameter map[string]interface{} `json:"parameter"`
	Legal     ThingLegal             `json:"legal"`
	//	Permission  ThingPermission  `json:"permission"`
}

func (thing *Thing) GenId() {
	thing.Id.Uuid = "urn:uuid:" + uuid.New().String()
}

func NewThing() *Thing {

	t := new(Thing)
	t.GenId()
	return t
}

// first do it - load all the complete Things into the slice
//func FlattenHeritageOrder(thing Thing, theOrder []Thing) ([]Thing, error) {

// the list consists of the following
//   - walk up the directory tree and add all init.yml
//     files from every directory creating Things called
//     after the directory they live in - the name starts
//     with the context dir.
//   - every Thing has 'Thing' as most ancient parent
//}

func ValidateJSONThing(schemaBytes []byte, contentBytes []byte) (bool, error) {
	schemaLoader := gojsonschema.NewStringLoader(string(schemaBytes))
	contentLoader := gojsonschema.NewStringLoader(string(contentBytes))

	result, err := gojsonschema.Validate(schemaLoader, contentLoader)
	if err != nil {
		return false, fmt.Errorf("Error validating the document: %s\n", err)
	}

	if result.Valid() {
		return true, nil
	} else {
		//TODO proper error handling
		log.Print("Invalid document:\n")
		for _, e := range result.Errors() {
			log.Printf("- %s\n", e)
		}
		return false, nil
	}
}

// We do only accept JSON compatible YAML anyway. TSENTSAK-YAML is defined to
// be an object/map and has only strings as keys.
func ValidateThing(schemaBytes []byte, contentBytes []byte) (bool, error) {

	JSONSchemaBytes, err := yaml.YAMLToJSON(schemaBytes)
	if err != nil {
		return false, fmt.Errorf("Parsing the YAML schema failed: %s.\n", err)
	}
	JSONContentBytes, err := yaml.YAMLToJSON(contentBytes)
	if err != nil {
		return false, fmt.Errorf("Parsing the YAML thing failed: %s.\n", err)
	}

	return ValidateJSONThing(JSONSchemaBytes, JSONContentBytes)
}

// Currently we only allow nice and small files with max one document inside..
func ReadYAMLDocumentFromFile(fileName string) ([]byte, error) {

	var contentBytes [][]byte
	startDocument := false

	fh, err := os.Open(fileName)
	if err != nil {
		return []byte(""), err
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		l := []byte(scanner.Text())
		b := false
		if len(l) > 2 && bytes.Equal([]byte("---"), l[0:3]) {
			b = true
		}
		if startDocument {
			if b {
				break
			}
			contentBytes = append(contentBytes, l)
		} else if b {
			startDocument = true
		}
	}
	err = scanner.Err()
	if err == nil && len(contentBytes) < 1 {
		err = errors.New("Unable to parse sensible data from file.")
	}
	return bytes.Join(contentBytes, []byte("\n")), err
}

func ParseThing(yamlContent []byte) (Thing, error) {

	var t Thing
	err := yaml.Unmarshal(yamlContent, &t)
	if t.Id.Uuid == "" {
		t.GenId()
	}
	return t, err
}

func ParseThingFromFile(fileName string) (Thing, error) {

	yamlContent, err := ReadYAMLDocumentFromFile(fileName)
	if err != nil {
		return *NewThing(), err
	}
	return ParseThing(yamlContent)
}

// Wrap and hide the external lib
func Marshal(o interface{}) ([]byte, error) {

	resultBytes, err := yaml.Marshal(o)

	return resultBytes, err
}

func SerializeThing(thing *Thing) ([]byte, error) {

	// Make sure every Thing always has its UUID set
	if thing.Id.Uuid == "" {
		thing.GenId()
	}

	resultBytes, err := Marshal(thing)

	return resultBytes, err
}

func SerializeThingToFile(thing *Thing, fileName string) error {

	thingBytes, err := SerializeThing(thing)
	if err != nil {
		return err
	}
	tbs := [][]byte{[]byte("---"), thingBytes}
	ntbs := bytes.Join(tbs, []byte("\n"))
	return os.WriteFile(fileName, ntbs, 0644)
}

func WriteThingFile(thing *Thing, url string, context string, hasContext bool, overwrite bool) (string, string, error) {

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
	return dir, file, SerializeThingToFile(thing, path)
}

func CreateNewThingFile(url string, context string, hasContext bool) (*Thing, error) {

	t := NewThing()
	_, _, e := WriteThingFile(t, url, context, hasContext, false)
	return t, e
}
