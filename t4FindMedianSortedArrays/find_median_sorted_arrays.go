package t4FindMedianSortedArrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//设nums1的长度小于nums2的长度
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	//设一个指针从nums1左侧移动
	//nums1长度小于等于nums2的长度处理方案
	//最小移动开始位置
	low := 0
	//最大移动结束位置
	high := len(nums1)
	//数组连接在一起时的中间位置index
	k := (len(nums1) + len(nums2) + 1) >> 1
	//初始化nums1数组的中间位置
	nums1Mid := 0
	//初始化nums2 数组的中间位置
	nums2Mid := 0
	for low <= high {
		//分界线右侧是mid,分界线左侧是mid-1
		nums1Mid = low + (high-low)>>1
		nums2Mid = k - nums1Mid
		//nums1中的分界线划多了要向左边回移
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] {
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] {
			//nums1中的分界线划少了,要向右边移动
			low = nums1Mid + 1
		} else {
			//找到合适的就退出
			break
		}
	}
	midLeft := 0
	midRight := 0
	//确定mid左侧的数据 有三种情况
	//1.nums1Mid的位置为0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		//2.nums2Mid的位置为0
		midLeft = nums1[nums1Mid-1]
	} else {
		//其余情况取最大值
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}

	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	//确定mid右侧的数据 有三种情况
	//1.nums1Mid的位置为该数组最大长度
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		//1.nums2Mid的位置为该数组最大长度
		midRight = nums1[nums1Mid]
	} else {
		//其余情况取最小值
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func max(m int, n int) int {
	if m >= n {
		return m
	}
	return n
}

func min(m int, n int) int {
	if m <= n {
		return m
	}
	return n
}
