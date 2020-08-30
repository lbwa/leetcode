package easy657

// 关键在于量化机器人的执行步骤
func judgeCircle(moves string) bool {
	horizontal, vertical := 0, 0
	for _, move := range moves {
		switch move {
		case 'U':
			vertical++
		case 'D':
			vertical--
		case 'L':
			horizontal++
		case 'R':
			horizontal--
		}
	}
	return (horizontal == 0) && (vertical == 0)
}
