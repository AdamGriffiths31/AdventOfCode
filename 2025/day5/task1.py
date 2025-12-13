def read_input(filename="input.txt"):
    with open(filename, "r") as f:
        content = f.read()
    parts = content.strip().split("\n\n")
    database = [[int(x) for x in line.strip().split("-")] for line in parts[0].split("\n")]
    ingredients = [line.strip() for line in parts[1].split("\n")]
    return database, ingredients


if __name__ == "__main__":
    database, ingredients = read_input()

    result = 0
    for ingredient in ingredients:
        num = int(ingredient)
        for start, end in database:
            if start <= num <= end:
                result += 1
                break
    print(result)
