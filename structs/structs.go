package structs


import (
	"fmt"
)

type Student struct {
	firstName, lastName string
	age                 int
}

func (s Student) String() string {
	return fmt.Sprintf("Student: %s %s, Age: %d", s.firstName, s.lastName, s.age)
}

type Teacher struct {
	firstName, lastName string
	university          University
}

type University struct {
	yearBased     int
	infoShort 	  string
}

type Professor struct {
	firstName, lastName string
	age 			    int
	gradeWork           string
	AdditionalInfo // Struct embedding
}

type AdditionalInfo struct {
	phoneNumber, email string
}

func ShowStructsExamples() {
    student1 := Student{"John", "Doe", 20}
	student2 := Student{
		firstName: "Jane",
		lastName:  "Smith",
		age:       22,
	}
	student3 := Student{
		firstName: "Alice",
	}

	anotStudent := struct {
		age int
		groupId int
		professorName string
	} {
		age: 30,
		groupId: 1,
		professorName: "Dr. Smith",
	}

	fmt.Println(student1) // Student: John Doe, Age: 20
	fmt.Println(student2) // Student: Jane Smith, Age: 22
	fmt.Println(student3) // Student: Alice , Age: 0
	fmt.Println("Anonymous Student:", anotStudent) // Anonymous Student: {30 1 Dr. Smith}


	studPointer := &Student{"Bob", "Brown", 25}
	fmt.Printf("Type of pointer: %T\n", studPointer) // Type of pointer: *struct { firstName, lastName string; age int }
	fmt.Println("Pointer:", studPointer) // Pointer to Student: &{Bob Brown 25}
	studPointer.age = 30 // Change value via pointer
	fmt.Println("Pointer to Student:", studPointer) // Pointer to Student: &{Bob Brown 30}
}

func ShowStructsExamples2() {
    university := University{yearBased: 2023, infoShort: "MIT"}
	teacher := Teacher{"Alice", "Johnson", university}
	teacher2 := Teacher{
		firstName: "Bob",
		lastName:  "Smith",
		university: University{
			yearBased: 2022,
			infoShort: "Harvard",
		},
	}
	fmt.Println("Teacher:", teacher) // Teacher: {Alice Johnson {2023 MIT}}
	fmt.Println("Teacher2:", teacher2) // Teacher2: {Bob Smith {2022 Harvard}}

	professorKim := Professor{
		firstName: "Kim",
		lastName:  "Lee",
		age: 45,
		gradeWork: "Physics",
		AdditionalInfo: AdditionalInfo{
			phoneNumber: "123-456-7890",
			email:       "kim.lee@gdfdf.com",
		},
	};

	fmt.Println("Professor Kim additional info: ", professorKim.email, professorKim.phoneNumber)
	fmt.Println("Professor Kim: ", professorKim) //Professor Kim: {Kim Lee 45 Physics {123-456-7890 kim.lee@gdfdf.com}};


	// structs comparison
	student1 := Student{"John", "Doe", 20}
	student2 := Student{"John", "Doe", 20}
	fmt.Println("Student1 == Student2:", student1 == student2) // Student1 == Student2: true
}