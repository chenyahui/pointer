package pointer

import "golang.org/x/exp/constraints"

var True = Bool(true)
var False = Bool(false)

func Bool(x bool) *bool {
	return &x
}

func Int[T constraints.Integer](x T) *int {
	val := int(x)
	return &val
}

func Int32[T constraints.Integer](x T) *int32 {
	val := int32(x)
	return &val
}

func Uint32[T constraints.Integer](x T) *uint32 {
	val := uint32(x)
	return &val
}

func Int64[T constraints.Integer](x T) *int64 {
	val := int64(x)
	return &val
}

func Uint64[T constraints.Integer](x T) *uint64 {
	val := uint64(x)
	return &val
}

func Int8[T constraints.Integer](x T) *int8 {
	val := int8(x)
	return &val
}

func UInt8[T constraints.Integer](x T) *uint8 {
	val := uint8(x)
	return &val
}

func Int16[T constraints.Integer](x T) *int16 {
	val := int16(x)
	return &val
}

func Uint16[T constraints.Integer](x T) *uint16 {
	val := uint16(x)
	return &val
}

func Float32(x float32) *float32 {
	return &x
}

func String(x string) *string {
	return &x
}
