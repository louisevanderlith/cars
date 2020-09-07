package resources

import "github.com/louisevanderlith/husk/hsk"

func (src *Source) FetchVehicleInfo(k hsk.Key) (interface{}, error) {
	res, err := src.get("vehicle", "info", k.String())

	if err != nil {
		return nil, err
	}

	return res, nil
}
