package hugio

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"path/filepath"

	"github.com/stretchr/testify/assert"
	"github.com/spf13/afero"
	"github.com/gohugoio/hugo/hugofs"
	iofs "io/fs"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	f, _ := os.Create("branch_coverage_copyFile.txt")
	defer f.Close()

	w := bufio.NewWriter(f)
	fmt.Fprintf(w, "%v", ba.getAnalysis())
	w.Flush()

	os.Exit(exitCode)
}


func TestForAssignments(t *testing.T) {

	t.Run("test for function 'Copy'", func(t *testing.T) {
		fs := hugofs.Os
		sfp := filepath.Join(jekyllRoot, entry.Name())
		dfp := filepath.Join(dest, entry.Name())
		testCases := [3]struct {
			from string
			to 	 string
			expectedBool bool
		}{
			{from: "a string", expectedStr: "a string", expectedBool: true},
			{
				input:        StringerImplementation{"This is a Stringer implementation."},
				expectedStr:  "This is a Stringer implementation.",
				expectedBool: true,
			},
			{input: 2, expectedStr: "", expectedBool: false},
		}

		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("TestCase: %v", testCase), func(t *testing.T) {
				strOut, boolOut := ToString(testCase.input)
				assert.Equal(t, testCase.expectedStr, strOut)
				assert.Equal(t, testCase.expectedBool, boolOut)
			})
		}
	})

}