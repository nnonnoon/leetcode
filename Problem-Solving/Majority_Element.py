nums = [3,2,3]

check = {}
count = 0

for i in range(len(nums)):
    if nums[i] not in check:
        check[nums[i]] = 1
    else:
        check[nums[i]] = check[nums[i]] + 1


for key in check:
    if check[key] > (len(nums)//2):
        print(key)

