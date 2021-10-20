package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type projectConfig struct {
	Name         string
	localPath    string
	RepoURL      string
	StaticAssets bool
}

func setupParseFlags(w io.Writer, args []string) (projectConfig, error) {
	conf := projectConfig{}
	// This is how a local FlagSet object is setup instead of using the global
	// flagset which is difficult to write unit tests for.
	fs := flag.NewFlagSet("scaffold-gen", flag.ContinueOnError)
	// This function sets the output of the flagset to the io.Writer object
	// that is passed in as a parameter to the function
	fs.SetOutput(w)
	fs.StringVar(&conf.Name, "n", "", "Project name")
	fs.StringVar(&conf.localPath, "d", "", "Project location on disk")
	fs.StringVar(&conf.RepoURL, "r", "", "Project remote repository URL")
	fs.BoolVar(&conf.StaticAssets, "s", false, "Project will have static assets or not")
	// args := []{"-n", "myproject", "-d", /my/path -r github.com/amitsaha/test-project -s"}
	err := fs.Parse(args)
	if err != nil {
		return conf, err
	}
	if fs.NArg() != 0 {
		return conf, errors.New("no positional parameters expected")
	}
	return conf, err
}

// generateScaffold for now will only just display a message and will be used
// as the entrypoint function for creating the scaffold structure in Milestones
// 3 and 4
func generateScaffold(w io.Writer, conf projectConfig) {
	//fmt.Printf()
	fmt.Fprintf(w, "Generating scaffold for project %s in %s\n", conf.Name, conf.localPath)
}

func validateConf(conf projectConfig) []error {
	var validationErrors []error
	if len(conf.Name) == 0 {
		validationErrors = append(validationErrors, errors.New("Project name cannot be empty"))
	}
	if len(conf.localPath) == 0 {
		validationErrors = append(validationErrors, errors.New("Project path cannot be empty"))
	}
	if len(conf.RepoURL) == 0 {
		validationErrors = append(validationErrors, errors.New("Project repository URL cannot be empty"))
	}
	return validationErrors
}

func main() {
	conf, err := setupParseFlags(os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	errors := validateConf(conf)
	if len(errors) != 0 {
		for _, e := range errors {
			fmt.Println(e)
		}
		os.Exit(1)
		// log.Fatal("exited because of something")
	}

	generateScaffold(os.Stdout, conf)
}
