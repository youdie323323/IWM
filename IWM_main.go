package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Object struct {
	XMLName xml.Name `xml:"object"`
	Type    int      `xml:"type,attr"`
	X       int      `xml:"x,attr"`
	Y       int      `xml:"y,attr"`
	Param   []Param  `xml:"param,omitempty"`
	Event   []Event  `xml:"event,omitempty"`
	Obj     []Slots  `xml:"obj,omitempty"`
}

type Slots struct {
	Type  Block   `xml:"type,attr"`
	X     int     `xml:"x,attr"`
	Y     int     `xml:"y,attr"`
	Slot  int     `xml:"slot,attr"`
	Param []Param `xml:"param"`
}

type Param struct {
	Key string `xml:"key,attr"`
	Val string `xml:"val,attr"`
}

type Event struct {
	EventIndex Listener     `xml:"eventIndex,attr"`
	Param      []Param      `xml:"param"`
	Events     []EventInner `xml:"event"`
}

type EventInner struct {
	EventIndex int     `xml:"eventIndex,attr"`
	Param      []Param `xml:"param"`
}

func NewObject[T Block | int](x int, y int, ObjectType T, params []Param) *Object {
	return &Object{
		Type:  int(ObjectType),
		X:     x,
		Y:     y,
		Param: params,
		Event: make([]Event, 0),
		Obj:   make([]Slots, 0),
	}
}

func (t *Object) BuildXMLObject() string {
	buf, err := xml.Marshal(t)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return string(buf)
}

func main() {
	o, n, err := Bright[[]*Object](filepath.Join("assets", "nerd.png"), 810, 624)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	file, err := os.OpenFile(fmt.Sprintf("numobject=%d.txt", n), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("Cannot create file", err)
		return
	}
	defer file.Close()
	file.WriteString(BuildAllObject(o))
}
