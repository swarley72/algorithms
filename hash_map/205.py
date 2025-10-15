def isIsomorphic(s: str, t: str) -> bool:

    if len(s) != len(t):
        return False

    hash_map = {}
    hash_map2 = {}

    for i in range(len(s)):
        c1 = s[i]
        c2 = t[i]

        if c1 in hash_map:
            if hash_map[c1] != c2:
                return False
        if c2 in hash_map2:
            if hash_map2[c2] != c1:
                return False
        
        hash_map[c1] = c2
        hash_map2[c2] = c1
    return True
