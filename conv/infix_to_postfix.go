package conv

import (
	"main/stack"
	"strings"
)

var mapper = map[rune][2]int{
	'^': {6, 5},
	'*': {3, 4},
	'/': {3, 4},
	'+': {1, 2},
	'-': {1, 2},
	'(': {7, 0},
	')': {0, 0},
}

func ConvertInfixToPostfix(expr string) string {
	var sb strings.Builder
	st := stack.New()
	st.Push('(')

	for _, ch := range expr {
		if _, ok := mapper[ch]; !ok {
			sb.WriteRune(ch)
		} else {
			topElem := st.Pop()
			if ch == ')' {
				for topElem != '(' {
					sb.WriteRune(topElem)
					topElem = st.Pop()
				}
			} else if icp(ch) > isp(topElem) {
				st.Push(topElem)
				st.Push(ch)
			} else {
				for icp(ch) <= isp(topElem) {
					sb.WriteRune(topElem)
					topElem = st.Pop()
				}
				st.Push(topElem)
				st.Push(ch)
			}
		}
	}

	for st.Size() > 1 {
		sb.WriteRune(st.Pop())
	}

	return sb.String()
}

func isp(ch rune) int {
	return mapper[ch][1]
}

func icp(ch rune) int {
	return mapper[ch][0]
}
