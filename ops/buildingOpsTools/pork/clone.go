package pork

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	ref    string
	create bool
)

var CloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "clone repository from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must supply the repository")
		}
		if err := CloneRepository(args[0], ref, create); err != nil {
			log.Fatalln("error while cloning repository: ", err)
		}
	},
}

func init() {
	CloneCmd.PersistentFlags().StringVar(&ref, "ref", "", "specific reference to check out")
	CloneCmd.PersistentFlags().BoolVar(&create, "create", false, "create the reference if it does not exist")
}

func CloneRepository(repo, ref string, shouldCreate bool) error {
	return nil
}
