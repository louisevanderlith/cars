package logic

import (
	"github.com/louisevanderlith/mango/control"
)

func GetMenu(path string) *control.Menu {
	return getItems(path)
}

func getItems(path string) *control.Menu {
	result := control.NewMenu(path)

	result.AddItem("/create", "Upload your Car", "fa-arrow-up", nil)

	return result
}
