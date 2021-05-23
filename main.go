package main

import (
	"fmt"

	"main/model"
)

func main() {
	var cost model.Cost
	cl, _ := model.CreateList(10.00)
	fmt.Println(cl)
	newCost := cost.AddIngredient("flour", 0)
	newCost = cost.AddIngredient("sugar", 0)
	fmt.Println("NewCost", newCost)
	// tNow := time.Now()
	// //time.Time to Unix Timestamp
	// // tUnix := tNow.Unix()
	// // fmt.Printf("timeUnix %d\n", tUnix)
	// // timeT := time.Unix(tUnix, 0)
	// // fmt.Printf("time.Time: %s\n", timeT)

}
