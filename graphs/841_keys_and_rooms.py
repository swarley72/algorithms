def can_visit_all_rooms(rooms: list[list[int]]) -> bool:
    visited = set()
    visited.add(0)
    stack = [0]
    while stack:
        key = stack.pop()

        for room in rooms[key]:
            if room not in visited:
                visited.add(room)
                stack.append(room)

    return len(visited) == len(rooms)
