package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	dataset := make(map[string][]string)
	datacourse := make(map[string][]string)
	var highestMentorCourse string
	var highestMentorScore float64

	f, err := os.Open("survey.csv")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	fcsv := csv.NewReader(f)
	r, err := fcsv.ReadAll()
	if err != nil {
		log.Println(err)
	}

	for i := 1; i < len(r); i++ {
		dataset["MentorSatisfaction"] = append(dataset["MentorSatisfaction"], r[i][4])
		dataset["LearningSatisfaction"] = append(dataset["LearningSatisfaction"], r[i][3])
		dataset["CourseName"] = append(dataset["CourseName"], r[i][2])
	}

	for i := 0; i < len(dataset["CourseName"]); i++ {
		datacourse[dataset["CourseName"][i]] = append(datacourse[dataset["CourseName"][i]], dataset["MentorSatisfaction"][i])
	}

	for name, score := range datacourse {
		avrg := Average(score)
		// fmt.Println(name, avrg)  // log rata-rata course
		if avrg > highestMentorScore {
			highestMentorScore = avrg
			highestMentorCourse = name
		}
	}

	fmt.Printf("Rata-rata kepuasan mentor: %.2f\n", Average(dataset["MentorSatisfaction"]))
	fmt.Printf("Rata-rata kepuasan belajar: %.2f\n", Average(dataset["LearningSatisfaction"]))
	fmt.Printf("Course dengan rata-rata kepuasan mentor tertinggi: %s\n", highestMentorCourse)
}

func Sum(arr []string) float64 {
	var sum float64

	for _, v := range arr {
		a, _ := strconv.Atoi(v)
		sum += float64(a)
	}
	return sum
}

func Average(arr []string) float64 {
	sum := Sum(arr)
	len := float64(len(arr))

	return sum / len

}
