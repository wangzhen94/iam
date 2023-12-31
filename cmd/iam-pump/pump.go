// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// pump is iam analytics purger that moves the data generated by your iam-authz-server nodes to any back-end.
// It is primarily used to display your analytics data in the iam operating system.
package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/wangzhen94/iam/internal/pump"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	pump.NewApp("iam-pump").Run()
}
