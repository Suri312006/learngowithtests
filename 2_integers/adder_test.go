package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
    sum := Add(2, 2)
    expected := 4

    if sum != expected {
        // %d prints integer while %q is for stirngs
        t.Errorf("Expected '%d' but got '%d'", expected, sum)
    }
}


//examples 

// THESE EXAMPLES GET COMPILED TOOO!!!!!!!!!!!!!
func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}
