package main

import (
	"fmt"
	"time"
)

type Person struct{
	Name string
	YearOfBirth int
	PhoneNumber string
}

type Student struct{
	Person
	Faculty string
	Speciality string
	Marks [4]int
}

func main(){
	student:=Student{Person{"Nadya",1997,"+375(33)6814107",},"FAIS","IP",[4]int{6,7,8,9},} 
	fmt.Printf("Age - %v\n",GetAge(student.Person.YearOfBirth))
	fmt.Printf("Average Mark - %v\n",student.GetAverageMark())
}

func GetAge(year int) int{
	return time.Now().Year()-year
}
func (student Student) GetAverageMark() float64{
	sum:=0
	for _, v:= range student.Marks{
		sum+=v
	}
	return float64(sum)/float64(len(student.Marks))
}