package pxelinux

import (
	"fmt"
	"io"
	"log"

	"github.com/k8sp/auto-install/bootstrapper/cmd"
	"github.com/k8sp/auto-install/config"
	"github.com/topicai/candy"
)

func Pxelinux_install(tmpl string, c *config.Cluster){
	const (
		centos = "centos"
		ubuntu = "ubuntu"
	)
	
	linuxdis := config.LinuxDistro()   
	if linuxdis == ubuntu 
	{
		cmd.Run("apt-get","update")
		cmd.Run("apt-get", "-y", "install", "pxelinux", "syslinux-common")
		cmd.Copy("/srv/tftp/", "/usr/lib/PXELINUX/pxelinux.0")
		cmd.Copy("/srv/tftp/", "/usr/lib/syslinux/modules/bios/ldlinux.c32")
	}
	else if linuxdis == centos 
	{
		cmd.Run("yum", "-y", "install", "syslinux")
		cmd.Copy("/var/lib/tftpboot/", "/usr/share/syslinux/pxelinux.0")
	}
	else
	{
		log.Panicf("Unsupported OS: %s", linuxdis)
	}
	
}

//add the Copy function to cmd Package
func Copy(dst string, src string){
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
