package utils

import "strings"

func ToIPv6(value string) string {

	if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
		value = value[1 : len(value)-1]
	}

	if strings.Contains(value, "::") {

		var tmp = strings.Split(value, "::")
		var prefix = strings.Split(tmp[0], ":")
		var suffix = strings.Split(tmp[1], ":")
		var complete = []string{"0000", "0000", "0000", "0000", "0000", "0000", "0000", "0000"}

		for p := 0; p < len(prefix); p++ {

			var chunk = prefix[p]
			var c = p

			if len(chunk) == 1 {
				complete[c] = "000" + chunk
			} else if len(chunk) == 2 {
				complete[c] = "00" + chunk
			} else if len(chunk) == 3 {
				complete[c] = "0" + chunk
			} else if len(chunk) == 4 {
				complete[c] = chunk
			}

		}

		for s := 0; s < len(suffix); s++ {

			var c = 8 - len(suffix) + s
			var chunk = suffix[s]

			if len(chunk) == 1 {
				complete[c] = "000" + chunk
			} else if len(chunk) == 2 {
				complete[c] = "00" + chunk
			} else if len(chunk) == 3 {
				complete[c] = "0" + chunk
			} else if len(chunk) == 4 {
				complete[c] = chunk
			}

		}

		return "[" + strings.Join(complete, ":") + "]"

	} else if strings.Contains(value, ":") {

		var chunks = strings.Split(value, ":")
		var complete = []string{"0000", "0000", "0000", "0000", "0000", "0000", "0000", "0000"}

		for c := 0; c < len(chunks); c++ {

			var chunk = chunks[c]

			if len(chunk) == 1 {
				complete[c] = "000" + chunk
			} else if len(chunk) == 2 {
				complete[c] = "00" + chunk
			} else if len(chunk) == 3 {
				complete[c] = "0" + chunk
			} else if len(chunk) == 4 {
				complete[c] = chunk
			}

		}

		return "[" + strings.Join(complete, ":") + "]"

	}

	return ""

}
