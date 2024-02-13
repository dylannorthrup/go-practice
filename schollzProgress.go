package main

import (
	"fmt"
	"time"

	ansi "github.com/k0kubun/go-ansi"
	pb "github.com/schollz/progressbar/v3"
)

var (
	bar     *pb.ProgressBar
	barOpts []pb.Option
)

func schollzInit() {
	barOpts = make([]pb.Option, 0)
}

func setBarTheme() pb.Theme {
	var theme pb.Theme
	theme.Saucer = "[green]=[reset]"
	theme.SaucerHead = "[cyan]>[reset]"
	theme.SaucerPadding = "."
	theme.BarStart = "["
	theme.BarEnd = "]"
	return theme
}

func RunSchollz() {
	fmt.Println("Howdy")

	schollzInit()
	barTheme := setBarTheme()
	barOpts = append(barOpts, pb.OptionSetTheme(barTheme))
	barOpts = append(barOpts, pb.OptionSetWriter(ansi.NewAnsiStdout()))
	barOpts = append(barOpts, pb.OptionEnableColorCodes(true))
	barOpts = append(barOpts, pb.OptionUseANSICodes(true))
	barOpts = append(barOpts, pb.OptionSetDescription("[cyan]Seconds[reset] "))
	barOpts = append(barOpts, pb.OptionFullWidth())
	barOpts = append(barOpts, pb.OptionSetElapsedTime(false))
	// barOpts = append(barOpts, pb.OptionSetX())
	// barOpts = append(barOpts, pb.OptionSetX())
	// barOpts = append(barOpts, pb.OptionShowBytes(true)) // Show Bytes/sec
	// barOpts = append(barOpts, pb.OptionShowIts())  // Show iterations/sec

	bar = pb.NewOptions(100, barOpts...)
	for i := 0; i < 100; i++ {
		_ = bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}
	_ = bar.Finish()

	fmt.Println("")
	fmt.Println("schollzBar function execution complete")
}
