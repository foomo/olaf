package cmd

import "flag"

func FlagAddr() *string {
	return flag.String("addr", ":8080", "address to listen to")
}
