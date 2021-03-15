A = ["bella","label","roller"]

res = []
count = {}

for i in range(len(A)):
    for j in A[i]:
        if j not in count:
            count[j] = 0
        count[j] +=1
        if count[j] == 3:
            res.append(j)
            count[j] = 0
print(res)