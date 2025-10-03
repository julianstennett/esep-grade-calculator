package esepunittests

import "testing"

func makeUniformGC(g int) *GradeCalculator {
	gc := NewGradeCalculator()
	gc.AddGrade("a", g, Assignment)
	gc.AddGrade("e", g, Exam)
	gc.AddGrade("s", g, Essay)
	return gc
}

func TestGradeTypeString(t *testing.T) {
	cases := map[GradeType]string{
		Assignment: "assignment",
		Exam:       "exam",
		Essay:      "essay",
		GradeType(123): "",
	}

	for gt, want := range cases {
		if got := gt.String(); got != want {
			t.Fatalf("GradeType.String() for %v: want %q, got %q", gt, want, got)
		}
	}
}

func TestComputeAverageEmptyAndNonEmpty(t *testing.T) {
	if got := computeAverage([]Grade{}); got != 0 {
		t.Fatalf("computeAverage(empty) = %d; want 0", got)
	}

	grades := []Grade{
		{Name: "g1", Grade: 90, Type: Assignment},
		{Name: "g2", Grade: 100, Type: Assignment},
	}
	if got := computeAverage(grades); got != 95 {
		t.Fatalf("computeAverage([90,100]) = %d; want 95", got)
	}
}

func TestNewGradeCalculatorAndAddGradeSlices(t *testing.T) {
	gc := NewGradeCalculator()
	if gc == nil {
		t.Fatal("NewGradeCalculator returned nil")
	}

	gc.AddGrade("ass", 77, Assignment)
	gc.AddGrade("exam1", 88, Exam)
	gc.AddGrade("essay1", 99, Essay)

	if len(gc.assignments) != 1 {
		t.Fatalf("assignments len = %d; want 1", len(gc.assignments))
	}
	if len(gc.exams) != 1 {
		t.Fatalf("exams len = %d; want 1", len(gc.exams))
	}
	if len(gc.essays) != 1 {
		t.Fatalf("essays len = %d; want 1", len(gc.essays))
	}
	gc2 := NewGradeCalculator()
	gc2.AddGrade("bad", 50, GradeType(99))
	if len(gc2.assignments) != 0 || len(gc2.exams) != 0 || len(gc2.essays) != 0 {
		t.Fatalf("AddGrade with invalid type should not add anything; got %d, %d, %d",
			len(gc2.assignments), len(gc2.exams), len(gc2.essays))
	}
}

func TestCalculateNumericalGradeRoundingDown(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 89, Assignment)
	gc.AddGrade("e1", 90, Exam)
	gc.AddGrade("s1", 90, Essay)

	if got := gc.calculateNumericalGrade(); got != 89 {
		t.Fatalf("calculateNumericalGrade() = %d; want 89 (checks truncation)", got)
	}

	if got := gc.GetFinalGrade(); got != "B" {
		t.Fatalf("GetFinalGrade() = %q; want \"B\"", got)
	}
}

func TestGetFinalGradeBoundaries(t *testing.T) {
	cases := []struct {
		uniformGrade int
		wantLetter   string
	}{
		{90, "A"},
		{95, "A"},
		{89, "B"},
		{80, "B"},
		{79, "C"},
		{70, "C"},
		{69, "D"},
		{60, "D"},
		{59, "F"},
		{0, "F"},
	}

	for _, c := range cases {
		gc := makeUniformGC(c.uniformGrade)
		got := gc.GetFinalGrade()
		if got != c.wantLetter {
			t.Fatalf("uniform grade %d => GetFinalGrade() = %q; want %q", c.uniformGrade, got, c.wantLetter)
		}
	}
}

func TestNoGradesReturnsF(t *testing.T) {
	gc := NewGradeCalculator()
	// no grades added -> averages 0 -> weighted 0 -> F
	if got := gc.GetFinalGrade(); got != "F" {
		t.Fatalf("GetFinalGrade() with no grades = %q; want \"F\"", got)
	}
}
