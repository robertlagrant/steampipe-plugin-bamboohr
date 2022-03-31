package bamboohr

import (
	"testing"
)

func TestPreSignPayRate(t *testing.T) {
	salary, currency, _ := parsePayRate("$40000.00")
	if salary != "40000.00" {
		t.Log("salary should be 40000.00")
		t.Fail()
	}

	if currency != "$" {
		t.Log("currency should be $")
		t.Fail()
	}
}

func TestPostSignPayRate(t *testing.T) {
	salary, currency, _ := parsePayRate("80001.23 GBP")
	if salary != "80001.23" {
		t.Log("salary should be 80001.23")
		t.Fail()
	}

	if currency != "GBP" {
		t.Log("currency should be GBP")
		t.Fail()
	}
}

func TestNoDecimalPayRate(t *testing.T) {
	salary, currency, _ := parsePayRate("80001")

	if salary != "80001" {
		t.Log("salary should be 80001")
		t.Fail()
	}

	if currency != "" {
		t.Log("currency should be empty")
		t.Fail()
	}
}
