package main

import "rbac-service/cmd"

func main()  {
	if err := cmd.Execute(); err != nil {
		return
	}
}
