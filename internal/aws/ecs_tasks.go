package aws

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func GetFirstTaskArn(cluster string) (string, error) {
	cmd := exec.Command("aws", "ecs", "list-tasks", "--cluster", cluster, "--service-name", cluster)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("ECSタスクの取得に失敗しました: %v", err)
	}

	var taskData map[string]interface{}
	if err := json.Unmarshal(output, &taskData); err != nil {
		return "", fmt.Errorf("タスクデータの解析に失敗しました: %v", err)
	}

	taskArns := taskData["taskArns"].([]interface{})
	if len(taskArns) == 0 {
		return "", fmt.Errorf("タスクが見つかりませんでした")
	}

	return taskArns[0].(string), nil
}