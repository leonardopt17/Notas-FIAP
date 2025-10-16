package main

import "fmt"

var FIRST_SEMESTER_WEIGHT = 0.4
var SECOND_SEMESTER_WEIGHT = 0.6

var CPS_AND_SPRINTS_WEIGHT = 0.4
var GS_WEIGHT = 0.6

var MAX_MISSING_PR = 0.25

type Semester struct {
	Cp1     float64 `json:"cp1"`
	Cp2     float64 `json:"cp2"`
	Sprint1 float64 `json:"sprint1"`
	Sprint2 float64 `json:"sprint2"`
	Gs      float64 `json:"gs"`
}

type Class struct {
	Name           string   `json:"name"`
	FirstSemester  Semester `json:"first_semester"`
	SecondSemester Semester `json:"second_semester"`
	Classes        int      `json:"classes"`
	Misses         int      `json:"misses"`
}

func (c Class) CalculateFinalGrade() float64 {
	firstSemesterFinalGrade := CalculateSemesterGrade(c.FirstSemester)

	secondSemesterCpsAndSprintsGrade := (c.SecondSemester.Cp1 + c.SecondSemester.Cp2 + c.SecondSemester.Sprint1 + c.SecondSemester.Sprint2) / 4
	secondSemesterFinalGrade := secondSemesterCpsAndSprintsGrade*CPS_AND_SPRINTS_WEIGHT + c.SecondSemester.Gs*GS_WEIGHT // c.SecondSemester.Gs is alaways 0

	return firstSemesterFinalGrade*FIRST_SEMESTER_WEIGHT + secondSemesterFinalGrade*SECOND_SEMESTER_WEIGHT
}

func (c Class) CalculateMissingGrade() float64 {
	firstSemesterFinalGrade := CalculateSemesterGrade(c.FirstSemester)
	secondSemesterCpsAndSprintsGrade := (c.SecondSemester.Cp1 + c.SecondSemester.Cp2 + c.SecondSemester.Sprint1 + c.SecondSemester.Sprint2) / 4

	required := (6 - firstSemesterFinalGrade*FIRST_SEMESTER_WEIGHT - secondSemesterCpsAndSprintsGrade*CPS_AND_SPRINTS_WEIGHT*GS_WEIGHT) / (GS_WEIGHT * GS_WEIGHT)

	switch {
	case required > 10:
		return 10
	case required < 0:
		return 0
	default:
		return required
	}
}

func CalculateSemesterGrade(s Semester) float64 {
	semesterCpsAndSprintsGrade := (s.Cp1 + s.Cp2 + s.Sprint1 + s.Sprint2) / 4
	semesterFinalGrade := semesterCpsAndSprintsGrade*CPS_AND_SPRINTS_WEIGHT + s.Gs*GS_WEIGHT

	return semesterFinalGrade
}

func (c Class) GetMissingGrade() int {
	maxMisses := int(float64(c.Classes) * MAX_MISSING_PR)
	remaining := maxMisses - c.Misses

	if remaining < 0 {
		return 0
	}
	return remaining
}

func (c Class) ToString() string {
	firstSemesterFinalGrade := CalculateSemesterGrade(c.FirstSemester)
	secondSemesterFinalGrade := CalculateSemesterGrade(c.SecondSemester)
	remainingMisses := c.GetMissingGrade()

	missingGrade := c.CalculateMissingGrade()

	hasPassedMessage := "  Nota necessária na GS: " + fmt.Sprintf("%.2f", missingGrade) + " ❌\n"

	if missingGrade <= 0 {
		hasPassedMessage = "  PASSEI PORRA CARALHO !!! ✅\n"
	}

	return c.Name + "\n" +
		"  Média do primeiro semestre: " + fmt.Sprintf("%.2f", firstSemesterFinalGrade) + "\n" +
		"  Média do segundo semestre: " + fmt.Sprintf("%.2f", secondSemesterFinalGrade) + "\n" +
		"  Nota final: " + fmt.Sprintf("%.2f", c.CalculateFinalGrade()) + "\n" +
		hasPassedMessage +
		"  Faltas restantes permitidas: " + fmt.Sprintf("%d", remainingMisses) + "\n"
}
