package main

import (
	"errors"
	"gitlab.com/gomidi/midi/v2/smf"
	"strconv"
)

type Note struct {
	AbsTicks uint64
	Velocity uint8
	Key      uint8
	Value    uint32
}

func CheckMidiTracks(f string) ([]smf.Track, error) {
	ff, err := smf.ReadFile(f)
	if err != nil {
		return nil, errors.New("error on read midi")
	}
	_, ok := ff.TimeFormat.(smf.MetricTicks)
	if !ok {
		return nil, errors.New("time format is not MetricTicks")
	}
	return ff.Tracks, nil
}

// return: objects, numobject, error
func midi[T []*Object | ObjectsGroup](f string, fps float64, track, volume int, x, y int, MinusOffset int, OnAudioEndEvent []EventInner) (T, int, error) {
	ff, err := smf.ReadFile(f)
	if err != nil {
		return nil, 0, errors.New("error on read midi")
	}
	ticks, ok := ff.TimeFormat.(smf.MetricTicks)
	if !ok {
		return nil, 0, errors.New("time format is not MetricTicks")
	}
	t := ff.Tracks[track]
	var absTicks uint64
	var channel uint8
	var velocity uint8
	var key uint8
	startedNotes := make([]Note, 0, len(t))
	notes := make([]Note, 0, len(t))
	for _, ev := range t {
		absTicks += uint64(ev.Delta)
		switch {
		case ev.Message.GetNoteStart(&channel, &key, &velocity):
			if key > 49 {
				key = key % 49
			}
			startedNotes = append(startedNotes, Note{
				AbsTicks: absTicks,
				Velocity: velocity,
				Key:      key,
			})
		case ev.Message.GetNoteEnd(&channel, &key):
			if key > 49 {
				key = key % 49
			}
			delIdx := -1
			for i, v := range startedNotes {
				if v.Key == key {
					v.Value = uint32(absTicks - v.AbsTicks)
					notes = append(notes, v)
					delIdx = i
				}
			}
			if delIdx >= 0 {
				startedNotes = startedNotes[:delIdx+copy(startedNotes[delIdx:], startedNotes[delIdx+1:])]
			}
		}
	}
	e := new(IWMEvent)
	var prevAbsTicks uint64
	var totalFrames int
	var num int
	var all []*Object
	for _, note := range notes {
		delayInFrames := (float64(note.AbsTicks-prevAbsTicks) * 60000.0 / float64(uint64(ticks)*120)) / (1000.0 / fps)
		o := NewObject(x, y, OBJECT_TYPE_BLOCK, []Param{
			{Key: "scale", Val: "1"},
			{Key: "layer", Val: "2"},
			{Key: "tileset", Val: "0"},
			{Key: "blend_color", Val: "16777215"},
		})
		var defaultEventInner []EventInner = []EventInner{
			{
				EventIndex: EVENT_PLAY_SOUND,
				Param: []Param{
					{Key: "sound", Val: "18"},
					{Key: "volume", Val: strconv.Itoa(volume)},
					{Key: "pitch", Val: strconv.FormatFloat(AllIwmNotes[note.Key], 'f', -1, 64)},
				},
			},
		}
		if OnAudioEndEvent != nil {
			defaultEventInner = append(defaultEventInner, OnAudioEndEvent...)
		}
		o.AppendEvent(e.OnTimerEvent(EVENT_ON_CREATE, (totalFrames+int(delayInFrames)+1)-MinusOffset, defaultEventInner, 1))
		all = append(all, o)
		prevAbsTicks = note.AbsTicks
		totalFrames += int(delayInFrames)
		num++
	}
	return all, num, nil
}

var AllIwmNotes = map[uint8]float64{ // key: pitch
	0:  0.1768,
	1:  0.1873,
	2:  0.1984,
	3:  0.2102,
	4:  0.2227,
	5:  0.236,
	6:  0.25,
	7:  0.2649,
	8:  0.2806,
	9:  0.2973,
	10: 0.315,
	11: 0.3337,
	12: 0.3536,
	13: 0.3746,
	14: 0.3968,
	15: 0.4204,
	16: 0.4454,
	17: 0.4719,
	18: 0.5,
	19: 0.5297,
	20: 0.5612,
	21: 0.5946,
	22: 0.63,
	23: 0.6674,
	24: 0.7071,
	25: 0.7491,
	26: 0.7937,
	27: 0.8409,
	28: 0.8909,
	29: 0.9439,
	30: 1,
	31: 1.0595,
	32: 1.1225,
	33: 1.1892,
	34: 1.2599,
	35: 1.3348,
	36: 1.4142,
	37: 1.4983,
	38: 1.5874,
	39: 1.6818,
	40: 1.7818,
	41: 1.8877,
	42: 2,
	43: 2.1189,
	44: 2.2449,
	45: 2.3874,
	46: 2.5198,
	47: 2.6697,
	48: 2.8284,
	49: 2.9966,
}
