package utils

import (
	"testing"
)

func TestAddition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "success",
			args: args{
				nums: []int{5, 7},
			},
			wantResult: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Addition(tt.args.nums); gotResult != tt.wantResult {
				t.Errorf("Sum() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSubtraction(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "success",
			args: args{
				nums: []int{15, 8},
			},
			wantResult: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Subtraction(tt.args.nums); gotResult != tt.wantResult {
				t.Errorf("Subtraction() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMultiplication(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "success",
			args: args{
				nums: []int{5, 4, 3},
			},
			wantResult: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Multiplication(tt.args.nums); gotResult != tt.wantResult {
				t.Errorf("Multiplication() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
