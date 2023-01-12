package offer

//面试题17：包含所有字符的最短字符串
//题目：输入两个字符串s和t，请找出字符串s中包含字符串t的所有字符的最短子字符串。例如，输入的字符串s为"ADDBANCAD"，
//字符串t为"ABC"，则字符串s中包含字符'A'、'B'和'C'的最短子字符串是"BANC"。如果不存在符合条件的子字符串，则返回空字符串""。
//如果存在多个符合条件的子字符串，则返回任意一个。

func MinWindow(s, t string) string {
	m := make(map[byte]int, len(t))
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}
	count := len(t)

	ret := ""
	for left, right := 0, 0; right < len(s); right++ {
		cRight := s[right]
		if _, ok := m[cRight]; !ok {
			continue
		}
		m[cRight]--
		count--

		for m[cRight] < 0 || count == 0 {
			if count == 0 && (ret == "" || right-left+1 < len(ret)) {
				ret = s[left : right+1]
			}
			cLeft := s[left]
			if _, ok := m[cLeft]; ok {
				m[cLeft]++
				count++
			}
			left++
		}
	}

	return ret
}
