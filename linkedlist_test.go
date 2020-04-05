package linkedCycle

import "testing"

func TestDetectCycleWithHash(t *testing.T) {
	type args struct {
		l *List
	}

	// prepare lists
	var ls []*List
	ls = append(ls, FromArray([]int{1, 2}, -1))
	ls = append(ls, FromArray([]int{1, 2}, 0))
	ls = append(ls, FromArray([]int{}, -1))
	ls = append(ls, FromArray([]int{1, 2, 3}, 1))

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"small list", args{ls[0]}, false},
		{"small list", args{ls[1]}, true},
		{"empty list", args{ls[2]}, false},
		{"small list", args{ls[3]}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectCycleWithHash(tt.args.l); got != tt.want {
				t.Errorf("DetectCycleWithHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetectCycleTwoPointers(t *testing.T) {
	type args struct {
		l *List
	}
	// prepare lists
	var ls []*List
	ls = append(ls, FromArray([]int{1, 2}, -1))
	ls = append(ls, FromArray([]int{1, 2}, 0))
	ls = append(ls, FromArray([]int{}, -1))
	ls = append(ls, FromArray([]int{1, 2, 3}, 1))

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"small list", args{ls[0]}, false},
		{"small list", args{ls[1]}, true},
		{"empty list", args{ls[2]}, false},
		{"small list", args{ls[3]}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectCycleTwoPointers(tt.args.l); got != tt.want {
				t.Errorf("DetectCycleTwoPointers() = %v, want %v", got, tt.want)
			}
		})
	}
}

const arrSz = 1000

func BenchmarkDetectCycleWithHash(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	var arr []int
	for i := 0; i < arrSz; i++ {
		arr = append(arr, i)
	}
	l := FromArray(arr, arrSz/2) // point in the middle
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		got := DetectCycleWithHash(l)
		if got != true {
			b.Errorf("incorrect result")
		}
	}
}

func BenchmarkDetectCycleTwoPointers(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	var arr []int
	for i := 0; i < arrSz; i++ {
		arr = append(arr, i)
	}
	l := FromArray(arr, arrSz/2) // point in the middle
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		got := DetectCycleTwoPointers(l)
		if got != true {
			b.Errorf("incorrect result")
		}
	}
}
