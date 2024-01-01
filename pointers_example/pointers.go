package main

import "fmt"

/* Advantages of using pointeres
- avoid Unnecessary value copies
- directly mutate values
*/

func main() {
	age := 32

	agePointer := &age
	// to dereference *agePointer
	// to get pointer &age

	fmt.Println("Age: ", *agePointer)
	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)

	mutateValue(&age)
	fmt.Println("Age: ", age)

}

func getAdultYears(age *int) int {
	return *age - 18
}

func mutateValue(age *int) {
	*age = *age - 10
}
