package main

import (
	"strings"
)

type ObjectsGroup []*Object

// AppendEventToAllObject will append event to all object
func AppendEventToAllObject[T ObjectsGroup | []*Object](Block T, Events []Event) T {
	for _, value := range Block {
		value.Event = append(value.Event, Events...)
	}
	return Block
}

func BuildAllObject[T ObjectsGroup | []*Object](Block T) string {
	var XML *strings.Builder = new(strings.Builder)
	for _, obj := range Block {
		XML.WriteString(obj.BuildXMLObject())
	}
	return XML.String()
}
