const std = @import("std");
const fs = std.fs;
const print = std.debug.print;

pub fn main() !void {
    print("Hello, world!\n", .{});

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

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
    } else |err| switch (err) {
        error.EndOfStream => {
            if (line.items.len > 0) {
                line_no += 1;
                print("{d}--{s}\n", .{ line_no, line.items });
            }
        },
        else => return err, // propagate error
    }

    print("Total lines: {d}\n", .{ line_no });
}