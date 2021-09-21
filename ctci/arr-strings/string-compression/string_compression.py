
def compress(s: str):
    """Return a run length encoded string.

    Eg compress(aabcccccaaa) -> a2b1c5a3
    """
    counts = []
    for i, c in enumerate(s):
        if i == 0:
            counts.append([c, 1])
            continue
        if counts[-1][0] == c:
            counts[-1][1] += 1
        else:
            counts.append([c, 1])
    return "".join([f"{item[0]}{item[1]}" for item in counts])
