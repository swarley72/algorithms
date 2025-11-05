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

    if len(stack) == 0:
        return invalid_idx
    if len(stack) == 1:
        return -1
    if len(stack) > 2:
        return -1
    if invalid_idx is not None:
        return -1
    return stack[-1][1]


def is_valid_sequence(s: str) -> bool:
    """Проверяет, является ли последовательность валидной."""
    balance = 0
    for char in s:
        balance += 1 if char == '(' else -1
        if balance < 0:
            return False
    return balance == 0


def brute_force_solution(s: str) -> int:
    """Эталонное решение для сравнения."""
    if len(s) <= 1:
        return -1

    if is_valid_sequence(s):
        return -1

    for i in range(len(s)):
        modified = list(s)
        modified[i] = ')' if s[i] == '(' else '('
        if is_valid_sequence(''.join(modified)):
            return i

    return -1


# Comprehensive edge case tests
edge_cases = [
    # Базовые случаи
    ("(", -1, "single open"),
    (")", -1, "single close"),
    ("((", -1, "two opens"),
    ("))", -1, "two closes (нужно 2 замены)"),

    # Простые исправимые
    ("())", 2, "one extra close at end"),
    ("())(", -1, "close and open extra"),
    ("()(", -1, "one extra open at end"),
    (")()", -1, "close at start"),

    # Множественные ошибки (нельзя исправить одной заменой)
    ("))))", -1, "four closes"),
    ("((((", -1, "four opens"),
    ("))((", -1, "two closes, two opens"),
    ("(()))(", -1, "valid pair + errors"),
    (")()(", -1, "alternating errors"),

    # Исправимые с одной лишней )
    ("())", 2, "valid + one close"),
    ("()))", -1, "valid + two closes"),
    ("()())", 4, "two valid + one close"),
    ("()()()))", 6, "three valid + one close"),
    ("())()", -1, "close in middle"),
    (")())", -1, "close at start + close at end"),

    # Исправимые с двумя лишними (
    ("((()", 0, "two opens + valid OR last open"),
    ("(()(", -1, "two opens, one valid"),
    ("(()(()", 4, "nested + two opens"),
    ("(()()(", 5, "open + two valid + open"),

    # Сложные паттерны
    ("()()()(", -1, "many valid + one open"),
    ("()()())", 6, "many valid + one close"),
    ("((()))", -1, "already valid - should not happen per problem"),
    ("(()())", -1, "already valid - should not happen per problem"),

    # Длинные последовательности
    ("((((((((()", -1, "nine opens, one close"),
    ("()()()()())", 10, "five valid + one close"),
    ("((()()()())", -1, "one open + five valid"),

    # Специальные паттерны
    (")()()(", -1, "starts with close, ends with open"),
    (")))(((", -1, "three closes, three opens"),
    ("()(())", -1, "already valid"),
    ("(()(())", -1, "one extra open"),

    # Edge cases с чередованием
    ("()()(())))", -1, "valid pairs + two closes"),
    ("((((()()", -1, "many opens + valid"),
    ("()()((", -1, "two valid + two opens"),

    # Минимальные исправимые случаи
    ("())", 2, "shortest fixable with )"),
    ("((()", 0, "shortest fixable with ( (maybe)"),

    # Проверка позиций возврата
    (")))", -1, "three closes - need 2+ changes"),
    ("(((", -1, "three opens - need 2+ changes"),
    ("()()()())", 8, "last position close"),
    ("(()()()()", -1, "first position open + valid"),
]


print("=" * 80)
print("EDGE CASE TESTING")
print("=" * 80)
print()

failed = []
for i, (input_str, expected, description) in enumerate(edge_cases, 1):
    result = valid_parentheses(input_str)
    brute_result = brute_force_solution(input_str)

    # Проверяем согласованность с брутфорс решением
    matches_brute = (result == brute_result) or (result != -1 and brute_result != -1)

    # Если результат не -1, проверяем что замена действительно дает валидную последовательность
    actually_valid = False
    if result is not None and result != -1 and 0 <= result < len(input_str):
        modified = list(input_str)
        modified[result] = ')' if input_str[result] == '(' else '('
        actually_valid = is_valid_sequence(''.join(modified))

    status = "✅" if result == expected else "❌"

    result_str = str(result) if result is not None else "None"
    print(f"Test {i:2d}: {status} '{input_str:20s}' expected={expected:2>3} got={result_str:>3}")
    print(f"         {description}")

    if result != expected:
        print(f"         FAILED! Brute force says: {brute_result}")
        if result != -1 and not actually_valid:
            print(f"         ERROR: Returned index {result} but replacement doesn't make valid sequence!")
        failed.append((i, input_str, expected, result, description))
    elif result != -1 and not actually_valid:
        print(f"         WARNING: Returned {result} but replacement doesn't work!")
        failed.append((i, input_str, expected, result, f"{description} (invalid replacement)"))
    elif not matches_brute and result != expected:
        print(f"         WARNING: Differs from brute force: {brute_result}")

    print()

print("=" * 80)
print(f"SUMMARY: {len(edge_cases) - len(failed)}/{len(edge_cases)} tests passed")
print("=" * 80)

if failed:
    print("\nFAILED TESTS:")
    for test_num, input_str, expected, result, desc in failed:
        print(f"  Test {test_num}: '{input_str}' - {desc}")
        print(f"    Expected: {expected}, Got: {result}")

        # Показываем что дает брутфорс
        brute_result = brute_force_solution(input_str)
        print(f"    Brute force: {brute_result}")

        # Если брутфорс нашел решение, покажем его
        if brute_result != -1:
            modified = list(input_str)
            modified[brute_result] = ')' if input_str[brute_result] == '(' else '('
            print(f"    Brute force fix: '{input_str}' -> '{''.join(modified)}'")
else:
    print("\n🎉 All tests passed!")
