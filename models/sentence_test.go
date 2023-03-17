package models

import (
	"reflect"
	"testing"
)

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
			name: "success",
			args: args{
				str: "bcde ab ac",
			},
			wantSentence: &Sentence{
				Verb:         "bcde",
				Names:        []string{"ab", "ac"},
				SubSentences: nil,
			},
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
