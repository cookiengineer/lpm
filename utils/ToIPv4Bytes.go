package utils

import "math"
import "strconv"
import "strings"

func ToIPv4Bytes(value string, prefix uint8) []byte {

	value = ToIPv4(value)

	if value != "" {

		var bytes []byte = make([]byte, 0)

		tmp := strings.Split(value, ".")

		for t := 0; t < len(tmp); t++ {

			number, err := strconv.ParseUint(tmp[t], 10, 8)

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
