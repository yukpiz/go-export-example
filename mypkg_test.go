package mypkg_test

import (
	"testing"

	mypkg "github.com/yukpiz/go-export-example"
)

func Test_hello(t *testing.T) {
	mypkg.Export_hello()
}
