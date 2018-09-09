package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"nuxctl/nuagex"

	"github.com/spf13/cobra"
)

// LabFPath is a path to the lab definition file
var LabFPath string

func init() {
	rootCmd.AddCommand(createLabCmd)

	createLabCmd.Flags().StringVarP(&CredFPath, "credentials", "c", "user_creds.yml", "Path to the user credentials file")

	createLabCmd.Flags().StringVarP(&LabFPath, "lab-configuration", "l", "lab.yml", "Path to the Lab configuration file")
}

var createLabCmd = &cobra.Command{
	Use:    "create-lab",
	Short:  "Create NuageX lab (environment)",
	Long:   `Create NuageX lab using the lab definition supplied in various formats`,
	PreRun: loginUser,
	Run:    createLab,
}

func createLab(cmd *cobra.Command, args []string) {
	lab.Conf(LabFPath)

	j, err := json.Marshal(lab)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("Sending request to create a lab...\n")
	nuagex.CreateLab(&user, j)
	fmt.Printf("Lab has been successfully queued for creation!\n")
}
