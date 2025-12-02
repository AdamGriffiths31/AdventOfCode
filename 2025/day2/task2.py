def read_input(filename="input.txt"):
    with open(filename, "r") as f:
        data = f.read().strip().split(",")
    return data


def walk_range(start, end):
    total = 0
    for i in range(int(start), int(end) + 1):
        s = str(i)
        new_string = (s + s)[1:-1]
        if s in new_string:
            total += i
    return total


if __name__ == "__main__":
    data = read_input()
    result = 0
    for item in data:
        start, end = item.split("-")
        result += walk_range(start, end)
    print(result)
