package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func IWM_Image2Bright_test() {
	o, n, err := Bright[[]*Object](filepath.Join("assets", "img.png"), 810, 624)
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
