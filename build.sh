#!/bin/bash

echo "#0 安装编译链"
brew install swiftwasm/tap/carton
# https://github.com/swiftwasm/swift/releases
swift_wasm_toolchain_version="wasm-5.3-SNAPSHOT-2020-10-25-a"
carton sdk install $swift_wasm_toolchain_version

echo "#1 编译 wasm 库"

cd ./plugins/swiftwasm

swift_cc=${HOME}/Library/Developer/Toolchains/swift-${swift_wasm_toolchain_version}.xctoolchain/usr/bin/swift
swift_imports_flag="-Xlinker --allow-undefined"
swift_exports_flag="-Xlinker --export=sum \
-Xlinker --export=allocate \
-Xlinker --export=deallocate \
-Xlinker --export=concatenate \
-Xlinker --export=fetch"
swift_build_flag="--triple wasm32-unknown-wasi \
-c debug \
$swift_imports_flag \
$swift_exports_flag"
$swift_cc build $swift_build_flag


cd $OLDPWD

echo "#2 快捷方式 wasm 库"

ln -vsf ./plugins/swiftwasm/.build/wasm32-unknown-wasi/debug/swiftwasm ./binary.wasm

echo "#3 简单测试刚被编译的 wasm 库"

wasmer ./binary.wasm

echo "#4 终于可以被 golang 编写的 host (宿主)主程序调用"

go run ./



