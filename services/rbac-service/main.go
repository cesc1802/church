package main

import "services.rbac-service/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		return
	}
}
