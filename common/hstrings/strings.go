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
	"fmt"
	"regexp"
	"strings"
	"sync"
	
	"github.com/gohugoio/hugo/compare"
)



var _ compare.Eqer = StringEqualFold("")

// StringEqualFold is a string that implements the compare.Eqer interface and considers
// two strings equal if they are equal when folded to lower case.
// The compare.Eqer interface is used in Hugo to compare values in templates (e.g. using the eq template function).
type StringEqualFold string

func (s StringEqualFold) EqualFold(s2 string) bool {
	return strings.EqualFold(string(s), s2)
}

func (s StringEqualFold) String() string {
	return string(s)
}

func (s StringEqualFold) Eq(s2 any) bool {
	switch ss := s2.(type) {
	case string: // branch id = 3
		ba.reachedBranch(3)
		return s.EqualFold(ss)
	case fmt.Stringer: // branch id = 4
		ba.reachedBranch(4)
		return s.EqualFold(ss.String())
	}
	// branch id = 5
	ba.reachedBranch(5)
	return false 
}

// EqualAny returns whether a string is equal to any of the given strings.
func EqualAny(a string, b ...string) bool {
	for _, s := range b {
		if a == s { // branch id = 6 (if condition evaluates to true at least once)
			ba.reachedBranch(6)
			return true
		}
		// (else)
		// branch id = 7 (if condition evaluates to false at least once)
		ba.reachedBranch(7)
	}
	// branch id = 8 (if condition always evaluates to false)
	ba.reachedBranch(8)
	return false 
}

// regexpCache represents a cache of regexp objects protected by a mutex.
type regexpCache struct {
	mu sync.RWMutex
	re map[string]*regexp.Regexp
}

func (rc *regexpCache) getOrCompileRegexp(pattern string) (re *regexp.Regexp, err error) {
	var ok bool

	if re, ok = rc.get(pattern); !ok { // branch id = 9
		ba.reachedBranch(9)
		re, err = regexp.Compile(pattern)
		if err != nil { // branch id = 10
			ba.reachedBranch(10)
			return nil, err
		}
		// (else)
		// branch id = 11
		ba.reachedBranch(11)
		rc.set(pattern, re)
	}
	// (else)
	// branch id = 12

	return re, nil
}

func (rc *regexpCache) get(key string) (re *regexp.Regexp, ok bool) {
	rc.mu.RLock()
	re, ok = rc.re[key]
	rc.mu.RUnlock()
	return
}

func (rc *regexpCache) set(key string, re *regexp.Regexp) {
	rc.mu.Lock()
	rc.re[key] = re
	rc.mu.Unlock()
}

var reCache = regexpCache{re: make(map[string]*regexp.Regexp)}

// GetOrCompileRegexp retrieves a regexp object from the cache based upon the pattern.
// If the pattern is not found in the cache, the pattern is compiled and added to
// the cache.
func GetOrCompileRegexp(pattern string) (re *regexp.Regexp, err error) {
	return reCache.getOrCompileRegexp(pattern)
}

// InSlice checks if a string is an element of a slice of strings
// and returns a boolean value.
func InSlice(arr []string, el string) bool {
	for _, v := range arr {
		if v == el { // branch id = 13 (if condition evaluates to true at least once)
			ba.reachedBranch(13)
			return true
		}
		// (else)
		// branch id = 14 (if condition evaluates to false at least once)
		ba.reachedBranch(14)
	}
	// (else)
	// branch id = 15 (if condition always evaluates to false)
	ba.reachedBranch(15)
	return false
}

// InSlicEqualFold checks if a string is an element of a slice of strings
// and returns a boolean value.
// It uses strings.EqualFold to compare.
func InSlicEqualFold(arr []string, el string) bool {
	for _, v := range arr { 
		if strings.EqualFold(v, el) { // branch id = 16 (if condition evaluates to true at least once)
			ba.reachedBranch(16)
			return true
		}
		// (else)
		// branch id = 17 (if condition evaluates to false at least once)
		ba.reachedBranch(17)
	}
	// (else)
	// branch id = 18 (if condition always evaluates to false)
	ba.reachedBranch(18)
	return false
}

// ToString converts the given value to a string.
// Note that this is a more strict version compared to cast.ToString,
// as it will not try to convert numeric values to strings,
// but only accept strings or fmt.Stringer.
func ToString(v any) (string, bool) {
	switch vv := v.(type) {
	case string: // branch id = 0
		ba.reachedBranch(0)
		return vv, true
	case fmt.Stringer: // branch id = 1
	ba.reachedBranch(1)
		return vv.String(), true
	}
	// branch id = 2
	ba.reachedBranch(2)
	return "", false 
}

type Tuple struct {
	First  string
	Second string
}


/*
	Code added for Assignment 1:
*/

// Contains the data relevant to branch coverage analysis.
type BranchAnalyzer struct {
	// Name of the analyzed file.
	filename string
	// Boolean array where index number correspond to a branch (idx = branchId).
	// A value is set to true if the corresponding branch is reached.
	branches  [19]bool
	// Functions subject to analysis. Each function instance contains 
	// starting and ending branch ids (the branch ids that are reachable in the function body).
	functions [6]Function
}

type Function struct {
	name string
	startBranchId int8
	untilId int8
}

var ba = BranchAnalyzer{
	filename: "strings.go",
	branches: [19]bool{},
	functions: [6]Function{
		{name: "ToString", startBranchId: 0, untilId: 3},
		{name: "Eq", startBranchId: 3, untilId: 6},
		{name: "EqualAny", startBranchId: 6, untilId: 9},
		{name: "getOrCompileRegexp", startBranchId: 9, untilId: 12},
		{name: "InSlice", startBranchId: 13, untilId: 16},
		{name: "InSlicEqualFold", startBranchId: 16, untilId: 19},
	},
}

// Called when the branch corresponding to parameter id.
func (ba *BranchAnalyzer) reachedBranch(id int) {
	ba.branches[id] = true
}


// Returns formatted branch analysis, both per-function and file wise.
func (ba *BranchAnalyzer) getAnalysis() string {
	var sb strings.Builder
	totalCovered := 0
	totalBranches := len(ba.branches)
	sb.WriteString("Branch coverage for file '" + ba.filename + "':\n")
	for _, f := range ba.functions {
		numCoveredPerFunc := 0
		totalBranchesPerFunc := f.untilId - f.startBranchId
		for _, b := range ba.branches[f.startBranchId:f.untilId] { if b { numCoveredPerFunc++ } }
		totalCovered += numCoveredPerFunc
		sb.WriteString(fmt.Sprintf(
			"  function '%v'\n    - %v covered branches\n    - %v total branches\n    - %0.2f%% branch coverage\n",
			f.name,
			numCoveredPerFunc,
			totalBranchesPerFunc,
			100 * float32(numCoveredPerFunc) / float32(totalBranchesPerFunc),
		))
	}
	sb.WriteString(fmt.Sprintf(
		"  Total for functions under analysis\n    - %v covered branches\n    - %v total branches\n    - %0.2f%% branch coverage\n",
		totalCovered,
		totalBranches,
		100 * float32(totalCovered) / float32(totalBranches),
	))

	return sb.String()
}