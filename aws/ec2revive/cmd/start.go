/*
Copyright © 2023 De Thar Htun

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/cobra"

	"github.com/dtherhtun/Learning-go/aws/ec2revive/revive"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		startProcess()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func startProcess() {
	data := &revive.Instances{}
	if err := data.Load("./state.yaml"); err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client := revive.Client{
		EC2Client: ec2.NewFromConfig(cfg),
	}

	var instanceIDs []string
	//var snapshotIDs []string

	//for i := 0; i < len(data.InstanceList); i++ {
	//	instanceIDs = append(instanceIDs, data.InstanceList[i].ID)
	//	snapshotIDs = append(snapshotIDs, data.InstanceList[i].Snap)
	//}

	for i := 0; i < len(data.InstanceList); i++ {
		instanceIDs = append(instanceIDs, data.InstanceList[i].ID)
		restoreResult, err := client.RestoreVolume(ctx, data.InstanceList[i].Snap, data.InstanceList[i].AZ)
		if err != nil {
			fmt.Println(err)
		}
		if err := data.UpdateVolumeByInstanceID(data.InstanceList[i].ID, *restoreResult.VolumeId); err != nil {
			fmt.Println(err)
		}
	}
	device_name := "/dev/sdf"

	for i := 0; i < len(data.InstanceList); i++ {
		attachResult, err := client.AttachVolume(ctx, data.InstanceList[i].ID, device_name, data.InstanceList[i].VID)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("volumeID [%s] is successfully attached to the instanceID [%s]\n", *attachResult.VolumeId, *attachResult.InstanceId)
	}

	startResult, err := client.StartInstance(ctx, instanceIDs)
	for _, res := range startResult.StartingInstances {
		fmt.Printf("instanceID [%s] is running.\n", *res.InstanceId)
	}

	if err := data.Update("./state.yaml"); err != nil {
		fmt.Println(err)
	}
}
