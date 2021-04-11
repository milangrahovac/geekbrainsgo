package fibonacci

var cache = map[int64]int64{0: 0, 1: 1}

func Fib(i int64) int64 {
	if fb, ok := cache[i]; ok {
		return fb
	}

	fbm1, ok := cache[i-1]
	if !ok {
		fbm1 = Fib(i - 1)
	}

	fbm2, ok := cache[i-2]
	if !ok {
		fbm2 = Fib(i - 2)
	}

	cache[i] = fbm1 + fbm2
	return cache[i]
}
