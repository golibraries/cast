package cast

import (
	"fmt"
	"math/big"
	"strconv"
)

// ToInt casts an interface to an int type.
func ToInt(i any) int {
	v, _ := ToIntE(i)
	return v
}

// ToIntE casts an interface to an int type.
func ToIntE(a any) (int, error) {
	a = indirect(a)

	switch v := a.(type) {
	case int:
		return v, nil
	case int64:
		return int(v), nil
	case int32:
		return int(v), nil
	case int16:
		return int(v), nil
	case int8:
		return int(v), nil
	case uint:
		return int(v), nil
	case uint64:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint8:
		return int(v), nil
	case float64:
		return int(v), nil
	case float32:
		return int(v), nil
	case *big.Int:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", a, a)
		}
		return int(v.Int64()), nil
	case *big.Float:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", a, a)
		}
		n, _ := v.Int64()
		return int(n), nil
	case string:
		n, err := strToInt(v)
		if err == nil {
			return int(n), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", a, a)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", a, a)
	}
}

// ToInt64 casts an interface to an int64 type.
func ToInt64(i any) int64 {
	v, _ := ToInt64E(i)
	return v
}

// ToInt64E casts an interface to an int64 type.
func ToInt64E(a any) (int64, error) {
	a = indirect(a)

	switch v := a.(type) {
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case int32:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case uint:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case float32:
		return int64(v), nil
	case *big.Int:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int64", a, a)
		}
		return v.Int64(), nil
	case *big.Float:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int64", a, a)
		}
		n, _ := v.Int64()
		return n, nil
	case string:
		n, err := strToInt(v)
		if err == nil {
			return n, nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", a, a)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", a, a)
	}
}

// ToInt32 casts an interface to an int32 type.
func ToInt32(i any) int32 {
	v, _ := ToInt32E(i)
	return v
}

// ToInt32E casts an interface to an int32 type.
func ToInt32E(a any) (int32, error) {
	a = indirect(a)

	switch v := a.(type) {
	case int:
		return int32(v), nil
	case int64:
		return int32(v), nil
	case int32:
		return v, nil
	case int16:
		return int32(v), nil
	case int8:
		return int32(v), nil
	case uint:
		return int32(v), nil
	case uint64:
		return int32(v), nil
	case uint32:
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case uint8:
		return int32(v), nil
	case float64:
		return int32(v), nil
	case float32:
		return int32(v), nil
	case *big.Int:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int32", a, a)
		}
		return int32(v.Int64()), nil
	case *big.Float:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int32", a, a)
		}
		n, _ := v.Int64()
		return int32(n), nil
	case string:
		n, err := strToInt(v)
		if err == nil {
			return int32(n), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", a, a)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", a, a)
	}
}

// ToInt16 casts an interface to an int16 type.
func ToInt16(i any) int16 {
	v, _ := ToInt16E(i)
	return v
}

// ToInt16E casts an interface to an int16 type.
func ToInt16E(a any) (int16, error) {
	a = indirect(a)

	switch v := a.(type) {
	case int:
		return int16(v), nil
	case int64:
		return int16(v), nil
	case int32:
		return int16(v), nil
	case int16:
		return v, nil
	case int8:
		return int16(v), nil
	case uint:
		return int16(v), nil
	case uint64:
		return int16(v), nil
	case uint32:
		return int16(v), nil
	case uint16:
		return int16(v), nil
	case uint8:
		return int16(v), nil
	case float64:
		return int16(v), nil
	case float32:
		return int16(v), nil
	case *big.Int:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int16", a, a)
		}
		return int16(v.Int64()), nil
	case *big.Float:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int16", a, a)
		}
		n, _ := v.Int64()
		return int16(n), nil
	case string:
		n, err := strToInt(v)
		if err == nil {
			return int16(n), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", a, a)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", a, a)
	}
}

// ToInt8 casts an interface to an int8 type.
func ToInt8(i any) int8 {
	v, _ := ToInt8E(i)
	return v
}

