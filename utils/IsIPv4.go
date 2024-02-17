package utils

import "strconv"
import "strings"

func IsIPv4(value string) bool {

	if strings.Contains(value, ".") {

		chunks := strings.Split(value, ".")

		if len(chunks) == 4 {

			var valid bool = true

			for c := 0; c < len(chunks); c++ {

				_, err := strconv.ParseUint(chunks[c], 10, 8)

				if err != nil {
					valid = false
					break
				}

			}

			return valid

		}

	}

	return false

}
