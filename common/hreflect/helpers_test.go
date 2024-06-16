// Copyright 2019 The Hugo Authors. All rights reserved.
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

package hreflect

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
	"github.com/stretchr/testify/assert"
)

func TestIsTruthful(t *testing.T) {
	c := qt.New(t)

	c.Assert(IsTruthful(true), qt.Equals, true)
	c.Assert(IsTruthful(false), qt.Equals, false)
	c.Assert(IsTruthful(time.Now()), qt.Equals, true)
	c.Assert(IsTruthful(time.Time{}), qt.Equals, false)
}

func TestGetMethodByName(t *testing.T) {
	c := qt.New(t)
	v := reflect.ValueOf(&testStruct{})
	tp := v.Type()

	c.Assert(GetMethodIndexByName(tp, "Method1"), qt.Equals, 0)
	c.Assert(GetMethodIndexByName(tp, "Method3"), qt.Equals, 2)
	c.Assert(GetMethodIndexByName(tp, "Foo"), qt.Equals, -1)
}

func TestIsContextType(t *testing.T) {
	c := qt.New(t)
	type k string
	ctx := context.Background()
	valueCtx := context.WithValue(ctx, k("key"), 32)
	c.Assert(IsContextType(reflect.TypeOf(ctx)), qt.IsTrue)
	c.Assert(IsContextType(reflect.TypeOf(valueCtx)), qt.IsTrue)
}

func BenchmarkIsContextType(b *testing.B) {
	type k string
	b.Run("value", func(b *testing.B) {
		ctx := context.Background()
		ctxs := make([]reflect.Type, b.N)
		for i := 0; i < b.N; i++ {
			ctxs[i] = reflect.TypeOf(context.WithValue(ctx, k("key"), i))
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if !IsContextType(ctxs[i]) {
				b.Fatal("not context")
			}
		}
	})

	b.Run("background", func(b *testing.B) {
		var ctxt reflect.Type = reflect.TypeOf(context.Background())
		for i := 0; i < b.N; i++ {
			if !IsContextType(ctxt) {
				b.Fatal("not context")
			}
		}
	})
}

func BenchmarkIsTruthFul(b *testing.B) {
	v := reflect.ValueOf("Hugo")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !IsTruthfulValue(v) {
			b.Fatal("not truthful")
		}
	}
}

type testStruct struct{}

func (t *testStruct) Method1() string {
	return "Hugo"
}

func (t *testStruct) Method2() string {
	return "Hugo"
}

func (t *testStruct) Method3() string {
	return "Hugo"
}

func (t *testStruct) Method4() string {
	return "Hugo"
}

func (t *testStruct) Method5() string {
	return "Hugo"
}

func BenchmarkGetMethodByName(b *testing.B) {
	v := reflect.ValueOf(&testStruct{})
	methods := []string{"Method1", "Method2", "Method3", "Method4", "Method5"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, method := range methods {
			_ = GetMethodByName(v, method)
		}
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

func TestForAssignments(t *testing.T) {

	t.Run("test for function 'IsUint'", func(t *testing.T) {

		testCases := [7]struct {
			input    reflect.Kind
			expected bool
		}{

			{input: reflect.Uint16, expected: true},
			{input: reflect.Uint32, expected: true},
			{input: reflect.Uint8, expected: true},
			{input: reflect.Uint64, expected: true},
			{input: reflect.Int, expected: false},
			{input: reflect.Bool, expected: false},
			{input: reflect.Chan, expected: false},
		}

		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("TestCase: %v", testCase), func(t *testing.T) {
				boolOut := IsUint(testCase.input)
				assert.Equal(t, testCase.expected, boolOut)
			})
		}
	})

	// t.Run("test for function 'IsInt'", func(t *testing.T) {

	// 	testCases := [10]struct {
	// 		input    reflect.Kind
	// 		expected bool
	// 	}{

	// 		{input: reflect.Int8, expected: true},
	// 		{input: reflect.Int16, expected: true},
	// 		{input: reflect.Int32, expected: true},
	// 		{input: reflect.Int64, expected: true},
	// 		{input: reflect.Int, expected: true},
	// 		{input: reflect.Bool, expected: false},
	// 		{input: reflect.Chan, expected: false},
	// 		{input: reflect.Uint16, expected: false},
	// 		{input: reflect.Uint32, expected: false},
	// 		{input: reflect.Uint8, expected: false},
	// 	}

	// 	for _, testCase := range testCases {
	// 		t.Run(fmt.Sprintf("TestCase: %v", testCase), func(t *testing.T) {
	// 			boolOut := IsInt(testCase.input)
	// 			assert.Equal(t, testCase.expected, boolOut)
	// 		})
	// 	}
	// })

	// t.Run("test for function 'IsFloat'", func(t *testing.T) {

	// 	testCases := [8]struct {
	// 		input    reflect.Kind
	// 		expected bool
	// 	}{

	// 		{input: reflect.Float32, expected: true},
	// 		{input: reflect.Float64, expected: true},
	// 		{input: reflect.Uint8, expected: false},
	// 		{input: reflect.Uint16, expected: false},
	// 		{input: reflect.Int, expected: false},
	// 		{input: reflect.Int8, expected: false},
	// 		{input: reflect.Bool, expected: false},
	// 		{input: reflect.Chan, expected: false},
	// 	}

	// 	for _, testCase := range testCases {
	// 		t.Run(fmt.Sprintf("TestCase: %v", testCase), func(t *testing.T) {
	// 			boolOut := IsFloat(testCase.input)
	// 			assert.Equal(t, testCase.expected, boolOut)
	// 		})
	// 	}
	// })
}
