package handles

import "github.com/louisevanderlith/droxolite/menu"

func FullMenu() *menu.Menu {
	m := menu.NewMenu()

	m.AddItem(menu.NewItem("a", "#home", "Home", nil))
	m.AddItem(menu.NewItem("e", "/create", "Sell your Vehicle", nil))

	return m
}
