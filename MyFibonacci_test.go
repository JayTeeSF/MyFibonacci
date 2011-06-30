package MyFibonacci
import "testing"

func TestFibber(t *testing.T) {
  fib := Fibber()

  expecteds := []int{1,2,3,5,8}

  for i := 0; i < len(expecteds); i++ {
    got := fib.next()
    if got != expecteds[i] {
      t.Errorf("expected %d, got %d", expecteds[i], got)
    }
  }
}
