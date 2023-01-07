package main

func check(s string) (int, int) {
	queue := make([]rune, 0)

	for _, r := range s {
		switch r {
		case '(':
			queue = append(queue, ')')
		case '[':
			queue = append(queue, ']')
		case '{':
			queue = append(queue, '}')
		case '<':
			queue = append(queue, '>')
		default:
			if r != queue[len(queue)-1] {
				switch r {
				case ')':
					return 3, 0
				case ']':
					return 57, 0
				case '}':
					return 1197, 0
				case '>':
					return 25137, 0
				}
			}

			queue = queue[:len(queue)-1]
		}
	}

	result := 0

	for i := len(queue) - 1; i >= 0; i-- {
		result *= 5

		switch queue[i] {
		case ')':
			result += 1
		case ']':
			result += 2
		case '}':
			result += 3
		case '>':
			result += 4
		}
	}

	return 0, result
}
