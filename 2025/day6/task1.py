def read_input(filename="input.txt"):
    with open(filename, "r") as f:
        data = [line.rstrip('\n') for line in f.readlines()]
    return data


if __name__ == "__main__":
    data = read_input()
    
    rows = [line.split() for line in data]
    
    num_cols = len(rows[0])
    cols = []
    for col_idx in range(num_cols):
        col = []
        for row in rows:
            if col_idx < len(row):
                col.append(row[col_idx])
        cols.append(col)
    
    result = 0
    for i, col in enumerate(cols):
        operator = col[-1]
        numbers = [int(x) for x in col[:-1]]
        
        if operator == '*':
            val = 1
            for n in numbers:
                val *= n
        elif operator == '+':
            val = sum(numbers)
        
        print(f"Col {i}: {numbers} {operator} = {val}")
        result += val
    
    print(f"Result: {result}")
