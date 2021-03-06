package model

import (
	"errors"
	"time"
)

type Cost struct {
	MaxSpend    float32
	Ingredients []Ingredient
}

type Ingredient struct {
	Name  string
	Price float32
}

// var total
// var total map[string]*Cost
var recipeTotal []*Cost

var d = time.Now()
var day, _, _ = d.Date()

// budget per shopping trip

// fmt.Println(d.Date())
// fmt.Printf("YEAR = %v\n", year)
// fmt.Printf("month = %v\n", month)
// fmt.Printf("day = %v\n", day)
var total map[time.Month]*Cost

func InitalizeCostCalculator() {
	total = make(map[time.Month]*Cost)
}

func init() {
	InitalizeCostCalculator()
}

func (c Cost) CurrentTotal() float32 {
	var sum float32
	for _, ingredient := range c.Ingredients {
		sum += ingredient.Price
	}
	return sum
}

var fitsBudgetErr = errors.New("This costs too much money")

var alreadyAddedErr = errors.New("This was already added")

func (c *Cost) AddIngredient(name string, price float32) error {
	// fmt.Println(c)
	if c.CurrentTotal()+price > c.MaxSpend {
		return fitsBudgetErr
	}

	newItem := Ingredient{Name: name, Price: price}
	c.Ingredients = append(c.Ingredients, newItem)

	return nil
}

func (c *Cost) RemoveIngredient(name string) {
	for i := range c.Ingredients {
		if c.Ingredients[i].Name == name {
			c.Ingredients = append(c.Ingredients[:i], c.Ingredients...)
			break
		}
	}
}

func CreateList(maxSpend float32) (*Cost, error) {
	var newTotal *Cost
	return newTotal, nil
}

func GetList() *Cost {
	return nil
}
