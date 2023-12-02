package main

import (
	"encoding/binary"
	"encoding/hex"
	"math"
)

/* Special thanks: onbo */

func HexadecimalToFloat64(f string) (float64, error) {
	byteSlices := make([]byte, 8)
	for bytes := 0; bytes < 8; bytes++ {
		byteVal, err := hex.DecodeString(f[bytes*2 : bytes*2+2])
		if err != nil {
			return 0, err
		}
		byteSlices[bytes] = byteVal[0]
	}
	return math.Float64frombits(binary.LittleEndian.Uint64(byteSlices)), nil
}

func Float64ToHexadecimal(f float64) string {
	byteSlices := make([]byte, 8)
	binary.LittleEndian.PutUint64(byteSlices, math.Float64bits(f))
	return hex.EncodeToString(byteSlices)
}
