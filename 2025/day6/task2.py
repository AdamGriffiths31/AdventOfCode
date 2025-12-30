def read_input(filename="input.txt"):
    with open(filename) as f:
        return [line.rstrip('\n') for line in f]


if __name__ == "__main__":
    lines = read_input()
    number_lines, op_line = lines[:-1], lines[-1]
    width = max(len(line) for line in lines)
    number_lines = [line.ljust(width) for line in number_lines]
    op_line = op_line.ljust(width)

    total = product = 0
    adding = True

    for i, op in enumerate(op_line):
        if op in '+*':
            if not adding:
                total += product
            adding = op == '+'
            product = 1
        
        digits = ''.join(line[i] for line in number_lines if line[i] != ' ')
        if digits:
            if adding:
                total += int(digits)
            else:
                product *= int(digits)

    if not adding:
        total += product

    print(total)
