package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"nuxctl/nuagex"

	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

// TemplateID is the ID of a NuageX Template
var TemplateID string

// outTemplateFPath is a path to the file that will receive NuageX Template configuration dumped with `dump-template command`
var outTemplateFPath string

func init() {
	rootCmd.AddCommand(dumpTemplateCmd)

	dumpTemplateCmd.Flags().StringVarP(&TemplateID, "template-id", "i", "", "Template ID.")
	dumpTemplateCmd.MarkFlagRequired("template-id")

	dumpTemplateCmd.Flags().StringVarP(&CredFPath, "credentials", "c", "user_creds.yml", "Path to the user credentials file.")

	dumpTemplateCmd.Flags().StringVarP(&outTemplateFPath, "file", "f", "dumptemplate.yml", "Path to the local YAML file that will receive template configuration.")
}

var dumpTemplateCmd = &cobra.Command{
	Use:   "dump-template",
	Short: "Dump NuageX template configuration in a file.",
	Long:  `Dump the existing NuageX template configuration in a file.`,
	Run:   dumpTemplate,
}

func dumpTemplate(cmd *cobra.Command, args []string) {
	loginUser(cmd, args)

	fmt.Println("Retrieving NuageX Template configuration...")

	t, _, err := nuagex.GetTemplate(&user, TemplateID)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// set Template field to has ID value
	// and empty ID to exclude it from marshalling to YAML
	t.Template = t.ID
	t.ID = ""

	ty, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("Writing Template configuration to '%s' file...\n", outTemplateFPath)

	err = ioutil.WriteFile(outTemplateFPath, ty, 0644)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("Template configuration has been successfully written to '%s' file!\n", outTemplateFPath)
}
