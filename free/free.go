package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

type St struct {
	A int
	B int
}

func consumer(p *St) {
	fmt.Println(p)
	select {}
}

func main() {
	b := f()
	go consumer(b)
	a := unsafe.Pointer(b)
	p := uintptr(a)
	fmt.Printf("b is at %0x\n", &b)
	fmt.Printf("a is at %0x\n", &a)
	fmt.Printf("p is at %0x\n", &p)
	runtime.Free(p)
}

func f() *St {

	p := &St{A: 2, B: 3}
	return p
}

// to see the escape analysis:
// go build -gcflags -m free.go

// to activate prints of memory allocation/free, see
// ~/pkg/go1.3.3/go/src/pkg/runtime/malloc.goc
