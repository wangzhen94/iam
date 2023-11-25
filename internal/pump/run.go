// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pump

import (
	"github.com/wangzhen94/iam/internal/pump/config"
)

// Run runs the specified pump server. This should never exit.
func Run(cfg *config.Config, stopCh <-chan struct{}) error {

	server, err := createPumpServer(cfg)
	if err != nil {
		return err
	}

	return server.PrepareRun().Run(stopCh)
}
