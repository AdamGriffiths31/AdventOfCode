const std = @import("std");

const alloc = std.heap.page_allocator;

pub fn main() !void {
    const max_bytes_per_line = 4096;
    var file = std.fs.cwd().openFile("input.txt", .{}) catch {
        return;
    };
    defer file.close();
    var buffered_reader = std.io.bufferedReader(file.reader());
    const reader = buffered_reader.reader();
    var returnValue: i32 = 0;
    while (try reader.readUntilDelimiterOrEofAlloc(alloc, '\n', max_bytes_per_line)) |line| {
        defer alloc.free(line);
        returnValue += findFirstAndLast(line);
    }
    std.debug.print("{d}", .{returnValue});
}

fn findFirstAndLast(line: []u8) i32 {
    var firstValue: ?u8 = null;
    var lastValue: ?u8 = null;
    for (line) |lineValue| {
        if (isDigit(lineValue)) {
            const digit = lineValue - '0';
            if (firstValue == null) {
                firstValue = digit;
            }
            lastValue = digit;
        }
    }

    return 10 * firstValue.? + lastValue.?;
}

fn isDigit(char: u8) bool {
    return char >= '0' and char <= '9';
}
