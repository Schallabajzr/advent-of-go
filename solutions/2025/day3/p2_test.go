package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestP2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "test",
			input: "987654321111111\n811111111111119\n234234234234278\n818181911112111",
			want:  "3121910778619",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := pt2(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.want, res)
		})
	}
}
