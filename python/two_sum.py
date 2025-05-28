# Two Sum – Find two numbers that add up to a target.

# Example:

# Input: nums = [2,7,11,15], target = 9
# Output: [2,7]

def two_sum(nums, target):
    # complexity: O(n^2)
    for i in range(len(nums)):
        for j in range(i,len(nums)):
            if nums[i]+nums[j]==target:
                return [nums[i], nums[j]]
    return 0


# compliment = target-num (9-2=7)--so find if 7 in list
def two_sum_linear(nums, target):
    seen = set()
    for num in nums:
        compliment = target-num
        if compliment in seen:
            return [compliment, num]
        
        seen.add(num)
    return 0


if __name__ == "__main__":
    nums = [3,4,5,7]
    target = 10

    sums = two_sum(nums, target)
    print(sums)

    sums = two_sum_linear(nums, target)
    print(sums) 