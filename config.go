package main

import (
	"encoding/json"
	"io/ioutil"
)

type BattlestationsConfig struct {
	Addr string
}

func ParseConfig(path string) (error, BattlestationsConfig) {

	var config BattlestationsConfig

	configBytes, err := ioutil.ReadFile(path)

	if err == nil {
		err = json.Unmarshal(configBytes, &config)
	}

	return err, config

}
