package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type CLI struct {
	bc *Blockchain
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
}

func (cli *CLI) printUsage() {
	fmt.Print("Usage is: .....")
}

func (cli *CLI) Run() {
	cli.validateArgs()

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	historyCmd := flag.NewFlagSet("history", flag.ExitOnError)
	addData := addCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Printf("error: %v", err)
		}
	case "history":
		err := historyCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Printf("error: %v", err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addCmd.Parsed() {
		if *addData == "" {
			addCmd.Usage()
			os.Exit(1)
		}
		cli.add(*addData)
	}

	if historyCmd.Parsed() {
		cli.history()
	}
}

func (cli *CLI) add(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Done!")
}

func (cli *CLI) history() {
	bci := cli.bc.Manager()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW Validation: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
