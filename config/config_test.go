package config

import (
	"reflect"
	"testing"

	"github.com/novapo/go-mailer/config"
)

func TestFromFile(t *testing.T) {
	c, err := config.FromFile("../config.json.example")

	if err != nil {
		t.Error(err)
	}

	assertEqual(t, "John", c.Smtp.Username)
	assertEqual(t, "Doe", c.Smtp.Password)
	assertEqual(t, "smtp.example.com", c.Smtp.Host)

	assertEqual(t, "john.doe@example.com", c.Rcpt)

	assertEqual(t, "c_name", c.FormData.Name)
	assertEqual(t, "c_email", c.FormData.Email)
	assertEqual(t, "c_message", c.FormData.Message)

	assertEqual(t, 8080, c.Port)
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	actualType := reflect.TypeOf(actual)
	expectedValue := reflect.ValueOf(expected)
	if expectedValue.Type().ConvertibleTo(actualType) {
		if reflect.DeepEqual(actual, expectedValue.Convert(actualType).Interface()) {
			return
		}
	}
	t.Errorf("Expect '%s' but got '%s'", expected, actual)
}
