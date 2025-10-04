def maxProfit(prices: list[int]) -> int:
	l, r = 0, 1
	res = 0

	while r < len(prices):
		if prices[l] < prices[r]:
			new_price = prices[r] - prices[l]
			res = max(res, new_price)
		else:
			l = r
		r += 1

	return res
