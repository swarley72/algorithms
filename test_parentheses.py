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


# Тесты
tests = [
    ("()))", 2),
    ("())", -1),
    ("((((", -1),
    ("()()()))", 6),
    ("(", -1),
    (")", -1),
    ("()(", -1),
    ("(()()(", 5),
    ("(()", -1),
    ("))", 0),
    ("((((", -1),
    (")(", -1),
    ("))((", -1),
]

print("Running tests...\n")
for i, (input_str, expected) in enumerate(tests, 1):
    result = valid_parentheses(input_str)
    status = "✅" if result == expected else "❌"
    print(f"Test {i:2d}: {status} input='{input_str:12s}' expected={expected:2d} got={result:2d}")
    if result != expected:
        print(f"         FAILED!")
