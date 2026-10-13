package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator(LetterGrade)

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator(LetterGrade)

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator(LetterGrade)

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 45, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 40, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator(LetterGrade)

	gradeCalculator.AddGrade("open source assignment", 75, Assignment)
	gradeCalculator.AddGrade("exam 1", 72, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator(LetterGrade)

	gradeCalculator.AddGrade("open source assignment", 65, Assignment)
	gradeCalculator.AddGrade("exam 1", 62, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// Each kind of grade is averaged on its own before the weights are applied.
// Assignments average 90, exams 80, essays 80: 45 + 28 + 12 = 85, a B.
func TestGetGradeAveragesEachGradeType(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator(LetterGrade)

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("unit testing assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 90, Exam)
	gradeCalculator.AddGrade("exam 2", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)
	gradeCalculator.AddGrade("essay on code review", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// Assignments count for 50%, exams for 35% and essays for 15%.
func TestGetGradeWeightsEachGradeType(t *testing.T) {
	testCases := []struct {
		name       string
		assignment int
		exam       int
		essay      int
		expected   string
	}{
		{"only assignments earn points: 50", 100, 0, 0, "F"},
		{"only exams earn points: 35", 0, 100, 0, "F"},
		{"only essays earn points: 15", 0, 0, 100, "F"},
		{"assignments and exams: 85", 100, 100, 0, "B"},
		{"assignments and essays: 65", 100, 0, 100, "D"},
		{"exams and essays: 50", 0, 100, 100, "F"},
	}

	for _, testCase := range testCases {
		gradeCalculator := NewGradeCalculator(LetterGrade)

		gradeCalculator.AddGrade("assignment", testCase.assignment, Assignment)
		gradeCalculator.AddGrade("exam", testCase.exam, Exam)
		gradeCalculator.AddGrade("essay", testCase.essay, Essay)

		actual_value := gradeCalculator.GetFinalGrade()

		if testCase.expected != actual_value {
			t.Errorf("%s: expected GetGrade to return '%s'; got '%s' instead", testCase.name, testCase.expected, actual_value)
		}
	}
}

// The lowest score of each letter, and the score one point below it.
func TestGetGradeLetterBoundaries(t *testing.T) {
	testCases := []struct {
		score    int
		expected string
	}{
		{100, "A"},
		{90, "A"},
		{89, "B"},
		{80, "B"},
		{79, "C"},
		{70, "C"},
		{69, "D"},
		{60, "D"},
		{59, "F"},
		{0, "F"},
	}

	for _, testCase := range testCases {
		gradeCalculator := NewGradeCalculator(LetterGrade)

		gradeCalculator.AddGrade("assignment", testCase.score, Assignment)
		gradeCalculator.AddGrade("exam", testCase.score, Exam)
		gradeCalculator.AddGrade("essay", testCase.score, Essay)

		actual_value := gradeCalculator.GetFinalGrade()

		if testCase.expected != actual_value {
			t.Errorf("score %d: expected GetGrade to return '%s'; got '%s' instead", testCase.score, testCase.expected, actual_value)
		}
	}
}

// A weighted score that lands exactly on 60 must be a D.
// 23*0.50 + 97*0.35 + 97*0.15 = 60.
func TestGetGradeExactBoundaryFromMixedGrades(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator(LetterGrade)

	gradeCalculator.AddGrade("open source assignment", 23, Assignment)
	gradeCalculator.AddGrade("exam 1", 97, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 97, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeWithNoGrades(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator(LetterGrade)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// A kind of grade with nothing in it counts as 0. One perfect assignment is 50 points.
func TestGetGradeWithOneGradeTypeOnly(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator(LetterGrade)

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestAddGradeIgnoresUnknownGradeType(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator(LetterGrade)

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)
	gradeCalculator.AddGrade("not a real kind of grade", 0, GradeType(99))

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	testCases := []struct {
		gradeType GradeType
		expected  string
	}{
		{Assignment, "assignment"},
		{Exam, "exam"},
		{Essay, "essay"},
	}

	for _, testCase := range testCases {
		actual_value := testCase.gradeType.String()

		if testCase.expected != actual_value {
			t.Errorf("Expected String to return '%s'; got '%s' instead", testCase.expected, actual_value)
		}
	}
}

// Pass/Fail evaluation: passing requires a C or better, which is a score of 70 or more.
func TestPassFailReturnsPassForPassingGrades(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator(PassFail)

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestPassFailReturnsFailForFailingGrades(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := NewGradeCalculator(PassFail)

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 45, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 40, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// A C passes and a D does not.
func TestPassFailRequiresCOrBetter(t *testing.T) {
	testCases := []struct {
		score    int
		expected string
	}{
		{100, "Pass"},
		{90, "Pass"},
		{80, "Pass"},
		{70, "Pass"},
		{69, "Fail"},
		{60, "Fail"},
		{0, "Fail"},
	}

	for _, testCase := range testCases {
		gradeCalculator := NewGradeCalculator(PassFail)

		gradeCalculator.AddGrade("assignment", testCase.score, Assignment)
		gradeCalculator.AddGrade("exam", testCase.score, Exam)
		gradeCalculator.AddGrade("essay", testCase.score, Essay)

		actual_value := gradeCalculator.GetFinalGrade()

		if testCase.expected != actual_value {
			t.Errorf("score %d: expected GetGrade to return '%s'; got '%s' instead", testCase.score, testCase.expected, actual_value)
		}
	}
}

func TestPassFailWithNoGrades(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := NewGradeCalculator(PassFail)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// Asking for letter grades still returns letters.
func TestLetterGradeModeReturnsLetters(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator(LetterGrade)

	gradeCalculator.AddGrade("open source assignment", 75, Assignment)
	gradeCalculator.AddGrade("exam 1", 72, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
