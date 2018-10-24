package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	commands = []*cmd{
		cmdNew,
		cmdRenew,
		cmdVersion,
	}
	exitMu     sync.Mutex
	exitStatus = 0
)

func main() {
	flag.Usage = usage
	flag.Parse() // catch -h argument
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	args := flag.Args()

	if len(args) < 1 {
		usage()
	}

	if args[0] == "help" {
		help(args[1:])
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Runnable() {
			addFlags(&cmd.flag)
			cmd.flag.Usage = func() { cmd.Usage() }
			cmd.flag.Parse(args[1:])
			cmd.run(cmd.flag.Args())
			exit()
			return
		}
	}

	fatalf("Unknown subcommand %q.\nRun 'ngx help' for usage.\n", args[0])
}

type cmd struct {
	run       func(args []string)
	flag      flag.FlagSet
	UsageLine string
	Short     string
	Long      string
}

func (c *cmd) Name() string {
	name := c.UsageLine
	i := strings.IndexRune(name, ' ')
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *cmd) Usage() {
	help([]string{c.Name()})
	os.Exit(2)
}

func (c *cmd) Runnable() bool {
	return c.run != nil
}

var logf = log.Printf

func errorf(format string, args ...interface{}) {
	logf(format, args...)
	setExitStatus(1)
}

func fatalf(format string, args ...interface{}) {
	errorf(format, args...)
	exit()
}

func setExitStatus(n int) {
	exitMu.Lock()
	if exitStatus < n {
		exitStatus = n
	}
	exitMu.Unlock()
}

func exit() {
	os.Exit(exitStatus)
}

func addFlags(f *flag.FlagSet) {
	f.StringVar(&configDir, "configDir", configDir, "")
	f.StringVar(&directoryURL, "directoryURL", directoryURL, "")
	f.StringVar(&resourceURL, "resourceURL", resourceURL, "")
	f.StringVar(&siteConfDir, "siteConfDir", siteConfDir, "")
	f.StringVar(&siteRootDir, "siteRootDir", siteRootDir, "")
	f.IntVar(&allowRenewDays, "allowRenewDays", allowRenewDays, "")
}
