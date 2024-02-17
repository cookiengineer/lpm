package main

import "github.com/cookiengineer/lpm"
import "fmt"

func main() {

	trie := lpm.NewTrie()

	trie.Insert("192.168.0.0/24")
	trie.Insert("192.169.0.0/16")
	trie.Insert("192.169.128.0/17")
	trie.Insert("192.169.0.0/24")

	fmt.Println("parent of 192.168.0.123 is:", trie.Search("192.168.0.123/32"))
	fmt.Println("parent of 192.169.0.123 is:", trie.Search("192.169.0.123/32"))

	fmt.Println("")
	fmt.Println("trie:")
	trie.Print()

}
