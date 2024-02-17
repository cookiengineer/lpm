package utils

import "strings"

func ToIPv6Scope(value string) string {

	value = ToIPv6(value)

	if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
		value = value[1 : len(value)-1]
	}

	var scope string = "public"

	private_ipv6s := []string{

		// RFC3513
		"0000:0000:0000:0000:0000:0000:0000:0000",
		"0000:0000:0000:0000:0000:0000:0000:0001",
		"fe80:0000:0000:0000",
	}

	for p := 0; p < len(private_ipv6s); p++ {

		if strings.HasPrefix(value, private_ipv6s[p]) {
			scope = "private"
			break
		}

	}

	return scope

}
