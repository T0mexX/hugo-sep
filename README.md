
# Report for Assignment 1

###### Group Members
- Alessio Leonardo Tomei (lto223)
- Marco Trapasso (mtr237)
- Norah Elisabeth Milanesi (nmi245)

## Project chosen: [HUGO](https://github.com/gohugoio/hugo)


###### Lines of code 

![](readme_images/lines_of_code.png)

&nbsp;  
###### Programming language (for test purposes): [*Golang*](https://go.dev/)
![](readme_images/golang_logo.png)


&nbsp;  
&nbsp;  
_______
_______
&nbsp;  
## Coverage measurement

### Existing tool
We made use of ***Golang*** built in testing tools.
We runned the following command to get *statement coverage output* in file `.cover.out`.
```
go test ./... -coverprofile .cover.out ./...
``` 
<br>

We can then format the output on the console with the following command.
```
go tool cover -func .cover.out
```
![](readme_images/coverage_console_output.png) <br><br>

We can alternatevely use the following command to open a html page where we can visually check the statement coverage for each file.
```
go tool cover -html .cover.out
```
>***NOTE:*** red statements are not reached, while green statements are. The following is just an example section.



![](readme_images/html_coverage_example.png) <br><br>

From the html GUI we were able to identify which packages/files lacked ***statement coverage*** and, consequently, ***branch coverage***.
We chose to improve coverage of the packages `hstrings` with file `strings.go` and `hreflect` with the file `helpers.go`.

