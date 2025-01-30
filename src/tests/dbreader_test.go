package tests

import (
	"strings"
	"io/ioutil"
	"testing"

	"src/dbreader"
)

func TestJSONReader(t *testing.T) {
	inputFile := "data/test_recipes.json"
	expectedFile := "data/test_recipes.xml"

	var reader dbreader.DBReader
	var jsonReader = dbreader.JSONReader{}
	reader = &jsonReader
	var err error
	err = reader.Read(&inputFile)
	if err != nil {
		t.Fatalf("JSONReader.Read() failed: %v", err)
	}

	var expectedBytes []byte
	expectedBytes, _ = ioutil.ReadFile(expectedFile)
	
	var expected string
	var actual string
	expected = strings.TrimSpace(string(expectedBytes))
	actual = strings.TrimSpace(reader.ToString())

	if actual != expected {
		t.Errorf("JSONReader.ToString() mismatch.\nExpected:\n%s\nGot:\n%s", expected, actual)
	}
}

func TestXMLReader(t *testing.T) {
	inputFile := "data/test_recipes.xml"
	expectedFile := "data/test_recipes.json"

	var reader dbreader.DBReader
	var xmlReader = dbreader.XMLReader{}
	reader = &xmlReader
	var err error
	err = reader.Read(&inputFile)
	if err != nil {
		t.Fatalf("XMLReader.Read() failed: %v", err)
	}

	var expectedBytes []byte
	expectedBytes, _ = ioutil.ReadFile(expectedFile)

	var expected string
	expected = strings.TrimSpace(string(expectedBytes))
	var actual string
	actual = strings.TrimSpace(reader.ToString())

	if actual != expected {
		t.Errorf("XMLReader.ToString() mismatch.\nExpected:\n%s\nGot:\n%s", expected, actual)
	}
}