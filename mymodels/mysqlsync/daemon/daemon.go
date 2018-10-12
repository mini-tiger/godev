package godaemon

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//var godaemon = flag.Bool("d", false, "run app as a daemon with -d=true or -d true.")
var godaemon = flag.String("d", "false", "run app as a daemon with -d true or -d false.")
var cfgfile = flag.String("c", "cfg.json", "configuration file")

func init() {

	if !flag.Parsed() {
		flag.Parse()
	}

	if strings.Contains(*godaemon, "true") {
		cmd := exec.Command(os.Args[0])
		//fmt.Println(flag.NFlag())
		//fmt.Println(flag.NArg())
		if flag.NFlag() >= 1 {
			cmd = exec.Command(os.Args[0])
			cmd.Args = append(cmd.Args, fmt.Sprintf("-s %s", *cfgfile))
		}
		cmd.Start()
		fmt.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
		fmt.Println(cmd.Args)
		*godaemon = "false"
		os.Exit(0)
	} else {
		cmd := exec.Command(os.Args[0])

		cmd.Args = append(cmd.Args, fmt.Sprintf("-s %s", *cfgfile))
		cmd.Start()
	}
}
