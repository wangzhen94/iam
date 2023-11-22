package main

import (
	"github.com/wangzhen94/iam/internal/authzserver"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	authzserver.NewApp("iam-authz-server").Run()
}
