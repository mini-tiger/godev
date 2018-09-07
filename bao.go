package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"time"
	"net"
	"log"
	"os"
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

	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(fd, oldState)

	// excute command
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	termWidth, termHeight, err := terminal.GetSize(fd)
	if err != nil {
		panic(err)
	}

	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Request pseudo terminal
	if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
		log.Fatal(err)
	}

	session.Run("top")
}
