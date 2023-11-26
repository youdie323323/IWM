package main

import (
	"strings"
)

type ObjectsGroup []*Object

func BuildAllObject[T ObjectsGroup | []*Object](Block T) string {
	var XML *strings.Builder = new(strings.Builder)
	for _, obj := range Block {
		XML.WriteString(obj.BuildXMLObject())
	}
	return XML.String()
}
