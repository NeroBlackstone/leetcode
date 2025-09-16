// https://leetcode.com/problems/increasing-triplet-subsequence/description/
function increasingTriplet(nums: number[]): boolean {
    let littleOne = nums[0]!;
    let biggerOne =Number.MAX_VALUE;
    for (let i =1;i<nums.length;i++){
        const n = nums[i]!;
        if(n > biggerOne){
            return true;
        }else if(n >littleOne){
            biggerOne=n;
        }else{
            littleOne=n;
        }
    }
    return false
};