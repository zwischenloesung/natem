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

type ThingTarget struct {
	Url       string `json:"url"`
	Checksum  string `json:"checksum"`
	Tag       string `json"tag"`
	Timestamp string `json:"timestamp"`
	Geotag    string `json:"geotag"`
}

type ThingRelation struct {
	BeA     []string         `json:"be_a"`
	Have    []string         `json:"have"`
	Know    []string         `json:"know"`
	Show    []string         `json:"show"`
	Include []NameUrlVersion `json:"include"`
}

type ThingId struct {
	Uuid    string   `json:"uuid"`
	Name    string   `json:"name"`
	Version string   `json:"version"`
	Url     []string `json:"url"`
}

type ThingAttribution struct {
	Author    []NameUrl `json:"author"`
	Reference []NameUrl `json:"reference"`
}

type ThingPermission struct {
	// it definitely makes sense to define only one user, but who are
	// we to decide / enforce it here?
	Owner    []string `json:"owner"`
	Editor   []string `json:"editor"`
	Consumer []string `json:"consumer"`
}

type Thing struct {
	// usually we have either "null" or "one" target, but there might
	// exist exceptions..
	Target      []ThingTarget    `json:"target"`
	Relation    ThingRelation    `json:"relation"`
	Parameter   interface{}      `json:"parameter"`
	Id          ThingId          `json:"id"`
	Schema      []NameUrlVersion `json:"schema"`
	Attribution ThingAttribution `json:"attribution"`
	Permission  ThingPermission  `json:"permission"`
}

func (t *Thing) GenId() {
	t.Id.Uuid = "urn:uuid:" + uuid.New().String()
}

func NewThing() *Thing {

	t := new(Thing)
	t.GenId()
	return t
}

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

	var thing Thing
	err := yaml.Unmarshal(yamlContent, &thing)
	if thing.Id.Uuid == "" {
		thing.GenId()
	}
	return thing, err
}

func ParseThingFromFile(fileName string) (Thing, error) {

	yamlContent, err := ReadYAMLDocumentFromFile(fileName)
	if err != nil {
		return *NewThing(), err
	}
	return ParseThing(yamlContent)
}

func SerializeThing(theThing *Thing) ([]byte, error) {

	fmt.Println("theThing address: ", &theThing)
	// Make sure every Thing always has its UUID set
	if theThing.Id.Uuid == "" {
		theThing.GenId()
	}

	resultBytes, err := yaml.Marshal(theThing)

	return resultBytes, err
}

func SerializeThingToFile(theThing *Thing, fileName string) error {

	thingBytes, err := SerializeThing(theThing)
	if err != nil {
		return err
	}
	tbs := [][]byte{[]byte("---"), thingBytes}
	ntbs := bytes.Join(tbs, []byte("\n"))
	return os.WriteFile(fileName, ntbs, 0644)
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
