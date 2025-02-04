package aws

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func GetTaskTarget(cluster, taskArn string) (string, error) {
	cmd := exec.Command("aws", "ecs", "describe-tasks", "--cluster", cluster, "--task", taskArn)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("タスク詳細の取得に失敗しました: %v", err)
	}

	var descData map[string]interface{}
	if err := json.Unmarshal(output, &descData); err != nil {
		return "", fmt.Errorf("タスク詳細データの解析に失敗しました: %v", err)
	}

	containers := descData["tasks"].([]interface{})[0].(map[string]interface{})["containers"].([]interface{})
	for _, container := range containers {
		c := container.(map[string]interface{})
		if strings.Contains(c["name"].(string), "php") {
			return fmt.Sprintf("%s_%s", strings.Split(c["runtimeId"].(string), "-")[0], c["runtimeId"].(string)), nil
		}
	}

	return "", fmt.Errorf("ターゲットが見つかりませんでした")
}