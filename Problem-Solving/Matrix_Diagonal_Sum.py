mat = [[1,1,1,1],
              [1,1,1,1],
              [1,1,1,1],
              [1,1,1,1]]

row = 0
left = 0
right = len(mat[0])-1
total = 0 

while row < len(mat):
    if right != left:
        total += mat[row][right]

    total += mat[row][left]
    row += 1
    left += 1
    right -= 1


print(total)