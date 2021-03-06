package main

// #include <stdlib.h>
//
// extern int fetchCode(void *, int);
import "C"
import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/wasmerio/go-ext-wasm/wasmer"
)

func main() {
	// Reads the WebAssembly module as bytes.
	bytes, err := wasmer.ReadBytes("binary.wasm")
	if err != nil {
		panic(err)
	}

	// Compile bytes into wasm binary
	module, err := wasmer.Compile(bytes)
	if err != nil {
		panic(err)
	}

	// Get current wasi version and corresponded import objects
	wasiVersion := wasmer.WasiGetVersion(module)
	if wasiVersion == 0 {
		// wasiVersion is unknow, use Latest instead
		wasiVersion = wasmer.Latest
	}
	importObject := wasmer.NewDefaultWasiImportObjectForVersion(wasiVersion)

	// Make new C imports
	imports := wasmer.NewImports()
	imports, err = imports.AppendFunction("fetch_code", fetchCode, C.fetchCode)
	if err != nil {
		panic(err)
	}

	// Extend the import objects with C imports
	importObject.Extend(*imports)

	// Instantiates the WebAssembly module using derived import objects.
	instance, err := module.InstantiateWithImportObject(importObject)
	if err != nil {
		panic(err)
	}
	defer importObject.Close()
	defer imports.Close()
	defer instance.Close()

	fmt.Println(instance)
	start(&instance)
	hello(&instance)
	sum(&instance)
	concatenate(&instance, "Hello1", "World2!")

	fmt.Println(fetchCodeOnBinary(&instance, 2))
	fmt.Println(fetchCodeOnBinary(&instance, 5))
}

func start(instance *wasmer.Instance) {
	// Gets start function from the WebAssembly instance.
	start := instance.Exports["_start"]
	// Calls that exported function with Go standard values. The WebAssembly types are inferred and values are casted automatically.
	result, err := start()
	if err != nil {
		panic(err)
	}
	// To ensure the start function doesn't return any values
	fmt.Println(result.GetType() == wasmer.TypeVoid)
}

func hello(instance *wasmer.Instance) {
	// Gets hello function from the WebAssembly instance.
	hello := instance.Exports["hello"]
	// Calls that exported function with Go standard values. The WebAssembly types are inferred and values are casted automatically.
	result, err := hello()
	if err != nil {
		panic(err)
	}
	// To ensure the hello function doesn't return any values
	fmt.Println(result.GetType() == wasmer.TypeVoid)
}

func sum(instance *wasmer.Instance) {
	// Gets sum function from the WebAssembly instance.
	sum := instance.Exports["sum"]
	// Calls that exported function with Go standard values. The WebAssembly types are inferred and values are casted automatically.
	result, err := sum(1, 2)
	if err != nil {
		panic(err)
	}
	// To ensure the hello function doesn't return any values
	fmt.Println(result.ToI32())
}

func concatenate(instance *wasmer.Instance, s1, s2 string) {
	// Allocate memory function
	allocate := instance.Exports["allocate"]
	memoryPointerS1 := func() int32 {
		// run allocate for s2
		pointer, err := allocate(len(s1) + 1)
		if err != nil {
			panic(err)
		}
		// copy s2 to memory
		memoryS1 := instance.Memory.Data()[pointer.ToI32():]
		copy(memoryS1, s1)
		memoryS1[len(s1)] = 0
		return pointer.ToI32()
	}()
	memoryPointerS2 := func() int32 {
		// run allocate for s2
		pointer, err := allocate(len(s2) + 1)
		if err != nil {
			panic(err)
		}
		// copy s2 to memory
		memoryS2 := instance.Memory.Data()[pointer.ToI32():]
		copy(memoryS2, s2)
		memoryS2[len(s2)] = 0
		return pointer.ToI32()
	}()
	// Get concatenate function from the WebAssembly instance.
	concatenate := instance.Exports["concatenate"]
	// Calls that exported function with memory pointers.
	result, err := concatenate(memoryPointerS1, memoryPointerS2)
	if err != nil {
		panic(err)
	}
	// The result is another memory pointer is another string
	fmt.Println(result)

	output, _ := convertToString(instance, result)
	fmt.Println(output)
}

func convertToString(instance *wasmer.Instance, output wasmer.Value) (string, int32) {
	memory := instance.Memory.Data()[output.ToI32():]
	var builder strings.Builder
	counter := 0
	for memory[counter] != 0 {
		builder.WriteByte(memory[counter])
		counter++
	}
	return builder.String(), int32(counter)
}

func fetchCodeOnBinary(instance *wasmer.Instance, input int32) int32 {
	// Gets start function from the WebAssembly instance.
	fetch := instance.Exports["fetch"]

	// Calls that exported function with Go standard values. The WebAssembly types are inferred and values are casted automatically.
	result, err := fetch(input)
	if err != nil {
		panic(err)
	}

	return result.ToI32()
}

//export fetchCode
func fetchCode(context unsafe.Pointer, input C.int) C.int {
	return input * 2
}
