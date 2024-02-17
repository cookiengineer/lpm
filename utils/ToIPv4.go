package utils

import "strconv"
import "strings"

func ToIPv4(value string) string {

	if strings.Contains(value, ".") {

		var tmp1 []string = strings.Split(value, ".")
		var tmp2 []uint8
		var tmp3 []string

		for t := 0; t < len(tmp1); t++ {

			num, err := strconv.ParseUint(tmp1[t], 10, 8)

			if err == nil {

				if num >= 0 && num <= 255 {
					tmp2 = append(tmp2, uint8(num))
				}

			}

		}

		for t := 0; t < len(tmp2); t++ {

			var str = strconv.FormatUint(uint64(tmp2[t]), 10)

			if str != "" {
				tmp3 = append(tmp3, str)
			}

		}

		if len(tmp3) == 4 {
			return strings.Join(tmp3, ".")
		}

	}

	return ""

}
