package helpers

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestSliceInt(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	type testCase[V constraints.Integer] struct {
		name       string
		args       args
		wantResult []V
		wantErr    bool
	}
	tests := []testCase[int]{
		{
			name: "test int",
			args: args{
				s:   "1,2,3,4",
				sep: ",",
			},
			wantResult: []int{1, 2, 3, 4},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := SliceIntFromString[int](tt.args.s, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("SliceIntFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("SliceIntFromString() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSliceInt64(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	type testCase[V constraints.Integer] struct {
		name       string
		args       args
		wantResult []V
		wantErr    bool
	}
	tests := []testCase[int64]{
		{
			name: "test int64",
			args: args{
				s:   "1,2,3,4",
				sep: ",",
			},
			wantResult: []int64{1, 2, 3, 4},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := SliceIntFromString[int64](tt.args.s, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("SliceIntFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("SliceIntFromString() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSliceIntToString(t *testing.T) {
	type args[V constraints.Integer] struct {
		values []V
		sep    string
	}
	type testCase[V constraints.Integer] struct {
		name       string
		args       args[int]
		wantResult string
	}
	tests := []testCase[int]{
		{
			name: "test int",
			args: args[int]{
				values: []int{1, 2, 3, 4, 5},
				sep:    ",",
			},
			wantResult: "1,2,3,4,5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := SliceIntToString[int](tt.args.values, tt.args.sep)
			if gotResult != tt.wantResult {
				t.Errorf("SliceIntToString() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
