package config

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data struct {
	Field1 string
	Field2 string
	Field3 string
}

func TestGenerate(t *testing.T) {

	tempDir := t.TempDir()
	outputDir := path.Join(tempDir, "test")
	inputDir := path.Join("test/input")
	err := Generate(inputDir, outputDir, Data{
		Field1: "Field1",
		Field2: "Field2",
		Field3: "Field3",
	})
	assert.Nil(t, err)

	actual1, err := os.ReadFile(path.Join(outputDir, "template1.txt"))
	assert.Nil(t, err)
	expected1, err := os.ReadFile("test/output/template1.txt")
	assert.Nil(t, err)
	assert.Equal(t, string(actual1), string(expected1))

	actual2, err := os.ReadFile(path.Join(outputDir, "template2.txt"))
	assert.Nil(t, err)
	expected2, err := os.ReadFile("test/output/template2.txt")
	assert.Nil(t, err)
	assert.Equal(t, string(actual2), string(expected2))

	actual3, err := os.ReadFile(path.Join(outputDir, "dir", "template3.txt"))
	assert.Nil(t, err)
	expected3, err := os.ReadFile("test/output/dir/template3.txt")
	assert.Nil(t, err)
	assert.Equal(t, string(actual3), string(expected3))
}

func TestGenerate_Delims(t *testing.T) {

	tempDir := t.TempDir()
	outputDir := path.Join(tempDir, "test_delims")
	inputDir := path.Join("test_delims/input")
	err := GenerateWithDelims(inputDir, outputDir, Data{
		Field1: "Field1",
		Field2: "Field2",
		Field3: "Field3",
	},
		"{{{",
		"}}}",
	)
	assert.Nil(t, err)

	actual1, err := os.ReadFile(path.Join(outputDir, "template1.txt"))
	assert.Nil(t, err)
	expected1, err := os.ReadFile("test/output/template1.txt")
	assert.Nil(t, err)
	assert.Equal(t, string(actual1), string(expected1))

	actual2, err := os.ReadFile(path.Join(outputDir, "template2.txt"))
	assert.Nil(t, err)
	expected2, err := os.ReadFile("test/output/template2.txt")
	assert.Nil(t, err)
	assert.Equal(t, string(actual2), string(expected2))

	actual3, err := os.ReadFile(path.Join(outputDir, "dir", "template3.txt"))
	assert.Nil(t, err)
	expected3, err := os.ReadFile("test/output/dir/template3.txt")
	assert.Nil(t, err)
	assert.Equal(t, string(actual3), string(expected3))
}
