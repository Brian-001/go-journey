package main
import "fmt"
type Student struct{
	Name string
	Age int
	Grade string
}


func main() {

	//Create a struc instance
	alice := Student{
		Name: "Alice",
		Age: 20,
		Grade: "A",
	}

	fmt.Println("Student", alice) // Student {Alice 20 A}

	//Accessing and modifying fields
	fmt.Println("Name", alice.Name)
	fmt.Println("Age", alice.Age)
	alice.Grade = "B+"
	fmt.Println("Grade", alice.Grade)

	//Creating a struct with zero values
	var bob Student
	fmt.Println("Bob (zero values):", bob) //Bob (zero values): { 0 }


	//Using a pointer to a struct
	dave := &Student{
		Name: "Dave",
		Age: 20,
		Grade: "B",
	}

	fmt.Println("Dave (pointer)", *dave) //Dave (pointer) {Dave 20 B}
	dave.Age = 22
	fmt.Println("Dave age after update", *dave) //Dave age after update {Dave 22 B}

	charlie :=&Student{
		Name: "Charlie",
		Age: 21,
		Grade: "B",
	}
	fmt.Println("Charlie:", charlie)

	//Struct with a slice field
	type Classroom struct{
		Name string
		Students []Student
	}

	class := Classroom{
		Name: "CS101",
		Students: []Student{alice, *charlie},
	}
	fmt.Println("Classroom:", class)
}