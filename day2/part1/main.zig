const std = @import("std");
////
const INPUT_FILE = "small_input.txt";
////
const COUNT_FILE_LINES = if (std.mem.eql(u8, INPUT_FILE, "input.txt")) 1000 else 6;
const INT_SIZE = if (std.mem.eql(u8, INPUT_FILE, "input.txt")) i19 else i6;

pub fn main() !void {
    const file = std.fs.cwd().openFile(INPUT_FILE, .{}) catch |err| {
        std.log.err("Failed to open file: {s}", .{@errorName(err)});
        return;
    };
    defer file.close();

    std.debug.print("{d}\n", .{std.math.maxInt(INT_SIZE)});

    var buffer: [128]u8 = undefined;
    var fba = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fba.allocator();
    var safe: usize = 0;
    while (file.reader().readUntilDelimiterOrEofAlloc(allocator, '\n', std.math.maxInt(usize)) catch |err| {
        std.log.err("Failed to read line: {s}", .{@errorName(err)});
        return;
    }) |line| outer: {
        defer allocator.free(line);
        var iter = std.mem.splitScalar(u8, line, ' ');
        var prev: INT_SIZE = try std.fmt.parseInt(INT_SIZE, iter.next().?, 0);
        var curr: INT_SIZE = undefined;
        var increasing = false;
        var first_run: bool = true;
        while (iter.next()) |char| {
            curr = try std.fmt.parseInt(INT_SIZE, char, 0);
            std.debug.print("{d} {d} {}\n", .{ prev, curr, increasing });
            if (curr == prev) {
                break :outer;
            }
            if (@abs(curr - prev) > 3) {
                break :outer;
            } else if (first_run) {
                if (prev < curr) {
                    increasing = true;
                    first_run = false;
                }
            } else if (increasing) {
                if (prev >= curr) break :outer;
            } else if (!increasing) {
                if (prev <= curr) break :outer;
            }
            prev = curr;
        }
        safe += 1;
    }
    std.debug.print("{d}\n", .{safe});
}
