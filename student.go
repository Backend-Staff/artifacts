package main

import "fmt"

// () :=

type Student struct {
	ID int
	Name string
	Age int16
	Grades []int
	Averge float64
}

func  store (id int, name string, age int16, grades []int) *Student {
	return &Student {
		ID : id,
		Name :name,
		Age : age,
		Grades : grades,
	}	 
}

func allStudent() []*Student {
 
	return []*Student{}
}


func main(){
	

	
	s1 := store(1, "ahmed", 12, []int{1 , 34 , 34})
	students := allStudent()
	students = append(students, s1)
		for _, student := range students {
			fmt.Println("Student Name ", student.Name)
	}
}