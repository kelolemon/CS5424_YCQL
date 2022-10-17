package common

import "github.com/mitchellh/mapstructure"

func ToCqlMap(value interface{}) (map[string]interface{}, error) {
	structMap := map[string]interface{}{}

	decoderConfig := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &structMap,
		TagName:  "cql",
	}

	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		return nil, err
	}

	if err := decoder.Decode(value); err != nil {
		return nil, err
	}

	return structMap, nil
}

func ToCqlStruct(structMap map[string]interface{}, result interface{}) error {
	decoderConfig := &mapstructure.DecoderConfig{
		Result:  &result,
		TagName: "cql",
	}

	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		return err
	}
	
	err = decoder.Decode(structMap)
	return err
}
