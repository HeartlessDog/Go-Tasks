package main

import (
	"fmt"
	"time"
)

type Informer interface{
	GetAge() int
	GetAverageMark() float64
}

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
	students:= Init() 
	for i, v := range students {
		fmt.Printf("%v | %v\n", i+1,v)
	}
	GetInfo(students[0])
}

func Init() [4]Student{
	return [4]Student{
		Student{Person{"Nadya",1997,"+375(33)6161616",},"FAIS","IP",[4]int{6,7,8,9},},
		Student{Person{"Vanya",1998,"+375(33)6464646",},"FAIS","IP",[4]int{9,8,7,6},},
		Student{Person{"Kostya",1998,"+375(29)2121212",},"FAIS","IT",[4]int{6,7,8,9},},
		Student{Person{"Zhenya",1998,"+375(29)8118811",},"FAIS","IT",[4]int{9,8,7,6},},	
	}	
}

func (person Person)GetAge() int{
	return time.Now().Year()-person.YearOfBirth
}

func (student Student) GetAverageMark() float64{
	sum:=0
	for _, v:= range student.Marks{
		sum+=v
	}
	return float64(sum)/float64(len(student.Marks))
}

func (student Student) String() string{
	return fmt.Sprintf("%v, %v, %v, %v, %v, Marks:%v", student.Name, student.YearOfBirth,student.PhoneNumber,student.Faculty,student.Speciality,student.Marks)
}

func GetInfo(info Informer) {
	fmt.Println("--------------------------------------------------------")
	fmt.Println(info)
	fmt.Printf("Age - %v\n",info.GetAge())
	fmt.Printf("Average Mark - %v\n",info.GetAverageMark())	
	fmt.Println("--------------------------------------------------------")
}