package main

import (
	"flag"
	"os"
	"log"
	"runtime/pprof"
	"io"
	"fmt"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")


func main() {
	//runTest()
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		var err error
		memFile, err := os.Create(*memprofile)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("start write heap profile....")
			pprof.WriteHeapProfile(memFile)
			defer memFile.Close()
		}
	}

	fi,err := os.Open("ec_analytics2.test.exe")
	if err != nil{
		panic(err)
	}
	defer fi.Close()



	buf := make([]byte,1024)

	tt := make([]int, 10000000)
	i := 0
	for {
		if i == 100 {
			break
		}
		tt[i] = i
		i++
	}
fmt.Println(i)

	for{
		n,err := fi.Read(buf)
		if err != nil && err != io.EOF{panic(err)}
		if 0 ==n {break}
	}

}


