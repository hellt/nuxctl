package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"nuxctl/nuagex"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

// labID is the ID of a NuageX Lab (which is also the hostname name)
var labID string

// outLabFPath is a path to the file that will receive NuageX Lab configuration dumped with `dump-lab command`
var outLabFPath string

func init() {
	rootCmd.AddCommand(dumpLabCmd)

	dumpLabCmd.Flags().StringVarP(&labID, "lab-id", "i", "", "Lab ID. Seen as the variable portion in the lab hostname.")
	dumpLabCmd.MarkFlagRequired("lab-id")

	dumpLabCmd.Flags().StringVarP(&CredFPath, "credentials", "c", "user_creds.yml", "Path to the user credentials file.")

	dumpLabCmd.Flags().StringVarP(&outLabFPath, "file", "f", "dumplab.yml", "Path to the local YAML file that will receive lab configuration.")
}

var dumpLabCmd = &cobra.Command{
	Use:   "dump-lab",
	Short: "Dump NuageX lab (environment) configuration in a file.",
	Long:  `Dump the existing NuageX lab configuration in a file.`,
	Run:   dumpLab,
}

func dumpLab(cmd *cobra.Command, args []string) {
	loginUser(cmd, args)

	fmt.Println("Retrieving NuageX Lab configuration...")

	l, err := nuagex.DumpLab(&user, labID)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println("Parsing Lab configuration...")
	ly, err := yaml.Marshal(&l)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("Writing Lab configuration to '%s' file...\n", outLabFPath)

	err = ioutil.WriteFile(outLabFPath, ly, 0644)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("Lab configuration has been successfully written to '%s' file!\n", outLabFPath)

	// remove private network portion from the dumped lab file
	commentPrivateNetwork(outLabFPath)
}

func commentPrivateNetwork(fp string) {
	tmpPath := "tmp_dumplab.yml"
	tmpFile, err := os.Create(tmpPath)
	if err != nil {
		log.Fatal(err)
	}
	defer tmpFile.Close()

	file, err := os.Open(outLabFPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inNetsSection := false    // if inside `networks:` section
	inNetPrivSection := false // if inside private network section
	for scanner.Scan() {
		s := scanner.Text()
		if s == "networks:" {
			inNetsSection = true
		}

		if inNetsSection && (s != "networks:") && !(strings.HasPrefix(s, "-") || strings.HasPrefix(s, "  ")) {
			inNetsSection = false
		}

		if inNetsSection && s == "- name: private" {
			inNetPrivSection = true
			tmpFile.WriteString("# nuxctl: private network has been commented out, since it should not be a part of the lab configuration file used in `create-lab` command.\n# nuxctl: network `private` is always implicitly created by NuageX for every lab.\n")
		}

		if inNetPrivSection {
			if strings.HasPrefix(s, "  ") || s == "- name: private" {
				s = fmt.Sprintf("# %s", s) // add comment symbol for private network children
			} else {
				inNetPrivSection = false
			}

		}

		tmpFile.WriteString(fmt.Sprintf("%s\n", s))
	}

	err = os.Rename(tmpPath, outLabFPath)
	if err != nil {
		log.Fatal(err)
	}

}
