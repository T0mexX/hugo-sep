// Copyright 2024 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hstrings

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/stretchr/testify/assert"
)

func TestStringEqualFold(t *testing.T) {
	c := qt.New(t)

	s1 := "A"
	s2 := "a"

	c.Assert(StringEqualFold(s1).EqualFold(s2), qt.Equals, true)
	c.Assert(StringEqualFold(s1).EqualFold(s1), qt.Equals, true)
	c.Assert(StringEqualFold(s2).EqualFold(s1), qt.Equals, true)
	c.Assert(StringEqualFold(s2).EqualFold(s2), qt.Equals, true)
	c.Assert(StringEqualFold(s1).EqualFold("b"), qt.Equals, false)
	c.Assert(StringEqualFold(s1).Eq(s2), qt.Equals, true)
	c.Assert(StringEqualFold(s1).Eq("b"), qt.Equals, false)
}

func TestGetOrCompileRegexp(t *testing.T) {
	c := qt.New(t)

	re, err := GetOrCompileRegexp(`\d+`)
	c.Assert(err, qt.IsNil)
	c.Assert(re.MatchString("123"), qt.Equals, true)
	// re1, err1 := getOrCompileRegexp(`.v;dfvlfvb;sa`)
	// c.Assert(err1, qt.isNotNil)
}

func BenchmarkGetOrCompileRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetOrCompileRegexp(`\d+`)
	}
}

func BenchmarkCompileRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		regexp.MustCompile(`\d+`)
	}
}

/*
	Code added for Assignment 1.
*/

// Main function for tests. It allows to execute
// statements before and/or after all tests are executed.
// m.Run() runs all tests in the file.
func TestMain(m *testing.M) {
	exitCode := m.Run()
	f, _ := os.Create("branch_coverage.txt")
	defer f.Close()

	w := bufio.NewWriter(f)
	fmt.Fprintf(w, "%v", ba.getAnalysis())
	w.Flush()

	os.Exit(exitCode)
}

type StringerImplementation struct{ str string }

func (si StringerImplementation) String() string { return si.str }

func TestForAssignments(t *testing.T) {

	t.Run("test for function 'ToString'", func(t *testing.T) {
		testCases := [3]struct {
			input        any
			expectedStr  string
			expectedBool bool
		}{
			{input: "a string", expectedStr: "a string", expectedBool: true},
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

	t.Run("test for function 'Eq'", func(t *testing.T) {
		funReceiver := StringEqualFold("a string")
		testCases := [5]struct {
			input    any
			expected bool
		}{
			{input: "a string", expected: true},
			{input: "a string but the wrong one", expected: false},
			{input: "a string", expected: true},
			{input: StringerImplementation{"a string but the wrong one"}, expected: false},
			{input: 4, expected: false},
		}

		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("TestCase: %v", testCase), func(t *testing.T) {
				boolOut := funReceiver.Eq(testCase.input)
				assert.Equal(t, testCase.expected, boolOut)
			})
		}
	})
	
	t.Run("test for function 'InSlice'", func(t *testing.T) {
		testCases := [6]struct {
			array_str  []string
			target_str string
			expected   bool
		}{
			{array_str: []string{"a", "string", "jennifer"}, target_str: "jennifer", expected: true},
			{array_str: []string{"eh", "io volevo", "te"}, target_str: "te", expected: true},
			{array_str: []string{"a", "string", "jennifer"}, target_str: "big", expected: false},
			{array_str: []string{}, target_str: "big", expected: false},
			{array_str: []string{}, target_str: "", expected: false},
			{array_str: []string{"     "}, target_str: "", expected: false},
		}

		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("TestCase: %v", testCase), func(t *testing.T) {
				boolOut := InSlice(testCase.array_str, testCase.target_str)
				assert.Equal(t, testCase.expected, boolOut)
			})
		}
	})

	t.Run("test for function 'EqualAny'", func (t *testing.T) {
		testCases:= [5]struct {
			input1 string
			input2 string
			input3 string
			input4 string
			expected bool
		} {
			{input1: "random1", input2: "random1", expected: true},
			{input1: "random1", input2: "raandom", input3: "random1", input4: "random2", expected: true},
			{input1: "raandom1", input2: "raandom", input3: "random1", input4: "random2", expected: false},
			{input1: "random1", input2: "1", input3: "2", input4:"3", expected: false},
			{input1: "random1", input2: "random1", input3: "2", input4:"3", expected: true},
		}

		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("TestCase: %v", testCase), func (t *testing.T) {
				boolOut := EqualAny(testCase.input1, testCase.input2, testCase.input3, testCase.input4)
				assert.Equal(t, testCase.expected, boolOut)
			})
		}
	})
	
	t.Run("test for function 'InSliceEqualFold'", func(t *testing.T) {
		testCases := [6]struct {
			array_str  []string
			target_str string
			expected   bool
		}{
			{array_str: []string{"a", "string", "jennifer"}, target_str: "jennifer", expected: true},
			{array_str: []string{"eh", "io volevo", "te"}, target_str: "te", expected: true},
			{array_str: []string{"a", "string", "jennifer"}, target_str: "big", expected: false},
			{array_str: []string{}, target_str: "big", expected: false},
			{array_str: []string{}, target_str: "", expected: false},
			{array_str: []string{"     "}, target_str: "", expected: false},
		}

		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("TestCase: %v", testCase), func(t *testing.T) {
				boolOut := InSlicEqualFold(testCase.array_str, testCase.target_str)
				assert.Equal(t, testCase.expected, boolOut)
			})
		}
	})
}
