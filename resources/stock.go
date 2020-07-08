package resources

func (src *Source) FetchStockCar(key string) (interface{}, error) {
	return src.get("stock", "car", key)
}

func (src *Source) FetchStockCars(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("stock", "cars", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}

func (src *Source) FetchStockPart(key string) (interface{}, error) {
	return src.get("stock", "part", key)
}

func (src *Source) FetchStockParts(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("stock", "part", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}
