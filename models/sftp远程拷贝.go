package main

import (
	"github.com/pkg/sftp"
	"time"
	"log"
	"io/ioutil"
	"golang.org/x/crypto/ssh"
	"fmt"
	"os"
	"path"
)



func main() {
	var (
		err        error
		sftpClient *sftp.Client
	)
	//start := time.Now()
	sftpClient, err = connect("root","root","192.168.43.11",22)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	_, errStat := sftpClient.Stat("/root/")
	if errStat != nil {
		log.Fatal("/root/" + " remote path not exists!")
	}

	//backupDirs, err := ioutil.ReadDir("c:\\1.log")
	//if err != nil {
	//	log.Fatal("c:\\1.log  local path not exists!")
	//}
	uploadFile(sftpClient,"C:\\1.log","/root/1.log")
}
func connect(user, password, host string, port int) (*sftp.Client, error) {
	var (
	auth         []ssh.AuthMethod
	addr         string
	clientConfig *ssh.ClientConfig
	sshClient    *ssh.Client
	sftpClient   *sftp.Client
	err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
	User:            user,
	Auth:            auth,
	Timeout:         30 * time.Second,
	HostKeyCallback: ssh.InsecureIgnoreHostKey(), //ssh.FixedHostKey(hostKey),
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
	return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
	return nil, err
	}
	return sftpClient, nil
	}


func uploadFile(sftpClient *sftp.Client, localFilePath string, remotePath string) {
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		fmt.Println("os.Open error : ", localFilePath)
		log.Fatal(err)

	}
	defer srcFile.Close()

	//var remoteFileName = path.Base(localFilePath)

	dstFile, err := sftpClient.Create(path.Join("/root/", "1.log"))
	if err != nil {
		fmt.Println("sftpClient.Create error : ", path.Join("/root/", "1.log"))
		log.Fatal(err)

	}
	defer dstFile.Close()

	ff, err := ioutil.ReadAll(srcFile)
	if err != nil {
		fmt.Println("ReadAll error : ", localFilePath)
		log.Fatal(err)

	}
	dstFile.Write(ff)
	fmt.Println(localFilePath + "  copy file to remote server finished!")
}

