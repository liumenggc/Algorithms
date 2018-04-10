package twosum

func twoSum(nums []int, target int) []int{
	//m 负责保存map[int]int 的序列号.
	m := make(map[int]int, len(nums))

	//
	for i, b := range nums{
		//通过查询map,获取a = target - b的序列号.
		if j, ok := m[target-b]; ok {
			//说明在i之前， 存在nums[j]==a.
			return []int{j,i}
			// 顺序时j, i,因为j > i
		}

		//把i  和i的值，存入map
		m[nums[i]] = i
	}
}
