def sliding_window(some_collection):
    begin = 0
    window_state = 0
    result = float("-inf")

    for end in range(len(some_collection)):

        window_size = end - begin + 1
        if ... # window condition
            begin += 1        
