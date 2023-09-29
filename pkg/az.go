package az

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

var jsonFlag bool
var detailFlag bool
var hiddenFlag bool

type FileInfo struct {
	Name     string    `json:"name"`
	Size     int64     `json:"size"`
	Modified time.Time `json:"modified"`
	Type     string    `json:"type"`
	Path     string    `json:"path"`
}

func searchForFiles(_ *cobra.Command, args []string) {
	pattern := "*"
	dirPath := "."

	var files []FileInfo

	if len(args) == 1 {
		pattern = args[0]
	} else if len(args) == 2 {
		pattern = args[0]
		dirPath = args[1]
	}

	fileInfo, err := os.Stat(dirPath)

	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	if !fileInfo.IsDir() {
		fmt.Printf("%s is not a directory", dirPath)
	}

	err = filepath.WalkDir(dirPath, func(path string, info os.DirEntry, err error) error {
		if path == dirPath {
			return nil
		}

		if err != nil {
			return err
		}

		// Get filename (last element of path)
		name := filepath.Base(path)

		// Skip hidden files/directories
		if !hiddenFlag && name[0] == '.' {
			if info.IsDir() {
				return filepath.SkipDir
			} else {
				return nil
			}
		}

		// Check if the file matches the pattern
		matched, _ := filepath.Match(pattern, info.Name())
		if matched {
			var fileInfo FileInfo

			if detailFlag {
				fileInfo = FileInfo{
					Name:     info.Name(),
					Size:     fileInfo.Size,
					Modified: fileInfo.Modified,
					Path:     path,
				}
				if info.IsDir() {
					fileInfo.Type = "dir"
				} else {
					fileInfo.Type = "file"
				}

			} else {
				fileInfo.Path = path
			}

			files = append(files, fileInfo)
		}

		return nil
	})

	if jsonFlag {
		if detailFlag {
			jsonFiles, _ := json.Marshal(files)
			fmt.Printf("%s", string(jsonFiles))
		} else {
			var pathList []map[string]string
			for _, file := range files {
				path := map[string]string{"path": file.Path}
				pathList = append(pathList, path)
			}
			jsonPathList, _ := json.Marshal(pathList)
			fmt.Printf("%s", string(jsonPathList))
		}
	} else {
		if detailFlag {
			for _, e := range files {
				fmt.Printf("%s\n\tName: %s\n\tSize: %s\n\tDate modified: %s\n\tFile type: %s\n\n", e.Path, filepath.Base(e.Path), humanize.BigBytes(big.NewInt(e.Size)), e.Modified, e.Type)
			}
		} else {
			for _, e := range files {
				fmt.Printf("%s\n", e.Path)
			}
		}
	}

	if err != nil {
		fmt.Printf("%s", err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "az [pattern] [path]",
	Short: "Search for files matching a pattern",
	Args:  cobra.RangeArgs(0, 2),
	Run:   searchForFiles,
}

func Execute() {

	rootCmd.PersistentFlags().BoolVarP(&jsonFlag, "json", "j", false, "output results in JSON format")
	rootCmd.PersistentFlags().BoolVarP(&detailFlag, "detail", "d", false, "results return path, filename, size, date modified, and file type")
	rootCmd.PersistentFlags().BoolVarP(&hiddenFlag, "hidden", "H", false, "search hidden files and directories")

	rootCmd.Flags().SortFlags = false

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
