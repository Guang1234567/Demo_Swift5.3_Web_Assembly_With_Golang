// swift-tools-version:5.3
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
        name: "swiftwasm",
        dependencies: [
            // Dependencies declare other packages that this package depends on.
            // .package(url: /* package url */, from: "1.0.0"),
        ],
        targets: [
            // Targets are the basic building blocks of a package. A target can define a module or a test suite.
            // Targets can depend on other targets in this package, and on products in packages this package depends on.
            .target(
                    name: "swiftwasm",
                    dependencies: ["exported_functions", "imported_functions"]),
            .target(name: "c_header_exported_imported",
                    dependencies: []),
            .target(
                    name: "exported_functions",
                    dependencies: ["c_header_exported_imported"]),
            .target(
                    name: "imported_functions",
                    dependencies: ["c_header_exported_imported"]),
            .testTarget(
                    name: "swiftwasmTests",
                    dependencies: ["swiftwasm"]),
        ]
)
