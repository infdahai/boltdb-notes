package bolt

// maxMapSize represents the largest mmap size supported by Bolt.
// bolt 支持最大mmap大小
const maxMapSize = 0xFFFFFFFFFFFF // 256TB

// maxAllocSize is the size used when creating array pointers.
// 2^32 大概 4GB
const maxAllocSize = 0x7FFFFFFF

// Are unaligned load/stores broken on this arch?
// 似乎是要在arm上对齐所需要的trade off
var brokenUnaligned = false
