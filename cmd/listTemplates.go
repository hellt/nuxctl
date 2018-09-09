package cmd

import (
	"fmt"
	"log"
	"nuxctl/nuagex"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listTemplatesCmd)

	listTemplatesCmd.Flags().StringVarP(&CredFPath, "credentials", "c", "user_creds.yml", "Path to the user credentials file.")
}

var listTemplatesCmd = &cobra.Command{
	Use:   "list-templates",
	Short: "Display NuageX Lab templates.",
	Long:  `Outputs to console the list of availalbe NuageX Lab templates.`,
	Run:   listTemplates,
}

func listTemplates(cmd *cobra.Command, args []string) {
	loginUser(cmd, args)

	fmt.Println("Retrieving NuageX Lab Templates...")

	templates, err := nuagex.GetTemplates(&user, labID)
	if err != nil {
		log.Fatalf("%v", err)
	}

	printTemplates(templates)
}

type byTemplateName []*nuagex.Template

func (x byTemplateName) Len() int           { return len(x) }
func (x byTemplateName) Less(i, j int) bool { return x[i].Name < x[j].Name }
func (x byTemplateName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func printTemplates(t []*nuagex.Template) {
	const format = "%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	sort.Sort(byTemplateName(t))
	fmt.Fprintf(tw, format, "ID", "Name")
	fmt.Fprintf(tw, format, "------------------------", "------------------------")
	for _, t := range t {
		fmt.Fprintf(tw, format, t.ID, t.Name)
	}
	tw.Flush()
}
