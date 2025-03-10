package goexec

import (
	"fmt"
	. "github.com/janpfeifer/gonb/common"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

// The tests here uses the sample code and utility functions defined in `parser_test.go`.

func TestCreateGoFileFromLines(t *testing.T) {
	// Test cursor positioning in generated cellLines.
	s := newEmptyState(t)
	defer func() {
		err := s.Finalize()
		require.NoError(t, err, "Failed to finalized state")
	}()
	fmt.Println(s.MainPath())

	content := sampleCellCode
	cellLines := strings.Split(content, "\n")
	skipLines := MakeSet[int]()
	for ii, line := range cellLines {
		if strings.HasPrefix(line, "!") {
			skipLines.Insert(ii)
		}
	}

	cursorInCell := Cursor{38, 27} // "func (k *Kg) Gain(lasagna K_g) {"
	cursorLine := cellLines[cursorInCell.Line]
	cursorInFile, fileToCellLines, err := s.createGoFileFromLines(s.MainPath(), cellLines, skipLines, cursorInCell)
	require.NoErrorf(t, err, "Failed createGoFileFromLines(%q)", s.MainPath())

	// Read generated contents:
	contentBytes, err := os.ReadFile(s.MainPath())
	require.NoErrorf(t, err, "Failed os.ReadFile(%q)", s.MainPath())
	content = string(contentBytes)
	require.Contains(t, content, "func main() {")
	require.NotContains(t, content, "echo nonono", "Line should have been filtered out, since it is in skipLine.")

	numCellLines := len(cellLines)
	fileLines := strings.Split(content, "\n")
	numFileLines := len(fileLines)
	require.Equal(t, numCellLines+5, numFileLines, "Number of lines of generated main.go")
	require.Equal(t, cursorLine, fileLines[cursorInFile.Line], "Cursor line remains the same.")

	for ii, newLine := range fileLines {
		if ii >= numFileLines-8 {
			// Content of cellLines change (an indentation is added) so we skip these.
			break
		}
		cellLineIdx := fileToCellLines[ii]
		if cellLineIdx == NoCursorLine {
			continue
		}
		if cellLines[cellLineIdx] == "%%" {
			// The "%%" is mapped to `func main() { \n flags.Parse()\n`, we also skip these.
			continue
		}
		require.Equalf(t, cellLines[cellLineIdx], newLine, "Line mapping look wrong: file line %d --> cell line %d", ii, cellLineIdx)
	}
}
