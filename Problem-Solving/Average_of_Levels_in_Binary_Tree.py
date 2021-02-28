res = []
q = []

while q:
    t_q = []
    t_r = []

    for i in q :
        t_r.append(i.val)

        if i.left:
            t_q.append(i.left)
        if i.right:
            t_q.append(i.right)
    q = t_q
    res.append(t_r)

print([sum(i)/len(i) for i in res])

