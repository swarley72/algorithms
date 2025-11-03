"""
Definition of Interval:

"""

import heapq


class Interval(object):
    def __init__(self, start, end):
        self.start = start
        self.end = end


class Solution:
    def minMeetingRooms(self, intervals: list[Interval]) -> int:
        intervals.sort(key=lambda x: x.start)

        min_heap = []
        for meet in intervals:
            if not min_heap:
                heapq.heappush(min_heap, meet.end)
            else:
                if min_heap[0] <= meet.start:
                    heapq.heappop(min_heap)
                heapq.heappush(min_heap, meet.end)
        return len(min_heap)
