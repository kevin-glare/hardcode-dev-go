package hw

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestSortInts(t *testing.T) {
	got := []int{1, 5, 2, 4, 3}
	sort.Ints(got)
	want := []int{1, 2, 3, 4, 5}
	if !checkSliceSort(got, want) {
		t.Errorf("получили %d, ожидалось %d", got, want)
	}
}

func TestSortStrings(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "Тест №1",
			args: []string{"a", "b", "c", "d"},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "Тест №2",
			args: []string{""},
			want: []string{""},
		},
		{
			name: "Тест №3",
			args: []string{"d", "c", "b", "a"},
			want: []string{"a", "b", "c", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args
			sort.Strings(got)
			if !checkSliceSort(got, tt.want) {
				t.Errorf("получили %v, ожидалось %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSortInts-16    	 6055674	       202.3 ns/op

func BenchmarkSortInts(b *testing.B) {
	data := sampleDataInts()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000)
		res := sort.SearchInts(data, n)
		_ = res
	}
}

// BenchmarkSortFloat64s-16    	 5183310	       214.8 ns/op

func BenchmarkSortFloat64s(b *testing.B) {
	data := sampleDataFloats()

	for i := 0; i < b.N; i++ {
		n := rand.Float64()
		res := sort.SearchFloat64s(data, n)
		_ = res
	}
}

// Пытался сделать универсальный генератор, но не вышло.

//func sampleData[N int | float64] (t string) []N {
//	rand.Seed(time.Now().UnixNano())
//	var data []N
//
//	if t == "int" {
//		data = append(data, rand.Intn(1000))
//	} else {
//		for i := 0; i < 1_000_000; i++ {
//			data = append(data, rand.Float64())
//		}
//	}
//
//	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
//
//	return data
//}

func sampleDataFloats() []float64 {
	rand.Seed(time.Now().UnixNano())
	var data []float64

	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Float64())
	}

	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })

	return data
}

func sampleDataInts() []int {
	rand.Seed(time.Now().UnixNano())

	var data []int

	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Intn(1000))
	}

	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })

	return data
}

// Стало интересно потрогать дженерики, кажется, что тут хорошо подходят, апнул версию го.

func checkSliceSort[V int | string](slice1, slice2 []V) bool {
	for i, val := range slice1 {
		if val != slice2[i] {
			return false
		}
	}

	return true
}
