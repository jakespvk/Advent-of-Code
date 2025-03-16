const std = @import("std");
////
const INPUT_FILE = "input.txt";
////
const COUNT_FILE_LINES = if (std.mem.eql(u8, INPUT_FILE, "input.txt")) 1000 else 6;
const INT_SIZE = if (std.mem.eql(u8, INPUT_FILE, "input.txt")) u18 else u5;

pub fn main() !void {
    const file = std.fs.cwd().openFile(INPUT_FILE, .{}) catch |err| {
        std.log.err("Failed to open file: {s}", .{@errorName(err)});
        return;
    };
    defer file.close();

    std.debug.print("{d}\n", .{std.math.maxInt(INT_SIZE)});

    var left_nums: [COUNT_FILE_LINES]INT_SIZE = undefined;
    var right_nums: [COUNT_FILE_LINES]INT_SIZE = undefined;
    var buffer: [128]u8 = undefined;
    var fba = std.heap.FixedBufferAllocator.init(&buffer);
    const allocator = fba.allocator();
    var idx: usize = 0;
    while (file.reader().readUntilDelimiterOrEofAlloc(allocator, '\n', std.math.maxInt(usize)) catch |err| {
        std.log.err("Failed to read line: {s}", .{@errorName(err)});
        return;
    }) |line| {
        defer allocator.free(line);
        left_nums[idx] = try std.fmt.parseInt(INT_SIZE, line[0 .. std.mem.indexOf(u8, line, " ") orelse unreachable], 0);
        right_nums[idx] = try std.fmt.parseInt(INT_SIZE, line[1 + (std.mem.lastIndexOf(u8, line, " ") orelse unreachable) ..], 0);
        std.debug.print("{d} {d}\n", .{ left_nums[idx], right_nums[idx] });
        idx += 1;
    }
    var sum: u64 = 0;
    for (left_nums) |num| {
        const num_ptr: *const [1]INT_SIZE = &num;
        const count_usize = std.mem.count(INT_SIZE, &right_nums, num_ptr);
        if (count_usize != 0) sum += @as(usize, num) * count_usize;
    }

    std.debug.print("\nfinal sum: {d}\n", .{sum});
}
