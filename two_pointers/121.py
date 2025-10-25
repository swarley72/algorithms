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



def maxProfit(self, prices: list[int]) -> int:
	res = 0

	current_min = prices[0]

	for i in range(1, len(prices)):
		price = prices[i]
		res = max(res, price - current_min) 
		current_min = min(current_min, price)

	return res
