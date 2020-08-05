package bouncing_balls

func BouncingBall(h, bounce, window float64) int {
	if h < 0 || bounce < 0 || bounce >= 1 || window >= h {
		return -1
	}

	ballPasses := 0
	for {
		ballPasses++
		if h = h * bounce; h <= window {
			break
		}
		ballPasses++
	}

	return ballPasses
}
