package main

import (
	"fmt"
	"os"
	"qden-regional-ticket-scripts/internal/aws"
	"qden-regional-ticket-scripts/internal/config"
	"qden-regional-ticket-scripts/internal/input"
)

func main() {
	// プロファイル選択
	selectedProfile := input.SelectProfile()

	// 選択したプロファイルを環境変数に設定
	os.Setenv("AWS_PROFILE", selectedProfile)

	// 接続先コンテナ選択（api/backend）
	container := input.GetUserInput("接続先コンテナを選択", "api", []string{"api", "backend"})

	// 自治体選択
	municipality := input.GetMunicipalityInput("自治体を選択または検索")
	accountID := input.MunicipalityMap[municipality]

	// ユーザー名を取得
	conf := config.LoadConfig()
	user := input.GetUserInput("ユーザー名を入力", conf.Username, nil)

	// 環境選択（production/staging）
	env := input.GetUserInput("接続環境を選択", "staging", []string{"staging", "production"})
	if env == "staging" {
		accountID = "283372850953"
		if municipality != "corda" { // corda環境の場合は'-staging'は不要
			municipality += "-staging"
		}	}

	// AWS Assume Role処理
	roleArn := fmt.Sprintf("arn:aws:iam::%s:role/%s-%s-role", accountID, user, municipality)
	if env == "staging" {
		roleArn = fmt.Sprintf("arn:aws:iam::%s:role/%s-staging-role", accountID, user)
	}

	fmt.Println("AWS STS AssumeRole 実行中...:", roleArn)
	creds, err := aws.AssumeRole(roleArn, "Cli-Session", user)
	if err != nil {
		fmt.Println("AssumeRole エラー:", err)
		os.Exit(1)
	}

	// クレデンシャルを環境変数に設定
	os.Setenv("AWS_ACCESS_KEY_ID", creds.AccessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", creds.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", creds.SessionToken)

	// ECSタスク取得
	cluster := fmt.Sprintf("qden-vc-%s-%s", municipality, container)
	fmt.Println("ECSタスク検索中...", cluster)
	taskArn, err := aws.GetFirstTaskArn(cluster)
	if err != nil {
		fmt.Println("ECSタスク取得エラー:", err)
		os.Exit(1)
	}

	// ターゲット取得
	target, err := aws.GetTaskTarget(cluster, taskArn)
	if err != nil {
		fmt.Println("ターゲット取得エラー:", err)
		os.Exit(1)
	}

	// RDSエンドポイント取得
	endpoint, err := aws.GetRdsEndpoint()
	if err != nil {
		fmt.Println("RDSエンドポイント取得エラー:", err)
		os.Exit(1)
	}
	fmt.Println("取得したエンドポイント:", endpoint)

	// SSMセッション開始（ポートフォワード）
	err = aws.StartPortForwardingSession(cluster, target, endpoint, "3306", "13306")
	if err != nil {
		fmt.Println("ポートフォワードセッション開始エラー:", err)
		os.Exit(1)
	}

	fmt.Println("ポートフォワードが正常に終了しました。")
}