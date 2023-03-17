package main

import (
	"reflect"
	"testing"
)

func TestParseSentenceWithFollowing(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name         string
		args         args
		wantSentence *Sentence
		wantErr      bool
	}{
		{
			name: "success",
			args: args{
				str: "bcde ab ac",
			},
			wantSentence: &Sentence{
				Verb:         "bcde",
				Names:        []string{"ab", "ac"},
				SubSentences: nil,
			},
			wantErr: false,
		},
		{
			name: "success multiple",
			args: args{
				str: "bcde ab ac abcd a b",
			},
			wantSentence: &Sentence{
				Verb:      "bcde",
				Names:     []string{"ab", "ac"},
				following: "abcd a b",
			},
			wantErr: false,
		},
		{
			name: "success composite",
			args: args{
				str: "abcd bcde ab ac abcd a b",
			},
			wantSentence: &Sentence{
				Verb:      "abcd",
				following: "bcde ab ac abcd a b",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSentences, err := ParseSentenceWithFollowing(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSentence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSentences, tt.wantSentence) {
				t.Errorf("ParseSentence() = %v, want %v", gotSentences, tt.wantSentence)
			}
		})
	}
}

func TestParseSentence(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name         string
		args         args
		wantSentence *Sentence
		wantErr      bool
	}{
		{
			name: "success composite",
			args: args{
				str: "abcd bcde ab ac abcd a b",
			},
			wantSentence: &Sentence{
				Verb: "abcd",
				// following: "bcde ab ac abcd a b",
				SubSentences: []*Sentence{
					{
						Verb:  "bcde",
						Names: []string{"ab", "ac"},
					},
					{
						Verb:  "abcd",
						Names: []string{"a", "b"},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSentence, err := ParseSentence(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSentence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSentence, tt.wantSentence) {
				t.Errorf("ParseSentence() = %v, want %v", gotSentence, tt.wantSentence)
			}
		})
	}
}
