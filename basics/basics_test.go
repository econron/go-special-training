package basics

import (
	"testing"
	"reflect"
)

func TestStringDistance(t *testing.T) {
	got := StringDistance("foo", "foh")
	want := 1
	if got != want {
		t.Fatalf("want %v, got %v:", want, got)
	}
}

func TestStringDistanceByTable(t *testing.T){
	tests := []struct {
		name string
		lhs string
		rhs string
		want int
	}{
		{name: "lhsはrhsより長い", lhs: "foo", rhs: "fo", want: -1},
		{name: "lhsはrhsより短い", lhs: "fo", rhs: "foo", want: -1},
		{name: "lhsとrhsが同じ", lhs: "foo", rhs: "foo", want: 0},
		{name: "lhsとrhsが1文字異なる", lhs: "foo", rhs: "foh", want: 1},
		{name: "lhsとrhsが2文字異なる", lhs: "foo", rhs: "fgh", want: 2},
	}

	for _, tc := range tests {
		got := StringDistance(tc.lhs, tc.rhs)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("%s: expected %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func BenchmarkStringDistance(b *testing.B){
	for i := 0; i < b.N; i++ {
		StringDistance("foo", "foh")
	}
}