#           a1----a2      a1----b1  a1----b1      a1----b1
#  a2----b2           a2----b2          a2----b2           a2----b2

#  определение перескаются ли интервалы max(a1, a2) <= min(b1, b2)
def is_overlaping(interval_1, interval_2):
    a1, b1 = interval_1
    a2, b2 = interval_2
    return max(a1, a2) <= min(b1, b2)

#  объединение интервалов [min(a1, a2), max(b1, b2)]
def merge_intervals(interval_1, interval_2):
    a1, b1 = interval_1
    a2, b2 = interval_2
    return [min(a1, a2), max(b1, b2)]

#  площадь пересечения [max(a1, a2), min(b1, b2)]

