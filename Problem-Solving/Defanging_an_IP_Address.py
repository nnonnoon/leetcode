s = ""
for i in address:
    if i == ".":
        i = "[.]"
        s += i
    else:
        s += i
print(s)