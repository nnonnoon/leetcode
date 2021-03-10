arr1 = [1,2,3,12,20]
arr2 = [6,7,8,9,10]

res = []

i = 0
j = 0

while i < len(arr1) and j < len(arr2):
    if arr1[i] < arr2[i]:
        res.append(arr1[i])
        i += 1
    else:
        res.append(arr2[j])
        j += 1

while i < len(arr1):
    res.append(arr1[i])
    i+= 1


while j < len(arr2):
    res.append(arr2[j])
    i+= 1

print(res)



