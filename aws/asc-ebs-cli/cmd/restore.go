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
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/cobra"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:     "restore -sid <snapshot id> -zone <zone name>",
	Short:   "Restore EBS from snapshot",
	Aliases: []string{"r"},
	RunE: func(cmd *cobra.Command, args []string) error {
		ssId, err := cmd.Flags().GetString("sid")
		if err != nil {
			return err
		}
		if ssId == "" {
			return errors.New("empty snapshot id")
		}
		zone, err := cmd.Flags().GetString("zone")
		if err != nil {
			return err
		}
		return restoreAction(ssId, zone)
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)

	restoreCmd.PersistentFlags().StringP("sid", "i", "", "snapshot id")
	restoreCmd.PersistentFlags().StringP("zone", "z", "ap-southeast-1a", "zone name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func restoreAction(ssId, zone string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	ebs := ebsAction{
		ec2Client: ec2.NewFromConfig(cfg),
	}

	vid, err := ebs.restoreSnapshot(ssId, zone)
	if err != nil {
		return err
	}

	fmt.Printf("Volume [%s] was successfully restored from snapshot [%s].\n", *vid, ssId)
	return nil
}
