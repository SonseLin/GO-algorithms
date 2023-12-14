package main

import "fmt"

// passed 55/61 tests except for 1000+ values

func maxArea(height []int) int {
	maxArea := 0
	secondMax := 0
	savedVI := 0
	savedVJ := 0
	for i, vi := range height[0 : len(height)-1] {
		measure := vi
		if savedVI == vi && i < len(height)-1 {
			continue
		} else {
			savedVI = vi
		}
		for j, vj := range height[1:] {
			j++
			if savedVJ == vj && j < len(height)-1 && savedVJ == height[j+1] {
				continue
			} else {
				savedVJ = vj
			}
			if i != j {
				if vj > secondMax {
					secondMax = vj
				}
				if vi > vj {
					measure = vj
				}
				measure *= j - i
				if maxArea < measure {
					maxArea = measure
				}
				measure = vi
			}
		}
		secondMax = 0
	}
	return maxArea
}

func main() {
	test_cases := [][]int{
		{1, 1},
		{4, 6, 4, 6, 4, 4, 4, 6},
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{4, 3, 2, 1, 4},
		{1, 2},
		{2, 1},
		{4, 3, 2, 1, 4},
		{1, 0, 0, 0, 0, 0, 0, 2, 2},
		{1, 1000, 1000, 6, 2, 5, 4, 8, 3, 7},
		{8361, 5302, 8672, 2400, 5150, 3527, 9216, 6713, 2902, 310, 555, 9176, 311, 9968, 5705, 3983, 7992, 8553, 6953, 9541, 5828, 1750, 6731, 3552, 5274, 7303, 3724, 5387, 9504, 1900, 937, 1146, 7266, 7943, 7911, 9055, 8046, 7180, 6516, 7810, 686, 5210, 1956, 4540, 7540, 2083, 1579, 4260, 2450, 2527, 6524, 5723, 6766, 777, 5694, 6018, 2880, 3653, 6011, 8172, 5943, 2862, 6594, 2902, 9887, 5878, 3065, 8197, 9195, 4560, 3428, 2209, 475, 852, 9488, 3368, 4319, 6230, 1975, 5829, 9474, 4490, 2067, 6048, 9136, 5344, 6022, 1787, 5553, 140, 5130, 524, 3450, 4008, 721, 6154, 5598, 8219, 4614, 3404, 8232, 9023, 4552, 7711, 6057, 5324, 8578, 3595, 4663, 4, 3703, 1429, 7921, 3085, 3694, 1461, 8932, 2632, 7046, 801, 6043, 617, 7565, 3469, 1627, 1464, 3050, 7982, 6702, 5467, 8604, 5515, 9155, 3260, 5040, 313, 8885, 929, 4103, 7947, 1139, 702, 1047, 2889, 1439, 3945, 4738, 2462, 8491},
	}
	for tc := range test_cases {
		fmt.Println(maxArea(test_cases[tc]))
	}
}
