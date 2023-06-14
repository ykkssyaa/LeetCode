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
