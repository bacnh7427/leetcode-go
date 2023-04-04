func decodeString(s string) string {
	left := []string{}
	ans := ""
	c := 0
	tem := []int{0}
	for i := range s {
		if s[i:i+1] == "[" {
			left = append(left, "")
			c = 1
			continue
		}

		if s[i:i+1] == "]" {
			tt := Str(left[len(left)-1], tem[len(tem)-1])
			left = left[:len(left)-1]
			tem = tem[:len(tem)-1]
			c = 1
			if len(left) == 0 {
				ans += tt
			} else {
				left[len(left)-1] += tt
			}
			continue
		}

		if s[i:i+1] >= "0" && s[i:i+1] <= "9" {
			k, _ := strconv.Atoi(s[i : i+1])
			if c == 0 {
				tem[len(tem)-1] = 10*tem[len(tem)-1] + k
			} else {
				c = 0
				tem = append(tem, k)
			}
			continue
		} else {
			if len(left) > 0 {
				left[len(left)-1] += s[i : i+1]
			} else {
				ans += s[i : i+1]
			}
			c = 1
		}
	}
	return ans
}

func Str(s string, i int) string {
	ans := ""
	for j := 0; j < i; j++ {
		ans += s
	}
	return ans
}