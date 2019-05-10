# Visual Studio Code Plugins

* Go
* Go Autotest
* Markdown Preview Enhanced
* Code Spell Checker
* Markdownlint
* SQL Server

## Learnings

### If oDATA/GraphQL implementations are not performant, need to think about data modeling.

### UNPHAT Approach for Problem Solving

* Understand
* Enumerate
* Paper (white papers)
* History
* Advantages vs Disadvantages
* Think sober

**GO:**

To begin with, look at the documentation for the following packages:

* bufio
* bytes
* errors
* flag
  * Helps to provide options for command line go application. It will do type casting which will ease accepting inputs.
* fmt
* io
* log
* math
* sort
* strconv
* strings
* time

Tobin Harding has really nice implementations for go. 
Follow - tcharding on github

sync pattern in GO for concurrency management.

https://blog.golang.org/subtests
https://medium.com/go-walkthrough

### NOSQL

* Suitable for Availability.
* Major companies use them for analytics purposes.
* Redundancy Factor and Quorum mechanism is used to achieve distributed consensus while returning data.
* ACID properties are not promised to be maintained.
* Not read optimized but suitable for write intensive applications.
* Hard to setup relationship between tables.
* No updates. Only Delete + Insert supported.

### Random Forest

* Quantitative evaluation techniques
* Gini Impurity
* Gini Gain
