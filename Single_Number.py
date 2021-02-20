nums = [4,1,2,1,2]

tmp = {}

for i in nums:
    if i not in tmp:
        tmp[i] = 1
    else:
        tmp[i] = tmp[i] + 1

for key in tmp:
    if tmp[key] == 1 :
        print(key)
        break
    