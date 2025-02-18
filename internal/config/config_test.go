package config_test

import (
	"testing"

	"github.com/mohammadne/zanbil/internal/config"
)

func TestLoadDefaults(t *testing.T) {
	_, err := config.LoadDefaults(true)
	if err != nil {
		t.Error(err)
	}
}
