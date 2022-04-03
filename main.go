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

var categorys = []string{"Household", "Food", "Drinks"}

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
	"8. Exit Program",
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

	displayMainMenu()
	fmt.Scanln(&input1)

	for !validSelection(option1, input1) {
		fmt.Println("Please enter valid choice of number at next screen.")
		time.Sleep(2 * time.Second)
		displayMainMenu()
		fmt.Scanln(&input1)
	}

	for validSelection(option1, input1) {
		switch input1 {
		case 1:
			fmt.Println("\nShopping List Contents:")
			for index, value := range items {
				fmt.Printf("Category: %v - Item: %v, Quantity: %v, Unit Cost %g\n", categorys[value.Category], index, value.Quantity, value.Unit_Cost)
			}
			fmt.Println("")
			input1 = -1
			fmt.Scanln()
			//done

		case 2:
			var input2 int
			option2 := makeMenuRange(1, len(ReportMenu))
			genReportMenu()
			fmt.Scanln(&input2)

			for !validSelection(option2, input2) {
				fmt.Println("Please enter valid choice of number at next screen.")
				time.Sleep(2 * time.Second)
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
						fmt.Printf("%v cost: %g\n", categorys[j], totalCatCost[j])
					}
					fmt.Println("")
					input1 = -1
					fmt.Println("Press enter to go back")
					fmt.Scanln()
					genReportMenu()
					fmt.Scanln(&input2)
				} else if input2 == 2 {
					for index, valueOut := range categorys {
						for key, valueIn := range items {
							if index == int(valueIn.Category) {
								fmt.Printf("Category: %v - Item: %v  Quantity: %v  Unit Cost: %g\n", valueOut, key, valueIn.Quantity, valueIn.Unit_Cost)
							}
						}
					}
					break
				} else if input2 == 3 {
					displayMainMenu()
					fmt.Scanln(&input1)
				}
			}
			input1 = -1
			fmt.Println("\nPress enter to return to main menu.")
			fmt.Scanln()

		case 3:
			var itemName, itemCat string
			var itemQuantity int
			var itemCost float64
			fmt.Println("Enter name of item you wish to add?")
			fmt.Scanln(&itemName)
			fmt.Println("What category does it belong to?")
			fmt.Scanln(&itemCat)
			fmt.Println("How many units would you like to add?")
			fmt.Scanln(&itemQuantity)
			fmt.Println("How much does it cost per unit?")
			fmt.Scanln(&itemCost)

			index, foundCategory := findCategory(categorys, itemCat)
			if !foundCategory && index == -1 {
				fmt.Println("Category not found. Please create new category.")
			} else if foundCategory && index == 1 {
				for key, value := range items {
					if key == itemName {
						fmt.Printf("Item %v, would you like to change the Quantity to %v and Unit Cost to %v?", value, value.Quantity, value.Unit_Cost)
						//not ready
					}
				}
			}

		case 4:
			var modItem, newName, modCat string
			var modQuantity int
			var modPrice float64
			fmt.Println("\nModify Items.")
			fmt.Println("Which item would you like to modify?")
			fmt.Scanln(&modItem)
			for key, value := range items {
				if key == modItem {
					fmt.Printf("Current item name is %v - Category is %v - Quantity is %v - Unit Cost is %g\n", modItem, categorys[value.Category], strconv.Itoa(value.Quantity), value.Unit_Cost)
					fmt.Println("\nEnter new name. Press enter if there are no changes.")
					fmt.Scanln(&newName)
					fmt.Println("Enter new Category. Press enter if there are no changes")
					fmt.Scanln(&modCat)
					fmt.Println("Enter new Quantity. Press enter if there are no changes.")
					fmt.Scanln(&modQuantity)
					fmt.Println("Enter new Unit Cost. Press enter if there are no changes.")
					fmt.Scanln(&modPrice)

					if newName == "" {
						fmt.Println("No changes made to item Name.")
					}

					if modCat == "" {
						fmt.Println("No changes made to Item Category")
					}

					if modQuantity == 0 {
						fmt.Println("No changes made to item Quantity")
					}

					if modPrice == 0.0 {
						fmt.Println("No changes made to item Unit Cost")
					}
				}
			}
			input1 = -1
			fmt.Println("\nPress enter to return to main menu.")
			fmt.Scanln()
			//done
		case 5:
			var deleteItem string
			fmt.Println("\nDelete Item")
			fmt.Printf("Enter item to delete: ")
			fmt.Scanln(&deleteItem)
			//Run through range using only the names of products
			if _, ok := items[deleteItem]; ok {
				delete(items, deleteItem)
				fmt.Printf("\nDeleted %v", deleteItem)
			} else {
				fmt.Println("\nItem not found. Nothing to delete.")
				//done
			}
			input1 = -1
			fmt.Println("\nPress enter to return to main menu.")
			fmt.Scanln()

		case 6:
			fmt.Println("\nPrint Current Data")
			if len(items) != 0 {
				for key, value := range items {
					fmt.Printf("%v - %v\n", key, value)
				}
			} else {
				fmt.Println("No data found.")
			}
			input1 = -1
			fmt.Println("\nPress enter to return to main menu")
			fmt.Scanln()
			//done

		case 7:
			var newCat string
			fmt.Println("Add New Category Name")
			fmt.Println("What is the New Category Name to add?")
			fmt.Scanln(&newCat)
			key, newName := findCategory(categorys, newCat)
			if newName == false {
				categorys = append(categorys, newCat)
				fmt.Printf("New Category: %v added at index %v", newCat, strconv.Itoa(len(categorys)))
			} else if newName == true {
				fmt.Printf("Category: %v already found at index %v", newCat, strconv.Itoa(key))
			}
			input1 = -1
			fmt.Println("\nPress enter to return to main menu")
			fmt.Scanln()
		case 8:
			fmt.Println("Exiting Program. Hope to see you again.")
			return
		}
		displayMainMenu()
		fmt.Scanln(&input1)
	}
}
