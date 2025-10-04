from collections import defaultdict

def isValidSudoku(board: list[list[str]]) -> bool:
	rows = defaultdict(set)
	cols = defaultdict(set)
	boxes = defaultdict(set)

	for r in range(9):
		for c in range(9):
			if board[r][c] == ".":
				continue
			n = board[r][c]
			if n in rows[r] or n in cols[c] or n in boxes[(r // 3, c // 3)]:
				return False
			rows[r].add(n)
			cols[c].add(n)
			boxes[(r//3, c//3)].add(n)

	return True
