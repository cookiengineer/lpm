package main

import "github.com/cookiengineer/lpm"
import "fmt"

func main() {

	root := lpm.NewNode("0.0.0.0", 0)
	node := lpm.NewNode("192.168.0.123", 32)

	fmt.Println(root.ContainsNode(node))

	node1 := lpm.NewNode("192.169.0.0", 16)
	node2 := lpm.NewNode("192.169.128.0", 17)

	fmt.Println(node1.ContainsNode(node2))

}
