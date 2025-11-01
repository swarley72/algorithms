# была задачка по игре в слова на 5 букв,
# типа где угадал позицию и букву +, где только букву *, иначе минус.
# Типа если загадали атлас,
# а попытка была сахар то вернуть **-+- (с учётом количества 'а').


def five(guess: str, word: str): ...


assert five("сахар", "атлас") == "**-+-"
assert five("атлас", "атлас") == "+++++"  # полное совпадение
assert five("книги", "атлас") == "-----"  # нет совпадений
assert five("книга", "атлас") == "----*"  # одно совпадение
assert five("ааааа", "бабан") == "*+---"  # 'а' встречается 2 раза в "бабан"
assert five("ллллл", "атлас") == "*----"  # 'л' встречается 1 раз


def five_success(guess: str, word: str):
    from collections import Counter

    counter = Counter(word)
    res = ""
    for i in range(5):
        if word[i] == guess[i] and word[i] in counter:
            res += "+"
            counter[word[i]] -= 1
            if counter[word[i]] == 0:
                del counter[word[i]]
        elif guess[i] in counter:
            res += "*"
            counter[guess[i]] -= 1
            if counter[guess[i]] == 0:
                del counter[guess[i]]
        else:
            res += "-"

    return res
