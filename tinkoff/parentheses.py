# Дана неправильная скобочная последовательность
# из открывающих и закрывающих скобок.
# Нужно вывести, можно ли заменить одну из открывающих скобок
# на закрывающую, или наоборот,
# так чтобы получилась правильная скобочная последовательность.
# Если можно, то вывести индекс скобки, которую надо заменить.
# Если нельзя, то вывести -1. Примеры:


def valid_parentheses(s: str) -> int: ...


assert valid_parentheses("()))") == 2
assert valid_parentheses("())") == -1
assert valid_parentheses("((((") == -1
assert valid_parentheses("))") == 0
assert valid_parentheses("()()()))") == 6
assert valid_parentheses("(") == -1
assert valid_parentheses(")") == -1  # одна ) → нельзя исправить
assert valid_parentheses("()(") == -1  # лишняя ( в конце
assert valid_parentheses("(()()(") == 5  # лишняя ( в конце длинной строки
assert valid_parentheses("(()") == -1  # лишняя ( в начале
assert valid_parentheses("))") == 0  # две лишние )
assert valid_parentheses("((((") == -1  # четыре лишние (
assert valid_parentheses(")(") == -1  #
assert valid_parentheses("))((") == -1  # две ошибки


def valid_parentheses_success(s: str) -> int:
    stack = []
    invalid_idx = None

    for i, p in enumerate(s):
        if p == "(":
            stack.append((p, i))
        else:
            if stack:
                stack.pop()
            else:
                if invalid_idx is None:
                    invalid_idx = i
                    stack.append(("(", i))
                else:
                    return -1

    if not stack:
        return invalid_idx

    if len(stack) == 2 and invalid_idx is None:
        return stack[-1][1]

    return -1
