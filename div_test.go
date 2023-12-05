package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	dividend    string
	divisor     string
	q           string
	r           string
	maxDecCount int
	rounding    int
}

func TestDiv(t *testing.T) {

	cases := []TestCase{
		{"123.123", "3", "41", "041", 3, 2},
		{"987654321987654321.987654321987654321", "123456789123456789", "8", "000", 3, 2},

		{"1000000000000000000000", "2", "500000000000000000000", "000", 3, 2},
		{"0", "123123123", "0", "000", 3, 2},
		{"123123123.123123123", "1", "123123123", "123", 3, 2},

		{"36.1", "4", "9", "025", 3, 2},
		{"36.1", "4", "9", "02", 2, 0},
		{"36.1", "4", "9", "03", 2, 1},

		{"83703.168", "678", "123", "456", 3, 2},
		{"83703.168", "678", "123", "46", 2, 1},
		{"83703.168", "678", "123", "46", 2, 0},

		{"38612.9032", "30", "1287", "096", 3, 2},

		{"21", "7", "3", "000", 3, 2},
		{"21.0", "7", "3", "000", 3, 2},
		{"5.5", "5", "1", "100", 3, 2},
		{"5.0055", "5", "1", "0011", 4, 2},
		{"5.0055", "5", "1", "0011", 4, 2},
		{"5555.0055", "11", "505", "000500000000", 12, 2},

		{"738.7407407367402", "6", "123", "1234567894567", 13, 2},
		{"738.7407407367402", "6", "123", "123456789457", 12, 1},

		{"861.8641975265", "7", "123", "1234567895", 10, 1},
		{"861.8641975265", "7", "123", "123456789500", 12, 1},
		{"861.8641975265", "7", "123", "123456790", 9, 1},

		{"984.792", "8", "123", "099", 3, 2},
		{"984.792", "8", "123", "10", 2, 1},

		{"991.992", "8", "123", "999", 3, 1},
		{"991.992", "8", "124", "00", 2, 1},
	}

	for _, test_ip := range cases {
		t.Run(fmt.Sprintf("%s/%s=%s.%s", test_ip.dividend, test_ip.divisor, test_ip.q, test_ip.r), func(t *testing.T) {
			divid, divis := get_input(test_ip.dividend, test_ip.divisor)
			q, r, err := run_division2(divid, divis, test_ip.maxDecCount, test_ip.rounding)
			if err != nil {
				t.Log(err.Error())
				t.Fail()
			}
			assert.Equal(t, test_ip.q, q.String())
			assert.Equal(t, test_ip.r, r)
		})
	}
}

func TestRounding(t *testing.T) {
	type TC struct {
		a, ans string
		mode   int
		ov     bool
	}

	cases := []TC{
		{"0025", "003", 1, false},
		{"095", "10", 1, false},
		{"995", "00", 1, true},

		{"0026", "003", 1, false},
		{"096", "10", 1, false},
		{"996", "00", 1, true},

		{"0024", "002", 1, false},
		{"094", "09", 1, false},
		{"994", "99", 1, false},

		{"0000", "000", 2, false},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprint(tc.a, tc.mode), func(t *testing.T) {
			ans, ov := rounding(tc.a, tc.mode)
			assert.Equal(t, tc.ans, ans)
			assert.Equal(t, tc.ov, ov)
		})
	}
}
