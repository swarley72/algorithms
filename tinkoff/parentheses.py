# Дана неправильная скобочная последовательность
# из открывающих и закрывающих скобок.
# Нужно вывести, можно ли заменить одну из открывающих скобок
# на закрывающую, или наоборот,
# так чтобы получилась правильная скобочная последовательность.
# Если можно, то вывести индекс скобки, которую надо заменить.
# Если нельзя, то вывести -1. Примеры:


def valid_parentheses(s: str) -> int:
    invalid_idx = None
    stack = []

    for i, n in enumerate(s):
        if n == "(":
            stack.append((n, i))
        else:
            if stack:
                prev, _ = stack.pop()
                if prev == "(":
                    continue
                elif invalid_idx is None:
                    invalid_idx = i
                else:
                    return -1
            else:
                if invalid_idx is None:
                    invalid_idx = i
                    stack.append(("(", i))
                else:
                    return -1
    # Если последовательность валидна
    if len(stack) == 0 and invalid_idx is None:
        return -1
    # Если дошли сюда значит была одна замена
    if len(stack) == 0:
        return invalid_idx
    # Если в стеке одна открывающая скобка значит точно невалидная
    if len(stack) == 1:
        return -1
    # Если в стеке больше 2 открывающих заначит точно невалидная
    if len(stack) > 2:
        return -1
    # Если в стеке две открывающие и мы уже использовали замену то точно не валидная
    if invalid_idx is not None:
        return -1
    # Если дошли до сюда значит в стеке 2 открывающие
    return stack[-1][1]


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
assert valid_parentheses("))((") == -1  # две ошибки
assert valid_parentheses("(())") == -1  # две ошибки
