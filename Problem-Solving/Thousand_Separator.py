n = str(1234)

result = ""
count = 0

for i in range(len(n)-1, -1, -1):
    if count == 3 :
        result = "." + result
        count = 0
    result = n[i] + result 
    count += 1

print(result)