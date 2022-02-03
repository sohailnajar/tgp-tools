package main

import (
	"findold"
	"flag"
	"os"
)

/*
./findold -age 10 /dir
*/
func main() {
	var age int
	flag.IntVar(&age, "age", 0, "age in days")
	flag.Parse()
	fsys := os.DirFS("testdata/backup")
	// fmt.Println(findold.LastModDate("../testdata/backup"))
	findold.PrintOldFiles(fsys, age)

}
