class Node:
    def __init__(self, data=None):
        self.data = data
        self.next = None

class SLinkedList:
    def __init__(self):
        self.head = None

    def Atbegining(self, data_in):
        NewNode = Node(data_in)
        NewNode.next = self.head
        self.head = NewNode


llist = SLinkedList()
llist.Atbegining(2)
llist.Atbegining(1)
llist.Atbegining(3)
llist.Atbegining(5)
llist.RemoveNode(6)


llist.LListprint()

