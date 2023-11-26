package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func IWM_Image2Object_Test() {
	Objects, numObjects, err := Image[ObjectsGroup](filepath.Join("assets", "img.png"), 794, 608, make([]int, 0), 0.5, []Event{
		{
			EventIndex: EVENT_ON_CREATE,
			Events: []EventInner{
				{
					EventIndex: EVENT_DESTROY,
				},
			},
		},
	})
	if err != nil {
		fmt.Printf("%s\r\n", err)
		return
	}
	file, err := os.OpenFile(fmt.Sprintf("numobject=%d.txt", numObjects), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("Cannot create file", err)
		return
	}
	defer file.Close()
	file.WriteString(BuildAllObject(Objects))
}
