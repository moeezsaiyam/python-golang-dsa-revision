import sys

def max_el_in_array(li):
    return max(li)


def max_el_in_array_loop(li):
    maxx= -sys.maxsize -1
    for el in li:
        if maxx < el:
            maxx = el


    return maxx

print(max_el_in_array([2,5,1,8,20,-30]))
print(max_el_in_array_loop([2,5,1,8,20,-30]))