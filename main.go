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
	categorys := []string{"Household", "Food", "Drinks"}

	displayMainMenu()
	fmt.Scanln(&input1)

	for !validSelection(option1, input1) {
		fmt.Println("Please enter valid choice of number.")
		time.Sleep(3 * time.Second)
		displayMainMenu()
		fmt.Scanln(&input1)
		break
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
			} else if foundCategory && index == 0 {
				for key, value := range items {
					if key == itemName {
						fmt.Printf("Item %v, would you like to change the Quantity to %v and Unit Cost to %v?", value, value.Quantity, value.Unit_Cost)

					}
				}
			}
		case 4:
			var modItem, modCat string
			var modQuantity int
			var modPrice float64
			fmt.Println("\nModify Items.")
			fmt.Println("Which item would you like to modify?")
			fmt.Scanln(&modItem)
			for key, value := range items {
				if modItem != key {
					fmt.Println("Sorry item not found. Please add new item before modifying.")
					break
				} else {
					if modItem == key {
						fmt.Printf("Current item name is %v - Category is %v - Quantity is %v - Unit Cost is %g", modItem, categorys[value.Category], strconv.Itoa(value.Quantity), value.Unit_Cost)
						fmt.Scanln(&modCat)
						fmt.Println("Enter new Quantity. Press enter for no change.")
						fmt.Scanln(&modQuantity)
						fmt.Println("Enter new Unit Cost. Press enter for no change.")
						fmt.Scanln(&modPrice)
						//Havent add the comments afer modifying
					}
				}
			}
		case 5:
			var deleteItem string
			fmt.Println("\nDelete Item")
			fmt.Printf("Enter item to delete: ")
			fmt.Scanln(&deleteItem)
			for key, _ := range items {
				if key != deleteItem {
					fmt.Println("\nItem not found. Nothing to delete.")
				} else {
					delete(items, key)
					fmt.Printf("Deleted %s", deleteItem)
					break
					//not working try again
				}
				input1 = -1
				fmt.Println("\nPress enter to exit program")
				fmt.Scanln()
			}
		case 6:
			fmt.Println("\nPrint Current Data")
			if len(items) != 0 {
				for key, value := range items {
					fmt.Println(key + "{" + strconv.Itoa(value.Category) + " " + strconv.Itoa(value.Quantity) + " " + fmt.Sprintf("%g", value.Unit_Cost) + "}")
				}
			} else {
				fmt.Println("No data found.")
			}
			input1 = -1
			fmt.Println("\nPress enter to exit program")
			fmt.Scanln()

		case 7:
			var newCat string
			fmt.Println("Add New Category Name")
			fmt.Println("What is the New Category Name to add?")
			fmt.Scanln(&newCat)
			_, newName := findCategory(categorys, newCat)
			if newName == false {
				categorys = append(categorys, newCat)
				fmt.Printf("New Category: %v "+"added at index "+strconv.Itoa(len(categorys)), newCat)
			} else if newName == true {
				fmt.Printf("Category: %v already found at index "+strconv.Itoa(len(categorys)), newCat)
			} else {
				fmt.Println("No input found")
				break
			}
			input1 = -1
		case 8:
			fmt.Println("Exiting Program. Hope to see you again.")
			return
		}
		displayMainMenu()
		fmt.Scanln(&input1)
	}
}
