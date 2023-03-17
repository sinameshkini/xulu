package main

import "testing"

func TestNameToNumber(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name       string
		args       args
		wantNumber int
	}{
		{
			name: "success",
			args: args{
				name: "cccc",
			},
			wantNumber: 4,
		},
		{
			name: "success multi",
			args: args{
				name: "cccd",
			},
			wantNumber: 32,
		},
		{
			name: "success 2",
			args: args{
				name: "aabbc",
			},
			wantNumber: 29,
		},
		{
			name: "success 3",
			args: args{
				name: "ab",
			},
			wantNumber: 5,
		},
		{
			name: "success 4",
			args: args{
				name: "c",
			},
			wantNumber: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNumber := NameToNumber(tt.args.name); gotNumber != tt.wantNumber {
				t.Errorf("NameToNumber() = %v, want %v", gotNumber, tt.wantNumber)
			}
		})
	}
}
