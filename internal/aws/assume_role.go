package aws

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type AWSCredentials struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

// AssumeRole AWSのAssumeRoleを実行してクレデンシャルを取得
func AssumeRole(roleArn, sessionName, externalID string) (*AWSCredentials, error) {
	cmd := exec.Command("aws", "sts", "assume-role",
		"--role-arn", roleArn,
		"--role-session-name", sessionName,
		"--external-id", externalID,
	)

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("AssumeRoleの実行に失敗しました: %v", err)
	}

	var creds map[string]interface{}
	if err := json.Unmarshal(output, &creds); err != nil {
		return nil, fmt.Errorf("クレデンシャル情報の解析に失敗しました: %v", err)
	}

	awsCreds := creds["Credentials"].(map[string]interface{})
	return &AWSCredentials{
		AccessKeyID:     awsCreds["AccessKeyId"].(string),
		SecretAccessKey: awsCreds["SecretAccessKey"].(string),
		SessionToken:    awsCreds["SessionToken"].(string),
	}, nil
}