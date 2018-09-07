package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"time"
	"net"
	"log"
	"strings"
	"os"
	"io"
	"sync"
)

type sshinfo struct {
	IP, Username string
	Passwd       string
	Port         int
	client       *ssh.Client
	Session      *ssh.Session
	Result       string
}

func New_ssh(port int, args ...string) *sshinfo {
	temp := new(sshinfo)
	temp.Port = port
	temp.IP = args[0]
	temp.Username = args[1]
	temp.Passwd = args[2]
	return temp

}
func (cli *sshinfo) connect() error {
	auth := make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(cli.Passwd))

	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	clientConfig := &ssh.ClientConfig{
		User:            cli.Username,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// connet to ssh
	addr := fmt.Sprintf("%s:%d", cli.IP, cli.Port)

	client, err := ssh.Dial("tcp", addr, clientConfig)
	if err != nil {
		return err
	}

	// create session
	session, err := client.NewSession()
	if err != nil {
		defer cli.close_session()
		return err
	}
	cli.Session = session
	return nil
}
func (cli *sshinfo) close_session() {
	cli.Session.Close()
}

func main() {

	ssh := New_ssh(22, []string{"192.168.43.12", "root", "root"}...)
	fmt.Println(ssh)
	err := ssh.connect()
	if err != nil {
		log.Fatal(err)
	}
	//ssh.Session.Stdout=os.Stdout
	//ssh.Session.Stderr=os.Stderr
	//ssh.Session.Run("touch /root/1")
	//ssh.Session.Run("ls /; ls /tmp")
	//ssh.close_session() //todo session一次运行一次run

	terminal_run(ssh.Session)
	ssh.close_session()

}

func terminal_run(session *ssh.Session)  {

	//fd := int(os.Stdin.Fd())
	//oldState, err := terminal.MakeRaw(fd)
	//if err != nil {
	//	panic(err)
	//}
	//defer terminal.Restore(fd, oldState)

	// excute command
	w, err := session.StdinPipe()
	if err != nil {
		panic(err)
	}
	r, err := session.StdoutPipe()
	if err != nil {
		panic(err)
	}
	e, err := session.StderrPipe()
	if err != nil {
		panic(err)
	}

	//
	//termWidth, termHeight, err := terminal.GetSize(fd)
	//if err != nil {
	//	panic(err)
	//}

	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	// Request pseudo terminal
	if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}
	// Start remote shell
	//if err := session.Shell(); err != nil {
	//	log.Fatal("failed to start shell: ", err)
	//}

	in, out := MuxShell(w, r, e)
	if err := session.Shell(); err != nil {
		log.Fatal(err)
	}
	<-out //ignore the shell output
	in <- "ls /"
	in <- "ls /tmp"

	in <- "exit"
	//in <- "exit"

	fmt.Printf("%s\n%s\n", <-out, <-out)

	_, _ = <-out, <-out
	session.Wait()



}

func checkError(err error, info string) {
	if err != nil {
		fmt.Printf("%s. error: %s\n", info, err)
		os.Exit(1)
	}
}

func MuxShell(w io.Writer, r, e io.Reader) (chan<- string, <-chan string) {
	in := make(chan string, 3)
	out := make(chan string, 5)
	var wg sync.WaitGroup
	wg.Add(1) //for the shell itself
	go func() {
		for cmd := range in {
			wg.Add(1)
			w.Write([]byte(cmd + "\n"))
			wg.Wait()
		}
	}()

	go func() {
		var (
			buf [65 * 1024]byte
			t   int
		)
		for {
			n, err := r.Read(buf[t:])
			if err != nil {
				fmt.Println(err.Error())
				close(in)
				close(out)
				return
			}
			t += n
			result := string(buf[:t])
			if strings.Contains(result, "Username:") ||
				strings.Contains(result, "Password:") ||
				strings.Contains(result, "#") {
				out <- string(buf[:t])
				t = 0
				wg.Done()
			}
		}
	}()
	return in, out
}
