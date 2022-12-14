package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	is := assert.New(t)

	r := Intersect([]int{1, 2, 3, 4}, []int{2, 4, 6})
	is.Equal(r, []int{2, 4})

	r = Intersect([]string{"foo", "bar", "hello", "bar"}, []string{"foo", "bar"})
	is.Equal(r, []string{"foo", "bar"})

	r = Intersect([]string{"foo", "bar"}, []string{"foo", "bar", "hello", "bar"})
	is.Equal(r, []string{"foo", "bar", "bar"})
}

func TestIntersectString(t *testing.T) {
	is := assert.New(t)

	r := IntersectString([]string{"foo", "bar", "hello", "bar"}, []string{"foo", "bar"})
	is.Equal(r, []string{"foo", "bar"})
}

func TestDifference(t *testing.T) {
	is := assert.New(t)

	r1, r2 := Difference([]int{1, 2, 3, 4}, []int{2, 4, 6})
	is.Equal(r1, []int{1, 3})
	is.Equal(r2, []int{6})

	r1, r2 = Difference([]string{"foo", "bar", "hello", "bar"}, []string{"foo", "bar"})
	is.Equal(r1, []string{"hello"})
	is.Equal(r2, []string{})

	r1, r2 = Difference(map[string]string{"foo": "bar", "hello": "baz"}, map[string]string{"foo": "bar", "bar": "baz"})
	is.Equal(r1, map[string]string{"hello": "baz"})
	is.Equal(r2, map[string]string{"bar": "baz"})
}

func TestDifferenceString(t *testing.T) {
	is := assert.New(t)

	r1, r2 := DifferenceString([]string{"foo", "bar", "hello", "bar"}, []string{"foo", "bar"})
	is.Equal(r1, []string{"hello"})
	is.Equal(r2, []string{})
}

func TestDifferenceInt64(t *testing.T) {
	is := assert.New(t)

	r1, r2 := DifferenceInt64([]int64{1, 2, 3, 4}, []int64{4, 6})
	is.Equal(r1, []int64{1, 2, 3})
	is.Equal(r2, []int64{6})
}
