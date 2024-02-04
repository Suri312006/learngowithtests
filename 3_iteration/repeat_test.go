package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
    repeated := Repeat("a", 5)
    expected := "aaaaa"

    if repeated != expected {
        t.Errorf("expected %q but got %q ", expected, repeated)
    }
}

// benchmark to see how fast funciton is lmao
func BenchmarkRepeat(b *testing.B){
    for i :=0; i < b.N; i++ {
        Repeat("a", 10)
    }
}

func ExampleRepeat(){
    repeated := Repeat("lmao", 5)
    fmt.Println(repeated)
    // Output: lmaolmaolmaolmaolmao
}
