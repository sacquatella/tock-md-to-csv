package cmd

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var Verbose bool
var Folder string
var Output string

type FrontMatter struct {
	Title   string `yaml:"title"`
	SiteURL string `yaml:"site_url"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "md-to-csv",
	Short: "Générate a CSV file from markdown files",
	Long: `Générate a CSV file from markdown files. The CSV file will contain the title, the source and the text of each markdown file.
	
	Example usage:
	md-to-csv -f samples -c output.csv`,
	Run: computeMarkdownFiles,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	Verbose = false
	rootCmd.PersistentFlags().StringVarP(&Folder, "folder", "f", ".", "Folder containing markdown files")
	rootCmd.PersistentFlags().StringVarP(&Output, "csv", "c", "output.csv", "CSV file to generate")
}

func computeMarkdownFiles(cmd *cobra.Command, args []string) {

	files, err := ioutil.ReadDir(Folder)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du répertoire %s : %s", Folder, err)
		return
	}

	var records [][]string
	records = append(records, []string{"title", "source", "text"})

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			content, err := ioutil.ReadFile(Folder + "/" + file.Name())
			if err != nil {
				fmt.Println("Erreur lors de la lecture du fichier:", err)
				continue
			}

			parts := strings.SplitN(string(content), "---", 3)
			if len(parts) < 3 {
				fmt.Println("Fichier markdown mal formé:", file.Name())
				continue
			}

			var fm FrontMatter
			err = yaml.Unmarshal([]byte(parts[1]), &fm)
			if err != nil {
				fmt.Println("Erreur lors de l'analyse du front matter:", err)
				continue
			}

			text := strings.TrimSpace(parts[2])
			text = strings.ReplaceAll(text, "\n", " ")
			records = append(records, []string{fm.Title, fm.SiteURL, text})
		}
	}

	csvFile, err := os.Create(Output)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier CSV:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	writer.Comma = '|'

	defer writer.Flush()

	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Erreur lors de l'écriture dans le fichier CSV:", err)
			return
		}
	}

	fmt.Println("Fichier CSV créé avec succès.")
}
