package main

import (
	"fmt"
	"math"
	"strings"

	"lukechampine.com/uint128"
)

type Decimal struct {
	V uint128.Uint128
	F string
}

const (
	NEAR, UP, DOWN = 0, 1, 2
)

func main() {
	dividend, divisor := get_input("38612.9032", "30")
	man, dec, err := run_division2(dividend, divisor, 3, DOWN)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(man.String() + "." + dec)
}

func run_division2(dividend Decimal, divisor uint128.Uint128, maxDecimalCount, roundingMode int) (*uint128.Uint128, string, error) {

	var man uint128.Uint128
	var dec string

	divid, err := uint128.FromString(dividend.V.String() + dividend.F)
	if err != nil {
		return nil, "", err
	}
	q, r := divid.QuoRem(divisor)

	if dividend.F != "" {
		man = q.Div64(uint64(math.Pow10(len(dividend.F))))
		dec = q.String()[len(man.String()):]
	} else {
		man = q
	}

	if maxDecimalCount+1 > len(dec) {
		dec += div_remainder(r, divisor, maxDecimalCount+1-len(dec))
	} else if maxDecimalCount+1 < len(dec) {
		dec = dec[:maxDecimalCount+1]
	}

	var ov bool
	dec, ov = rounding(dec, roundingMode)
	if ov {
		man = man.Add64(1)
	}

	return &man, dec, nil

}

func rounding(dec string, mode int) (string, bool) {
	allowedDecCount := len(dec) - 1
	last := dec[allowedDecCount]
	dec = dec[:allowedDecCount]

	if (mode == NEAR && last > '5') || (mode == UP && last >= '5') {
		decint, _ := uint_from_string(dec)
		sum := decint.Add64(1)
		var ans string
		var ov bool
		if len(sum.String()) > allowedDecCount {
			ans = sum.String()[1:]
			ov = true
		} else {
			ans = sum.String()
		}
		return strings.Repeat("0", allowedDecCount-len(ans)) + ans, ov
	}
	return dec, false
}

func div_remainder(divid, divis uint128.Uint128, count int) string {

	ans := ""
	for !divid.IsZero() && len(ans) <= count {
		if divid.Cmp(divis) < 0 {
			var times int
			divid, times = div_mul_10(divid, divis)
			ans += strings.Repeat("0", times)
		}
		q, r := divid.QuoRem(divis)
		ans += q.String()
		divid = r
	}
	if len(ans) < count {
		ans += strings.Repeat("0", count-len(ans))
	}
	return ans[:count]
}
func div_mul_10(divid, divis uint128.Uint128) (uint128.Uint128, int) {
	times := 0
	divid = divid.Mul64(10)
	for divid.Cmp(divis) < 0 {
		divid = divid.Mul64(10)
		times += 1
	}
	return divid, times
}

func get_input(dividend, divisor string) (Decimal, uint128.Uint128) {
	divid := Decimal{}
	var divis uint128.Uint128

	index := strings.IndexByte(dividend, '.')
	if index != -1 {
		divid.V, _ = uint_from_string(dividend[:index])
		divid.F = dividend[index+1:]
	} else {
		divid.V, _ = uint_from_string(dividend)
		divid.F = ""
	}
	divis, _ = uint_from_string(divisor)
	return divid, divis
}

func uint_from_string(a string) (uint128.Uint128, error) {
	a = strings.TrimLeft(a, "0")
	return uint128.FromString(a)
}
