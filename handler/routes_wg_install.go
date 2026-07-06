package handler

import (
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/ngoduykhanh/wireguard-ui/store"
)

func WireGuardStatus(db store.IStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		installed := false
		moduleLoaded := false
		wgVersion := ""

		if _, err := exec.LookPath("wg"); err == nil {
			installed = true
			out, _ := exec.Command("wg", "--version").Output()
			wgVersion = strings.TrimSpace(string(out))
		}

		out, _ := exec.Command("modprobe", "wireguard").Output()
		_ = out
		out, _ = exec.Command("lsmod").Output()
		if strings.Contains(string(out), "wireguard") {
			moduleLoaded = true
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":        true,
			"installed":     installed,
			"module_loaded": moduleLoaded,
			"version":       wgVersion,
			"os":            runtime.GOOS,
		})
	}
}

func InstallWireGuard(db store.IStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		go func() {
			if err := installWG(); err != nil {
				log.Errorf("WireGuard install failed: %v", err)
			} else {
				log.Info("WireGuard installation/reinstall completed successfully")
			}
		}()
		return c.JSON(http.StatusOK, jsonHTTPResponse{true, "WireGuard installation started, please refresh to check status"})
	}
}

func installWG() error {
	if err := detectAndInstall(); err != nil {
		return err
	}
	loadModule()
	enableIPForward()
	return nil
}

func detectAndInstall() error {
	osRelease, err := exec.Command("cat", "/etc/os-release").Output()
	if err != nil {
		return err
	}
	content := string(osRelease)

	var cmd *exec.Cmd
	switch {
	case strings.Contains(content, "Ubuntu"), strings.Contains(content, "Debian"):
		exec.Command("apt-get", "update", "-qq").Run()
		cmd = exec.Command("apt-get", "install", "-y", "-qq", "wireguard")
	case strings.Contains(content, "CentOS"), strings.Contains(content, "Rocky"), strings.Contains(content, "Red Hat"), strings.Contains(content, "AlmaLinux"), strings.Contains(content, "Fedora"):
		cmd = exec.Command("dnf", "install", "-y", "-q", "wireguard-tools")
	case strings.Contains(content, "openSUSE"), strings.Contains(content, "SUSE"):
		cmd = exec.Command("zypper", "install", "-y", "wireguard-tools")
	case strings.Contains(content, "Arch"):
		cmd = exec.Command("pacman", "-S", "--noconfirm", "wireguard-tools")
	case strings.Contains(content, "Alpine"):
		cmd = exec.Command("apk", "add", "wireguard-tools")
	default:
		cmd = exec.Command("dnf", "install", "-y", "-q", "wireguard-tools")
	}

	if cmd != nil {
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Warnf("WireGuard install output: %s", string(out))
			// try alternative: install just wireguard-tools without kmod
			if strings.Contains(string(out), "No match") || strings.Contains(string(out), "no package") {
				alt := exec.Command("dnf", "install", "-y", "-q", "wireguard-tools")
				return alt.Run()
			}
			return err
		}
	}
	return nil
}

func loadModule() {
	exec.Command("modprobe", "wireguard").Run()
}

func enableIPForward() {
	exec.Command("sysctl", "-w", "net.ipv4.ip_forward=1").Run()
	exec.Command("sysctl", "-w", "net.ipv6.conf.all.forwarding=1").Run()
}


