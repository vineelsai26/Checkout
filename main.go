package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"vineelsai.com/checkout/src"
	"vineelsai.com/checkout/src/project"
	"vineelsai.com/checkout/src/utils"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "checkout:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		printUsage()
		return fmt.Errorf("missing command")
	}

	switch args[0] {
	case "init":
		return runInit(args[1:])
	case "deinit":
		return runDeinit(args[1:])
	case "version", "-v", "--version":
		src.GetVersion()
		return nil
	case "help", "-h", "--help":
		printUsage()
		return nil
	default:
		printUsage()
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func runInit(args []string) error {
	flags, options, config, excludes := newCommandFlagSet("init")
	noOpen := flags.Bool("no-open", false, "do not open the checked-out project in VS Code")
	if err := flags.Parse(args); err != nil {
		return err
	}
	applyConfig(config)
	options.Copy.Excludes = excludes.values
	options.OpenVS = !*noOpen

	projectDir := ""
	if flags.NArg() == 0 {
		var err error
		projectDir, err = project.PromptForName()
		if err != nil {
			return err
		}
	} else if flags.NArg() == 1 {
		var err error
		projectDir, err = project.FindProject(flags.Arg(0))
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("init accepts at most one project name")
	}

	fmt.Println("Found project at:", projectDir)
	fmt.Println("Checking out project...")
	return project.Init(projectDir, *options)
}

func runDeinit(args []string) error {
	flags, options, config, excludes := newCommandFlagSet("deinit")
	if err := flags.Parse(args); err != nil {
		return err
	}
	applyConfig(config)
	options.Copy.Excludes = excludes.values
	options.OpenVS = false

	if flags.NArg() < 1 || flags.NArg() > 2 {
		return fmt.Errorf("deinit requires a project name and optional project folder")
	}

	projectFolder := utils.DefaultProjectSource
	if flags.NArg() == 2 {
		projectFolder = flags.Arg(1)
	}

	fmt.Println("Checking project back in...")
	return project.DeInit(flags.Arg(0), projectFolder, *options)
}

func newCommandFlagSet(name string) (*flag.FlagSet, *project.Options, *runtimeConfig, *stringList) {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)

	options := &project.Options{}
	flags.BoolVar(&options.Copy.IncludeEnv, "include-env", false, "include .env and .env.* files")
	flags.BoolVar(&options.Copy.DryRun, "dry-run", false, "print planned copy/delete actions without changing files")

	config := &runtimeConfig{}
	flags.StringVar(&config.sourceDir, "source-dir", utils.ProjectSourceDir, "project source root")
	flags.StringVar(&config.checkoutRoot, "checkout-root", utils.ProjectCheckoutRootDir, "project checkout root")

	excludes := &stringList{}
	flags.Var(excludes, "exclude", "additional file, directory, or glob pattern to exclude; repeatable")
	return flags, options, config, excludes
}

type runtimeConfig struct {
	sourceDir    string
	checkoutRoot string
}

func applyConfig(config *runtimeConfig) {
	utils.ProjectSourceDir = config.sourceDir
	utils.ProjectCheckoutRootDir = config.checkoutRoot
}

type stringList struct {
	values []string
}

func (value *stringList) String() string {
	return strings.Join(value.values, ",")
}

func (value *stringList) Set(input string) error {
	for _, part := range strings.Split(input, ",") {
		part = strings.TrimSpace(part)
		if part != "" {
			value.values = append(value.values, part)
		}
	}
	return nil
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Checkout moves local projects between the source tree and %s.

Usage:
  checkout init [flags] [project]
  checkout deinit [flags] <project> [source-folder]
  checkout version

Flags:
  --include-env       include .env and .env.* files
  --exclude value     additionally exclude a file, directory, or glob pattern; repeatable
  --source-dir path   project source root
  --checkout-root path
                      project checkout root
  --dry-run           show planned copy/delete actions without changing files
  --no-open           init only: do not open the checked-out project in VS Code

Default excludes:
  %s

`, utils.ProjectCheckoutRootDir, strings.Join(utils.DefaultExcludedNames, ", "))
}
