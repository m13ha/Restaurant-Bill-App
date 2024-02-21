package main

import (
	"bufio"
	"fmt"
	"os"
	"restaurant_bill_app/internal/biller"
	"restaurant_bill_app/internal/getter"
	"restaurant_bill_app/internal/menu"
	"strings"
)

func main() {
	fmt.Println("Welcome")
	taskManager()
}

func taskManager() {
	reader := bufio.NewReader(os.Stdin)

	response, _ := getter.GetInput("TASK MENU ACTIONS: \n B - NEW BILL, S - SALES MENU, E - END PROGRAM:", reader)
	response = strings.ToLower(response)

	switch response {
	case "b":
		biller.BillManager(biller.NewBill())
		taskManager()
	case "s":
		menu.MenuManager()
		taskManager()
	case "a":
		fmt.Print("done")
	case "e":
		//do nothing
	default:
		fmt.Print("please enter a valid response \n")
		taskManager()
	}
}
