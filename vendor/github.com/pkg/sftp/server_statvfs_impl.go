// +build darwin linux

// fill in statvfs structure with OS specific values
// Statfs_t is different per-kernel, and only exists on some unixes (not Solaris for instance)

package sftp

import (
	"syscall"
)

func (p sshFxpExtendedPacketStatVFS) respond(svr *Server) responsePacket {
	stat := &syscall.Statfs_t{}
	if err := syscall.Statfs(p.Path, stat); err != nil {
		return statusFromError(p, err)
	}

	retPkt, err := statvfsFromStatfst(stat)
	if err != nil {
		return statusFromError(p, err)
	}
	retPkt.ID = p.ID

	return retPkt
}
