package linux

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestCreateMissingDirs(t *testing.T) {
	tempDir := t.TempDir()

	testDir1 := path.Join(tempDir, "test1")
	testDir2 := path.Join(tempDir, "test2")

	err := CreateMissingDirs(testDir1, testDir2)
	assert.NoError(t, err)

	fileInfo1, err := os.Stat(testDir1)
	assert.NoError(t, err)
	assert.True(t, fileInfo1.IsDir())

	fileInfo2, err := os.Stat(testDir2)
	assert.NoError(t, err)
	assert.True(t, fileInfo2.IsDir())

}
