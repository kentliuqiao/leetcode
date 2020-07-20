package main

func main() {
	
}

func maxArea(height []int) int {
	i,j:=0,len(height)-1
	res:=0
	for ;i<j;{
		tmp:=min(height[i], height[j]) * (j-i)
		res = max(res, tmp)
	}
	return res
}

func max(x,y int) int {
	if x<y {
		return y
	}
	return x
}

func min(x,y int) int {
	if x<y{
		return x
	}
	return  y
}
