ret, power = 0, 31
while n:
    ret += (n & 1) << power
    n = n >> 1
    power -= 1
print(ret)