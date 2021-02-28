nums1 = [1]
m = 1
nums2 = []
n = 0

nums1 = sorted(nums1 + nums2)

count = 0

for i in range(len(nums1)):
   if nums1[i] == 0:
       count += 1

while count:
    nums1.pop(0)
    count -= 1

print(nums1)