arr = [1,1,1]

new_arr = [0] + arr + [0]
size = len(new_arr)

for i in range(size-1):
    if new_arr[i]%2 == 1 and  new_arr[i+1]%2 == 1 and new_arr[i-1]%2 == 1:
        print("true")
        break
