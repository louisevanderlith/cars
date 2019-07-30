package logic

import "github.com/louisevanderlith/droxolite/bodies"

func GetMenu() *bodies.Menu {
	return getItems()
}

func getItems() *bodies.Menu {
	result := bodies.NewMenu()

	result.AddItem("/create", "Upload your Car", "fa-arrow-up", nil)

	return result
}
