package pngGen

import (
	"testing"
)

func Test_Generate(t *testing.T) {
	path, err := Generate(500, 500)
	if err != nil {
		t.Error(err)
	}
	_ = err
	t.Log("Sucess: ", path)
}
