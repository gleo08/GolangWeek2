package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Job      string `json:"job"`
	City     string `json:"city"`
	Salary   int    `json:"salary"`
	Birthday string `json:"birthdate"`
}

func (p *Person) String() string {
	return fmt.Sprintf("name: %s, email: %s, job: %s, city: %s, salary: %d, birthday: %s",
		p.Name, p.Email, p.Job, p.City, p.Salary, p.Birthday)
}

func main() {
	jsonFile, err := os.Open("person.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened person.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var people []Person
	json.Unmarshal(byteValue, &people)

	// 2.1 Group people by city
	fmt.Println("Group people by city: ")
	peopleByCity := GroupPeopleByCity(people)
	for key, value := range peopleByCity {
		fmt.Println("-", key)
		for _, person := range value {
			fmt.Println("  - ", (&person).Name)
		}
	}
	fmt.Println("------------------------------------------------------------")
	// 2.2 Group people by job
	fmt.Println("Group people by job: ")
	peopleByJob := GroupPeopleByJob(people)
	for key, value := range peopleByJob {
		fmt.Printf("%s : %d", key, value)
		fmt.Println("")
	}
	fmt.Println("------------------------------------------------------------")
	// 2.3 Top 5 jobs
	fmt.Println("Top 5 jobs by number:")
	top5Jobs := Top5JobsByNumber(peopleByJob)
	for _, job := range top5Jobs {
		fmt.Printf("%s : %d", job.Key, job.Value)
		fmt.Println("")
	}
	fmt.Println("------------------------------------------------------------")
	// 2.4 Top 5 cities
	fmt.Println("Top 5 cities by number:")
	top5Cities := Top5CitiesByNumber(people)
	for _, city := range top5Cities {
		fmt.Printf("%s : %d", city.Key, city.Value)
		fmt.Println("")
	}
	fmt.Println("------------------------------------------------------------")
	// 2.5 Top job in each city
	fmt.Println("Top job in each city: ")
	resultMap := make(map[string]map[string]int)
	jobByCity := GroupJobByCity(people)
	for key, value := range jobByCity {
		result := CountNumberEachJob(value)
		resultMap[key] = result
	}
	topJobEachCity := TopJobByNumberInEachCity(resultMap)
	for key, value := range topJobEachCity {
		fmt.Println("-", key)
		for k, v := range value {
			fmt.Printf("%s : %d", k, v)
			fmt.Println("")
		}
	}
	fmt.Println("------------------------------------------------------------")
	// 2.6 Average salary by job
	fmt.Println("Average salary by job")
	numberEachJob := GroupPeopleByJob(people)
	salaryEachJob := SalaryEachJob(people)
	averageSalaryByJob := AverageSalaryByJob(numberEachJob, salaryEachJob)
	for key, value := range averageSalaryByJob {
		fmt.Printf("%s : %d", key, value)
		fmt.Println("")
	}
	fmt.Println("------------------------------------------------------------")
	// 2.7 Top 5 cities average salary
	fmt.Println("Top 5 cities average salary")
	numberEachCity := CountPersonByCity(people)
	salaryEachCity := SalaryEachCity(people)
	averageSalaryByCity := FiveCitiesHasTopAverageSalary(numberEachCity, salaryEachCity)
	for _, a := range averageSalaryByCity {
		fmt.Printf("%s : %d", a.Key, a.Value)
		fmt.Println("")
	}
	fmt.Println("------------------------------------------------------------")
	// 2.8 Top 5 cities has top salary for developer
	fmt.Println("Top 5 cities has top salary for developer ")
	numberDeveloperEachCity := CountDeveloperByCity(people)
	salaryDeveloperEachCity := SalaryDeveloperByCity(people)
	averageSalaryDeveloperByCity := FiveCitiesHasTopSalaryForDeveloper(numberDeveloperEachCity, salaryDeveloperEachCity)
	for _, a := range averageSalaryDeveloperByCity {
		fmt.Printf("%s : %d", a.Key, a.Value)
		fmt.Println("")
	}
	fmt.Println("------------------------------------------------------------")
	// 2.9 Average age per job
	fmt.Println("Average age per job: ")
	sumAgeEachJob := SumAgeEachJob(people)
	numberPeopleEachJob := GroupPeopleByJob(people)
	averageAgePerJob := AverageAgePerJob(numberPeopleEachJob, sumAgeEachJob)
	for key, value := range averageAgePerJob {
		fmt.Printf("%s : %d", key, value)
		fmt.Println("")
	}
	fmt.Println("------------------------------------------------------------")
	// 2.10 Average age per city
	fmt.Println("Average age per city: ")
	sumAgeEachCity := SumAgeEachCity(people)
	numberPeopleEachCity := NumberPeopleEachCity(people)
	averageAgePerCity := AverageAgePerCity(numberPeopleEachCity, sumAgeEachCity)
	for key, value := range averageAgePerCity {
		fmt.Printf("%s : %d", key, value)
		fmt.Println("")
	}
	fmt.Println("------------------------------------------------------------")
}
