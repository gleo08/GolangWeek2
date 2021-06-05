package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GroupPeopleByCity(p []Person) (result map[string][]Person) {
	result = make(map[string][]Person)
	for _, person := range p {
		result[person.City] = append(result[person.City], person)
	}
	return result
}

func MaxValue(input map[string]int) (result map[string]int) {
	result = make(map[string]int)
	max := 0
	for _, value := range input {
		if value > max {
			max = value
		}
	}
	for key, value := range input {
		if value == max {
			result[key] = value
		}
	}
	return result
}

func GroupJobByCity(p []Person) (result map[string][]string) {
	result = make(map[string][]string)
	for _, person := range p {
		result[person.City] = append(result[person.City], person.Job)
	}
	return result
}

func TopJobByNumberInEachCity(input map[string]map[string]int) (result map[string]map[string]int) {
	result = make(map[string]map[string]int)
	for key, value := range input {
		result[key] = MaxValue(value)
	}
	return result
}

func CountNumberEachJob(input []string) (result map[string]int) {
	result = make(map[string]int)
	for _, job := range input {
		result[job]++
	}
	return result
}

func SalaryEachJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.Job] += person.Salary
	}
	return result
}

func GroupPeopleByJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.Job]++
	}
	return result
}

type element struct {
	Key   string
	Value int
}

func Top5JobsByNumber(input map[string]int) (result []element) {
	var ss []element
	for key, value := range input {
		ss = append(ss, element{key, value})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	result = ss[0:5]
	return result
}

func Top5CitiesByNumber(p []Person) (result []element) {
	tmpMap := make(map[string]int)
	for _, person := range p {
		tmpMap[person.City]++
	}
	var ss []element
	for key, value := range tmpMap {
		ss = append(ss, element{key, value})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	result = ss[0:5]
	return result
}

func AverageSalaryByJob(number map[string]int, salary map[string]int) (result map[string]int) {
	result = make(map[string]int)
	for key := range number {
		result[key] = salary[key] / number[key]
	}
	return result
}

func CountPersonByCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.City]++
	}
	return result
}

func SalaryEachCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.City] += person.Salary
	}
	return result
}

func FiveCitiesHasTopAverageSalary(number map[string]int, salary map[string]int) (result []element) {
	tmpMap := make(map[string]int)
	for key := range number {
		tmpMap[key] = salary[key] / number[key]
	}
	var ss []element
	for key, value := range tmpMap {
		ss = append(ss, element{key, value})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	result = ss[0:5]
	return result
}

func CountDeveloperByCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	job := "developer"
	for _, person := range p {
		if person.Job == job {
			result[person.City]++
		}
	}
	return result
}

func SalaryDeveloperByCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	job := "developer"
	for _, person := range p {
		if person.Job == job {
			result[person.City] += person.Salary
		}
	}
	return result
}

func FiveCitiesHasTopSalaryForDeveloper(number map[string]int, salary map[string]int) (result []element) {
	tmpMap := make(map[string]int)
	for key := range number {
		tmpMap[key] = salary[key] / number[key]
	}
	var ss []element
	for key, value := range tmpMap {
		ss = append(ss, element{key, value})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	result = ss[0:5]
	return result
}

func CalculateAge(birthDate string) (result int) {
	now := time.Now()
	ny := now.Year()
	nm := int(now.Month())
	nd := now.Day()
	birthDateSplit := strings.Split(birthDate, "-")
	tmp := make([]int, 0)
	for _, value := range birthDateSplit {
		if x, err := strconv.Atoi(value); err == nil {
			tmp = append(tmp, x)
		} else {
			fmt.Println(err)
		}
	}
	var age int
	if tmp[1] > nm {
		age = ny - tmp[0]
	}
	if tmp[1] == nm {
		if tmp[2] >= nd {
			age = ny - tmp[0]
		} else {
			age = ny - tmp[0] - 1
		}
	} else if tmp[1] < nm {
		age = ny - tmp[0] - 1
	}
	return age
}

func SumAgeEachJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.Job] += CalculateAge(person.Birthday)
	}
	return result
}

func AverageAgePerJob(number map[string]int, sumAge map[string]int) (result map[string]int) {
	result = make(map[string]int)
	for key := range number {
		result[key] = sumAge[key] / number[key]
	}
	return result
}

func NumberPeopleEachCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.City]++
	}
	return result
}

func SumAgeEachCity(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.City] += CalculateAge(person.Birthday)
	}
	return result
}

func AverageAgePerCity(number map[string]int, sumAge map[string]int) (result map[string]int) {
	result = make(map[string]int)
	for key := range number {
		result[key] = sumAge[key] / number[key]
	}
	return result
}
