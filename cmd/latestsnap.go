// Copyright © 2016 Nathaniel Dean <nathaniel.dean@alumni.purdue.edu>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/IronhandedLayman/icfp-origami/fsapi"
	"github.com/IronhandedLayman/icfp-origami/objs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// latestsnapCmd represents the latestsnap command
var latestsnapCmd = &cobra.Command{
	Use:   "latestsnap",
	Short: "Gets the latest snapshot and returns it to the console",
	Long:  `Gets the latest snapshot and returns it to the console`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Sending ping to folding server...\n")
		serv := fsapi.NewBasicServer(viper.GetString("website"), viper.GetString("ApiKey"))
		resp, err := serv.SnapshotListRequest()
		if err != nil {
			panic(fmt.Sprintf("Error while requesting snapshot list: %v", err))
		}
		fmt.Printf("Obtained snapshot list: %s\n", resp.Ok)
		var lsh *objs.SnapshotHash
		for _, sh := range resp.Snapshots {
			if lsh == nil || sh.SnapshotTime.After(lsh.SnapshotTime) {
				lsh = sh
			}
		}
		fmt.Printf("Obtaining snapshot %v::%s\n", lsh.SnapshotTime, lsh.SnapshotHash)
		blresp, berr := serv.GetBlob(lsh.SnapshotHash)
		if berr != nil {
			panic(fmt.Sprintf("Error obtaining snapshot blob: %v", berr))
		}
		fmt.Printf("Snapshot Blob: %s\n", blresp)
	},
}

func init() {
	RootCmd.AddCommand(latestsnapCmd)

}
