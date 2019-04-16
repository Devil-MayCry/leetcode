/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (36.23%)
 * Total Accepted:    559.7K
 * Total Submissions: 1.5M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output: true
 *
 *
 */

package main

type stack struct {
	s []rune
}

func newStack() *stack {
	s := &stack{}
	s.s = make([]rune, 0)
	return s
}

func (s *stack) push(b rune) {
	s.s = append(s.s, b)
}

func (s *stack) pop() rune {
	top := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return top
}

func (s *stack) peek() rune {
	return s.s[len(s.s)-1]
}

func (s *stack) len() int {
	return len(s.s)
}

func isValid(s string) bool {
	arr := []rune(s)
	if len(arr) == 0 {
		return true
	}
	if len(arr)%2 == 1 {
		return false
	}
	st := newStack()
	m := map[rune]int{'(': -3, ')': 3, '{': -2, '}': 2, '[': -1, ']': 1}
	for _, v := range arr {
		if m[v] < 0 {
			st.push(v)
			continue
		}
		if st.len() == 0 {
			return false
		}
		top := st.peek()
		if m[v]+m[top] == 0 {
			st.pop()
		} else {
			return false
		}
	}
	if st.len() == 0 {
		return true
	}
	return false
}

// func main() {
// 	fmt.Println(isValid(`"()[]{}"`))
// }
