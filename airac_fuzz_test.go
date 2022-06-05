// +build go1.18

package airac

import (
	"strconv"
	"strings"
	"testing"
)

//goland:noinspection GoUnusedExportedFunction
func FuzzFromString(f *testing.F) {
	for i := -999; i <= 9999; i++ {
		f.Add(strconv.Itoa(i))
	}

	f.Fuzz(func(t *testing.T, s string) {
		t.Parallel()

		a, err := FromString(s)
		if err != nil {
			t.SkipNow()
		}

		s2 := a.String()
		if strings.TrimSpace(s) != s2 {
			t.Errorf("%q != %q", s, s2)
		}
	})
}
