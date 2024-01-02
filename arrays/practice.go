package main

import "fmt"

type Product struct { 
	id string
	title string
	price float64
}

func main() {
	hobies := [3]string{"Sports","Reading","Coding"}

	fmt.Println(hobies)
	firstElement := hobies[0]
	fmt.Println("First Element: ",firstElement)
	fmt.Println("Second and Third Element list: ",hobies[1:3])

	mainHobies := hobies[:2]
	fmt.Println("First and Second Element: ",mainHobies)
	
	fmt.Println("Capacity of main hobbies: ",cap(mainHobies))
	fmt.Println("Len of main hobbies: ",len(mainHobies))

	courseGoal := []string{"Learn Go","Learn all the basics"}
	fmt.Println(courseGoal)

	courseGoal[1] = "Learn all the details"
	courseGoal = append(courseGoal,"Learn all the basics" )
	fmt.Println(courseGoal)

	
	products := []Product {  
	{ 
		"1",
		"first title",
		12.99,
	}, 
	{ 
		"2",
		"second title",
		12.00,
	},
	}
	fmt.Println(products)

	fmt.Println(products)

	newProduct := Product { 
		"3",
		"third title",
		 3,
	}

	products = append(products, newProduct)
	fmt.Println(products)


	fruits := []string { "apple","mango","orange"}
	vegetables := []string { "carrot","cabbage","potato"}

	foods :=[]string {}
	foods = append(foods, fruits...)
	foods = append(foods, vegetables...)
	fmt.Println(foods)




}