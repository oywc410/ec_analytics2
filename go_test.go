package ec_analytics2
import (
	"testing"
	"fmt"
)

func BenchmarkTemplateParallel(b *testing.B) {

	b.RunParallel(func(pb *testing.PB){

		for pb.Next() {
			fmt.Println("aaaa")
		}
	})
}