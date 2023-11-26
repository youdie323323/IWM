package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func IWM_Midi2Iwm_test() {
	tracks, err := CheckMidiTracks(filepath.Join("assets", "entrance.mid"))
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	for i, v := range tracks {
		fmt.Println("Track: " + strconv.Itoa(i+1) + " Key count: " + strconv.Itoa(len(v)))
	}
	fmt.Printf("Which Track You Want ?>")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	STrack, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}
	STrack--
	Objects, numObjects, err := midi[ObjectsGroup](filepath.Join("assets", "entrance.mid"), 25.0, STrack, 1, 50, 50, 0, []EventInner{
		{
			EventIndex: EVENT_DESTROY,
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
