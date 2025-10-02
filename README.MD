# ESEP Grade Calculator

This is a simple grade calculator library written in **Go**. It computes a student’s final grade by averaging assignments, exams, and essays, applying class-specific weights, and converting the result to a letter grade.

## Features
- Calculate weighted averages for assignments, exams, and essays.
- Support custom weighting for each grade type.
- Convert numeric grades to standard letter grades.
- Includes unit tests for accuracy.

## Requirements
- Go 1.24+

## Installation
Clone the repo and enter the project folder:
```bash
git clone https://github.com/YOUR-USERNAME/esep-grade-calculator.git
cd esep-grade-calculator
````

## Usagage
Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```