// ToInt8E casts an interface to an int8 type.
func ToInt8E(a any) (int8, error) {
	a = indirect(a)

	switch v := a.(type) {
	case int:
		return int8(v), nil
	case int64:
		return int8(v), nil
	case int32:
		return int8(v), nil
	case int16:
		return int8(v), nil
	case int8:
		return v, nil
	case uint:
		return int8(v), nil
	case uint64:
		return int8(v), nil
	case uint32:
		return int8(v), nil
	case uint16:
		return int8(v), nil
	case uint8:
		return int8(v), nil
	case float64:
		return int8(v), nil
	case float32:
		return int8(v), nil
	case *big.Int:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int8", v, v)
		}
		return int8(v.Int64()), nil
	case *big.Float:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int8", v, v)
		}
		n, _ := v.Int64()
		return int8(n), nil
	case string:
		n, err := strToInt(v)
		if err == nil {
			return int8(n), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", v, v)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", v, v)
	}
}

// ToUint casts an interface to a uint type.
func ToUint(i any) uint {
	v, _ := ToUintE(i)
	return v
}

// ToUintE casts an interface to a uint type.
func ToUintE(a any) (uint, error) {
	a = indirect(a)

	switch v := a.(type) {
	case int:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint(v), nil
	case int64:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint(v), nil
	case int32:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint(v), nil
	case int16:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint(v), nil
	case int8:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint(v), nil
	case uint:
		return v, nil
	case uint64:
		return uint(v), nil
	case uint32:
		return uint(v), nil
	case uint16:
		return uint(v), nil
	case uint8:
		return uint(v), nil
	case float64:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint(v), nil
	case float32:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint(v), nil
	case *big.Int:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint", a, a)
		}
		if v.Sign() < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint(v.Uint64()), nil
	case *big.Float:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint", a, a)
		}
		if v.Sign() < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		n, _ := v.Uint64()
		return uint(n), nil
	case string:
		n, err := strToInt(v)
		if err == nil {
			if n < 0 {
				return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
			}
			return uint(n), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint", a, a)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint", a, a)
	}
}

// ToUint64 casts an interface to a uint64 type.
func ToUint64(i any) uint64 {
	v, _ := ToUint64E(i)
	return v
}

// ToUint64E casts an interface to a uint64 type.
func ToUint64E(a any) (uint64, error) {
	a = indirect(a)

	switch v := a.(type) {
	case int:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint64(v), nil
	case int64:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint64(v), nil
	case int32:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint64(v), nil
	case int16:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint64(v), nil
	case int8:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case uint64:
		return v, nil
	case uint32:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case float64:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint64(v), nil
	case float32:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint64(v), nil
	case *big.Int:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", a, a)
		}
		if v.Sign() < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return v.Uint64(), nil
	case *big.Float:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", a, a)
		}
		if v.Sign() < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		n, _ := v.Uint64()
		return n, nil
	case string:
		n, err := strToInt(v)
		if err == nil {
			if n < 0 {
				return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
			}
			return uint64(n), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", a, a)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", a, a)
	}
}

// ToUint32 casts an interface to a uint32 type.
func ToUint32(i any) uint32 {
	v, _ := ToUint32E(i)
	return v
}

// ToUint32E casts an interface to a uint32 type.
func ToUint32E(a any) (uint32, error) {
	a = indirect(a)

	switch v := a.(type) {
	case int:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint32(v), nil
	case int64:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint32(v), nil
	case int32:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint32(v), nil
	case int16:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint32(v), nil
	case int8:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint32(v), nil
	case uint:
		return uint32(v), nil
	case uint64:
		return uint32(v), nil
	case uint32:
		return v, nil
	case uint16:
		return uint32(v), nil
	case uint8:
		return uint32(v), nil
	case float64:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint32(v), nil
	case float32:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint32(v), nil
	case *big.Int:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", a, a)
		}
		if v.Sign() < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint32(v.Uint64()), nil
	case *big.Float:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", a, a)
		}
		if v.Sign() < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		n, _ := v.Uint64()
		return uint32(n), nil
	case string:
		n, err := strToInt(v)
		if err == nil {
			if n < 0 {
				return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
			}
			return uint32(n), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", a, a)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", a, a)
	}
}

// ToUint16 casts an interface to a uint16 type.
func ToUint16(i any) uint16 {
	v, _ := ToUint16E(i)
	return v
}

