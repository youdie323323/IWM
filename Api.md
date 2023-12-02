
# Midi

```Func
  midi()
```

#### Convert midi into IWM

|    Arg    | Type     | Description                |
| :-------- | :------- | :------------------------- |
|    `f`    | `string` | **Required**. FilePath     |
|   `fps`   | `float64` | **Required**. Speed (25.0)    |
|  `track`  | `int` | **Required**. Track Number|
|  `volume`  | `int` | **Required**. Volume|
|  `x, y`  | `int` | **Required**. Objects Coordinate|
|  `MinusOffset`  | `int` | **PreRequired**. Reduce start offset|
|  `OnAudioEndEvent`  | `[]EventInner` | **Required**. Do Event on end|

Return Objects, Count, error
#### Usage: midi[[]*Object](filepath, fps, track, volume, x, y, MinusOffset, []EventInner{})
#### Example: `IWM_Midi2Iwm_test.go`
###

# Image - THIS POSSIBLE TO CRASH YOUR IWM, BE CAREFUL
```Func
  Image()
```

#### Convert image into IWM

|    Arg    | Type     | Description                |
| :-------- | :------- | :------------------------- |
|    `ImagePath`    | `string` | **Required**. FilePath     |
|   `width, height`   | `int` | **Required**. Size (794, 608)    |
|  `IgnoreColors`  | `[]int` | **PreRequired**. Ignore Color(B, G, R)|
|  `Scale`  | `float64` | **Required**. Scale (0.7)|
|  `EventOnObjects`  | `[]Event` | **PreRequired**. Add Event onto Objects (nil)|

Return Objects, Count, error
#### Usage: Image[ObjectsGroup](ImagePath, width, height, make([]int, 0), Scale, nil)
#### Example: `IWM_Image2Object_test.go`

# Bright
```Func
  Bright()
```

#### Convert image into IWM (using bright)

|    Arg    | Type     | Description                |
| :-------- | :------- | :------------------------- |
|    `ImagePath`    | `string` | **Required**. FilePath     |
|   `width, height`   | `int` | **Required**. Size (810, 624)    |

Return Objects, Count, error
#### Usage: Bright[[]*Object](ImagePath, width, height)
#### Example: `IWM_Image2Bright_test.go`
