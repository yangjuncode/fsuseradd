package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)
import "github.com/jessevdk/go-flags"

//go:embed 1001.xml
var userTempalteXML []byte

var opts struct {

	//start number
	Start int ` short:"s" long:"start" description:"start number"`
	//end number
	End int ` short:"e" long:"end" description:"end number"`

	//output dir
	OutputDir string ` short:"o" long:"output" description:"output dir"`

	Help bool ` short:"h" long:"help" description:"show this help message"`
	//remove files
	RemoveFiles bool `short:"r" long:"remove" description:"remove files"`
	Version     bool `short:"v" long:"version" description:"show version message"`
}

func main() {

	parser := flags.NewParser(&opts, flags.IgnoreUnknown)
	_, err := parser.Parse()

	if err != nil {
		panic(err)
	}

	start := opts.Start
	if start <= 0 {
		start = 1
	}
	for start <= opts.End {
		processingOneFile(start)
		start++
	}

	if opts.Help {
		parser.WriteHelp(os.Stdout)
		return
	}

	if opts.Version {
		fmt.Println("version: 1.0.0")
		return
	}

}

func processingOneFile(start int) {
	filePath := filepath.Join(opts.OutputDir, fmt.Sprintf("%d", start)+".xml")

	if len(opts.OutputDir) == 0 {
		filePath = fmt.Sprintf("%d", start) + ".xml"
	}
	if opts.RemoveFiles {
		fmt.Println("remove file:", filePath)
		_ = os.Remove(filePath)
		return
	}

	fmt.Println("create file:", filePath)
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("create file err:", err, filePath)
		return
	}

	newXML := bytes.ReplaceAll(userTempalteXML, []byte("1001"), []byte(fmt.Sprintf("%d", start)))
	_, err = f.Write(newXML)
	if err != nil {
		fmt.Println("write file err:", err, filePath)
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println("close file err:", err, filePath)
		return
	}

}
