package aws

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

// GetTaskTarget は、指定された ECS タスクのコンテナ情報を取得し、
// `php` を含むコンテナの `runtimeId` を返す
func GetTaskTarget(cluster, taskArn string) (string, error) {
	// AWS SDK のデフォルト設定をロード
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", fmt.Errorf("AWS設定のロードに失敗しました: %w", err)
	}

	// ECS クライアントを作成
	ecsClient := ecs.NewFromConfig(cfg)

	// ECS タスクの詳細を取得
	resp, err := ecsClient.DescribeTasks(context.TODO(), &ecs.DescribeTasksInput{
		Cluster: &cluster,
		Tasks:   []string{taskArn},
	})
	if err != nil {
		return "", fmt.Errorf("ECSタスクの詳細取得に失敗しました: %w", err)
	}

	// タスクが存在しない場合
	if len(resp.Tasks) == 0 {
		return "", fmt.Errorf("指定されたタスクが見つかりませんでした")
	}

	// コンテナ情報を検索
	for _, container := range resp.Tasks[0].Containers {
		if strings.Contains(*container.Name, "php") {
			if container.RuntimeId == nil {
				return "", fmt.Errorf("対象コンテナの runtimeId が見つかりませんでした")
			}
			runtimeID := *container.RuntimeId
			return fmt.Sprintf("%s_%s", strings.Split(runtimeID, "-")[0], runtimeID), nil
		}
	}

	return "", fmt.Errorf("ターゲットのコンテナが見つかりませんでした")
}