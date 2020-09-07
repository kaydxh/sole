// Copyright 2020 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package down

import (
	"fmt"

	"github.com/searKing/sole/internal/pkg/provider/viper"
	"github.com/spf13/cobra"
)

// represent the down command
func New() *cobra.Command {

	return &cobra.Command{
		Use:   "down",
		Short: "Migrates downgrade sql",
		Long: fmt.Sprintf(`Run this command when you downgrade %[1]s to an old minor version. For example,
downgrading %[1]s 0.8.0 to 0.7.0 requires running this command.

It is recommended to run this command close to the SQL instance (e.g. same subnet) instead of over the public internet.
This decreases risk of failure and decreases time required.

### WARNING ###

Before running this command on an existing database, create a back up!
`, viper.ServiceName),
		Run: controller(),
	}
}
