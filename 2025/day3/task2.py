def read_input(filename="input.txt"):
    with open(filename, "r") as f:
        data = [line.strip() for line in f.readlines()]
    return data

def find_joltage(data, num_digits=12):
    n = len(data)
    result = ""
    pos = 0
    for i in range(num_digits):
        max_pos = n - (num_digits - i)
        best_digit = ""
        best_idx = pos
        for j in range(pos, max_pos + 1):
            if data[j] > best_digit:
                best_digit = data[j]
                best_idx = j
        result += best_digit
        pos = best_idx + 1
    return int(result)

if __name__ == "__main__":
    data = read_input()
    result = 0
    for line in data:
        joltage = find_joltage(line)
        result += joltage
    print(result)
