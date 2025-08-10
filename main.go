package main

import "fmt"

type MenuItem struct {
	key		string
	item 	string
}

type Menu struct {
	Items	[]MenuItem
	idx		int
}

func (menu Menu) Show() {
	fmt.Print("\033[?25l")
	for {
		for _, item := range menu.Items {
			fmt.Println(item.item)
		}
		fmt.Print("\033[4A")
	}
	fmt.Print("\033[?25h")
}

func (menu *Menu) Add(key, item string) {
	newItem := MenuItem {
		key: key,
		item: item,
	}
	menu.Items = append(menu.Items, newItem)
}

func main() {
	menu := Menu{}
	menu.Add("1", "Option 1")
	menu.Add("2", "Option 2")
	menu.Add("3", "Option 3")
	menu.Add("4", "Option 4")

	menu.Show()
}
