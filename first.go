package main

import (
	"os/exec"
	"strings"
	"fmt"
	"unsafe"
)



func TrimAny(cm *string) *string {
	*cm = strings.Trim(*cm, " ")
	*cm = strings.Trim(*cm, "\t")
	*cm = strings.Trim(*cm, "\r\n")
	return cm
}

func getOSVersion() (cm *string) {
	cmd := exec.Command("/bin/bash", "-c", "lsb_release -r")
	buf, _ := cmd.CombinedOutput()
	cmd.Run()
	//cm = strings.Trim(string(buf), " ")
	cm = (*string)(unsafe.Pointer(&buf))
	if strings.Contains(*cm,"not found"){

	}else{
		*cm=strings.Trim(*cm,"Release:")
	}


	return TrimAny(cm)
}
func main()  {
	//fmt.Print(getOSName())
	fmt.Print(*getOSVersion())
}