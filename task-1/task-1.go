package main

import (
	"fmt"
)

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

func main() {
	var numo int
	fmt.Print("Please write the number of subjects you will take: ")
	fmt.Scan(&numo)
	fmt.Println()

	courseReport := make(map[string]float32)
	for k := 1; k <= numo; k++ {
		var course string
		var mark float32
		fmt.Println("Start by adding the course #", k, ":")
		fmt.Print("\t-----> Subject Name: ")
		fmt.Scan(&course)
		fmt.Print("\t-----> Mark: ")
		fmt.Scan(&mark)
		fmt.Println()
		courseReport[course] = mark
	}

	avg_mark := avgCalculator(courseReport)

	fmt.Println("Course Results:")
	fmt.Println("-----------------------------------")
	fmt.Println("    Course\tMark\tGrade")
	for course, mark := range courseReport {
		fmt.Printf("     %s\t%.1f\t%s\n", course, mark, markGrade(mark))
	}

	fmt.Println("-----------------------------------")
	fmt.Printf("\tAverage Mark: %.2f\n", avg_mark)
	fmt.Printf("\tAverage Grade: %s\n", markGrade(avg_mark))
	fmt.Println()
}
