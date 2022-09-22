package address

import "testing"

type testScenario struct {
	addressIn      string
	expectedReturn string
}

func TestAddressType(t *testing.T) {

	testScenarios := []testScenario{
		{"avenue ABC", "Avenue"},
		{"Avenue ABC", "Avenue"},
		{"AVENUE ABC", "Avenue"},
		{"Highway ABC", "Highway"},
		{"highway ABC", "Highway"},
		{"HIGHWAY ABC", "Highway"},
		{"road ABC", "Road"},
		{"Road ABC", "Road"},
		{"ROAD ABC", "Road"},
		{"street ABC", "Street"},
		{"Street ABC", "Street"},
		{"STREET ABC", "Street"},
		{"Square ABC", "Invalid Type"},
		{"", "Invalid Type"},
	}

	for _, scenario := range testScenarios {
		addressReceived := AddressType(scenario.addressIn)

		if addressReceived != scenario.expectedReturn {
			t.Errorf("Received type is %s is different of %s",
				addressReceived,
				scenario.expectedReturn)
		}
	}
}
