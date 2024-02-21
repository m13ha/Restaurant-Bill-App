package menu

import (
	"bufio"
	"fmt"
	"os"
	"restaurant_bill_app/internal/getter"
	"strconv"
	"strings"
	"unicode"
)

var Menu = map[string]float64{
	"beef":    450,
	"moi-moi": 250,
	"rice":    300,
	"beans":   300,
	"fish":    500,
}

func MenuManager() {
	reader := bufio.NewReader(os.Stdin)

	response, _ := getter.GetInput("SALES MENU ACTIONS: \n SM - SHOW MENU,  AT - ADD MENU-ITEM, EM - EXIT MENU:", reader)

	response = strings.ToLower(response)

	switch response {
	case "sm":
		ShowMenu()
		MenuManager()
	case "at":
		UpdateMenu()
		MenuManager()
	case "em":
		//do nothing
	default:
		fmt.Print("please enter a valid response \n")
		MenuManager()
	}
}

func ShowMenu() {
	details := "Todays Menu: \n"
	details += strings.ToTitleSpecial(unicode.SpecialCase{}, fmt.Sprintf("%-25v %v \n", "Item", "Price"))

	for k, v := range Menu {
		details += fmt.Sprintf("%-25v....%v \n", k+":", v)
	}

	fmt.Print(details)
}

func UpdateMenu() {
	reader := bufio.NewReader((os.Stdin))
	product, _ := getter.GetInput("Product Name:", reader)
	price, _ := getter.GetInput("Product Price:", reader)

	val, err := strconv.ParseFloat(price, 64)
	if err == nil {
		Menu[product] = val
	} else {
		fmt.Println("Please input a proper value")
		UpdateMenu()
	}
}
