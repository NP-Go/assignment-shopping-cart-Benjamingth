package main

import (
	"fmt"
	"strconv"
	"time"
)

type product struct {
	Category  int
	Quantity  int
	Unit_Cost float64
}

var items map[string]product

func init() {
	items = make(map[string]product)
	items["Cups"] = product{0, 5, 3}
	items["Cake"] = product{1, 3, 1}
	items["Sprite"] = product{2, 5, 2}
	items["Fork"] = product{0, 4, 3}
	items["Bread"] = product{1, 2, 2}
	items["Plates"] = product{0, 4, 3}
	items["Coke"] = product{2, 5, 2}
}

var mainMenu = []string{
	"1. View Entire Shopping List.",
	"2. Generate Shopping List Report",
	"3. Add Items",
	"4. Modify Items",
	"5. Delete Item",
	"6. Print Current Data",
	"7. Add New Category Name",
	"0. Exit Program",
}

var ReportMenu = []string{
	"1. Total Cost of each catagory",
	"2. List of item by catagory",
	"3. Main Menu",
}

func genReportMenu() {
	fmt.Println("\nGenerate Report")
	for i := range ReportMenu {
		fmt.Println(ReportMenu[i])
	}
}

func displayMainMenu() {
	fmt.Println("Shopping List Application")
	fmt.Println("=========================")
	for a := range mainMenu {
		fmt.Println(mainMenu[a])
	}
	fmt.Printf("Select your choice: ")
}

func makeMenuRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func findCategory(C []string, c string) (int, bool) {
	for i := range C {
		if C[i] == c {
			return i, true
		}
	}
	return -1, false
}

func validSelection(m []int, c int) bool {
	for i := range m {
		if m[i] == c {
			return true
		}
	}
	return false
}

func main() {

	var input1 int
	option1 := makeMenuRange(1, len(mainMenu))
	categorys := []string{"Household", "Food", "Drinks"}

	displayMainMenu()
	fmt.Scanln(&input1)

	for !validSelection(option1, input1) {
		fmt.Println("Please enter valid choice of number.")
		time.Sleep(3 * time.Second)
		displayMainMenu()
		fmt.Scanln(&input1)
	}

	for validSelection(option1, input1) {
		switch input1 {
		case 1:
			fmt.Println("\nShopping List Contents:")
			for index, value := range items {
				fmt.Println("Catagory: " + categorys[value.Category] + " - Item: " + index + " Quantity: " + strconv.Itoa(value.Quantity) + " Unit Cost: " + fmt.Sprintf("%v", value.Unit_Cost))
			}
			fmt.Println("")
			input1 = -1
			fmt.Scanln()

		case 2:
			var input2 int
			option2 := makeMenuRange(1, len(ReportMenu))
			genReportMenu()
			fmt.Scanln(&input2)

			for !validSelection(option2, input2) {
				fmt.Println("Please enter valid choice of number.")
				time.Sleep(3 * time.Second)
				genReportMenu()
				fmt.Scanln(&input2)
			}
			for validSelection(option2, input2) {
				if input2 == 1 {
					totalCatCost := make([]float64, len(categorys))
					for _, value := range items {
						for i := range categorys {
							if value.Category == i {
								totalCatCost[i] = totalCatCost[i] + value.Unit_Cost*float64(value.Quantity)
								break
							}
						}
					}
					fmt.Println("Total Cost By Category")
					for j := range categorys {
						fmt.Println(categorys[j] + " cost : " + fmt.Sprintf("%g", totalCatCost[j]))
					}
					fmt.Println("")
					input1 = -1
					fmt.Println("Press enter to go back")
					fmt.Scanln()
					genReportMenu()
					fmt.Scanln(&input2)
				} else if input2 == 2 {
					fmt.Println("List by Category")
					for key, value := range items {
						for k := range categorys {
							if value.Category == k {
								fmt.Println(categorys[value.Category] + ": " + key + " - Item: " + strconv.Itoa(value.Quantity) + " Unit Cost: " + fmt.Sprintf("%g", value.Unit_Cost))
								continue
							}
							fmt.Println("")
							input1 = -1
							fmt.Println("Press enter to go back")
							fmt.Scanln()
							genReportMenu()
							fmt.Scanln(&input2)
						}
					}
				} else if input2 == 3 {
					displayMainMenu()
					fmt.Scanln(&input1)
				}
			}
		case 3:
			var input30, input31 string
			var input32 int
			var input33 float64
			fmt.Println("Enter name of item you wish to add?")
			fmt.Scanln(&input30)
			fmt.Println("What category does it belong to?")
			fmt.Scanln(&input31)
			fmt.Println("How many units would you like to add?")
			fmt.Scanln(&input32)
			fmt.Println("How much does it cost per unit?")
			fmt.Scanln(&input33)

			a, foundCategory := findCategory(categorys, input31)
			if !foundCategory && a == -1 {
				fmt.Println("Category not found. Please create new category.")
			}
		case 4:
			var input40, input41 string
			var input42 int
			var input43 float64
			fmt.Println("Modify Items.")
			fmt.Println("Which item would you like to modify?")
			fmt.Scanln(&input40)
			for key, _ := range items {
				if input40 == key {
					fmt.Println("Current item name is " + input40 + " - Category is " + categorys[value.Category] + " Quantity is " + strconv.Itoa(value.Quantity) + " Unit Cost " + fmt.Sprintf("%g", value.Unit_Cost))
					fmt.Scanln(input41, input42, input43)
				}
			}
		}
	}
}
