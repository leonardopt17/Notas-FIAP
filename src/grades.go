package main

var grades = []Class{
	{
		Name:    "Nome da matéria",
		Classes: 40,
		Misses:  1,
		FirstSemester: Semester{
			Cp1:     9,
			Cp2:     8,
			Sprint1: 8.5,
			Sprint2: 8.5,
			Gs:      9.7,
		},
		SecondSemester: Semester{
			Cp1:     10,
			Cp2:     9,
			Sprint1: 9.8,
			Sprint2: 7,
			Gs:      0,
		},
	},
}
