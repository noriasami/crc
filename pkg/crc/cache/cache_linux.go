package cache

import (
	"github.com/crc-org/crc/pkg/crc/machine/libvirt"
)

func NewMachineDriverLibvirtCache() *Cache {
	return New(libvirt.MachineDriverCommand, libvirt.MachineDriverDownloadURL, libvirt.MachineDriverVersion, getCurrentLibvirtDriverVersion)
}

func getCurrentLibvirtDriverVersion(executablePath string) (string, error) {
	return getVersionGeneric(executablePath, "version")
}
