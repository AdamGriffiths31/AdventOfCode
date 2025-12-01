def read_input(filename="input.txt"):
    with open(filename, "r") as f:
        data = [line.strip() for line in f.readlines()]
    return data

def handle_update(direction, current, amount):
    if direction == "R":
        return (current + amount) % 100
    else:
        return (current - amount) % 100

if __name__ == "__main__":
    data = read_input()
    current = 50
    count_of_zero = 0
    for line in data:
        direction = line[0]
        amount = int(line[1:])
        current = handle_update(direction, current, amount)
        if current == 0:
            count_of_zero += 1
    print(count_of_zero)
