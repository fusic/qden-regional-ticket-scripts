package aws

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/term"
)

func StartSSMSession(cluster, target string) error {
	cmd := exec.Command("aws", "ssm", "start-session",
		"--target", fmt.Sprintf("ecs:%s_%s", cluster, target),
	)

	fd := int(os.Stdin.Fd())
	if term.IsTerminal(fd) {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		return fmt.Errorf("ターミナルが必要です")
	}

	return cmd.Run()
}

func StartPortForwardingSession(cluster, target, endpoint, remotePort, localPort string) error {
	cmd := exec.Command("aws", "ssm", "start-session",
		"--target", fmt.Sprintf("ecs:%s_%s", cluster, target),
		"--document-name", "AWS-StartPortForwardingSessionToRemoteHost",
		"--parameters", fmt.Sprintf("host=\"%s\",portNumber=\"%s\",localPortNumber=\"%s\"", endpoint, remotePort, localPort),
	)

	fd := int(os.Stdin.Fd())
	if term.IsTerminal(fd) {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		return fmt.Errorf("ターミナルが必要です")
	}

	return cmd.Run()
}