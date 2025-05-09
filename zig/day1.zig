const std = @import("std");
const fs = std.fs;
const print = std.debug.print;

// References:
// https://cookbook.ziglang.cc/01-01-read-file-line-by-line.html

pub fn main() !void {
    print("Hello, world!\n", .{});

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var first_numbers = std.ArrayList(i64).init(allocator);
    defer first_numbers.deinit();
    var second_numbers = std.ArrayList(i64).init(allocator);
    defer second_numbers.deinit();

    const file = try fs.cwd().openFile("../input-1.txt", .{});
    defer file.close();

    // Wrap the file reader in a buffered reader.
    // Since it's usually faster to read a bunch of bytes at once.
    var buf_reader = std.io.bufferedReader(file.reader());
    const reader = buf_reader.reader();

    var line = std.ArrayList(u8).init(allocator);
    defer line.deinit();

    const writer = line.writer();
    var line_no: usize = 0;
    while (reader.streamUntilDelimiter(writer, '\n', null)) {
        // clear the line so we can reuse it
        defer line.clearRetainingCapacity();
        line_no += 1;

        print("{d}--{s}\n", .{ line_no, line.items });

        var iter = std.mem.tokenizeSequence(u8, line.items, " ");
        if (iter.next()) |num_str| {
            const num = try std.fmt.parseInt(i64, num_str, 10);
            try first_numbers.append(num);
        }
        if (iter.next()) |num_str| {
            const num = try std.fmt.parseInt(i64, num_str, 10);
            try second_numbers.append(num);
        }
    } else |err| switch (err) {
        error.EndOfStream => {
            if (line.items.len > 0) {
                line_no += 1;
                print("{d}--{s}\n", .{ line_no, line.items });
            }
        },
        else => return err, // propagate error
    }

    print("First numbers: ", .{});
    for (first_numbers.items) |num| {
        print("{d} ", .{num});
    }

    print("\nSecond numbers: ", .{});
    for (second_numbers.items) |num| {
        print("{d} ", .{num});
    }

    print("Total lines: {d}\n", .{line_no});
}
