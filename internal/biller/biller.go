package biller

import (
	"bufio"
	"fmt"
	"os"
	"restaurant_bill_app/internal/getter"
	"restaurant_bill_app/internal/menu"
	"strconv"
	"strings"
	"unicode"
)

type bill struct {
	name    string
	items   map[string]float64
	cashier string
}

func NewBill() *bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getter.GetInput("Customer Name:", reader)

	b := bill{
		name:    name,
		items:   map[string]float64{},
		cashier: "",
	}

	return &b
}

func BillManager(b *bill) {
	reader := bufio.NewReader(os.Stdin)
	response, _ := getter.GetInput("BILL MENU ACTIONS :- SH - SHOW BILL,  AD - ADD ITEM, SA - SAVE BILL:", reader)
	response = strings.ToLower(response)

	switch response {
	case "sh":
		b.ShowBill()
		BillManager(b)
	case "ad":
		b.AddItem()
		BillManager(b)
	case "sa":
		b.SignBill()
		b.Save()
	default:
		fmt.Print("please enter a valid response \n")
		BillManager(b)
	}
}

func (b *bill) AddItem() {
	menu.ShowMenu()
	reader := bufio.NewReader(os.Stdin)
	response, _ := getter.GetInput("please select a product from the menu above or type (n) - next:", reader)

	if response != "n" {
		pdval, exist := menu.Menu[response]

		if !exist {
			fmt.Println("That is not on the menu")
		} else {
			quantity, _ := getter.GetInput("Quanity:", reader)
			prVal, err := strconv.ParseFloat(quantity, 64)

			if err != nil {
				fmt.Println("Please enter a valid number:")
				b.AddItem()
			}

			b.items[response] = pdval * prVal
			fmt.Print("Item added to bill \n")
		}

	}

}

func (b *bill) SignBill() {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getter.GetInput("Cashier:", reader)
	b.cashier = name
}

func (b *bill) ShowBill() {
	details := "Bill Details: \n"
	total := 0.0

	details += fmt.Sprintf("Customer Name: %v \n", b.name)
	details += strings.ToTitleSpecial(unicode.SpecialCase{}, fmt.Sprintf("%-25v %v \n", "Item", "Price"))

	for k, v := range b.items {
		details += fmt.Sprintf("%-25v....%v \n", k+":", v)
		total += v
	}

	details += fmt.Sprintf("%-25v....%v \n", "Total:", total)
	details += fmt.Sprintf("Sold by %v \n", b.cashier)

	fmt.Print(details)
}

func (b *bill) Format() string {
	details := "Bill Details: \n"
	total := 0.0

	details += fmt.Sprintf("Customer Name: %v \n", b.name)
	details += strings.ToTitleSpecial(unicode.SpecialCase{}, fmt.Sprintf("%-25v %v \n", "Item", "Price"))

	for k, v := range b.items {
		details += fmt.Sprintf("%-25v....%v \n", k+":", v)
		total += v
	}

	details += fmt.Sprintf("%-25v....%v \n", "Total:", total)
	details += fmt.Sprintf("Sold by %v \n", b.cashier)

	return details
}

func (b *bill) Save() {
	data := []byte(b.Format())
	err := os.WriteFile("Receipts/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v's receipt saved to file", b.name)
}
