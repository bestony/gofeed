package atom_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mmcdole/gofeed/atom"
	"github.com/stretchr/testify/assert"
)

// Tests

func TestParser_Parse(t *testing.T) {
	files, _ := filepath.Glob("../testdata/parser/atom/*.xml")
	for _, f := range files {
		base := filepath.Base(f)
		name := strings.TrimSuffix(base, filepath.Ext(base))

		fmt.Printf("Testing %s... ", name)

		// Get actual source feed
		ff := fmt.Sprintf("../testdata/parser/atom/%s.xml", name)
		f, _ := ioutil.ReadFile(ff)

		// Parse actual feed
		fp := &atom.Parser{}
		actual, _ := fp.Parse(bytes.NewReader(f))

		// Get json encoded expected feed result
		ef := fmt.Sprintf("../testdata/parser/atom/%s.json", name)
		e, _ := ioutil.ReadFile(ef)

		// Unmarshal expected feed
		expected := &atom.Feed{}
		json.Unmarshal(e, expected)

		if assert.Equal(t, expected, actual, "Feed file %s.xml did not match expected output %s.json", name, name) {
			fmt.Printf("OK\n")
		} else {
			fmt.Printf("Failed\n")
		}
	}
}

func TestParser_ParseNItem(t *testing.T) {
	files, _ := filepath.Glob("../testdata/parser/atom/*.xml")
	for _, f := range files {
		base := filepath.Base(f)
		name := strings.TrimSuffix(base, filepath.Ext(base))

		fmt.Printf("Testing %s... ", name)

		// Get actual source feed
		ff := fmt.Sprintf("../testdata/parser/atom/%s.xml", name)
		f, _ := ioutil.ReadFile(ff)

		// Parse actual feed
		fp := &atom.Parser{}
		actual, _ := fp.ParseNItem(bytes.NewReader(f), 5)

		// Get json encoded expected feed result
		ef := fmt.Sprintf("../testdata/parser/atom/%s.json", name)
		e, _ := ioutil.ReadFile(ef)

		// Unmarshal expected feed
		expected := &atom.Feed{}
		json.Unmarshal(e, expected)
		if len(expected.Entries) > 5 {
			expected.Entries = expected.Entries[0:5]
		}

		if assert.Equal(t, expected, actual, "Feed file %s.xml did not match expected output %s.json", name, name) {
			fmt.Printf("OK\n")
		} else {
			fmt.Printf("Failed\n")
		}
	}
}

func TestParser_ParseNItem_Part(t *testing.T) {
	files, _ := filepath.Glob("../testdata/parser/part_5items/atom/*.xml")
	for _, f := range files {
		base := filepath.Base(f)
		name := strings.TrimSuffix(base, filepath.Ext(base))

		fmt.Printf("Testing %s... ", name)

		// Get actual source feed
		ff := fmt.Sprintf("../testdata/parser/part_5items/atom/%s.xml", name)
		f, _ := ioutil.ReadFile(ff)

		// Parse actual feed
		fp := &atom.Parser{}
		actual, _ := fp.ParseNItem(bytes.NewReader(f), 5)

		// Get json encoded expected feed result
		ef := fmt.Sprintf("../testdata/parser/part_5items/atom/%s.json", name)
		e, _ := ioutil.ReadFile(ef)

		// Unmarshal expected feed
		expected := &atom.Feed{}
		json.Unmarshal(e, expected)

		if assert.Equal(t, expected, actual, "Feed file %s.xml did not match expected output %s.json", name, name) {
			fmt.Printf("OK\n")
		} else {
			fmt.Printf("Failed\n")
		}
	}
}

// TODO: Examples
