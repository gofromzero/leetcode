// LeetCode Problem: Two Sum
// Author: Assistant
// Language: Rust

use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map: HashMap<i32, i32> = HashMap::new();
        
        for (i, &num) in nums.iter().enumerate() {
            let complement = target - num;
            
            if let Some(&index) = map.get(&complement) {
                return vec![index, i as i32];
            }
            
            map.insert(num, i as i32);
        }
        
        vec![] // This should never be reached for valid inputs
    }
}

// For testing purposes, we need to define the Solution struct
pub struct Solution;

fn main() {
    // Example usage
    let nums = vec![2, 7, 11, 15];
    let target = 9;
    let result = Solution::two_sum(nums, target);
    println!("Result: {:?}", result);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_two_sum() {
        assert_eq!(Solution::two_sum(vec![2, 7, 11, 15], 9), vec![0, 1]);
        assert_eq!(Solution::two_sum(vec![3, 2, 4], 6), vec![1, 2]);
        assert_eq!(Solution::two_sum(vec![3, 3], 6), vec![0, 1]);
    }
}