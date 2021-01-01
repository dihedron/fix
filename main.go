package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jessevdk/go-flags"
)

// Options contains the application options.
type Options struct {
	Simulate bool `short:"s" long:"simulate" description:"Run in simulation mode (no renaming)."`
	NoDirs   bool `short:"d" long:"skip-dirs" description:"Skip directory renaming."`
	NoFiles  bool `short:"f" long:"skip-files" description:"Skip MP3 file renaming."`
}

type rename struct {
	from string
	to   string
}

func main() {

	opts := Options{}
	dirs, _ := flags.Parse(&opts)
	if len(dirs) == 0 {
		dirs = append(dirs, ".")
	}

	renames := []rename{}

	if !opts.NoDirs {
		re := regexp.MustCompile(`^([^-]+\S)(?:-)(\S[^-]+)$`)
		for _, dir := range dirs {
			err := filepath.Walk(dir,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if info.IsDir() && info.Name() != "." && info.Name() != ".." && info.Name() != "albumart_backup" {
						groups := re.FindAllStringSubmatch(info.Name(), -1)
						if len(groups) > 0 {
							name := fmt.Sprintf("%s - %s", strings.ReplaceAll(groups[0][1], "_", " "), strings.ReplaceAll(groups[0][2], "_", " "))
							newPath := strings.Replace(path, info.Name(), name, 1)
							renames = append(renames, rename{
								from: path,
								to:   newPath,
							})
						}
					}
					return nil
				})
			if err != nil {
				log.Println(err)
			}
		}
	}

	if !opts.NoFiles {
		re := regexp.MustCompile(`^((\d{0,2})\.)(.*)\.(mp3)$`)
		for _, dir := range dirs {
			err := filepath.Walk(dir,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if info.Name() != "." && info.Name() != ".." {
						groups := re.FindAllStringSubmatch(info.Name(), -1)
						if len(groups) > 0 {
							name := fmt.Sprintf("%s - %s.%s", groups[0][2], strings.ReplaceAll(groups[0][3], "_", " "), groups[0][4])
							name = strings.ReplaceAll(name, "  ", " - ")
							newPath := strings.Replace(path, info.Name(), name, 1)
							renames = append(renames, rename{from: path, to: newPath})
						}
					}
					return nil
				})
			if err != nil {
				log.Println(err)
			}
		}
	}

	for i := len(renames) - 1; i >= 0; i-- {
		fmt.Printf("renaming %q to %q...", renames[i].from, renames[i].to)
		if !opts.Simulate {
			if err := os.Rename(renames[i].from, renames[i].to); err != nil {
				fmt.Println(" KO!")
			} else {
				fmt.Println(" OK!")
			}
		} else {
			fmt.Println(" SKIPPED!")
		}
	}
}
