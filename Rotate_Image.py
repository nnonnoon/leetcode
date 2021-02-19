matrix = [[1,2],[3,4]]


#-----Me-----#
result = []

for i in range(len(matrix)):
    tmp = []
    for j in range(len(matrix[i])):
        tmp.append(matrix[len(matrix)-1-j][i])
    result.append(tmp)
                
print(result)

#-----Solution-----#
n = len(matrix[0])
for i in range(n // 2 + n % 2):
    for j in range(n // 2):
        tmp = matrix[n - 1 - j][i]
        matrix[n - 1 - j][i] = matrix[n - 1 - i][n - j - 1]
        matrix[n - 1 - i][n - j - 1] = matrix[j][n - 1 -i]
        matrix[j][n - 1 - i] = matrix[i][j]
        matrix[i][j] = tmp