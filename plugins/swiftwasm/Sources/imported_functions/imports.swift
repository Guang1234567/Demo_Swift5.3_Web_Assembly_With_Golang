
import c_header_exported_imported

@_cdecl("fetch")
func fetchCodefromHost(input: Int32) -> Int32 {
    return c_fetchCode(input);
}