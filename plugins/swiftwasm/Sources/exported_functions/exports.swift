import c_header_exported_imported

// not support `Foundation` in wasm env
// import Foundation

@_cdecl("helloFromHostImpl")
func helloFromHostImpl() {
    print("The function is called from a host application")
}

@_cdecl("sum")
func sumFromHost(x: Int32, y: Int32) -> Int32 {
    return x + y
}

///
/// not support `String` type in wasm env
/// https://webassembly.github.io/spec/core/syntax/types.html#value-types
///
/*@_cdecl("concatenate")
func concatenateFromHost(s1: String, s2: String) -> String {
    return s1 + " " + s2
}*/

@_cdecl("allocate")
func allocate(size: Int) -> UnsafeMutableRawPointer {
    return UnsafeMutableRawPointer.allocate(byteCount: size, alignment: MemoryLayout<UInt8>.alignment)
}

@_cdecl("deallocate")
func deallocate(pointer: UnsafeMutableRawPointer, size: Int) {
    pointer.deallocate()
}

@_cdecl("concatenate")
func concatenateFromHost(s1: UnsafePointer<CChar>, s2: UnsafePointer<CChar>) -> UnsafePointer<CChar> {
    let str1 = String(cString: s1)
    let str2 = String(cString: s2)
    let str = str1 + " " + str2
    return str.withCString { (baseAddress) in
        let count = str.utf8.count + 1
        let result = UnsafeMutablePointer<Int8>.allocate(capacity: count)
        result.initialize(from: baseAddress, count: count)
        return UnsafePointer<CChar>(result)
    }
}
