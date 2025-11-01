# Дана неправильная скобочная последовательность
# из открывающих и закрывающих скобок.
# Нужно вывести, можно ли заменить одну из открывающих скобок
# на закрывающую, или наоборот,
# так чтобы получилась правильная скобочная последовательность.
# Если можно, то вывести индекс скобки, которую надо заменить.
# Если нельзя, то вывести -1. Примеры:


def valid_parentheses(seq: str) -> int:
    # Находим все незакрытые "(" и лишние ")"
    stack = []
    replace_idx = -1
    replace_used = False

    for i, s in enumerate(seq):
        if s == "(":
            stack.append((s, i))
        else:
            if not stack:
                if not replace_used:
                    stack.append(("*", i))
                    replace_idx = i
                else:
                    return -1
            else:
                last, _ = stack.pop()
                if last == "(":
                    continue
                elif last == "*" and not replace_used:
                    replace_used = True
                    continue
                else:
                    return -1
    if len(stack) == 2 and not replace_used:
        return stack[-1][1]

    if len(stack) > 0:
        return -1

    return replace_idx


assert valid_parentheses("()))") == 2
assert valid_parentheses("())") == -1
assert valid_parentheses("((((") == -1
assert valid_parentheses("()()()))") == 6
assert valid_parentheses("(") == -1
assert valid_parentheses(")") == -1  # одна ) → нельзя исправить
assert valid_parentheses("()(") == -1  # лишняя ( в конце
assert valid_parentheses("(()()(") == 5  # лишняя ( в конце длинной строки
assert valid_parentheses("(()") == -1  # лишняя ( в начале
assert valid_parentheses("))") == 0  # две лишние )
assert valid_parentheses("((((") == -1  # четыре лишние (
assert valid_parentheses(")(") == -1  #
# assert valid_parentheses("))((") == -1  # две ошибки
# assert valid_parentheses("(()())") == -1  # валидная
