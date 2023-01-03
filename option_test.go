package opt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func ExampleSome_Get() {
	val := Some(2)
	fmt.Println(val.Get())
	// Output: 2
}

func TestSome_Get(t *testing.T) {
	val := Some(2)
	result := val.Get()
	assert.Equal(t, 2, result)
}

func ExampleSome_GetOrElse() {
	val := Some("test")
	fmt.Println(val.GetOrElse("test2"))
	// Output: test
}

func TestSome_GetOrElse(t *testing.T) {
	val := Some("test")
	result := val.GetOrElse("test2")
	assert.Equal(t, "test", result)
}

func ExampleSome_GetOrZero() {
	val := Some(true)
	fmt.Println(val.GetOrZero())
	// Output: true
}

func TestSome_GetOrZero(t *testing.T) {
	val := Some(true)
	result := val.GetOrZero()
	assert.Equal(t, true, result)
}

func ExampleSome_Empty() {
	val := Some("value")
	fmt.Println(val.Empty())
	// Output: false
}

func TestSome_Empty(t *testing.T) {
	val := Some(time.Now())
	assert.Equal(t, false, val.Empty())
}

func ExampleNone_Get() {
	noneVal := None[string]()
	fmt.Println(noneVal.Empty())
	// noneVal.Get() panics!!
	// Output: true
}

func TestNone_Get(t *testing.T) {
	noneVal := None[int]()
	assert.PanicsWithValue(t, "None has no value", func() {
		noneVal.Get()
	}, "Attempting to get from none should panic")
}

func ExampleNone_GetOrElse() {
	val := None[int]()
	fmt.Println(val.GetOrElse(3))
	// Output: 3
}

func TestNone_GetOrElse(t *testing.T) {
	nameOpt := None[string]()
	val := nameOpt.GetOrElse("name")
	assert.Equal(t, "name", val)
}

func ExampleNone_GetOrZero() {
	nameOpt := None[string]()
	fmt.Println(nameOpt.GetOrZero())
	// Output:
}

func TestNone_GetOrZero(t *testing.T) {
	nameOpt := None[string]()
	name := nameOpt.GetOrZero()
	assert.Zero(t, name)
}

func ExampleNone_Empty() {
	nameOpt := None[string]()
	fmt.Println(nameOpt.Empty())
	// Output: true
}

func TestNone_Empty(t *testing.T) {
	strOpt := None[string]()
	assert.Equal(t, true, strOpt.Empty())
}

func ExampleMap_Some() {
	//with value
	valOpt := Some(3)
	twiceOpt := Map(func(t int) int {
		return t * 2
	}, valOpt)
	fmt.Println(twiceOpt.Get())
	// Output: 6
}

func TestMap_Some(t *testing.T) {
	intOpt := Some(2)
	twiceOpt := Map(func(t int) int {
		return t * 2
	}, intOpt)
	assert.Equal(t, 4, twiceOpt.Get())
}

func TestMap_None(t *testing.T) {
	noneOpt := None[int]()
	twiceOpt := Map(func(t int) int {
		return t * 2
	}, noneOpt)
	assert.Equal(t, true, twiceOpt.Empty())
}
