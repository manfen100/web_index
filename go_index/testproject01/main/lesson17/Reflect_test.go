package main

import (
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {

	var f int64 = 10

	t.Log(reflect.TypeOf(f))

	t.Log(reflect.ValueOf(f))

}
