package main

//func main() {
//
//}
func isValid(s string) bool {
	 m := map[byte]byte{
		'(':')',
		'[':']',
		'{':'}',
	}

	l:=len(s)
	if l%2!=0{
		return false
	}

	lps:=make([]byte, 0, l)

	for i:=0;i<l;i++{
		if _,ok:=m[s[i]];ok{
			lps = append(lps, s[i])
		} else {
			if len(lps) > 0 && m[lps[len(lps) - 1]] == s[i] {
				lps = lps[:len(lps) - 1]
			}
			return false
		}

	}

	if len(lps) == 0{
		return  true
	}
	return false
}
