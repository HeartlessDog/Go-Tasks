package main

import "fmt"

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
		fmt.Printf("%v | %v, %v, %v, %v, %v, Marks:%v\n", i+1,v.Name, v.YearOfBirth,v.PhoneNumber,v.Faculty,v.Speciality,v.Marks)
	}
}

func Init() [4]Student{
	return [4]Student{
		Student{Person{"Nadya",1997,"+375(33)6161616",},"FAIS","IP",[4]int{6,7,8,9},},
		Student{Person{"Vanya",1998,"+375(33)6464646",},"FAIS","IP",[4]int{9,8,7,6},},
		Student{Person{"Kostya",1998,"+375(29)2121212",},"FAIS","IT",[4]int{6,7,8,9},},
		Student{Person{"Zhenya",1998,"+375(29)8118811",},"FAIS","IT",[4]int{9,8,7,6},},	
	}	
}
