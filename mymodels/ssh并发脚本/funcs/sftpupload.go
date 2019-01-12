package funcs

import (
	"sync"
	"tjtools/logDiy"
	"tjtools/sshtools"
)

var wg *sync.WaitGroup = new(sync.WaitGroup)
var FailSftpHosts = make([]string, 0)

func UploadFile(sfile, dfile string) {

	for host, passwd := range HostPass.M {
		wg.Add(1)
		go BeginUpload(host, passwd.(string), sfile, dfile)
	}

	wg.Wait()
}

func BeginUpload(host, passwd, sFile, dFile string) {
	defer func() {
		wg.Done()
	}()
	if err := sshtools.SftpUpload(host, passwd, sFile, dFile); err == nil {
		logDiy.Logger().Printf("sftp success host:%s begin run script %s", host, dFile)

	} else {
		logDiy.Logger().Error("sftp Fail err:%s", err)
		FailSftpHosts = append(FailSftpHosts, host)
	}
}
