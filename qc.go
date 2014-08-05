package qc

import (
	"fmt"
	"math"
	"strings"

	"github.com/moznion/go-text-visual-width"
)

// Format string
func Format(s string) string {
	l := (int)(math.Ceil((float64)(visualwidth.Width(s) / 2)))

	format := `
＿人%s人＿
＞　%s　＜
￣Y^%sY^￣`

	return fmt.Sprintf(
		strings.TrimSpace(format),
		strings.Repeat("人", l),
		s,
		strings.Repeat("Y^", l),
	)
}