// ToUint16E casts an interface to a uint16 type.
func ToUint16E(a any) (uint16, error) {
	a = indirect(a)

	switch v := a.(type) {
	case int:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint16(v), nil
	case int64:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint16(v), nil
	case int32:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint16(v), nil
	case int16:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint16(v), nil
	case int8:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint16(v), nil
	case uint:
		return uint16(v), nil
	case uint64:
		return uint16(v), nil
	case uint32:
		return uint16(v), nil
	case uint16:
		return v, nil
	case uint8:
		return uint16(v), nil
	case float64:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint16(v), nil
	case float32:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint16(v), nil
	case *big.Int:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", a, a)
		}
		if v.Sign() < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint16(v.Uint64()), nil
	case *big.Float:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", a, a)
		}
		if v.Sign() < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		n, _ := v.Uint64()
		return uint16(n), nil
	case string:
		n, err := strToInt(v)
		if err == nil {
			if n < 0 {
				return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
			}
			return uint16(n), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", a, a)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", a, a)
	}
}

// ToUint8 casts an interface to a uint8 type.
func ToUint8(i any) uint8 {
	v, _ := ToUint8E(i)
	return v
}

// ToUint8E casts an interface to a uint type.
func ToUint8E(a any) (uint8, error) {
	a = indirect(a)

	switch v := a.(type) {
	case int:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint8(v), nil
	case int64:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint8(v), nil
	case int32:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint8(v), nil
	case int16:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint8(v), nil
	case int8:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint8(v), nil
	case uint:
		return uint8(v), nil
	case uint64:
		return uint8(v), nil
	case uint32:
		return uint8(v), nil
	case uint16:
		return uint8(v), nil
	case uint8:
		return v, nil
	case float64:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint8(v), nil
	case float32:
		if v < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint8(v), nil
	case *big.Int:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", a, a)
		}
		if v.Sign() < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		return uint8(v.Uint64()), nil
	case *big.Float:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", a, a)
		}
		if v.Sign() < 0 {
			return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
		}
		n, _ := v.Uint64()
		return uint8(n), nil
	case string:
		n, err := strToInt(v)
		if err == nil {
			if n < 0 {
				return 0, fmt.Errorf("unable to cast %#v of type %T to negative value", v, v)
			}
			return uint8(n), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", a, a)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", a, a)
	}
}

// ToFloat64 casts an interface to a float64 type.
func ToFloat64(i any) float64 {
	v, _ := ToFloat64E(i)
	return v
}

// ToFloat64E casts an interface to a float64 type.
func ToFloat64E(a any) (float64, error) {
	a = indirect(a)

	switch v := a.(type) {
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case *big.Int:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to float64", a, a)
		}
		n, _ := new(big.Float).SetInt(v).Float64()
		return n, nil
	case *big.Float:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to float64", a, a)
		}
		n, _ := v.Float64()
		return n, nil
	case string:
		n, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return n, nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", a, a)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", a, a)
	}
}

// ToFloat32 casts an interface to a float32 type.
func ToFloat32(i any) float32 {
	v, _ := ToFloat32E(i)
	return v
}

// ToFloat32E casts an interface to a float32 type.
func ToFloat32E(a any) (float32, error) {
	a = indirect(a)

	switch v := a.(type) {
	case int:
		return float32(v), nil
	case int64:
		return float32(v), nil
	case int32:
		return float32(v), nil
	case int16:
		return float32(v), nil
	case int8:
		return float32(v), nil
	case uint:
		return float32(v), nil
	case uint64:
		return float32(v), nil
	case uint32:
		return float32(v), nil
	case uint16:
		return float32(v), nil
	case uint8:
		return float32(v), nil
	case float64:
		return float32(v), nil
	case float32:
		return v, nil
	case *big.Int:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to float32", a, a)
		}
		n, _ := new(big.Float).SetInt(v).Float32()
		return n, nil
	case *big.Float:
		if v == nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to float32", a, a)
		}
		n, _ := v.Float32()
		return n, nil
	case string:
		n, err := strconv.ParseFloat(v, 32)
		if err == nil {
			return float32(n), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to float32", a, a)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float32", a, a)
	}
}

func strToInt(s string) (int64, error) {
	var foundZero bool
loop:
	for i := len(s); i > 0; i-- {
		switch s[i-1] {
		case '.':
			if foundZero {
				s = s[:i-1]
				break loop
			}
		case '0':
			foundZero = true
		default:
			break loop
		}
	}
	return strconv.ParseInt(s, 0, 0)
}
