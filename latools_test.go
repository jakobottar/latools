package latools

import (
	"math"
	"testing"
)

func TestLength(t *testing.T) {
	type args struct {
		v []float64
	}
	tests := []struct {
		name  string
		args  args
		wantD float64
	}{
		{
			name:  "zero",
			args:  args{[]float64{0, 0}},
			wantD: 0,
		},
		{
			name:  "sqrt(2)",
			args:  args{[]float64{1, 1}},
			wantD: math.Sqrt(2),
		},
		{
			name:  "nD",
			args:  args{[]float64{0, 0, 0, 0, 10, 0}},
			wantD: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotD := Length(tt.args.v); gotD != tt.wantD {
				t.Errorf("Length() = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}
