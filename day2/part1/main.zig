const std = @import("std");
////
const INPUT_FILE = "small_input.txt";
////
const INT_SIZE = if (std.mem.eql(u8, INPUT_FILE, "input.txt")) i8 else i6;

pub fn main() !void {
    const file = std.fs.cwd().openFile(INPUT_FILE, .{}) catch |err| {
        std.log.err("Failed to open file: {s}", .{@errorName(err)});
        return;
    };
    defer file.close();

    // std.debug.print("{d}\n", .{std.math.maxInt(INT_SIZE)});

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
        var prev: INT_SIZE = try std.fmt.parseInt(INT_SIZE, iter.next().?, 10);
        var curr: INT_SIZE = undefined;
        var increasing = false;
        var first_run: bool = true;
        var probation: bool = false;
        while (iter.next()) |num| {
            curr = try std.fmt.parseInt(INT_SIZE, num, 10);
            if (first_run) {
                first_run = false;
                if (prev < curr) {
                    increasing = true;
                }
            }
            if (curr == prev) {
                if (probation) break :outer else probation = true;
            } else if (@abs(curr - prev) > 3) {
                break :outer;
            } else if (increasing and prev >= curr) {
                if (probation) break :outer else probation = true;
            } else if (!increasing and prev <= curr) {
                if (probation) break :outer else probation = true;
            }
            prev = curr;
        }
        safe += 1;
    }
    std.debug.print("{d}\n", .{safe});
}
