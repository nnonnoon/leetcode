paths = [["B","C"],["D","B"],["C","A"]]

des = {}

for i in range(len(paths)):
    des[paths[i][0]] = paths[i][1]

path = paths[0][0]

while path in des:
    path = des[path]
print(path)
