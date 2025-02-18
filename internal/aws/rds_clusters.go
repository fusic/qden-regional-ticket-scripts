package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

// GetRdsEndpoint は、最初の RDS クラスタのエンドポイントを取得する
func GetRdsEndpoint() (string, error) {
	// AWS SDK のデフォルト設定をロード
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", fmt.Errorf("AWS設定のロードに失敗しました: %w", err)
	}

	// RDS クライアントを作成
	rdsClient := rds.NewFromConfig(cfg)

	// RDS クラスタ情報を取得
	resp, err := rdsClient.DescribeDBClusters(context.TODO(), &rds.DescribeDBClustersInput{})
	if err != nil {
		return "", fmt.Errorf("RDSクラスタ情報の取得に失敗しました: %w", err)
	}

	// クラスタが存在しない場合
	if len(resp.DBClusters) == 0 {
		return "", fmt.Errorf("RDSクラスタが見つかりませんでした")
	}

	// 最初のクラスタのエンドポイントを取得
	firstCluster := resp.DBClusters[0]
	if firstCluster.Endpoint == nil {
		return "", fmt.Errorf("エンドポイントが見つかりませんでした")
	}

	return *firstCluster.Endpoint, nil
}