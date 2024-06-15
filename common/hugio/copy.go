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

package hugio

import (
	"fmt"
	"io"
	iofs "io/fs"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

// Contains the data relevant to branch coverage analysis.
type BranchAnalyzer struct {
	// Name of the analyzed file.
	filename string
	// Boolean array where index number correspond to a branch (idx = branchId).
	// A value is set to true if the corresponding branch is reached.
	branches  [18]bool
	// Functions subject to analysis. Each function instance contains 
	// starting and ending branch ids (the branch ids that are reachable in the function body).
	functions [2]Function
}

type Function struct {
	name string
	startBranchId int8
	untilId int8
}

var ba = BranchAnalyzer{
	filename: "copy.go",
	branches: [18]bool{},
	functions: [2]Function{
		{name: "CopyFile", startBranchId: 0, untilId: 6},
		{name: "CopyDir", startBranchId: 6, untilId: 18},
	},
}

func (ba *BranchAnalyzer) reachedBranch(id int) {
	ba.branches[id] = true
}

// CopyFile copies a file.
func CopyFile(fs afero.Fs, from, to string) error {
	sf, err := fs.Open(from)
	if err != nil {
		ba.reachedBranch(0)
		return err
	}
	defer sf.Close()
	df, err := fs.Create(to)
	if err != nil {
		ba.reachedBranch(1)
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err != nil {
		ba.reachedBranch(2)
		return err
	}
	si, err := fs.Stat(from)
	if err != nil {
		ba.reachedBranch(3)
		err = fs.Chmod(to, si.Mode())

		if err != nil {
			ba.reachedBranch(4)
			return err
		}
	}
	ba.reachedBranch(5)
	return nil
}

// CopyDir copies a directory.
func CopyDir(fs afero.Fs, from, to string, shouldCopy func(filename string) bool) error {
	fi, err := fs.Stat(from)
	if err != nil {
		ba.reachedBranch(6)
		return err
	}

	if !fi.IsDir() {
		ba.reachedBranch(7)
		return fmt.Errorf("%q is not a directory", from)
	}

	err = fs.MkdirAll(to, 0o777) // before umask
	if err != nil {
		ba.reachedBranch(8)
		return err
	}

	d, err := fs.Open(from)
	if err != nil {
		ba.reachedBranch(9)
		return err
	}
	entries, _ := d.(iofs.ReadDirFile).ReadDir(-1)
	for _, entry := range entries {
		fromFilename := filepath.Join(from, entry.Name())
		toFilename := filepath.Join(to, entry.Name())
		if entry.IsDir() {
			ba.reachedBranch(10)
			if shouldCopy != nil && !shouldCopy(fromFilename) {
				ba.reachedBranch(11)
				continue
			}
			if err := CopyDir(fs, fromFilename, toFilename, shouldCopy); err != nil {
				ba.reachedBranch(12)
				return err
			}
			ba.reachedBranch(13)
		} else {
			ba.reachedBranch(14)
			if err := CopyFile(fs, fromFilename, toFilename); err != nil {
				ba.reachedBranch(15)
				return err
			}
			ba.reachedBranch(16)
		}

	}
	ba.reachedBranch(17)
	return nil
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
