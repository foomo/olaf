package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/foomo/olaf/backend/cmd"
	"github.com/foomo/olaf/backend/services/helloworld"
)

func main() {
	flagAddr := cmd.FlagAddr()
	flag.Parse()
	as := &helloworld.AdminService{}
	asProxy := helloworld.NewDefaultAdminServiceGoTSRPCProxy(as, []string{"http://localhost:3000"})
	fmt.Println(http.ListenAndServe(*flagAddr, asProxy))
}
