package constants

import (
	"path/filepath"

	"github.com/crc-org/crc/pkg/crc/version"
)

const (
	OcExecutableName           = "oc"
	PodmanRemoteExecutableName = "podman"
	TrayExecutableName         = "Red Hat OpenShift Local.app"
	DaemonAgentLabel           = "com.redhat.crc.daemon"
	QemuGuestAgentPort         = 1234
)

var (
	TapSocketPath        = filepath.Join(CrcBaseDir, "tap.sock")
	DaemonHTTPSocketPath = filepath.Join(CrcBaseDir, "crc-http.sock")
)

func TrayExecutablePath() string {
	if version.IsInstaller() {
		return filepath.Clean(filepath.Join(version.InstallPath(), "..", "MacOS", "crc-tray"))
	}
	// Should not be reached, tray is only supported on installer builds
	return filepath.Clean(filepath.Join(BinDir(), "Red Hat OpenShift Local"))
}
