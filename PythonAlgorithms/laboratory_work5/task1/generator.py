# def generator_values(start, stop, step):
#     cur = start
#     result = []
#
#     while cur < stop:
#         result.append(cur)
#         cur += step
#
#     return result

def generator_values(start, stop, step):
    cur = start
    result = []

    while cur < stop:
        result.append(cur)
        cur += step

    return result