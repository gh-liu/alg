package sort

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{
			name:  "int",
			array: []int{1, 5, 3, 4, 2, 7, 6},
			want:  []int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeIter(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{
			name:  "int",
			array: []int{1, 5, 3, 4, 2, 7, 6},
			want:  []int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeIter(tt.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeIter() = %v, want %v", got, tt.want)
			}
		})
	}
}
