package main

import (
	"fmt"
)

func (t *Object) AppendEvent(event []Event) {
	t.Event = append(t.Event, event...) // イベントを追加
}

func (t *Object) AppendTag(slot []Slots) {
	t.Obj = append(t.Obj, slot...) // タグを追加
}

func (t *Object) AppendParam(param []Param) {
	t.Param = append(t.Param, param...)
}

type IWMEvent struct{}

// OnTimerEvent will make timer event
func (e IWMEvent) OnTimerEvent(EventListener Listener, frames int, innerEvents []EventInner, timerNumber int) []Event {
	var event []Event
	event = append(event, e.OnTimer(innerEvents, timerNumber)...)
	event = append(event, e.SetTimer(EventListener, frames, timerNumber)...)
	return event
}

// SetTimer is Do OnTimer event
func (_ IWMEvent) SetTimer(EventListener Listener, frames int, timerNumber int, P ...Param) []Event {
	var a = []Event{
		{
			EventIndex: EventListener,
			Events: []EventInner{
				{
					EventIndex: EVENT_SET_TIMER,
					Param: []Param{
						{Key: "frames", Val: fmt.Sprintf("%d", frames)},
						{Key: "timer_number", Val: fmt.Sprintf("%d", timerNumber)},
					},
				},
			},
		},
	}
	if len(P) > 0 {
		for _, v := range P {
			a[0].Events[0].Param = append(a[0].Events[0].Param, v)
		}
	}
	return a
}

// OnTimer is Set OnTimer event
func (_ IWMEvent) OnTimer(innerEvents []EventInner, timerNumber int) []Event {
	return []Event{
		{
			EventIndex: EVENT_ON_TIMER,
			Param: []Param{
				{Key: "timer_number", Val: fmt.Sprintf("%d", timerNumber)}, // タイマー番号を設定
			},
			Events: innerEvents,
		},
	}
}

// OnTrigger is Do Event when player touch trigger
func (_ IWMEvent) OnTrigger(innerEvents []EventInner, triggerNumber int) []Event {
	return []Event{
		{
			EventIndex: EVENT_ON_TRIGGER,
			Param: []Param{
				{Key: "trigger_number", Val: fmt.Sprintf("%d", triggerNumber)},
			},
			Events: innerEvents,
		},
	}
}

// OnKeyPressEvent will make key press event
func (_ IWMEvent) OnKeyPressEvent(Key any, innerEvents []EventInner) []Event {
	switch Key.(type) {
	case KeysLeftRight:
		return []Event{
			{
				EventIndex: EVENT_ON_LEFT_RIGHT_PRESSED,
				Param: []Param{
					{Key: "key", Val: fmt.Sprintf("%d", Key)},
				},
				Events: innerEvents, // イベントを追加
			},
		}
	case KeysUpDown:
		return []Event{
			{
				EventIndex: EVENT_ON_UP_DOWN_PRESSED,
				Param: []Param{
					{Key: "key_up_down", Val: fmt.Sprintf("%d", Key)},
				},
				Events: innerEvents,
			},
		}
	}
	return []Event{}
}

// IntervalEvent will make interval event, Offset is start after, Frames is interval
func (_ IWMEvent) IntervalEvent(Frames, Offset int, innerEvents []EventInner) []Event {
	return []Event{
		{
			EventIndex: EVENT_ON_METRONOME_TICK,
			Param: []Param{
				{Key: "frames", Val: fmt.Sprintf("%d", Frames)},
				{Key: "offset", Val: fmt.Sprintf("%d", Offset)},
			},
			Events: innerEvents,
		},
	}
}

// SoundEvent will return sound event
func (_ IWMEvent) SoundEvent(SoundType, Volume, Pitch int) []EventInner {
	return []EventInner{
		{
			EventIndex: EVENT_PLAY_SOUND,
			Param: []Param{
				{Key: "volume", Val: fmt.Sprintf("%d", Volume)},
				{Key: "pitch", Val: fmt.Sprintf("%d", Pitch)},
				{Key: "sound", Val: fmt.Sprintf("%d", SoundType)},
			},
		},
	}
}
