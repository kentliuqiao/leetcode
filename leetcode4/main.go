package main

import (
	"fmt"
	"io/ioutil"
	"os"
)
//
//func main() {
//
//	//m:=map[byte]int{}
//	//fmt.Println(m['a'])
//
//	aa:="C:\\Users\\pc-11\\go\\src\\resource-repo\\upload\\b2\\b2"
//	dir,_:=ioutil.ReadDir(aa)
//
//	fmt.Printf("%#v\n",dir)
//
//	if len(dir) == 0{
//		os.RemoveAll()
//		fmt.Println("dir")
//	} else {
//		fmt.Println("not")
//	}
//}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen % 2 == 1 {
		return float64(getKthElement(nums1, nums2, totalLen / 2 +1))
	}
	return  float64(getKthElement(nums1, nums2, totalLen / 2) + getKthElement(nums1, nums2, totalLen / 2+1)) / 2.0
}

func getKthElement(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2 + k - 1]
		}
		if index2 == len(nums2) {
			return nums1[index1 + k - 1]
		}
		if k == 1 {
			return min(nums2[index2], nums1[index1])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		if nums1[newIndex1] <= nums2[newIndex2] {
			k -= (newIndex1 - index1 +1)
			index1 = newIndex1+1
		} else {
			k -= (newIndex2 - index2 +1)
			index2 = newIndex2+1
		}
	}
}

func min(x,y int) int {
	if x < y {
		return x
	}
	return y
}
