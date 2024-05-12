package magic

import (
	"strconv"

	"golang.org/x/exp/constraints"
	tb "gopkg.in/telebot.v3"
)

type NumberConvertor struct {
	getter ItemGetter[string]
}

func (n NumberConvertor) Int() NumberFilter[int] {
	return newNumberFilter(joinConvertor(n.getter, strconv.Atoi))
}

func (n NumberConvertor) Int8() NumberFilter[int8] {
	return intConvertor[int8](n.getter, 8)
}

func (n NumberConvertor) Int16() NumberFilter[int16] {
	return intConvertor[int16](n.getter, 16)
}

func (n NumberConvertor) Int32() NumberFilter[int32] {
	return intConvertor[int32](n.getter, 32)
}

func (n NumberConvertor) Int64() NumberFilter[int64] {
	return intConvertor[int64](n.getter, 64)
}

func (n NumberConvertor) Uint() NumberFilter[uint] {
	return uintConvertor[uint](n.getter, 0)
}

func (n NumberConvertor) Uint8() NumberFilter[uint8] {
	return uintConvertor[uint8](n.getter, 8)
}

func (n NumberConvertor) Uint16() NumberFilter[uint16] {
	return uintConvertor[uint16](n.getter, 16)
}

func (n NumberConvertor) Uint32() NumberFilter[uint32] {
	return uintConvertor[uint32](n.getter, 32)
}

func (n NumberConvertor) Uint64() NumberFilter[uint64] {
	return uintConvertor[uint64](n.getter, 64)
}

func (n NumberConvertor) Float32() NumberFilter[float32] {
	return newNumberFilter(
		joinConvertor(n.getter, func(s string) (float32, error) {
			f, err := strconv.ParseFloat(s, 32)
			return float32(f), err
		}),
	)
}

func (n NumberConvertor) Float64() NumberFilter[float64] {
	return newNumberFilter(
		joinConvertor(n.getter, func(s string) (float64, error) {
			return strconv.ParseFloat(s, 64)
		}),
	)
}

func joinConvertor[N Number](
	getter ItemGetter[string],
	convertor func(string) (N, error),
) ItemGetter[N] {
	return func(ctx tb.Context) (N, bool) {
		s, ok := getter(ctx)
		if !ok {
			return 0, false
		}

		n, err := convertor(s)
		if err != nil {
			return 0, false
		}
		return n, true
	}
}

func intConvertor[N constraints.Signed](getter ItemGetter[string], bitSize int) NumberFilter[N] {
	c := numberConvert[int64, N](bitSize, strconv.ParseInt)
	return newNumberFilter(joinConvertor(getter, c))
}

func uintConvertor[N constraints.Unsigned](getter ItemGetter[string], bitSize int) NumberFilter[N] {
	c := numberConvert[uint64, N](bitSize, strconv.ParseUint)
	return newNumberFilter(joinConvertor(getter, c))
}

func numberConvert[T, N constraints.Integer](
	bitSize int,
	fn func(string, int, int) (T, error),
) func(string) (N, error) {
	return func(s string) (N, error) {
		i, err := fn(s, 10, bitSize)
		if err != nil {
			return 0, err
		}

		return N(i), nil
	}
}
