package puuid

import (
	"fmt"
	"testing"
)

func Test_Generate(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(Generate())
	}
}
