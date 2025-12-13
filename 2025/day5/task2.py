def read_input(filename="input.txt"):
    with open(filename, "r") as f:
        content = f.read()
    parts = content.strip().split("\n\n")
    database = [[int(x) for x in line.strip().split("-")] for line in parts[0].split("\n")]
    return database


if __name__ == "__main__":
    database = read_input()

    intervals = sorted(database)

    merged = [intervals[0][:]]
    for start, end in intervals[1:]:
        if start <= merged[-1][1] + 1:
            merged[-1][1] = max(merged[-1][1], end)
        else:
            merged.append([start, end])

    result = sum(end - start + 1 for start, end in merged)
    print(result)
