package piscine

func CanJump(steps []uint) bool {
	if len(steps) == 0 {
		return false
	}

	reach := 0
	for i := 0; i < len(steps); i++ {
		if i > reach {
			return false
		}
		reach = max(reach, i+int(steps[i]))

		if reach == len(steps)-1 || len(steps) == 1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
