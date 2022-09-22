package address

import "strings"

//AddressType check if address has a valid type
func AddressType(address string) string {
	validTypes := []string{"avenue", "highway", "road", "street"}

	lowerAddress := strings.ToLower(address)
	firstAddressWord := strings.Split(lowerAddress, " ")[0]

	validTypeAddress := false
	for _, typeAddress := range validTypes {
		if typeAddress == firstAddressWord {
			validTypeAddress = true
			break
		}
	}

	if validTypeAddress {
		return strings.Title(firstAddressWord)
	}

	return "Invalid Type"
}
