package systemctl

import (
	"errors"
	"os/exec"
)

const ConfigDir = "/etc/systemd/system/"

/*
参考文档
https://www.jb51.net/article/136559.htm
systemctl 提供了一组子命令来管理单个的 unit，其命令格式为：
systemctl [command] [unit]
command 主要有：
start：立刻启动后面接的 unit。
stop：立刻关闭后面接的 unit。
restart：立刻关闭后启动后面接的 unit，亦即执行 stop 再 start 的意思。
reload：不关闭 unit 的情况下，重新载入配置文件，让设置生效。
enable：设置下次开机时，后面接的 unit 会被启动。
disable：设置下次开机时，后面接的 unit 不会被启动。
status：目前后面接的这个 unit 的状态，会列出有没有正在执行、开机时是否启动等信息。
is-active：目前有没有正在运行中。
is-enable：开机时有没有默认要启用这个 unit。
kill ：不要被 kill 这个名字吓着了，它其实是向运行 unit 的进程发送信号。
show：列出 unit 的配置。
mask：注销 unit，注销后你就无法启动这个 unit 了。
unmask：取消对 unit 的注销。
*/

func Run(args ...string) (string, error) {
	cmd := exec.Command("systemctl", args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func Start(serviceName string) error {
	out, err := Run("start", serviceName)
	return dealErr(out, err)
}

func Restart(serviceName string) error {
	out, err := Run("restart", serviceName)
	return dealErr(out, err)
}

func Stop(serviceName string) error {
	out, err := Run("stop", serviceName)
	return dealErr(out, err)
}

func Reload(serviceName string) error {
	out, err := Run("stop", serviceName)
	return dealErr(out, err)
}

func Enable(serviceName string) error {
	out, err := Run("enable", serviceName)
	return dealErr(out, err)
}

func Disable(serviceName string) error {
	out, err := Run("disable", serviceName)
	return dealErr(out, err)
}

func Status(serviceName string) (string, error) {
	return Run("status", serviceName)
}

func Show(serviceName string) (string, error) {
	return Run("show", serviceName)
}

func IsActive(serviceName string) (bool, error) {
	out, err := Run("is-active", serviceName)
	if err != nil {
		return false, err
	}
	return out == "active\n", nil
}

func dealErr(out string, err error) error {
	if err != nil {
		if len(out) > 0 {
			return errors.New(out)
		}
		return err
	}
	return nil
}
