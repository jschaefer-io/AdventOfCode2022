package day11

func gcd(a, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b uint64, list ...uint64) uint64 {
	result := a * b / gcd(a, b)
	for i := 0; i < len(list); i++ {
		result = lcm(result, list[i])
	}
	return result
}
