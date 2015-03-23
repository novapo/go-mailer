package config

import (
	"testing"

	"github.com/novapo/go-mailer/config"
)

func TestFromFile(t *testing.T) {
	c, err := config.FromFile("../config.json.example")

	if err != nil {
		t.Error(err)
	}

	if c.Smtp.Username != "John" {
		t.Errorf("Expect 'John' but got %s", c.Smtp.Username)
	}

	if c.Smtp.Password != "Doe" {
		t.Errorf("Expect 'Doe' but got %s", c.Smtp.Password)
	}

	if c.Smtp.Host != "smtp.example.com" {
		t.Errorf("Expect 'smpt.example.com' but got %s", c.Smtp.Host)
	}
}
