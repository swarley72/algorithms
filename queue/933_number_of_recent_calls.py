from collections import deque

class RecentCounter:
    def __init__(self):
        self.queue = deque()

    def ping(self, t: int) -> int:
        self.queue.append(t)
        window_size = t - 3000
        while self.queue and self.queue[0] < window_size:
            self.queue.popleft()

        return len(self.queue)
