package utils

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v1"
)

type Keys map[string]string

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadKeys(path string) Keys {
	fileData, err := ioutil.ReadFile(path)
	PanicIf(err)

	keyHolder := Keys{}
	err = yaml.Unmarshal([]byte(fileData), &keyHolder)
	PanicIf(err)

	return keyHolder
}

func MustMarshal(data interface{}) string {
	out, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return string(out)
}
