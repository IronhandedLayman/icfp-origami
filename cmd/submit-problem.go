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
	"time"

	"github.com/IronhandedLayman/icfp-origami/fsapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var timeToSubmit string

var submitProblemCmd = &cobra.Command{
	Use:   "submit-problem",
	Short: "Submit a problem to the fold server",
	Long:  `Using a schedule time and solution spec file, submit problem with solution to the fold server`,
	Run: func(cmd *cobra.Command, args []string) {
		//attempt to load file
		probBuf, err := ioutil.ReadFile(args[0])
		if err != nil {
			panic(fmt.Sprintf("Error while attempting to load problem file: %v", err))
		}

		tts, err := time.Parse("2006-01-02T15:04", timeToSubmit)
		if err != nil {
			panic(fmt.Sprintf("Illegal time for --schedule: [[%s]] <<%v>>", timeToSubmit, err))
		}
		fmt.Printf("Attempting to submit problem for presentation at %s...\n", tts.Format("2006-01-02T15:04"))

		serv := fsapi.NewBasicServer(viper.GetString("website"), viper.GetString("ApiKey"))
		resp, err := serv.ProblemSubmission(string(probBuf), tts)
		if err != nil {
			panic(fmt.Sprintf("Error while submitting problem: %v", err))
		}
		fmt.Printf("Response: %s\n", resp)
	},
}

func init() {
	RootCmd.AddCommand(submitProblemCmd)

	submitProblemCmd.PersistentFlags().StringVar(&timeToSubmit, "schedule", "", "When to schedule the problem's release, in 2006-01-02T15:04 format.")

	if berr := viper.BindPFlags(submitProblemCmd.PersistentFlags()); berr != nil {
		panic(fmt.Sprintf("Error binding flags: %v", berr))
	}

}
