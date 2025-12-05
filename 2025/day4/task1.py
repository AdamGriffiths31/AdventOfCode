def read_input(filename="input.txt"):
    with open(filename, "r") as f:
        data = [[char for char in line.strip()] for line in f.readlines()]
    return data


if __name__ == "__main__":
    grid = read_input()
    directions = [
        (-1, -1), (-1, 0), (-1, 1),
        (0, -1),           (0, 1),
        (1, -1),  (1, 0),  (1, 1),
    ]

    res = 0
    for row, line in enumerate(grid):
        for col, char in enumerate(line):
            if char == "@":
                count_of_close_paper = 0
                for dr, dc in directions:
                    new_row, new_col = row + dr, col + dc
                    if 0 <= new_row < len(grid) and 0 <= new_col < len(grid[0]):
                        neighbor = grid[new_row][new_col]
                        if neighbor == "@":
                            count_of_close_paper += 1
                if count_of_close_paper < 4:
                    res += 1
    print(res)
