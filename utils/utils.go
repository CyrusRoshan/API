package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

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

func KeyStore() func(string) string {
	keyfilePath, _ := filepath.Abs("./keys/keys.yaml")
	if _, err := os.Stat(keyfilePath); err == nil {
		keyHolder := ReadKeys(keyfilePath)
		return func(key string) string {
			return keyHolder[key]
		}
	} else {
		return os.Getenv
	}
}
