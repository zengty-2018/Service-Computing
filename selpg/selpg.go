package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

type selpg_arg struct {
	startPage  int
	endPage    int
	inFilename string
	pageLen    int
	pageType   string

	printDest string
}

func main() {
	arg := new(selpg_arg)
	Init(arg)

	check(arg)

	/*fmt.Println(arg.startPage)
	fmt.Println(arg.endPage)
	fmt.Println(arg.printDest)
	fmt.Println(arg.inFilename)
	fmt.Println(arg.pageLen)
	fmt.Println(arg.pageType)
	*/
	process_input(arg)
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:", os.Args[0])
	fmt.Fprintf(os.Stderr, "-s(start_page) -e(end_page) [-f | -l(lines_per_page)] [-d(dest)] [in_filename]")
	pflag.PrintDefaults()
}

func Init(arg *selpg_arg) {
	var exist_f bool

	pflag.BoolVarP(&exist_f, "fExist", "f", false, "exist of f")
	pflag.Usage = usage
	pflag.IntVarP(&(arg.startPage), "start", "s", -1, "start page")
	pflag.IntVarP(&(arg.endPage), "end", "e", -1, "end page")
	pflag.IntVarP(&(arg.pageLen), "line", "l", 72, "page len")
	pflag.StringVarP(&(arg.printDest), "destionation", "d", "", "print destionation")
	//pflag.StringVarP(&(arg.pageType), "type", "f", "l", "type of print")

	pflag.Parse()

	othersArg := pflag.Args()
	//fmt.Println(othersArg)
	if len(othersArg) > 0 {
		arg.inFilename = othersArg[0]
	} else {
		arg.inFilename = ""
	}

	if exist_f {
		arg.pageType = "f"
	} else {
		arg.pageType = "l"
	}
}

func check(arg *selpg_arg) {
	if arg.startPage < 0 || arg.startPage > math.MaxInt32-1 {
		fmt.Fprintf(os.Stderr, "ERROR: start page invaild.\n")
		usage()
		os.Exit(1)
	}
	if arg.endPage < 0 || arg.endPage > math.MaxInt32-1 || arg.endPage < arg.startPage {
		fmt.Fprintf(os.Stderr, "ERROR: end page invaild.\n")
		usage()
		os.Exit(2)
	}
	if arg.pageLen < 0 || arg.pageLen > math.MaxInt32 {
		fmt.Fprintf(os.Stderr, "ERROR: page len invaild.\n")
		usage()
		os.Exit(3)
	}
	if !(arg.pageType == "f" || arg.pageType == "l") {
		fmt.Fprintf(os.Stderr, "ERROR: page type invaild.\n")
		usage()
		os.Exit(4)
	}
}

func process_input(arg *selpg_arg) {
	var fin *os.File
	if arg.inFilename != "" {
		var opErr error
		fin, opErr = os.Open(arg.inFilename)
		if opErr != nil {
			fmt.Fprintf(os.Stderr, "\nERROR! Can not open the input file: %s\n", arg.inFilename)
			os.Exit(9)
		}
	} else {
		fin = os.Stdin
	}
	finBuffer := bufio.NewReader(fin)

	var fout io.WriteCloser
	var cmd *exec.Cmd
	if arg.printDest != "" {
		cmd = exec.Command("lp", "-d", arg.printDest)
		var desErr error
		cmd.Stdout, desErr = os.OpenFile(arg.printDest, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		fout, desErr = cmd.StdinPipe()
		if desErr != nil {
			fmt.Fprintf(os.Stderr, "\nERROR! Can not open pipe to \"lp -d%s\"\n", arg.printDest)
			os.Exit(5)
		}
		cmd.Start()
		cmd.Wait()
	} else {
		fout = os.Stdout
	}

	if arg.pageType == "l" {
		lineCtr := 0
		pageCtr := 1
		for {
			line, crc := finBuffer.ReadString('\n')

			if crc != nil {

				break
			}
			lineCtr++
			if lineCtr > arg.pageLen {
				pageCtr++
				lineCtr = 1
			}

			if (pageCtr >= arg.startPage) && (pageCtr <= arg.endPage) {
				_, err := fout.Write([]byte(line))
				if err != nil {
					fmt.Fprintf(os.Stderr, "Write ERROR!")
					os.Exit(6)
				}
			}
		}
	} else {
		pageCtr := 1
		for {
			page, crc := finBuffer.ReadString('\f')

			if crc != nil {
				fmt.Println(page)
				fmt.Fprintf(os.Stderr, "Can not find \\f!")
				break
			}
			pageCtr++
			if (pageCtr >= arg.startPage) && (pageCtr <= arg.endPage) {
				_, err := fout.Write([]byte(page))
				if err != nil {
					fmt.Fprintf(os.Stderr, "Write ERROR!")
					os.Exit(7)
				}
			}
		}
	}

	fin.Close()
	fout.Close()
}
