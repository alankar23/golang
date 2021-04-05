package main

//  Improting go packages
import (
	"fmt"
)

// Creating a student object
type Student struct {
	name   string
	grades []int
	pass   bool
	age    int
	gender string
}

// Function to return student percentage
func (s Student) getPercentage() float64 {
	sum := 0
	for _, grade := range s.grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.grades))
}

// Returns boolen true or flase for male student gender
func (s Student) isMale() bool {
	if s.gender == "male" {
		return true
	} else {
		return false
	}
}

// Returns student total marks obtained
func (s Student) getSum() int {
	sum := 0
	for _, grade := range s.grades {
		sum += grade
	}
	return sum
}

// Returns whether student passed or not in boolen
func (s *Student) isPass() bool {
	for _, grade := range s.grades {
		if grade < 35 {
			s.pass = false
		} else {
			s.pass = true
		}
	}
	return s.pass
}

func main() {
	Student1 := Student{"Alankar", []int{54, 65, 87, 85, 35}, true, 15, "male"}
	fmt.Println("Student Name :", Student1.name)
	percent := fmt.Sprintf("%.2f", Student1.getPercentage())
	fmt.Println("Percentage :", percent+"%")
	fmt.Println("Is Male :", Student1.isMale())
	fmt.Println("Total Marks :", Student1.getSum())
	fmt.Println("Passed :", Student1.isPass())
}