##### Statement Coverage Before Improvements
Total statement coverage [[complete file](https://github.com/T0mexX/hugo-sep/blob/master/covers/initial/statement_cover_list.txt)].

![](readme_images/total_statement_coverage.png)

&nbsp;  
Statement coverage for package `hstrings`.

![](readme_images/statement_coverage_list_before_hstrings.png)

&nbsp;  
Statement coverage for package `hreflect`.

![](readme_images/statement_coverage_list_before_hreflect.png)


&nbsp;  
&nbsp;
_______
_______ 
&nbsp;  
## Test Coverage and Hierarchy

### Our Coverage Tool
>NORAH SAYS: I don't understand the packet file sentence. To be fixed
Our own coverage tool focuses on *branch coverage*. We assigned a branch id that uniquely identifies the branch inside the packet (which usually means the file), so that each packet and its tests can be run independently.

>**PLEASE NOTE:** Our initial goal was to improve the coverage of the functions `Eq`, `EqualAny`, `ToString`, `InSlice`, `InSliceEqualFold`, and `getOrCompileRegexp` in the `strings.go` file. However, we found that the `getOrCompileRegexp` function contains an unreachable branch (`if err != nil` condition is never met because `Regexp.Compile()` always returns a `nil` error). As a result, we decided to replace `getOrCompileRegexp` with a sixth function from a different file(`helpers.go`). Despite this change, we have kept the branch setters in the original function, even though the related data is no longer used.


The `BranchAnalyzer` keeps track of: the name of the file that is being analyzed, the number of total branches and the `Function`s, which branch coverage is being calculated. `Function` is a struct containing the function name, the starting branch id and the final branch id. For example, if the *first* function has 3 branches, the `startBranchId` will be 0 while the `untilId` will be 3.
The logic is at follows.

```go
type BranchAnalyzer struct {
	// Name of the analyzed file.
	filename string
	// Boolean array where index number correspond to a branch (idx = branchId).
	// A value is set to true if the corresponding branch is reached.
	branches  [*]bool
	// Functions subject to analysis. Each function instance contains 
	// starting and ending branch ids (the branch ids that are reachable in the function body).
	functions [*]Function
}

type Function struct {
	name string
	startBranchId int8
	untilId int8
}

func (ba *BranchAnalyzer) reachedBranch(id int) {
	ba.branches[id] = true
}
```

&nbsp;  
The following function is used to format the resulting coverage.

```go

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
```

&nbsp;  
And the following `TestMain` function to output the results to a file. The name makes it the entry point for the package tests execution.

```go
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
```

&nbsp;  
___
&nbsp;  
 
### Tests Hierarchy
All tests related to this assignment and in the same package are defined hierarchically in the same test group.

```
func TestForAssignments(t *testing.T) {
	t.Run('test for function <testName1>', <testFunction1>)
	...
	...
	t.Run('test for function <testName2>', <testFunction2>)
	...
	...
}
```


&nbsp;  
&nbsp;
_______
_______
&nbsp;
![](readme_images/Marco__2_-removebg.png)
## Alessio
###### Setting Up
We set up our `BranchAnalyzer` [[commit](https://github.com/T0mexX/hugo-sep/commit/b2c03cb40f90bf92bbbe7aae49b229a3927ee393)]. In this case the total branches are 19 and the number of function analyzed 6.

```go
var ba = BranchAnalyzer{
	filename: "strings.go",
	branches: [19]bool{},
	functions: [6]Function{
		{name: "ToString", startBranchId: 0, untilId: 3},
		{name: "Eq", startBranchId: 3, untilId: 6},
		...
	},
}
```

&nbsp;  
Then we add the branch flag setters to the selected functions.

***Function 1:*** `ToString` &nbsp;  
***File:*** `common/hstrings/strings.go`

```go
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
```

&nbsp;  
&nbsp;  
***Function 2:*** `Eq` &nbsp;  
***File:*** `common/hstrings/strings.go`

```go
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
```

&nbsp; 
###### Coverage Results Before Improvements
As we can see from both our own *branch coverage*, and the external *statement coverage* tools, there is 1 branch out of 6 (*16.67%*) covered.

![](readme_images/strings_coverage_before_alessio.png)

&nbsp;   
![](readme_images/ToString_statement_coverage.png)
![](readme_images/Eq_statement_coverage.png)

&nbsp;  
###### Tests
Consider also the following declarations, that are needed to perform the tests.
>NORAH SAYS: Maybe we can explain why you implemented this struct
```go
type StringerImplementation struct{ str string }
func (si StringerImplementation) String() string { return si.str }
```

&nbsp;  
***Function 1:*** `ToString` [[commit](https://github.com/T0mexX/hugo-sep/commit/69cc0189e4c4f6af55c2d58bac4af44fc2495ac5)]&nbsp;  
***File:*** `common/hstrings/strings.go`

```go
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
```

&nbsp;  
&nbsp;  
***Function 2:*** `Eq` [[commit](https://github.com/T0mexX/hugo-sep/commit/ef6119620725c841ea83af37731adf627bbde815)]&nbsp;  
***File:*** `common/hstrings/strings.go`

```go
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
```

&nbsp;   
Tests result of functions `ToString` and `Eq`:

![](readme_images/verbose_tests_strings_alessio.png)

&nbsp;  
###### Coverage Results After Improvements
Considering these 2 functions, we went from *1/6* (*16.67%*) to *6/6* (*100%*) branches covered. Improving these 2 functions branch coverage concerned about passing parameter of different types. By defining test cases with parameter of type `string`, `Stringer` and a third different type (in our case `int`), we were able to reach all branches.

![](readme_images/strings_coverage_after_alessio.png)

&nbsp;   
![](readme_images/ToString_statement_coverage_after.png)
![](readme_images/Eq_statement_coverage_after.png)

&nbsp;  
&nbsp;  
## Marco
###### Setting Up 
The `BranchAnalyzer` and the flag setters are already set up [[commit](https://github.com/T0mexX/hugo-sep/commit/fd3a355808d73476661b655fafe999ec984622a5)].

```go
var ba = BranchAnalyzer{
	filename: "strings.go",
	branches: [19]bool{},
	functions: [6]Function{
		...
		{name: "InSlice", startBranchId: 13, untilId: 16},
		{name: "InSlicEqualFold", startBranchId: 16, untilId: 19},
	},
}
```

&nbsp;  
***Function 1:*** `InSlice`           
***File:*** `common/hstrings/strings.go`

```go
func  InSlice(arr []string, el string) bool {
	for  _, v  :=  range arr {
		if v == el { // branch id = 13 (if condition evaluates to true at least once)
			ba.reachedBranch(13)
			return  true
		}
		// (else)
		// branch id = 14 (if condition evaluates to false at least once)
		ba.reachedBranch(14)
	}
	// (else)
	// branch id = 15 (if condition always evaluates to false)
	ba.reachedBranch(15)
	return  false
}
```

&nbsp;  
&nbsp;  
***Function 2:*** `InSliceEqualFold`      
***File:*** `common/hstrings/strings.go`

```go
func  InSlicEqualFold(arr []string, el string) bool {
	for  _, v  :=  range arr {
		if strings.EqualFold(v, el) { // branch id = 16 (if condition evaluates to true at least 	once)
			ba.reachedBranch(16)
			return  true
		}
		// (else)
		// branch id = 17 (if condition evaluates to false at least onece 
		ba.branch(17)
	}
	// (else)
	// branch id = 18 (if condition always evaluates to false)
	ba.reachedBranch(18)
	return  false

}
```

&nbsp;   
###### Coverage Result Before Improvements
As we can see from both our own *branch coverage*, and the external *statement coverage* tools, there are 0 branches out of 6 (*0%*) covered.

![](readme_images/marco_string.png)

&nbsp;   
![](readme_images/marco_statement_cover.png)

&nbsp;  
###### Tests
***Function 1:*** `InSlice` [[commit](https://github.com/gohugoio/hugo/commit/6f60dc6125af5db5f8221185e82453280c7250ae)].    
***File:*** `common/hstrings/strings.go`

```go
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
```

&nbsp;  
&nbsp;   
***Function 2:*** `InSliceEqualFold` [[commit](https://github.com/gohugoio/hugo/commit/6f60dc6125af5db5f8221185e82453280c7250ae)].  
***File:*** `common/hstrings/strings.go`

```go
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
```

&nbsp;  
In the red boxes below we show the outcome of the tests for `InSlice`  and `InSliceEqualFold`.
  
![](readme_images/marco_verbose_tests_string.png)

&nbsp;  
###### Coverage Result After Improvement
Considering these 2 functions, we went from *0/6* (*0%*) to *6/6* (*100%*) branches covered. The function `InSlice` checks if a string is contained in an Array of strings. Additionally, the function `InSliceEqualFold`, taking as parameters an Array of string and a string, creates an `EqualFold` object with those two parameters and checks if the string is contained in the Array. Thanks to this behavioural similarity, we used the same tests for both functions.
>NORAH SAYS: explain better what you are testing, which cases did you include?

![](readme_images/marco_strings_coverage_after.png)
 
&nbsp;  
![](readme_images/marco_statement_cover_final.png)

&nbsp;  
&nbsp;  
## Norah
###### Setting Up
***Function 1:*** `EqualAny` &nbsp;  
***File:*** `common/hstrings/strings.go`

The `BranchAnalyzer` and the flag setters for `common/hstrings/strings.go` are already set up [[commit](https://github.com/T0mexX/hugo-sep/commit/fd3a355808d73476661b655fafe999ec984622a5)].

```go
var ba = BranchAnalyzer{
	filename: "strings.go",
	branches: [19]bool{},
	functions: [6]Function{
		...
		{name: "EqualAny", startBranchId: 6, untilId: 9},
		...
	},
}
```

```go
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
```

&nbsp;  
&nbsp;  
***Function 2:*** `IsFloat` &nbsp;   
***File:*** `common/hreflect/helpers.go`

We set up our `BranchAnalyzer` for `common/hreflect/helpers.go` [[commit](https://github.com/T0mexX/hugo-sep/commit/fd3a355808d73476661b655fafe999ec984622a5)].

```go
var ba = BranchAnalyzer{
	filename: "helpers.go",
	branches: [6]bool{},
	functions: [3]Function{
		...
		{name: "IsFloat", startBranchId: 4, untilId: 6},
	},
}
```

Then add the flag setters in the function [[commit](https://github.com/T0mexX/hugo-sep/commit/97fc43e4f2f34f6b962e3d3f7fb4d5efacb2242e)].

```go
// IsFloat returns whether the given kind is a float.
func IsFloat(kind reflect.Kind) bool {
	switch kind {
	case reflect.Float32, reflect.Float64:
		ba.reachedBranch(4)
		return true // branch id = 4
	default:
		ba.reachedBranch(5)
		return false // branch id = 5
	}
}
```

&nbsp;
###### Coverage Result Before Improvements
***Function 1:*** `EqualAny` &nbsp;  
***File:*** `common/hstrings/strings.go`

As we can see from both our own *branch coverage*, and the external *statement coverage* tools, there are 0 branches out of 2 (*0%*) covered.

![](readme_images/EqualAny_Coverage_Before.png)

&nbsp;
![](readme_images/EqualAny_statement_coverage.png)

&nbsp;  
&nbsp;  
***Function 2:*** `IsFloat` &nbsp;  
***File:*** `common/hreflect/helpers.go` 

As we can see from both our own *branch coverage*, and the external *statement coverage* tools, there are 0 branches out of 2 (*0%*) covered.

![](readme_images/IsFloat_Coverage_Before.png)

&nbsp;
![](readme_images/IsFloat_statement_coverage.png)

&nbsp;   
###### Tests
***Function 1:*** `EqualAny` [[commit](https://github.com/T0mexX/hugo-sep/commit/95a766930486ea4433912cd7bad2480c1df21ba1)]&nbsp;  
***File:*** `common/hstrings/strings.go`

```go
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
```

![](readme_images/verbose_tests_equalAny.png)

&nbsp;  
&nbsp;   
***Function 2:*** `IsFloat` [[commit](https://github.com/T0mexX/hugo-sep/commit/97fc43e4f2f34f6b962e3d3f7fb4d5efacb2242e)]&nbsp;   
***File:*** `common/hreflect/helpers.go` 

```go
t.Run("test for function 'IsFloat'", func(t *testing.T) {

	testCases := [8]struct {
		input    reflect.Kind
		expected bool
	}{

		{input: reflect.Float32, expected: true},
		{input: reflect.Float64, expected: true},
		{input: reflect.Uint8, expected: false},
		{input: reflect.Uint16, expected: false},
		{input: reflect.Int, expected: false},
		{input: reflect.Int8, expected: false},
		{input: reflect.Bool, expected: false},
		{input: reflect.Chan, expected: false},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase: %v", testCase), func(t *testing.T) {
			boolOut := IsFloat(testCase.input)
			assert.Equal(t, testCase.expected, boolOut)
		})
	}
})
```

![](readme_images/verbose_tests_isFloat.png)

&nbsp;  
###### Coverage Results After Improvements
***Function 1:*** `EqualAny` &nbsp;  
***File:*** `common/hstrings/strings.go`

We went from *0* (*0%*) to *3/3* (*100%*) branches covered. The function takes multiple `String`s as parameters and checks if the first `String` provided is equal to any of the other input `String`s. To test the function we made a few test cases that check, given some input `String`s, if the return value is as expected.

![](readme_images/EqualAny_Coverage_After.png)

&nbsp;
![](readme_images/EqualAny_statement_coverage_after.png)

&nbsp;  
&nbsp;  
***Function 2:*** `IsFloat` &nbsp;   
***File:*** `common/hreflect/helpers.go`

We went from *0* (*0%*) to *3/3* (*100%*) branches covered. The function gets an input and then checks if, the given parameter, is of type `Float`. To test the function we made a few test cases that check, given different input types (`Uint`, `String`, `Bool`, `Int`, `Chan` and `Float`), that the outcome is as expected (ex: `Uint8` -> `False`, `Float8` -> `True`).

![](readme_images/IsFloat_Coverage_After.png)

&nbsp;
![](readme_images/IsFloat_statement_coverage_after.png)

&nbsp;  
&nbsp;  
## Extra functions
###### Setting Up

The `BranchAnalyzer` for `common/hreflect/helpers.go` is already set up [[commit](https://github.com/T0mexX/hugo-sep/commit/fd3a355808d73476661b655fafe999ec984622a5)].

```go
var ba = BranchAnalyzer{
	filename: "helpers.go",
	branches: [6]bool{},
	functions: [3]Function{
		{name: "IsUint", startBranchId: 0, untilId: 2},
		{name: "IsInt", startBranchId: 2, untilId: 4},
		...
	},
}
```

&nbsp;  
We add the flag setters in the functions [[commit](https://github.com/T0mexX/hugo-sep/commit/97fc43e4f2f34f6b962e3d3f7fb4d5efacb2242e)].
   
***Function 1:*** `IsUint` &nbsp;  
***File:*** `common/hreflect/helpers.go`

```go
// IsUint returns whether the given kind is an uint.
func IsUint(kind reflect.Kind) bool {
	switch kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ba.reachedBranch(0)
		return true	// branch id = 0 
	default:
		ba.reachedBranch(1)
		return false	// branch id = 1 
	}
}
```

&nbsp;  
&nbsp;   
***Function 2:*** `IsInt` &nbsp;  
***File:*** `common/hreflect/helpers.go`

```go
// IsInt returns whether the given kind is an int.
func IsInt(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ba.reachedBranch(2)
		return true // branch id = 2 
	default:
		ba.reachedBranch(3)
		return false // branch id = 3
	}
}
```

&nbsp;  
###### Coverage Results Before Improvements
As we can see from both our own *branch coverage*, and the external *statement coverage* tools, there are 0 branches out of 4 (*0%*) covered.

![](readme_images/IsInt_IsUint_Coverage_before.png)

&nbsp;
![](readme_images/IsUint_statement_coverage.png)
![](readme_images/IsInt_statement_coverage.png)

&nbsp;  
###### Tests
***Function 1:*** `IsUint` [[commit](https://github.com/T0mexX/hugo-sep/commit/fd3a355808d73476661b655fafe999ec984622a5)] &nbsp;    
***File:*** `common/hreflect/helpers.go`

```go
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
```

![](readme_images/verbose_tests_isUint.png)

&nbsp;  
&nbsp;   
***Function 2:*** `IsInt` [[commit](https://github.com/T0mexX/hugo-sep/commit/97fc43e4f2f34f6b962e3d3f7fb4d5efacb2242e)] &nbsp;   
***File:*** `common/hreflect/helpers.go`

```go
t.Run("test for function 'IsInt'", func(t *testing.T) {

	testCases := [10]struct {
		input    reflect.Kind
		expected bool
	}{

		{input: reflect.Int8, expected: true},
		{input: reflect.Int16, expected: true},
		{input: reflect.Int32, expected: true},
		{input: reflect.Int64, expected: true},
		{input: reflect.Int, expected: true},
		{input: reflect.Bool, expected: false},
		{input: reflect.Chan, expected: false},
		{input: reflect.Uint16, expected: false},
		{input: reflect.Uint32, expected: false},
		{input: reflect.Uint8, expected: false},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase: %v", testCase), func(t *testing.T) {
			boolOut := IsInt(testCase.input)
			assert.Equal(t, testCase.expected, boolOut)
		})
	}
})
```

![](readme_images/verbose_tests_isInt.png)

&nbsp;
###### Coverage Results After Improvements
***Function 1:*** `IsUint` &nbsp;   
***File:*** `common/hreflect/helpers.go`

We went from *0* (*0%*) to *3/3* (*100%*) branches covered. The function gets an input and then checks if, the given parameter, is of type `Uint`. To test the function we made a few test cases that check, given different input types (`Uint`, `Bool`, `Int` and `Chan`), that the outcome is as expected (ex: `Uint8` -> `True`, `Float8` -> `False`).

![](readme_images/isUnit_Coverage_After.png)

&nbsp;
![](readme_images/IsUint_statement_coverage_after.png)

&nbsp;  
&nbsp;     
***Function 2:*** `IsInt` &nbsp;      
***File:*** `common/hreflect/helpers.go`

We went from *0* (*0%*) to *3/3* (*100%*) branches covered. The function gets an input and then checks if, the given parameter, is of type `Int`. To test the function we made a few test cases that check, given different input types (`Uint`, `Bool`, `Int` and `Chan`), that the outcome is as expected (ex: `Bool` -> `False`, `Int` -> `True`).

![](readme_images/isInt_Coverage_After.png)

&nbsp;
![](readme_images/IsInt_statement_coverage_after.png)


&nbsp;  
&nbsp;
_______
_______ 
&nbsp;  
## Overall


### Package ``hstrings``
Statement coverage before improvements [[file](common/hstrings/original_cover/statement_cover_list.txt)].

![](readme_images/statement_coverage_list_before_hstrings.png)

&nbsp;  
Statement coverage after improvements [[file](common/hstrings/statement_cover_list.txt)].

![](readme_images/statement_coverage_list_after_hstrings.png)

&nbsp;  
### Package ``hreflect``
Statement coverage before improvements [[file](common/hreflect/initial/statement_cover_list.txt)].

![](readme_images/statement_coverage_list_before_hreflect.png)

&nbsp;  
Statement coverage after improvements [[file](common/hreflect/statement_cover_list.txt)]:

![](readme_images/statement_coverage_list_after_hreflect.png)

&nbsp;  
### All Packages
Statement coverage before improvements [[complete file](common/hstrings/original_cover/statement_cover_list.txt)].

![](readme_images/total_statement_coverage.png)

&nbsp;  
Statement coverage after improvements [[complete file](covers/final/statement_cover_list.txt)]:

![](readme_images/total_statement_coverage_after.png)

>***REMEMBER:*** The project has more than 200.000 lines of code.

&nbsp;  
&nbsp;
_______
_______ 
&nbsp;    
## Statement of individual contributions
### Alessio
- Improved coverage for functions `ToString` and `Eq` in file `common/hstrings/strings.go`.
- Added the `README.md` sections related to the above mentioned functions.


### Marco
- Improved coverage for functions `InSlice` and `InSliceEqualFold` in file `common/hstrings/strings.go`.
- Added the `README.md` sections related to the above mentioned functions.
 

### Norah
- Improved coverage for functions `EqualAny` and `IsFloat` in the files `common/hstrings/strings.go` and `common/hreflect/helpers.go`.
- Added the `README.md` sections related to the above mentioned functions.


### Everyone
- Wrote introductory `README.md` section.
- Created the structures for measuring ***branch coverage***.
- Defined the function that formats the ***branch coverage*** to `branch_coverage.txt` file.
- Improved coverage for functions `IsUint` and `IsInt` in file `common/hreflect/helpers.go`.
- Added the `README.md` sections related to the above mentioned functions.
