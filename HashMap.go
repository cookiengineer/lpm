package lpm

import "github.com/cookiengineer/lpm/utils"
import "encoding/hex"
import "fmt"
import "sort"
import "strconv"
import "strings"

func printHashMapNode(node *Node, hash string, prefix uint8, indent string) {
	fmt.Println(indent + "└─" + strconv.FormatUint(uint64(prefix), 10) + "─" +node.Hash() + " (" + node.String() + ")")
}

type HashMap struct {
	Mapv4 map[uint8]map[string]*Node `json:"mapv4"`
	Mapv6 map[uint8]map[string]*Node `json:"mapv6"`
}

func NewHashMap() HashMap {

	var hashmap HashMap

	hashmap.Mapv4 = make(map[uint8]map[string]*Node)
	hashmap.Mapv6 = make(map[uint8]map[string]*Node)

	return hashmap

}

func (hashmap *HashMap) Insert(value string) bool {

	node := ToNode(value)

	return hashmap.InsertNode(node)

}

func (hashmap *HashMap) InsertNode(value Node) bool {

	var result bool = false
	var node = value

	if node.Type == "ipv6" {

		prefix := node.Prefix
		hash := node.Hash()

		_, ok1 := hashmap.Mapv6[prefix]

		if ok1 == false {
			hashmap.Mapv6[prefix] = make(map[string]*Node)
		}

		_, ok2 := hashmap.Mapv6[prefix][hash]

		if ok2 == true {
			result = false
		} else {
			hashmap.Mapv6[prefix][hash] = &node
			result = true
		}

	} else if node.Type == "ipv4" {

		prefix := node.Prefix
		hash := node.Hash()

		_, ok1 := hashmap.Mapv4[prefix]

		if ok1 == false {
			hashmap.Mapv4[prefix] = make(map[string]*Node)
		}

		_, ok2 := hashmap.Mapv4[prefix][hash]

		if ok2 == true {
			result = false
		} else {
			hashmap.Mapv4[prefix][hash] = &node
			result = true
		}

	}

	return result

}

func (hashmap *HashMap) Print() {

	prefixesv4 := make([]uint8, 0)
	prefixesv6 := make([]uint8, 0)

	for prefix := range hashmap.Mapv4 {
		prefixesv4 = append(prefixesv4, prefix)
	}

	for prefix := range hashmap.Mapv6 {
		prefixesv6 = append(prefixesv6, prefix)
	}

	sort.Slice(prefixesv4, func(a int, b int) bool {
		return prefixesv4[a] < prefixesv4[b]
	})

	sort.Slice(prefixesv6, func(a int, b int) bool {
		return prefixesv6[a] < prefixesv6[b]
	})

	fmt.Println("ipv4")

	if len(prefixesv4) > 0 {

		for p := 0; p < len(prefixesv4); p++ {

			prefix := prefixesv4[p]

			for hash, node := range hashmap.Mapv4[prefix] {
				printHashMapNode(node, hash, prefix, "")
			}

		}

	} else {
		fmt.Println("(no entries)")
	}

	fmt.Println("")

	fmt.Println("ipv6")

	if len(prefixesv6) > 0 {

		for p := 0; p < len(prefixesv6); p++ {

			prefix := prefixesv6[p]

			for hash, node := range hashmap.Mapv6[prefix] {
				printHashMapNode(node, hash, prefix, "")
			}

		}

	} else {
		fmt.Println("(no entries)")
	}

}

func (hashmap *HashMap) Search(value string) *Node {

	var result *Node = nil

	if strings.Contains(value, "/") {

		address := value[0:strings.Index(value, "/")]
		prefix := value[strings.Index(value, "/")+1:]

		if utils.IsIPv6(address) {

			address = utils.ToIPv6(address)
			result = hashmap.SearchNode(ToNode(address + "/" + prefix))

		} else if utils.IsIPv4(address) {

			address = utils.ToIPv4(address)
			result = hashmap.SearchNode(ToNode(address + "/" + prefix))

		}

	} else {

		if utils.IsIPv6(value) {

			value = utils.ToIPv6(value)
			result = hashmap.SearchNode(ToNode(value + "/128"))

		} else if utils.IsIPv4(value) {

			value = utils.ToIPv4(value)
			result = hashmap.SearchNode(ToNode(value + "/32"))

		}

	}

	return result

}

func (hashmap *HashMap) SearchNode(value Node) *Node {

	var result *Node = nil

	if value.Type == "ipv6" {

		prefixes := make([]uint8, 0)

		for prefix := range hashmap.Mapv6 {
			prefixes = append(prefixes, prefix)
		}

		sort.Slice(prefixes, func(a int, b int) bool {
			return prefixes[a] > prefixes[b]
		})

		for p := 0; p < len(prefixes); p++ {

			prefix := prefixes[p]
			hash := hex.EncodeToString(utils.ToIPv6Bytes(value.Address, prefix))

			tmp, ok := hashmap.Mapv6[prefix][hash]

			if ok == true {
				result = tmp
				break
			}

		}

	} else if value.Type == "ipv4" {

		prefixes := make([]uint8, 0)

		for prefix := range hashmap.Mapv4 {
			prefixes = append(prefixes, prefix)
		}

		sort.Slice(prefixes, func(a int, b int) bool {
			return prefixes[a] > prefixes[b]
		})

		for p := 0; p < len(prefixes); p++ {

			prefix := prefixes[p]
			hash := hex.EncodeToString(utils.ToIPv4Bytes(value.Address, prefix))

			tmp, ok := hashmap.Mapv4[prefix][hash]

			if ok == true {
				result = tmp
				break
			}

		}

	}

	return result

}

func (hashmap *HashMap) SetNodes(values []Node) bool {

	var result bool = true

	for v := 0; v < len(values); v++ {

		check := hashmap.InsertNode(values[v])

		if check == false {
			result = false
		}

	}

	return result

}
