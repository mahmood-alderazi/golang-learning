package piscine

func CanJump(input []uint) bool {
	if len(input) == 1 {
		return true
	}

	if len(input) == 0 {
		return false
	}

	maxreach := 0

	for i := 0; i < len(input); i++ {
		if i > maxreach {
			return false
		}

		jump := i + int(input[i])

		if jump > maxreach {
			maxreach = jump
		}

		if maxreach >= len(input)-1 {
			return true
		}
	}
	return false
}
