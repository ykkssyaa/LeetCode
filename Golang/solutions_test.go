package main

import (
	"testing"
)

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

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 ", args{"   fly me   to   the moon  "}, 4},
		{"2 ", args{"luffy is still joyboy"}, 6},
		{"3 ", args{"Hello World"}, 5},
		{"1 ", args{"  a  a  aa aaa a    "}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toLowerCase1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1 ", args{"Hello World!"}, "hello world!"},
		{"2 ", args{"AaAa _aA"}, "aaaa _aa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerCase(tt.args.s); got != tt.want {
				t.Errorf("toLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tictactoe(t *testing.T) {
	type args struct {
		moves [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1. ", args{[][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}}, "A"}, // [[0,0],[2,0],[1,1],[2,1],[2,2]]
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tictactoe(tt.args.moves); got != tt.want {
				t.Errorf("tictactoe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lemonadeChange(t *testing.T) {
	type args struct {
		bills []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{5, 5, 5, 10, 20}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lemonadeChange(tt.args.bills); got != tt.want {
				t.Errorf("lemonadeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestPerimeter(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 1, 10}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPerimeter(tt.args.nums); got != tt.want {
				t.Errorf("largestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkStraightLine(t *testing.T) {
	type args struct {
		coordinates [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[][]int{{1, -8}, {2, -3}, {1, 2}}}, false},
		{"2", args{[][]int{{4, 8}, {-2, 8}, {1, 8}, {8, 8}, {-5, 8}, {0, 8}, {7, 8}, {5, 8}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkStraightLine(tt.args.coordinates); got != tt.want {
				t.Errorf("checkStraightLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{a: "11", b: "1"}, "100"},
		{"1", args{a: "1010", b: "1011"}, "10101"},
		{"1", args{a: "111", b: "11"}, "1010"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
