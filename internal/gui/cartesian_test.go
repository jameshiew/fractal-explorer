package gui

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_toCartesian(t *testing.T) {
	const (
		width  = 320
		height = 240
	)
	for _, tc := range []struct {
		x         int
		y         int
		expectedX float64
		expectedY float64
	}{
		{0, 0, -160, 120},
		{width, height, 160, -120},
	} {
		t.Run(fmt.Sprintf("%v, %v", tc.x, tc.y), func(t *testing.T) {
			x, y := ToCartesian(tc.x, tc.y, width, height)

			require.EqualValues(t, tc.expectedX, x)
			require.EqualValues(t, tc.expectedY, y)
		})
	}
}
