def valid_anagram(s1, s2):
    s1 = sorted(s1)
    s2 = sorted(s2)

    if s1 == s2:
        return True
    
    return False


def valid_anagram_linear(s1, s2):
    frequency = {}

    for ch in s1:
        frequency[ch] = frequency.get(ch, 0) + 1
    for ch in s2:
        frequency[ch] = frequency.get(ch, 0) - 1

    for value in frequency.values():
        if value!=0:
            return False
    return True
    

print(valid_anagram("okay","kayo"))
print(valid_anagram_linear("okay","kayo"))