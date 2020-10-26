import XCTest

import swiftwasmTests

var tests = [XCTestCaseEntry]()
tests += swiftwasmTests.allTests()
XCTMain(tests)
