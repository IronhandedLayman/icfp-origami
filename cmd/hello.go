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
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Sending ping to folding server...\n")
		client := http.Client{}
		reqaddr := fmt.Sprintf("http://%s/api/%s", viper.GetString("website"), "hello")
		req, mrerr := http.NewRequest("GET", reqaddr, nil)
		if mrerr != nil {
			panic("CODER ERROR: request malformed")
		}
		req.Header.Set("Expect", "")
		req.Header.Set("X-Api-Key", viper.GetString("ApiKey"))
		resp, err := client.Do(req)
		if err != nil {
			panic(fmt.Sprintf("Error while requesting: %v", err))
		}
		respBody, rerr := ioutil.ReadAll(resp.Body)
		if rerr != nil {
			panic(fmt.Sprintf("Error while reading response: %v", err))
		}
		fmt.Printf("Response: %s\n", string(respBody))
	},
}

func init() {
	RootCmd.AddCommand(helloCmd)
}