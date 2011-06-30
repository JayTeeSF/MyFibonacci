package MyFibonacci

type fibCommander struct {
  previous, current, next (func() int)
}

func Fibber() fibCommander {
  previous := 0
  current := 1

  return fibCommander {
    func() int {
      return previous
    },
    func() int {
      return current
    },
    func() int {
      var next = previous + current
      previous = current
      current = next
      return next
    },
  }
}
