# swiftwasm

Example of Swift5.3 + Web Assembly.

# Run

./build.sh

# Output

```bash

╰─ ./build.sh

#0 安装编译链
Updating Homebrew...
Warning: swiftwasm/tap/carton 0.7.1 is already installed and up-to-date
To reinstall 0.7.1, run `brew reinstall carton`
- checking Swift compiler path: /Users/xxx/.swiftenv/versions/wasm-5.3-SNAPSHOT-2020-10-25-a/usr/bin/swift
- checking Swift compiler path: /Users/xxx/.carton/sdk/wasm-5.3-SNAPSHOT-2020-10-25-a/usr/bin/swift
- checking Swift compiler path: /Users/xxx/Library/Developer/Toolchains/swift-wasm-5.3-SNAPSHOT-2020-10-25-a.xctoolchain/usr/bin/swift
Inferring basic settings...
- swift executable: /Users/xxx/Library/Developer/Toolchains/swift-wasm-5.3-SNAPSHOT-2020-10-25-a.xctoolchain/usr/bin/swift
SwiftWasm Swift version 5.3 (swiftlang-5.3.0)
Target: x86_64-apple-darwin19.6.0

Parsing package manifest: /Users/xxx/Library/Developer/Toolchains/swift-wasm-5.3-SNAPSHOT-2020-10-25-a.xctoolchain/usr/bin/swift package dump-package

SDK successfully installed!
#1 编译 wasm 库
[12/12] Linking swiftwasm
#2 快捷方式 wasm 库
./binary.wasm -> ./plugins/swiftwasm/.build/wasm32-unknown-wasi/debug/swiftwasm
#3 简单测试刚被编译的 wasm 库
Error: Can't instantiate module: LinkError([ImportNotFound { namespace: "wasi_snapshot_preview1", name: "proc_exit" }, ImportNotFound { namespace: "env", name: "fetch_code" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "args_sizes_get" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "args_get" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "sched_yield" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "fd_read" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "fd_write" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "fd_close" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "clock_time_get" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "poll_oneoff" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "fd_fdstat_get" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "path_open" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "path_filestat_get" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "path_rename" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "path_link" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "path_create_directory" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "fd_readdir" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "path_readlink" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "path_symlink" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "fd_prestat_get" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "fd_prestat_dir_name" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "path_remove_directory" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "random_get" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "path_unlink_file" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "environ_sizes_get" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "environ_get" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "fd_fdstat_set_flags" }, ImportNotFound { namespace: "wasi_snapshot_preview1", name: "fd_seek" }])
#4 终于可以被 golang 编写的 host (宿主)主程序调用
{0x468b510 0xc00000e200 map[_start:0x40ab440 allocate:0x40ab440 concatenate:0x40ab440 deallocate:0x40ab440 fetch:0x40ab440 hello:0x40ab440 sum:0x40ab440] 0xc000012010 <nil>}
Hello, world!
true
The function is called from a host application
true
3
2097840
Hello1 World2!
4
10

```
