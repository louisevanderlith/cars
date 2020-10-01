package resources

func (src *Source) ValidateVIN(vin string) (bool, error) {
	res, err := src.get("vin", "validate", vin)

	if err != nil {
		return false, err
	}

	if res == nil {
		return false, nil
	}

	return res.(bool), nil
}

func (src *Source) LookupVIN(vin string) (interface{}, error) {
	return src.get("vin", "lookup", vin)
}
