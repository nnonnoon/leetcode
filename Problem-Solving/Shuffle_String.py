s = "art"
indices = [1,0,2]

res = ""
tmp = {}

for i in range(len(indices)):
    tmp[indices[i]] = s[i] 


for i in range(len(tmp)):
   res += tmp[i]

print(res)
