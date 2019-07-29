package logic

import "github.com/louisevanderlith/droxolite/bodies"

func GetMenu(path string) *bodies.Menu {
	return getItems(path)
}

func getItems(path string) *bodies.Menu {
	result := bodies.NewMenu(path)

	result.AddItem("/create", "Upload your Car", "fa-arrow-up", nil)

	return result
}
