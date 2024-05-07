package util

import (
	"os"
	"testing"
)

func TestValidatePort(t *testing.T) {

	Config = new(ConfigType)
	Config.Port = ""
	validatePort()
	if Config.Port != ":3000" {
		t.Error("no port should get set to default")
	}

	Config.Port = "4000"
	validatePort()
	if Config.Port != ":4000" {
		t.Error("Port without : suffix should have it added")
	}

	os.Setenv("PORT", "5000")
	validatePort()
	if Config.Port != ":5000" {
		t.Error("Port value should be overwritten by env var, and it should be prefixed appropriately")
	}
}
