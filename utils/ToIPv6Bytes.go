package utils

import "math"
import "strconv"
import "strings"

func ToIPv6Bytes(value string, prefix uint8) []byte {

	value = ToIPv6(value)

	if value != "" {

		bytes := make([]uint8, 0)

		tmp := strings.Join(strings.Split(value[1:len(value)-1], ":"), "")

		for t := 0; t < len(tmp); t += 2 {

			hex := string(tmp[t : t+2])
			number, err := strconv.ParseUint(hex, 16, 8)

			if err == nil {
				bytes = append(bytes, uint8(number))
			} else {
				bytes = append(bytes, uint8(0))
			}

		}

		if prefix%8 == 0 {

			bytes_length := int(prefix / 8)
			bytes_mask := make([]byte, bytes_length)
			copy(bytes_mask, bytes)

			for len(bytes_mask) < len(bytes) {
				bytes_mask = append(bytes_mask, uint8(0))
			}

			return bytes_mask

		} else {

			bytes_length := int(math.Floor(float64(prefix / 8)))
			bytes_mask := make([]byte, bytes_length)
			copy(bytes_mask, bytes)

			var last_byte byte = bytes[bytes_length]
			last_byte_shift := int(int(prefix) - bytes_length*8)
			bytes_mask = append(bytes_mask, last_byte>>(8-last_byte_shift))

			for len(bytes_mask) < len(bytes) {
				bytes_mask = append(bytes_mask, uint8(0))
			}

			return bytes_mask

		}

	}

	return []byte{}

}
