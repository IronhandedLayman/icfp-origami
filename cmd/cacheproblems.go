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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cacheproblemsCmd represents the cacheproblems command
var cacheproblemsCmd = &cobra.Command{
	Use:   "cacheproblems",
	Short: "Retrieves all problem specifications for easier retrieval later",
	Long:  `Retrieves all problem specifications for easier retrieval later`,
	Run: func(cmd *cobra.Command, args []string) {
		serv := fsapi.NewBasicServer(viper.GetString("website"), viper.GetString("ApiKey"))

		snap, err := serv.LatestSnapshot()
		if err != nil {
			panic(fmt.Sprintf("Error retrieving latest snapshot:%v", err))
		}
		for _, ph := range snap.Problems {
			fmt.Printf("Cacheing Problem %d:<<%s>>\n", ph.ProblemId, ph.ProblemSpecHash)
			_, err = serv.GetBlob(ph.ProblemSpecHash)
			if err != nil {
				fmt.Printf("Warning: couldn't retrieve problem spec: %v\n", err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(cacheproblemsCmd)
}
