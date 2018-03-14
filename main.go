package main

import "flag"

func main() {
	// Config DB
	dbPathOrName := flag.String("db", "database", "path or name of the database")
	flag.Parse()

	bc := NewBlockchain(*dbPathOrName)
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()

}
