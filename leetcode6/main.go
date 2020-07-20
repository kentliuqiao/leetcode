package main

//func main() {
//
//}

func convert(s string, numRows int) string {
	if numRows < 2{
		return s
	}

	tmp:=make([]string, numRows)
	flag := -1
	curPos:=0
	for _, val:=range s {
		tmp[curPos] += string(val)
		if curPos == 0|| curPos==numRows -1 {
			flag = -flag
		}
		curPos += flag
	}

	res:=""
	for _, val:=range tmp{
		res+=val
	}
	return res
}