arr1 = [2,21,43,38,0,42,33,7,24,13,12,27,12,24,5,23,29,48,30,31]
arr2 = [2,42,38,0,43,21]

tmp = {}

for nums in arr2:
    tmp[nums] = 0
    for i in arr1:
        if i == nums:
            tmp[nums] +=1

# print(tmp)
arr1.sort()

tmp2 = {}

for j in arr1:
    if j not in tmp and j not in tmp2:
        tmp2[j] = 1
        continue
    if j in tmp2:
        tmp2[j] += 1

tmp.update(tmp2)

res = []

for key in tmp:
    for j in range(tmp[key]):
        res.append(key)

print(res)

# ans=[]
# for i in arr2:
#     while i in arr1:
#         arr1.remove(i)
#         ans.append(i)
# arr1.sort()
# ans = ans + arr1

# print(ans)





