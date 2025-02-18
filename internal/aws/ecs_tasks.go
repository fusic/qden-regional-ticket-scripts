package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

// GetFirstTaskArn は指定したECSクラスター内の最初のタスクARNを取得
func GetFirstTaskArn(cluster string) (string, error) {
	// AWS SDK のデフォルト設定をロード
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", fmt.Errorf("AWS設定のロードに失敗しました: %w", err)
	}

	// ECS クライアントを作成
	ecsClient := ecs.NewFromConfig(cfg)

	// ECS タスクのリストを取得
	resp, err := ecsClient.ListTasks(context.TODO(), &ecs.ListTasksInput{
		Cluster:     &cluster,
		ServiceName: &cluster, // サービス名を指定
		MaxResults:  aws.Int32(1),        // 最初のタスク1件のみ取得
	})
	if err != nil {
		return "", fmt.Errorf("ECSタスクの取得に失敗しました: %w", err)
	}

	// タスクが見つからない場合
	if len(resp.TaskArns) == 0 {
		return "", fmt.Errorf("タスクが見つかりませんでした")
	}

	// 最初のタスクARNを返す
	return resp.TaskArns[0], nil
}