s = "AB"

result = 0 
base = 1

for i in range(len(s) - 1, -1 , -1):
    result += (ord(s[i]) - ord("A") + 1) * base
    base *= 26

print(result)

