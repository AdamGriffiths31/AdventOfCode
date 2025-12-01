def read_input(filename="input.txt"):
    with open(filename, "r") as f:
        data = [line.strip() for line in f.readlines()]
    return data


def handle_update(direction, current, amount):
    if direction == "R":
        crosses = (current + amount) // 100
        return (current + amount) % 100, crosses
    else:
        if current == 0:
            crosses = amount // 100
        elif amount >= current:
            crosses = (amount - current) // 100 + 1
        else:
            crosses = 0
        return (current - amount) % 100, crosses


if __name__ == "__main__":
    data = read_input()
    current = 50
    count_of_zero = 0
    for line in data:
        direction = line[0]
        amount = int(line[1:])
        current, crosses = handle_update(direction, current, amount)
        count_of_zero += crosses
    print(count_of_zero)
