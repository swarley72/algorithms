# Дана строка из латинских заглавных букв цифр и скобоквнутри строки:
# Скобочные выражения всегда верные
# за цифрой всегда стоит или буква или скобочное выражениес
# cкобочные выражения могут быть вложенными
# Надо преобразовать в строку умножая выражения за цифрой
# 3AB2(Z3K) == AAABZKKKZKKK
# 2(Z3(KA)) == ZKAKAKAZKAKAKA


def decode_string(s: str) -> str: ...


def decode_string_success(s: str) -> str:
    stack = []
    i = 0

    while i < len(s):
        if s[i].isdigit():
            if s[i + 1] == "(":
                stack.append(s[i])
                stack.append(s[i + 1])
            else:
                n = int(s[i])
                stack.append(s[i + 1] * n)
            i += 2
        elif s[i].isalpha() or s[i] == "(":
            stack.append(s[i])
            i += 1
        else:
            substring = ""
            while stack[-1] != "(":
                substring = stack.pop() + substring
            stack.pop()
            n = int(stack.pop())
            stack.append(substring * n)
            i += 1

    return "".join(stack)


assert decode_string("3AB2(Z3K)") == "AAABZKKKZKKK"
assert decode_string("2(Z3(KA))") == "ZKAKAKAZKAKAKA"
assert decode_string("A") == "A"
assert decode_string("2(A)B") == "AAB"
assert decode_string("2(3A)") == "AAAAAA"
assert decode_string("2(2(A))") == "AAAA"
assert decode_string("2()") == ""
assert decode_string("3A2B4C") == "AAABBCCCC"
assert decode_string("2(AB)3C") == "ABABCCC"
assert decode_string("2(A2B)") == "ABBABB"
assert decode_string("3(A)2(B)") == "AAABB"
assert decode_string("AB3C") == "ABCCC"
