package main

import (
	"errors"
	"fmt"
)

type ControllerOption struct {
	//speed is move speed
	Speed string // use string cuz int cant be decimal number
	//OneKeyPressMove is single key press or a long key press
	OneKeyPressMove      bool
	TouchBlockReflection ControllerReflectionOption
}
type ControllerReflectionOption struct {
	Enable           bool
	ReflectBlockType []Block
}

// NewKeyController will make controller event for object
func (_ IWMEvent) NewKeyController(option ControllerOption) ([]Event, error) {
	var ControllerBasic []Event
	if option.OneKeyPressMove {
		ControllerBasic = []Event{
			{
				EventIndex: EVENT_ON_UP_DOWN_PRESSED,
				Param: []Param{
					{Key: "key_up_down", Val: fmt.Sprintf("%d", UP)},
				},
				Events: []EventInner{
					{
						EventIndex: EVENT_SET_SPEED,
						Param: []Param{
							{Key: "speed", Val: fmt.Sprintf("%s", option.Speed)},
							{Key: "additive", Val: fmt.Sprintf("%d", 0)},
						},
					},
					{
						EventIndex: EVENT_SET_ANGLE,
						Param: []Param{
							{Key: "angle", Val: fmt.Sprintf("%d", 90)},
							{Key: "additive", Val: fmt.Sprintf("%d", 0)},
						},
					},
				},
			},
			{
				EventIndex: EVENT_ON_UP_DOWN_PRESSED,
				Param: []Param{
					{Key: "key_up_down", Val: fmt.Sprintf("%d", DOWN)},
				},
				Events: []EventInner{
					{
						EventIndex: EVENT_SET_SPEED,
						Param: []Param{
							{Key: "speed", Val: fmt.Sprintf("%s", option.Speed)},
							{Key: "additive", Val: fmt.Sprintf("%d", 0)},
						},
					},
					{
						EventIndex: EVENT_SET_ANGLE,
						Param: []Param{
							{Key: "angle", Val: fmt.Sprintf("%d", 270)},
							{Key: "additive", Val: fmt.Sprintf("%d", 0)},
						},
					},
				},
			},
			{
				EventIndex: EVENT_ON_LEFT_RIGHT_PRESSED,
				Param: []Param{
					{Key: "key", Val: fmt.Sprintf("%d", LEFT)},
				},
				Events: []EventInner{
					{
						EventIndex: EVENT_SET_SPEED,
						Param: []Param{
							{Key: "speed", Val: fmt.Sprintf("%s", option.Speed)},
							{Key: "additive", Val: fmt.Sprintf("%d", 0)},
						},
					},
					{
						EventIndex: EVENT_SET_ANGLE,
						Param: []Param{
							{Key: "angle", Val: fmt.Sprintf("%d", 180)},
							{Key: "additive", Val: fmt.Sprintf("%d", 0)},
						},
					},
				},
			},
			{
				EventIndex: EVENT_ON_LEFT_RIGHT_PRESSED,
				Param: []Param{
					{Key: "key", Val: fmt.Sprintf("%d", RIGHT)},
				},
				Events: []EventInner{
					{
						EventIndex: EVENT_SET_SPEED,
						Param: []Param{
							{Key: "speed", Val: fmt.Sprintf("%s", option.Speed)},
							{Key: "additive", Val: fmt.Sprintf("%d", 0)},
						},
					},
					{
						EventIndex: EVENT_SET_ANGLE,
						Param: []Param{
							{Key: "angle", Val: fmt.Sprintf("%d", 0)},
							{Key: "additive", Val: fmt.Sprintf("%d", 0)},
						},
					},
				},
			},
		}
	} else {
		ControllerBasic = []Event{
			{
				EventIndex: EVENT_ON_UP_DOWN_PRESSED,
				Param: []Param{
					{Key: "key_up_down", Val: fmt.Sprintf("%d", UP)},
				},
				Events: []EventInner{
					{
						EventIndex: EVENT_FOLLOW_PATH,
						Param: []Param{
							{Key: "speed", Val: fmt.Sprintf("%s", option.Speed)},
							{Key: "path_nodes_y", Val: "2E010000010000000000000000000000000020C0"},
							{Key: "ease", Val: fmt.Sprintf("%d", 0)},
							{Key: "path_nodes_x", Val: "2E01000001000000000000000000000000000000"},
							{Key: "end_action", Val: fmt.Sprintf("%d", 0)},
						},
					},
				},
			},
			{
				EventIndex: EVENT_ON_UP_DOWN_PRESSED,
				Param: []Param{
					{Key: "key_up_down", Val: fmt.Sprintf("%d", DOWN)},
				},
				Events: []EventInner{
					{
						EventIndex: EVENT_FOLLOW_PATH,
						Param: []Param{
							{Key: "speed", Val: fmt.Sprintf("%s", option.Speed)},
							{Key: "path_nodes_y", Val: "2E01000001000000000000000000000000002040"},
							{Key: "ease", Val: fmt.Sprintf("%d", 0)},
							{Key: "path_nodes_x", Val: "2E01000001000000000000000000000000000000"},
							{Key: "end_action", Val: fmt.Sprintf("%d", 0)},
						},
					},
				},
			},
			{
				EventIndex: EVENT_ON_LEFT_RIGHT_PRESSED,
				Param: []Param{
					{Key: "key", Val: fmt.Sprintf("%d", LEFT)},
				},
				Events: []EventInner{
					{
						EventIndex: EVENT_FOLLOW_PATH,
						Param: []Param{
							{Key: "speed", Val: fmt.Sprintf("%s", option.Speed)},
							{Key: "path_nodes_y", Val: "2E01000001000000000000000000000000000000"},
							{Key: "ease", Val: fmt.Sprintf("%d", 0)},
							{Key: "path_nodes_x", Val: "2E010000010000000000000000000000000020C0"},
							{Key: "end_action", Val: fmt.Sprintf("%d", 0)},
						},
					},
				},
			},
			{
				EventIndex: EVENT_ON_LEFT_RIGHT_PRESSED,
				Param: []Param{
					{Key: "key", Val: fmt.Sprintf("%d", RIGHT)},
				},
				Events: []EventInner{
					{
						EventIndex: EVENT_FOLLOW_PATH,
						Param: []Param{
							{Key: "speed", Val: fmt.Sprintf("%s", option.Speed)},
							{Key: "path_nodes_y", Val: "2E01000001000000000000000000000000000000"},
							{Key: "ease", Val: fmt.Sprintf("%d", 0)},
							{Key: "path_nodes_x", Val: "2E01000001000000000000000000000000002040"},
							{Key: "end_action", Val: fmt.Sprintf("%d", 0)},
						},
					},
				},
			},
		}
	}
	if option.TouchBlockReflection.Enable {
		if !option.OneKeyPressMove {
			return nil, errors.New("DefaultKeyPressMove cant be used with TouchBlockReflection")
		}
		for _, v := range option.TouchBlockReflection.ReflectBlockType {
			ControllerBasic = append(ControllerBasic, Event{
				EventIndex: EVENT_ON_OBJECT_COLLISION,
				Param: []Param{
					{Key: "obj_id", Val: fmt.Sprintf("%d", v)},
				},
				Events: []EventInner{{
					EventIndex: EVENT_SET_ANGLE,
					Param: []Param{
						{Key: "angle", Val: fmt.Sprintf("%d", 180)},
						{Key: "additive", Val: fmt.Sprintf("%d", 1)},
					},
				}},
			})
		}
	}
	return ControllerBasic, nil
}
