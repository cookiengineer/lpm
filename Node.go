package lpm

import "github.com/cookiengineer/lpm/utils"
import "encoding/hex"
import "strconv"
import "strings"

type Node struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Scope   string `json:"scope"`
	Type    string `json:"type"`
	Prefix  uint8  `json:"prefix"`
	Children []*Node `json:"children"`
}

func NewNode(address string, prefix uint8) Node {

	var node Node

	node.SetAddress(address)
	node.SetPrefix(prefix)

	// Necessary for Trie
	node.Children = make([]*Node, 0)

	return node

}

func ToNode(value string) Node {

	var node Node
	var address string
	var prefix uint8

	if strings.HasPrefix(value, "[") && strings.Contains(value, "]/") {

		tmp1 := value[0 : strings.Index(value, "]/")+1]
		tmp2, err := strconv.ParseUint(strings.Split(value, "]/")[1], 10, 8)

		if err == nil {
			address = strings.ToLower(tmp1)
			prefix = uint8(tmp2)
		}

	} else if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {

		address = strings.ToLower(value)
		prefix = uint8(128)

	} else if strings.Contains(value, ":") && strings.Contains(value, "/") {

		tmp1 := strings.Split(value, "/")[0]
		tmp2, err := strconv.ParseUint(strings.Split(value, "/")[1], 10, 8)

		if err == nil {
			address = strings.ToLower(tmp1)
			prefix = uint8(tmp2)
		}

	} else if strings.Contains(value, ":") {

		address = "[" + strings.ToLower(value) + "]"
		prefix = uint8(128)

	} else if strings.Contains(value, ".") && strings.Contains(value, "/") {

		tmp1 := strings.Split(value, "/")[0]
		tmp2, err := strconv.ParseUint(strings.Split(value, "/")[1], 10, 8)

		if err == nil {
			address = tmp1
			prefix = uint8(tmp2)
		}

	} else if strings.Contains(value, ".") {

		address = value
		prefix = uint8(32)

	}

	if address != "" && prefix != 0 {
		node = NewNode(address, prefix)
	}

	return node

}

func (node *Node) Contains(value string) bool {

	var result bool = false

	if utils.IsIPv6(value) {

		value = utils.ToIPv6(value)
		tmp := NewNode(value[1:len(value)-1], 128)
		result = node.ContainsNode(tmp)

	} else if utils.IsIPv4(value) {

		value = utils.ToIPv4(value)
		tmp := NewNode(value, 32)
		result = node.ContainsNode(tmp)

	}

	return result

}

func (node *Node) ContainsNode(value Node) bool {

	var result bool = false

	if node.Type == "ipv4" && value.Type == "ipv4" {

		if value.Prefix > node.Prefix {

			bytes_subnet := utils.ToIPv4Bytes(node.Address, node.Prefix)
			bytes_value := utils.ToIPv4Bytes(value.Address, node.Prefix)

			if len(bytes_subnet) > 0 && len(bytes_value) > 0 {

				result = true

				for b := 0; b < len(bytes_subnet); b++ {

					if bytes_subnet[b] != bytes_value[b] {
						result = false
						break
					}

				}

			}

		}

	} else if node.Type == "ipv6" && value.Type == "ipv6" {

		if value.Prefix > node.Prefix {

			bytes_subnet := utils.ToIPv6Bytes(node.Address, node.Prefix)
			bytes_value := utils.ToIPv6Bytes(value.Address, node.Prefix)

			if len(bytes_subnet) > 0 && len(bytes_value) > 0 {

				result = true

				for b := 0; b < len(bytes_subnet); b++ {

					if bytes_subnet[b] != bytes_value[b] {
						result = false
						break
					}

				}

			}

		}

	}

	return result

}

func (node *Node) IsValid() bool {

	var result bool = false

	if node.Type == "ipv4" {

		if node.Address != "" && node.Address != "0.0.0.0" {

			if node.Prefix != 0 {
				result = true
			}

		}

	} else if node.Type == "ipv6" {

		if node.Address != "" && node.Address != "[0000:0000:0000:0000:0000:0000:0000:0000]" {

			if node.Prefix != 0 {
				result = true
			}

		}

	}

	return result

}

func (node *Node) SetName(value string) {

	value = strings.TrimSpace(value)

	if value != "" {
		node.Name = value
	}

}

func (node *Node) SetAddress(value string) {

	if utils.IsIPv6(value) {

		node.Address = utils.ToIPv6(value)
		node.Scope = utils.ToIPv6Scope(value)
		node.Type = "ipv6"

	} else if utils.IsIPv4(value) {

		node.Address = utils.ToIPv4(value)
		node.Scope = utils.ToIPv4Scope(value)
		node.Type = "ipv4"

	}

}

func (node *Node) SetPrefix(value uint8) {

	if node.Type == "ipv4" {

		if value >= 0 && value <= 32 {
			node.Prefix = value
		}

	} else if node.Type == "ipv6" {

		if value >= 0 && value <= 128 {
			node.Prefix = value
		}

	}

}

func (node *Node) SetScope(value string) {

	if value == "private" || value == "public" {
		node.Scope = value
	}

}

func (node *Node) Addresses() uint {

	var result uint

	if node.Prefix != 0 {

		if node.Type == "ipv4" {
			result = 1 << (32 - node.Prefix)
		} else if node.Type == "ipv6" {
			result = 1 << (128 - node.Prefix)
		}

	}

	return result

}

func (node *Node) Hash() string {

	var hash string

	if node.Type == "ipv4" {

		bytes := utils.ToIPv4Bytes(node.Address, node.Prefix)
		hash = hex.EncodeToString(bytes)

	} else if node.Type == "ipv6" {

		bytes := utils.ToIPv6Bytes(node.Address, node.Prefix)
		hash = hex.EncodeToString(bytes)

	}

	return hash

}

func (node *Node) String() string {

	var result string

	if node.Address != "" {
		result = node.Address + "/" + strconv.FormatUint(uint64(node.Prefix), 10)
	}

	return result

}
