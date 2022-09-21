/*
Read the file named applicants.txt once again. Mind one additional column, right after the last exam's result.
This column represents the special exam's score. For example, Willie McBride 76 45 79 80 100 Physics Engineering Mathematics(where 100 is the admission exam's score).

Choose the best score for a student in the ranking: either the mean score for the final exam(s) or the special exam's score.
Use the same set of finals for each Department as in the previous stage. Note that you may need to compare the values several times:
for example, if a student doesn't get accepted to the Department of the first priority,
compare the finals mean score and the special exam's score once again (but this time, for the second priority department).

???
Mean it I should use the special score already for the first exam too if not passed
???
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var orderedDepartments = []string{
	"Biotech",
	"Chemistry",
	"Engineering",
	"Mathematics",
	"Physics",
}

type applicantStruct struct {
	fullName string
	score    float64
}

type applicantPreferences struct {
	fullName     string
	scores       []float64
	specialScore float64
	departments  []string
}

func main() {
	var maxApplicants int
	_, _ = fmt.Scanln(&maxApplicants)

	applicants := readApplicantPreferences()

	exam := map[string][]int{
		"Biotech":     {0, 1},
		"Chemistry":   {1, 1},
		"Engineering": {2, 3},
		"Mathematics": {2, 2},
		"Physics":     {0, 2},
	}

	departments := map[string][]applicantStruct{
		"Biotech":     {},
		"Chemistry":   {},
		"Engineering": {},
		"Mathematics": {},
		"Physics":     {},
	}

	var used []string

	chooseFaculty(applicants, maxApplicants, departments, exam, used)
	prepareFinalOrder(departments)
	showAccepted(departments)
}

func readApplicantPreferences() []applicantPreferences {
	file, err := os.Open("applicants.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var a []applicantPreferences
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		phyScore, _ := strconv.ParseFloat(parts[2], 64)
		chemScore, _ := strconv.ParseFloat(parts[3], 64)
		mathScore, _ := strconv.ParseFloat(parts[4], 64)
		engScore, _ := strconv.ParseFloat(parts[5], 64)
		specialScore, _ := strconv.ParseFloat(parts[6], 64)

		scores := []float64{phyScore, chemScore, mathScore, engScore}

		a = append(a, applicantPreferences{
			parts[0] + " " + parts[1], scores, specialScore, parts[7:],
		})
	}
	return a
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func chooseFaculty(applicants []applicantPreferences, maxApplicants int, departments map[string][]applicantStruct, exam map[string][]int, used []string) {
	for i := 0; i < 3; i++ {
		for _, dep := range orderedDepartments {
			applicantsSorted := sortByDept(applicants, dep)
			for _, applicant := range applicantsSorted {
				if applicant.departments[i] == dep && len(departments[dep]) < maxApplicants && !contains(used, applicant.fullName) {
					score := getHighestScore((applicant.scores[exam[dep][0]]+applicant.scores[exam[dep][1]])/2, applicant.specialScore)

					departments[dep] = append(departments[dep], applicantStruct{applicant.fullName, score})

					used = append(used, applicant.fullName)
				}
			}
		}
	}
}

func getHighestScore(meanScore, specialScore float64) float64 {
	if meanScore > specialScore {
		return meanScore
	} else {
		return specialScore
	}
}

func sortByDept(a []applicantPreferences, dep string) []applicantPreferences {
	switch dep {
	case "Biotech":
		sort.Slice(a, func(i, j int) bool {
			maxScoreI := getHighestScore((a[i].scores[0]+a[i].scores[1])/2, a[i].specialScore)
			maxScoreJ := getHighestScore((a[j].scores[0]+a[j].scores[1])/2, a[j].specialScore)

			if maxScoreI != maxScoreJ {
				return maxScoreI > maxScoreJ
			}
			return strings.Split(a[i].fullName, " ")[0] < strings.Split(a[j].fullName, " ")[0]
		})
	case "Chemistry":
		sort.Slice(a, func(i, j int) bool {
			maxScoreI := getHighestScore(a[i].scores[1], a[i].specialScore)
			maxScoreJ := getHighestScore(a[j].scores[1], a[j].specialScore)

			if maxScoreI != maxScoreJ {
				return maxScoreI > maxScoreJ
			}
			return strings.Split(a[i].fullName, " ")[0] < strings.Split(a[j].fullName, " ")[0]
		})
	case "Engineering":
		sort.Slice(a, func(i, j int) bool {
			maxScoreI := getHighestScore((a[i].scores[3]+a[i].scores[2])/2, a[i].specialScore)
			maxScoreJ := getHighestScore((a[j].scores[3]+a[j].scores[2])/2, a[j].specialScore)

			if maxScoreI != maxScoreJ {
				return maxScoreI > maxScoreJ
			}
			return strings.Split(a[i].fullName, " ")[0] < strings.Split(a[j].fullName, " ")[0]
		})
	case "Mathematics":
		sort.Slice(a, func(i, j int) bool {
			maxScoreI := getHighestScore(a[i].scores[2], a[i].specialScore)
			maxScoreJ := getHighestScore(a[j].scores[2], a[j].specialScore)

			if maxScoreI != maxScoreJ {
				return maxScoreI > maxScoreJ
			}
			return strings.Split(a[i].fullName, " ")[0] < strings.Split(a[j].fullName, " ")[0]
		})
	case "Physics":
		sort.Slice(a, func(i, j int) bool {
			maxScoreI := getHighestScore((a[i].scores[0]+a[i].scores[2])/2, a[i].specialScore)
			maxScoreJ := getHighestScore((a[j].scores[0]+a[j].scores[2])/2, a[j].specialScore)

			if maxScoreI != maxScoreJ {
				return maxScoreI > maxScoreJ
			}
			return strings.Split(a[i].fullName, " ")[0] < strings.Split(a[j].fullName, " ")[0]
		})
	}
	return a
}

func prepareFinalOrder(departments map[string][]applicantStruct) {
	for _, dep := range orderedDepartments {
		sort.Slice(departments[dep], func(i, j int) bool {
			if departments[dep][i].score != departments[dep][j].score {
				return departments[dep][i].score > departments[dep][j].score
			}
			return departments[dep][i].fullName < departments[dep][j].fullName
		})
	}
}

func showAccepted(departments map[string][]applicantStruct) {
	for _, dep := range orderedDepartments {
		fmt.Println(dep)
		fileName, err := os.Create(strings.ToLower(dep) + ".txt")
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range departments[dep] {
			fmt.Printf("%s %.2f\n", v.fullName, v.score)
			_, err = fmt.Fprintf(fileName, "%s %.2f\n", v.fullName, v.score)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println()
	}
}
