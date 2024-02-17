package main

import "github.com/cookiengineer/lpm"
import "fmt"

func main() {

	hashmap := lpm.NewHashMap()

	hashmap.Insert("192.168.0.0/24")
	hashmap.Insert("192.169.128.0/17")
	hashmap.Insert("192.169.0.0/16")
	hashmap.Insert("192.169.0.0/24")

	fmt.Println("parent of 192.168.0.123 is:", hashmap.Search("192.168.0.123/32"))
	fmt.Println("parent of 192.169.0.123 is:", hashmap.Search("192.169.0.123/32"))

	fmt.Println("")
	fmt.Println("hashmap:")
	hashmap.Print()

}
