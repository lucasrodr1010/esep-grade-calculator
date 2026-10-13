# ESEP Grade Calculator

A small Go library that turns a student's assignment, exam, and essay grades into one final letter grade.

## How the final grade is calculated

1. Each grade is added with a name, a score from 0 to 100, and a type: `Assignment`, `Exam`, or `Essay`.
2. The scores of each type are averaged.
3. The three averages are combined with these weights:

| Grade type  | Weight |
|-------------|--------|
| Assignments | 50%    |
| Exams       | 35%    |
| Essays      | 15%    |

4. The weighted score is converted to a letter:

| Weighted score | Letter |
|----------------|--------|
| 90 and above   | A      |
| 80 to 89       | B      |
| 70 to 79       | C      |
| 60 to 69       | D      |
| below 60       | F      |

Passing requires a C or better, which is a weighted score of 70 or more.

## Letter grades or Pass/Fail

The calculator is created in one of two grading modes:

| Mode | `GetFinalGrade` returns |
|------|-------------------------|
| `LetterGrade` | `"A"`, `"B"`, `"C"`, `"D"` or `"F"` |
| `PassFail` | `"Pass"` for a C or better, `"Fail"` below that |

A kind of grade with no grades in it counts as 0.

## Project layout

| File | What it is |
|------|------------|
| `README.MD` | This file. |
| `grade-calculator/go.mod` | Go module file. Names the module `esep/grade-calculator` and the Go version it needs. |
| `grade-calculator/grade_calculator.go` | The library: the `GradeCalculator` type, `AddGrade`, and `GetFinalGrade`. |
| `grade-calculator/grade_calculator_test.go` | Unit tests, written with Go's built-in `testing` package. |

There are no third-party dependencies.

## Requirements

- Go 1.24 or newer (`go.mod` asks for 1.24.4).
- An older Go, 1.21 or newer, also works: it downloads the 1.24.4 toolchain by itself the first time you run a `go` command in the module folder.

Check your version with `go version`. Install Go from https://go.dev/dl/.

## Using the library

```go
calculator := NewGradeCalculator(LetterGrade)

calculator.AddGrade("open source assignment", 95, Assignment)
calculator.AddGrade("exam 1", 88, Exam)
calculator.AddGrade("essay on ai ethics", 91, Essay)

finalGrade := calculator.GetFinalGrade() // "A": 47 + 30 + 13 = 91, with integer math

passFail := NewGradeCalculator(PassFail)
passFail.AddGrade("open source assignment", 95, Assignment)
passFail.AddGrade("exam 1", 88, Exam)
passFail.AddGrade("essay on ai ethics", 91, Essay)

result := passFail.GetFinalGrade() // "Pass"
```

## Running the tests

All commands are run from the module folder. Note the dash in its name:

```
cd grade-calculator
```

Run every test:

```
go test ./...
```

Run every test and print one line per test:

```
go test -v ./...
```

Run one test by name:

```
go test -run TestGetGradeA ./...
```

## Code coverage

Print the percentage of statements the tests execute:

```
go test -cover ./...
```

Write a coverage profile, then read it per function or as highlighted source in a browser:

```
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out
```

`coverage.out` is a generated file. Do not commit it.
