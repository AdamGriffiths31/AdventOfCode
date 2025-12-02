def read_input(filename="input.txt"):
    with open(filename, "r") as f:
        data = f.read().strip().split(",")
    return data


def walk_range(start, end):
    total = 0
    for i in range(int(start), int(end) + 1):
        s = str(i)
        if len(s) % 2 == 1:
            continue
        mid = len(s) // 2
        left = int(s[:mid])
        right = int(s[mid:])
        if left == right:
            total += i
    return total


if __name__ == "__main__":
    data = read_input()
    result = 0
    for item in data:
        start, end = item.split("-")
        result += walk_range(start, end)
    print(result)
