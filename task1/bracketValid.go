package main

type stack struct {
	elements []string
}

func (s *stack) Put(item string) {
	s.elements = append(s.elements, item)
}

func (s *stack) Pop() (string, bool) {
	if len(s.elements) == 0 {
		return "", false
	}
	item := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return item, true
}

func checkMatch(str1 string, str2 string) bool {
	switch str1 {
	case "(":
		return str2 == ")"
	case "[":
		return str2 == "]"
	case "{":
		return str2 == "}"
	default:
		return false
	}
}

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	st := stack{}

	for _, v := range s {
		char := string(v)
		if char == "(" || char == "[" || char == "{" {
			st.Put(char)
		} else {
			top, ok := st.Pop()
			if !ok || !checkMatch(top, char) {
				return false
			}
		}
	}

	return len(st.elements) == 0
}
