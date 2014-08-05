package qc

import (
	"strings"
	"testing"

	"github.com/ToQoz/gopwt/assert"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in: "AAA",
			out: strings.TrimSpace(`
＿人人人＿
＞　AAA　＜
￣Y^Y^Y^￣`),
		},
		{
			in: "突然の死",
			out: strings.TrimSpace(`
＿人人人人人人＿
＞　突然の死　＜
￣Y^Y^Y^Y^Y^Y^￣`),
		},
		{
			in: "突然death",
			out: strings.TrimSpace(`
＿人人人人人人＿
＞　a突然death　＜
￣Y^Y^Y^Y^Y^Y^￣`),
		},
	}

	for _, test := range tests {
		s := Format(test.in)
		assert.OK(t, s == test.out)
	}
}
