package main

//func main() {
//		s:="babd"
//		fmt.Println(longestPalindrome(s))
//}

// Brute Force
func longestPalindrome1(s string) string {
	length:=len(s)
	if length<2{
		return s
	}
	maxLen:=1
	begin := 0
	for i:=0;i<length - 1;i++{
		for j:=i+1;j<length;j++{
			if j-i+1>maxLen && isPalindrome(s, i, j) {
				maxLen = j-i+1
				begin = i
			}
		}
	}
	return s[begin:begin+maxLen]
}

func isPalindrome(s string, i,j int) bool {
	for;i<j;{
		if s[i:i+1] != s[j:j+1] {
			return false
		}
		i++
		j--
	}
	return true
}

// DP，动态规划
func longestPalindrome2(s string) string {
	l:=len(s)
	if l <2 {
		return s
	}

	begin,maxLen:=0,0
	dp := make([][]bool, l)
	for i:=0;i<l;i++{
		dp[i]=make([]bool, l)
		dp[i][i] = true
	}
	for j:=1;j<l;j++{
		for i:=0; i<j;i++{
			if s[i:i]!=s[j:j]{
				dp[i][j]=false
			} else {
				if j-i<3{
					dp[i][j]=true
				} else{
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1>maxLen {
				begin = i
				maxLen = j-i+1
			}
		}
	}
	return s[begin:begin+maxLen]
}

// 中心扩散法
func longestPalindrome(s string) string {
	res := ""
	l := len(s)
	maxLen := 1
	longer :=0
	longgerStr:=""
	for i:=0;i<l-1;i++{
		oddStr:=centerSpread(s, i, i)
		oddLen := len(oddStr)
		evenStr := centerSpread(s, i, i+1)
		evenLen := len(evenStr)

		if oddLen > evenLen {
			longer = oddLen
			longgerStr=oddStr
		} else {
			longer = evenLen
			longgerStr=evenStr
		}
		if longer >= maxLen {
			res = longgerStr
			maxLen = longer
		}

	}
	return  res
}

func centerSpread(s string, left, right int) string {
	l:=len(s)
	i,j:=left, right
	for i>=0 && j<l{
		if s[i:i+1] == s[j:j+1] {
			i--
			j++
		} else {
			break
		}
	}

	return s[i+1:j]
}
