package main

import (
	"fmt"
	"net/url"
	"sort"
)

func main() {
	u,err:=url.ParseRequestURI("www.baidu.com")
	fmt.Printf("url=%+v, err=%+v", u, err)
}

func threeSum(nums []int) [][]int {
	l:=len(nums)
	res:=make([][]int, 0)
	sort.Ints(nums)
	
	for first:=0;first<l;first++{
		if first>0 && nums[first] == nums[first-1] {
			continue
		}
		third:=l-1
		target:=-1*nums[first]
		for second:=first+1;second<l;second++ {
			if second > first+1 && nums[second]== nums[second-1] {
				continue
			}
			for third>second && nums[second]+nums[third]>target{
				third--
			}
			if third==second{
				break
			}
			if nums[second]+nums[third]==target{
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}
