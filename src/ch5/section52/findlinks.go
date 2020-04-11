package section52

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// //NodeType ...
// type NodeType int32

// const (
// 	//ErrorNode ...
// 	ErrorNode NodeType = iota
// 	//TextNode ...
// 	TextNode
// 	//DocumentNode ...
// 	DocumentNode
// 	//ElementNode ...
// 	ElementNode
// 	//CommentNode ...
// 	CommentNode
// 	//DoctypeNode ...
// 	DoctypeNode
// )

// //Attribute ...
// type Attribute struct {
// 	Key, Val string
// }

// //Node ...
// type Node struct {
// 	Type                    NodeType
// 	Data                    string
// 	Attr                    []Attribute
// 	FirstChild, NextSibling *Node
// }

// //Parse ...
// func Parse(r io.Reader) (*Node, error) {
// 	return nil, nil
// }

//FindLinks ...
func FindLinks() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	for _, link := range Visit(nil, doc) {
		fmt.Println(link)
	}
}

//Visit appends to links each link found in n and returns the result.
func Visit(links []string, n *html.Node) []string {
	return nil
}
