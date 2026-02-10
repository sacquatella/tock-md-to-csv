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
var IsMkdoc bool
var BaseUrl string
var Recursive bool

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
	rootCmd.PersistentFlags().BoolVarP(&IsMkdoc, "ismkdoc", "m", false, "Folder is a Mkdocs source")
	rootCmd.PersistentFlags().StringVarP(&BaseUrl, "base", "u", "http://localhost:9000", "Base URL for Mkdocs files")
	rootCmd.PersistentFlags().BoolVarP(&Recursive, "recursive", "r", false, "Parcourt récursivement tous les sous-dossiers")
}

// computeMarkdownFiles reads markdown files from the specified folder and generates a CSV file with their content.
// Each markdown file should have a front matter section with a title and a site URL.
// The text content of the markdown file is also included in the CSV file.
// The CSV file will have three columns: title, source, and text.
// If the folder is an Mkdocs source, it will build the URL based on the current file path.
func computeMarkdownFiles(cmd *cobra.Command, args []string) {

	var err error
	var allFiles []string
	// on construit la liste des fichiers markdown dans le dossier spécifié
	if Recursive {
		//var allFiles []string
		err = filepath.Walk(Folder, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(info.Name()) == ".md" {
				allFiles = append(allFiles, path)
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Error reading directory %s: %s", Folder, err)
			return
		}
	} else {
		entries, err := os.ReadDir(Folder)
		if err != nil {
			fmt.Printf("Error reading directory %s : %s", Folder, err)
			return
		}
		for _, entry := range entries {
			if !entry.IsDir() {
				allFiles = append(allFiles, filepath.Join(Folder, entry.Name()))
			}
		}
	}

	// on traite les fichiers markdown
	// et on génère le fichier CSV
	var records [][]string
	records = append(records, []string{"title", "source", "text"})

	for _, file := range allFiles {

		if filepath.Ext(file) == ".md" {
			content, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			var fm FrontMatter
			parts := strings.SplitN(string(content), "---", 3)
			if len(parts) < 3 {
				fmt.Println("Malformed markdown file, no frontmatter:", file)
				continue
			}
			err = yaml.Unmarshal([]byte(parts[1]), &fm)
			if err != nil {
				fmt.Println("Error parsing front matter:", err, file)
				continue
			}

			if fm.Title == "" {
				// Si le titre est vide, on utilise le premier titre # trouvé dans le markdow
				fm.Title = extractFirstTitle(string(content))
				fmt.Printf("title is empty on file %s. Try to get it form markdown titles: %s \n", file, fm.Title)
			}
			text := strings.TrimSpace(string(content))
			text = strings.ReplaceAll(text, "\n", " ")
			text = strings.ReplaceAll(text, "\"", "'")
			if IsMkdoc {
				fm.SiteURL = buildMkdocUrl(Folder, file, BaseUrl)
			}

			records = append(records, []string{fm.Title, fm.SiteURL, text})
		}
	}

	csvFile, err := os.Create(Output)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	writer.Comma = '|'
	defer writer.Flush()

	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error writing to CSV file:", err)
			return
		}
	}

	fmt.Println("CSV file created successfully.")
}

// buildMkdocUrl Build html url base on current mkdocs file path
func buildMkdocUrl(baseFolder string, currentFolder string, baseUrl string) string {
	// remove folder name from baseFolder, ie /home/user/docs/ to /home/user
	baseFolder = strings.TrimSuffix(baseFolder, string(os.PathSeparator))

	relPath, err := filepath.Rel(baseFolder, currentFolder)
	if err != nil {
		return baseUrl
	}
	relPath = strings.TrimSuffix(relPath, filepath.Ext(relPath))
	relPath = strings.ReplaceAll(relPath, string(os.PathSeparator), "/")
	if !strings.HasSuffix(baseUrl, "/") {
		baseUrl += "/"
	}
	return baseUrl + relPath + "/"
}

// extractFirstTitle extracts the first title from a markdown string.
func extractFirstTitle(markdown string) string {
	lines := strings.Split(markdown, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "# ") || strings.HasPrefix(trimmed, "## ") {
			// Remove the # and space
			return strings.TrimSpace(trimmed[strings.Index(trimmed, " "):])
		}
	}
	return ""
}
