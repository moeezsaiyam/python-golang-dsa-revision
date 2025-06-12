import heapq


def merge_sorted_lists(li1,li2):
    li = li1+li2
    return sorted(li)

def merge_sorted_list_pointer(li1, li2):
    p1 = 0
    p2 = 0
    li = []

    while p1 < len(li1) and p2 < len(li2):
        if li1[p1] < li2[p2]:
            li.append(li1[p1])
            p1+=1
        else:
            li.append(li2[p2])
            p2+=1
        
    while p1 < len(li1):
        li.append(li1[p1])
        p1+=1
    
    while p2 < len(li2):
        li.append(li2[p2])
        p2+=1

    return li


def merge_lists_heap(li1, li2):
    li = heapq.merge(li1, li2)
    return list(li)
    

li1 = [1,2,4,6,8]
li2 = [1,3,4,5,7,12]

li = merge_sorted_lists(li1,li2)
print(li)

li = merge_sorted_list_pointer(li1,li2)
print(li)

li =  merge_lists_heap(li1, li2)
print(li)