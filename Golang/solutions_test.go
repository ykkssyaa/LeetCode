package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{"X"}, want: 10},
		{name: "1", args: args{"III"}, want: 3},
		{name: "1", args: args{"XXIII"}, want: 23},
		{name: "1", args: args{"LVIII"}, want: 58},
		{name: "1", args: args{"MCMXCIV"}, want: 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
