package com.feyfree.leetcode.p0213;


/**
 * 213. 打家劫舍 II
 * https://leetcode-cn.com/problems/house-robber-ii/
 *
 * @author leilei
 */
public class Solution {
    public int rob(int[] nums) {
        int length = nums.length;
        if (length == 1) {
            return nums[0];
        } else if (length == 2) {
            return Math.max(nums[0], nums[1]);
        }
        return Math.max(interval(nums, 0, length - 2), interval(nums, 1, length - 1));
    }

    public int interval(int[] nums, int start, int end) {
        int a = nums[start], b = Math.max(nums[start], nums[start + 1]);
        for (int i = start + 2; i <= end; i++) {
            int temp = b;
            b = Math.max(a + nums[i], b);
            a = temp;
        }
        return b;
    }
}