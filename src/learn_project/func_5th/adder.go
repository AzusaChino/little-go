package func_5th

func Adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func Adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, Adder2(base + v)
	}
}
