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
	"io/ioutil"

	"github.com/IronhandedLayman/icfp-origami/fsapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var probId int

// submit-solutionCmd represents the submit-solution command
var submitSolutionCmd = &cobra.Command{
	Use:   "submit-solution",
	Short: "Submit solution to a given problem to the fold system",
	Long:  `Submit solution to a given problem to the fold system`,
	Run: func(cmd *cobra.Command, args []string) {
		//attempt to load file
		solnBuf, err := ioutil.ReadFile(args[0])
		if err != nil {
			panic(fmt.Sprintf("Error while attempting to load problem file: %v", err))
		}

		fmt.Printf("Attempting to submit solution for problem %d...\n", probId)

		serv := fsapi.NewBasicServer(viper.GetString("website"), viper.GetString("ApiKey"))
		resp, err := serv.SolutionSubmission(probId, string(solnBuf))
		if err != nil {
			panic(fmt.Sprintf("Error while submitting solution: %v", err))
		}
		fmt.Printf("Response: %s\n", resp)
	},
}

func init() {
	RootCmd.AddCommand(submitSolutionCmd)

	submitSolutionCmd.PersistentFlags().IntVar(&probId, "ProbID", 0, "Problem ID of problem that solution solves")

	if berr := viper.BindPFlags(submitSolutionCmd.PersistentFlags()); berr != nil {
		panic(fmt.Sprintf("Error binding flags: %v", berr))
	}

}
