package conv

import "main/stack"

func CheckBalancedExpr(expr string) bool {
	st := stack.New()
	for _, ch := range expr {
		if ch == '(' || ch == '{' || ch == '[' {
			st.Push(ch)
		} else if ch == ')' || ch == '}' || ch == ']' {
			elem := st.Pop()
			if (ch == ')' && elem == '(') || (ch == '}' && elem == '{') || (ch == ']' && elem == '[') {
				continue
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return st.Size() == 0
}
