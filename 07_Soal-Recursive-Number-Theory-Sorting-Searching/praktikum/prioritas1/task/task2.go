package task

import (
	"cmp"
	"fmt"
	"slices"
)

type Student struct {
	Name         string
	MathScore    int
	ScienceScore int
	EnglishScore int
}

func GetInfo(students []Student) {
	bestMathScore := slices.MaxFunc(students, func(i, j Student) int {
		return cmp.Compare(i.MathScore, j.MathScore)
	})
	bestScienceScore := slices.MaxFunc(students, func(i, j Student) int {
		return cmp.Compare(i.ScienceScore, j.ScienceScore)
	})
	bestEnglishScore := slices.MaxFunc(students, func(i, j Student) int {
		return cmp.Compare(i.EnglishScore, j.EnglishScore)
	})

	bestAverageScoreName, bestAverageScore := AverageScore(students)

	fmt.Printf("Best student in math: %v with %v\n", bestMathScore.Name, bestMathScore.MathScore)
	fmt.Printf("Best student in science: %v with %v\n", bestScienceScore.Name, bestScienceScore.ScienceScore)
	fmt.Printf("Best student in english: %v with %v\n", bestEnglishScore.Name, bestEnglishScore.EnglishScore)
	fmt.Printf("Best student in average: %v with %v\n", bestAverageScoreName, bestAverageScore)

}

func AverageScores(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}

	return sum / len(s)
}

func AverageScore(students []Student) (string, int) {
	var nameMaxScore string
	var maxScore int

	studentScore := make(map[string]int)
	for _, student := range students {
		studentScore[student.Name] = (student.EnglishScore + student.MathScore + student.ScienceScore) / 3
	}

	for name, score := range studentScore {
		if maxScore < score {
			maxScore = score
			nameMaxScore = name
		}
	}

	return nameMaxScore, maxScore
}
