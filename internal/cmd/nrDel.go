/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"github.com/cmingou/nrsim/internal/api"
	"github.com/spf13/cobra"
)

// nrDelCmd represents the del command
var nrDelCmd = &cobra.Command{
	Use:   "del",
	Short: "Del NR",
	Long:  `A command to del NR.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("NR del called,\ngNB struct value is: \n%+v\n", nr)
		if gNbId == NonExisted {
			errLog.Printf("Please specify the worker id.")
			return
		}

		id := &api.IdMessage{Id: uint32(gNbId)}

		ctx, cancel := context.WithTimeout(context.Background(), GrpcConnectTimeout)
		defer cancel()

		client := GetCliServerClient()
		if _, err := client.DelGnb(ctx, id); err != nil {
			dealError(err)
		}
	},
}

func init() {
	nrCmd.AddCommand(nrDelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	nrDelCmd.Flags().IntVarP(&gNbId, "id", "i", NonExisted, "Id of NR")
}
