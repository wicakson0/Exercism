package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	numStep := 0

	if n < 1 {
		return 0, errors.New("invalid number")
	}

	for n > 1 {
		if n%2 == 0 {
			n = n / 2
			numStep++
		} else {
			n = (3 * n) + 1
			numStep++
		}
	}

	return numStep, nil
}
