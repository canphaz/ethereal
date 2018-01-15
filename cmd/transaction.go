// Copyright © 2017 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
)

var transactionStr string

// transactionCmd represents the transaction command
var transactionCmd = &cobra.Command{
	Use:     "transaction",
	Aliases: []string{"tx"},
	Short:   "Manage transactions",
	Long:    `Send transactions, obtain information about transactions, regas transactions etc.`,
}

func init() {
	RootCmd.AddCommand(transactionCmd)
}
func transactionFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&transactionStr, "transaction", "t", "", "raw transaction data or ID of the transaction")
}
