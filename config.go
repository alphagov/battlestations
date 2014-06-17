package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/alphagov/battlestations/github"
)

type BattlestationsConfig struct {
	Addr    string `json:"addr"`
	AuthKey string `json:"auth_key"`
	EncKey  string `json:"enc_key"`
	BaseURL string `json:"base_url"`

	Github github.Config `json:"github"`
}

func ParseConfig(path string) (error, BattlestationsConfig) {

	var config BattlestationsConfig

	configBytes, err := ioutil.ReadFile(path)

	if err == nil {
		err = json.Unmarshal(configBytes, &config)
	}

	return err, config

}
