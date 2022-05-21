package struct_sample

import "fmt"

type Person struct {
	Name    string
	Address string
}

type Student struct {
	Person
	Code string
}

func DoStructExtendSamples() {
	fmt.Println("\n\n\n構造体継承?複合?テスト")
	s1 := BulkGenStudentStruct()
	fmt.Println(s1, s1.Name, s1.Code)
	s2 := GenStudentStruct()
	fmt.Println(s2, s2.Person.Name, s2.Code)
}

func BulkGenStudentStruct() Student {
	s1 := Student{
		Person: Person{
			Name:    "testname",
			Address: "testaddress",
		},
		Code: "test_code",
	}
	return s1
}

func GenStudentStruct() Student {
	s2 := Student{
		Code: "test_code2",
	}
	s2.Name = "test_name2"
	s2.Address = "test_address2"
	return s2
}
