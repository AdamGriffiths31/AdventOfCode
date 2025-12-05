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
    looping = True
    while looping:
        paper_removed = 0
        loop_count = 0
        for row, line in enumerate(grid):
            for col, char in enumerate(line):
                if char == "@":
                    count_of_close_paper = 0
                    for dr, dc in directions:
                        new_row, new_col = row + dr, col + dc
                        if 0 <= new_row < len(grid) and 0 <= new_col < len(grid[0]):
                            neighbor = grid[new_row][new_col]
                            if neighbor == "@" or neighbor == "x":
                                count_of_close_paper += 1
                    if count_of_close_paper < 4:
                        grid[row][col] = "x"
                        paper_removed += 1
                        loop_count += 1
        res += paper_removed
        for r in range(len(grid)):
            for c in range(len(grid[r])):
                if grid[r][c] == "x":
                    grid[r][c] = "."
        if loop_count == 0:
            looping = False
        print(res)
