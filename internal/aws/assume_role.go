package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type AWSCredentials struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

// AssumeRole AWSのAssumeRoleを実行してクレデンシャルを取得
func AssumeRole(roleArn, sessionName, externalID string) (*AWSCredentials, error) {
	// AWS SDK のデフォルト設定をロード
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("AWS設定のロードに失敗しました: %w", err)
	}

	// STS クライアントを作成
	stsClient := sts.NewFromConfig(cfg)

	// AssumeRole を実行
	resp, err := stsClient.AssumeRole(context.TODO(), &sts.AssumeRoleInput{
		RoleArn:         &roleArn,
		RoleSessionName: &sessionName,
		ExternalId:      &externalID,
	})
	if err != nil {
		return nil, fmt.Errorf("AssumeRoleの実行に失敗しました: %w", err)
	}

	// クレデンシャルを取得
	creds := resp.Credentials
	if creds == nil {
		return nil, fmt.Errorf("AssumeRoleに成功しましたが、クレデンシャルが返されませんでした")
	}

	return &AWSCredentials{
		AccessKeyID:     *creds.AccessKeyId,
		SecretAccessKey: *creds.SecretAccessKey,
		SessionToken:    *creds.SessionToken,
	}, nil
}