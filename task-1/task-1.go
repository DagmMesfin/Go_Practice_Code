package main

import (
	"fmt"
)

type Student struct {
	name      string
	numcourse int
	courses   map[string]float32
	average   float32
	grade     string
}

func main() {

	var student Student

	fmt.Print("Write your name: ")
	fmt.Scan(&student.name)
	fmt.Print("Please write the number of subjects you will take: ")
	fmt.Scan(&student.numcourse)
	fmt.Println()

	for k := 1; k <= student.numcourse; k++ {
		var course string
		var mark float32
		fmt.Println("Start by adding the course #", k, ":")
		fmt.Print("\t-----> Subject Name: ")
		fmt.Scan(&course)
		fmt.Print("\t-----> Mark: ")
		fmt.Scan(&mark)
		fmt.Println()
		student.courses[course] = mark
	}

	student.average = avgCalculator(student.courses)
	student.grade = markGrade(student.average)

	fmt.Printf("Name: %s", student.name)
	fmt.Println("Course Results:")
	fmt.Println("-----------------------------------")
	fmt.Println("    Course\tMark\tGrade")
	for course, mark := range student.courses {
		fmt.Printf("     %s\t%.1f\t%s\n", course, mark, markGrade(mark))
	}

	fmt.Println("-----------------------------------")
	fmt.Printf("\tAverage Mark: %.2f\n", student.average)
	fmt.Printf("\tAverage Grade: %s\n", student.grade)
	fmt.Println()
}

func markGrade(mark float32) string {
	var grade string

	if mark >= 90 {
		grade = "A+"
	} else if mark >= 85 {
		grade = "A"
	} else if mark >= 80 {
		grade = "A-"
	} else if mark >= 75 {
		grade = "B+"
	} else if mark >= 70 {
		grade = "B"
	} else if mark >= 65 {
		grade = "B-"
	} else if mark >= 60 {
		grade = "C+"
	} else if mark >= 50 {
		grade = "C"
	} else if mark >= 45 {
		grade = "C-"
	} else if mark >= 40 {
		grade = "D"
	} else {
		grade = "F"
	}

	return grade
}

func avgCalculator(cr map[string]float32) float32 {
	var sumo float32
	for _, mark := range cr {
		sumo += mark
	}
	avg := float32(sumo) / float32(len(cr))

	return avg
}
