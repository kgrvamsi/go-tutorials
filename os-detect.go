package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
)

func main() {
	var m runtime.MemStats

	/*If that is not working, or it is too much time, you can add a periodic call to
	FreeOSMemory (no need to call runtime.GC() before, it is done by debug.FreeOSMemory() )

		Something like this: http://play.golang.org/p/mP7_sMpX4F

		package main

	import (
	    "runtime/debug"
	    "time"
	)

	func main() {
	    go periodicFree(1 * time.Minute)

	    // Your program goes here

	}

	func periodicFree(d time.Duration) {
	    tick := time.Tick(d)
	    for _ = range tick {
	        debug.FreeOSMemory()
	    }
	}
	Take into account that every call to FreeOSMemory will take some time (not much)
	and it can be partly run in parallel if GOMAXPROCS>1 since Go1.3.*/

	debug.FreeOSMemory()

	/*Then you can either render it to a dot file with graphical
	representation of the heap or convert it to hprof format. To render it to a dot file:

	$ go get github.com/randall77/hprof/dumptodot

	$ dumptodot heapdump mybinary > heap.dot

	and open heap.dot with Graphviz.*/

	f, err := os.Create("heapdump")
	if err != nil {
		panic(err)
	}

	debug.WriteHeapDump(f.Fd())

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.GOARCH)
	runtime.ReadMemStats(&m)

	fmt.Println(m.TotalAlloc)
	fmt.Println(m.Alloc)
	fmt.Println(m.Sys)
}
