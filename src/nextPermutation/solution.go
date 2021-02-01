package nextPermutation

// https://leetcode.com/problems/next-permutation/
func nextPermutation(nums []int)  {
	length:=len(nums)
	i:=length-2
	for i>=0 && nums[i]>=nums[i+1]{
		i--
	}
	if i>=0{
		j:=length-1
		for nums[i] >= nums[j]{
			j--
		}
		nums[i],nums[j]=nums[j],nums[i]
	}
	reverse(nums[i+1:])
}

func reverse(a []int){
	for i,n:=0,len(a);i<n/2;i++{
		a[i],a[n-1-i]=a[n-1-i],a[i]
	}
} 

//Runtime: 4 ms, faster than 19.82% of Go online submissions for Next Permutation.
//Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Next Permutation.