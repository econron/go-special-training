package main

import (
	"fmt"
	"st/pointer"
	"runtime"
)

func main() {

	printMemUsage()

	fmt.Println("Hello, World!")
	pointer.Pointer()

	anime := pointer.SetAnime("陰の実力者になりたくて", "2022年10月から", "逢沢大介")
	printMemUsage()
	pointer.PrintAnimeWithStructCopy(anime)
	printMemUsage()
	pointer.PrintAnimeWithoutStructCopy(&anime)
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// バイト単位で出力
	fmt.Printf("Alloc = %v MiB", m.Alloc / 1024 / 1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc / 1024 / 1024)
	fmt.Printf("\tSys = %v MiB", m.Sys / 1024 / 1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}