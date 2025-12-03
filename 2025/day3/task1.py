def read_input(filename="input.txt"):
    with open(filename, "r") as f:
        data = [line.strip() for line in f.readlines()]
    return data

def find_joltage(data):
    best = 0
    max_right = data[-1]
    for char in reversed(data[:-1]):
        joltage = int(char + max_right)
        if joltage > best:
            best = joltage
        if char > max_right:
            max_right = char
    return best

if __name__ == "__main__":
    data = read_input()
    result = 0
    for line in data:
        joltage = find_joltage(line)
        result += joltage
    print(result)
