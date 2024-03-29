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

var (
	problemId int
	rawSpec   bool
)

// problemspecCmd represents the problemspec command
var problemspecCmd = &cobra.Command{
	Use:   "problemspec",
	Short: "Gets the problem spec for the problem id and returns it to the console",
	Long:  `Gets the problem spec for the problem id and returns it to the console`,
	Run: func(cmd *cobra.Command, args []string) {
		serv := fsapi.NewBasicServer(viper.GetString("website"), viper.GetString("ApiKey"))

		if rawSpec {
			rblob, err := serv.GetRawProblemSpec(problemId)
			if err != nil {
				panic(fmt.Sprintf("Error retrieving problem #%d: %v", problemId, err))
			}
			fmt.Printf("Raw problem spec:\n%v\n", rblob)

		} else {
			pblob, err := serv.GetProblemSpec(problemId)
			if err != nil {
				panic(fmt.Sprintf("Error retrieving problem #%d: %v", problemId, err))
			}
			fmt.Printf("Problem spec:\n%v\n", pblob)
		}
	},
}

func init() {
	RootCmd.AddCommand(problemspecCmd)
	problemspecCmd.PersistentFlags().IntVar(&problemId, "id", 0, "Problem ID of the problem spec to retrieve.")
	problemspecCmd.PersistentFlags().BoolVar(&rawSpec, "raw", false, "Returns the raw problem spec unparsed.")
}
