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

func init() {

	if !flag.Parsed() {
		flag.Parse()
	}

	if strings.Contains(*godaemon, "true") {
		cmd := exec.Command(os.Args[0])
		if flag.NArg() >= 1 {
			cmd = exec.Command(os.Args[0], flag.Args()[1:]...)
		}
		cmd.Args = append(cmd.Args,"-c",os.Args[3])
		fmt.Println(cmd.Args)
		fmt.Println(os.Args)
		cmd.Start()
		fmt.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
		*godaemon = "false"
		os.Exit(0)
	}
}
