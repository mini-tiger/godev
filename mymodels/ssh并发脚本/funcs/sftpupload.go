package funcs

import "tjtools/sshtools"

func UploadFile(sfile, dfile string) {
	sshtools.SftpUpload()
}
