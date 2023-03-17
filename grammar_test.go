package main

import "testing"

func TestCheck(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				sentence: "abcd abcd aabbc ab a c ccd dede cccd cd",
			},
			wantCode: 861,
			wantErr:  false,
		},
		{
			name: "set 1",
			args: args{
				sentence: "abab ab de ee",
			},
			wantErr: true,
		},
		{
			name: "set 2",
			args: args{
				sentence: "dede abcd abd abddd ddada dac abcd de ed",
			},
			wantCode: 2656,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, err := grammerCheck(tt.args.sentence)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCode != tt.wantCode {
				t.Errorf("Check() = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}
