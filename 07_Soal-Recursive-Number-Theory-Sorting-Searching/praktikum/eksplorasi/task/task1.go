package task

func GetDiamondSeq(n int) int {
	return GetSequal(n) + GetSequal(n-1)
}

func GetSequal(n int) int {
	if n <= 1 {
		return n
	}

	return n + GetSequal(n-1)
}
