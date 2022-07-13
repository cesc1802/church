package main

import "shard/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		return
	}
}
