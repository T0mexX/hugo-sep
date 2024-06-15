package maps

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)


// func TestMain(m *testing.M) {
// 	exitCode := m.Run()
// 	f, _ := os.Create("branch_coverage.txt")
// 	defer f.Close()

// 	w := bufio.NewWriter(f)
// 	fmt.Fprintf(w, "%v", ba.getAnalysis())
// 	w.Flush()

// 	os.Exit(exitCode)
// }


func returnItem (testItem any) any {
	return testItem
}

func TestForAssignments(t *testing.T) {
	t.Run("test for function 'GetOrCreate'", func(t *testing.T) {
		testCases := [2]struct {
			key any
			expectedObject any
		}{
			{key: "hello", expectedObject: "hello"},
			{key: 12, expectedObject: 12},
			// {input: 2, expectedStr: "", expectedBool: false},
		}

		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("TestCase: %v", testCase), func(t *testing.T) {
				obj := GetOrCreate(testCase.key, returnItem)
				assert.Equal(t, testCase.expectedObject, obj)
			})
		}
	})

}