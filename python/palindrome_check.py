def palindrome(p):
    for i in range(int(len(p)/2)):
        if p[i] != p[len(p)-i-1]:
            return False
    
    return True


print(palindrome('12211'))