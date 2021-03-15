def binary_tree_path(cur, res):
    if cur.left is None and cur.right is None:
        ans.append('->'.join(res + [str(cur.val)]))
    if cur.left:
        binary_tree_path(cur.left, res + [str(cur.val)])
    if cur.right:
        binary_tree_path(cur.right, res + [str(cur.val)])
                
ans = []
binary_tree_path(root, [])


# test = ["noon", "hello", "jarawee"]
# x = "->".join(test)
# print(x)