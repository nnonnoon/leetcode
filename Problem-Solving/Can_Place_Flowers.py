flowerbed = [1,0,0,1,0,1,0,0,0,0]
n = 1

s = len(flowerbed)
bed = [0] + flowerbed + [0]

for i in range(1, s+1):
    if bed[i] == bed[i-1] == bed[i+1] == 0:
        bed[i] = 1
        n -= 1

if n <= 0: 
    print("True")
else:
    print("False")
