package main

import (
	"fmt"

	"github.com/like-org/common/app/interface/account/http"
)

func main() {

	addr := fmt.Sprintf("%s:%d", "0.0.0.0", 5051)

	appSrv := http.NewServer()

	appSrv.Start(addr)
}
