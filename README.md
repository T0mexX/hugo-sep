
# Report for Assignment 1

## Project chosen: [hugo](https://github.com/gohugoio/hugo)


Number of lines of code and the tool used to count it: <TODO>
##### Lines of code: 

![](readme_images/lines_of_code.png)

##### Programming language (for test purposes): *Golang*

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
***ADD IMAGE***

We can alternatevely use the following command to open a html page where we can visually check the statement coverage for each file.
```
go tool cover -html .cover.out
```
***ADD IMAGE***

From the html GUI we were able to identify which packages / files lacked ***statement coverage*** and consequently thos that also lacked ***branch coverage***.
We chose to improve coverage of the package `hstrings` with file `strings.go`.

##### Statement Coverage
Full statement coverage: 

Statement coverage ([complete file](covers/initial/cover_list.txt)):
![](readme_images/total_statement_coverage.png)


<Show the coverage results provided by the existing tool with a screenshot>

### Your own coverage tool

<The following is supposed to be repeated for each group member>

<Group member name>

<Function 1 name>

<Show a patch (diff) or a link to a commit made in your forked repository that shows the instrumented code to gather coverage measurements>

<Provide a screenshot of the coverage results output by the instrumentation>

<Function 2 name>

<Provide the same kind of information provided for Function 1>

## Coverage improvement

### Individual tests

<The following is supposed to be repeated for each group member>

<Group member name>

<Test 1>

<Show a patch (diff) or a link to a commit made in your forked repository that shows the new/enhanced test>

<Provide a screenshot of the old coverage results (the same as you already showed above)>

<Provide a screenshot of the new coverage results>

<State the coverage improvement with a number and elaborate on why the coverage is improved>

<Test 2>

<Provide the same kind of information provided for Test 1>

### Overall

<Provide a screenshot of the old coverage results by running an existing tool (the same as you already showed above)>

<Provide a screenshot of the new coverage results by running the existing tool using all test modifications made by the group>

## Statement of individual contributions

<Write what each group member did>
