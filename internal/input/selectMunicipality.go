package input

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/manifoldco/promptui"
)

// GetMunicipalityInput は自治体の名前と AWS アカウントID を取得
func GetMunicipalityInput(promptMessage string) (string, string) {
	// SSM から自治体リストを取得
	municipalityMap, err := fetchMunicipalitiesFromSSM()
	if err != nil {
		fmt.Println("自治体リストの取得に失敗しました:", err)
		os.Exit(1)
	}

	// 自治体名リストを作成
	municipalityNames := make([]string, 0, len(municipalityMap))
	for name := range municipalityMap {
		municipalityNames = append(municipalityNames, name)
	}

	// 検索機能
	searcher := func(input string, index int) bool {
		municipality := municipalityNames[index]
		return strings.Contains(strings.ToLower(municipality), strings.ToLower(input))
	}

	// ユーザーに自治体を選択させる
	prompt := promptui.Select{
		Label:             promptMessage,
		Items:             municipalityNames,
		Size:              10,  // 表示する最大件数
		Searcher:          searcher,
		StartInSearchMode: true, // デフォルトで検索モード
	}

	// ユーザーの選択
	_, selectedMunicipality, err := prompt.Run()
	if err != nil {
		fmt.Println("自治体の選択に失敗しました:", err)
		os.Exit(1)
	}

	// 選択された自治体のAWSアカウントIDを取得
	accountID := municipalityMap[selectedMunicipality]

	return selectedMunicipality, accountID
}

// fetchMunicipalitiesFromSSM は自治体名とAWSアカウントIDのマップを取得（ページネーション対応）
func fetchMunicipalitiesFromSSM() (map[string]string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("AWS設定の読み込みに失敗しました: %v", err)
	}

	ssmClient := ssm.NewFromConfig(cfg)
	municipalityMap := make(map[string]string)

	// ページネーション用
	var nextToken *string

	for {
		// SSM から `/municipality/` 配下のデータを取得
		out, err := ssmClient.GetParametersByPath(context.TODO(), &ssm.GetParametersByPathInput{
			Path:           aws.String("/municipality/"),
			Recursive:      aws.Bool(true),
			WithDecryption: aws.Bool(false),
			MaxResults:     aws.Int32(10), // 一度に取得する最大件数
			NextToken:      nextToken,
		})
		if err != nil {
			return nil, fmt.Errorf("SSMからのデータ取得失敗: %v", err)
		}

		// データをマップに追加
		for _, param := range out.Parameters {
			// `/municipality/<自治体名>` の `/municipality/` を除去
			nameParts := strings.Split(*param.Name, "/")
			var municipality string
			if len(nameParts) > 2 {
				municipality = nameParts[2]
			} else {
				municipality = nameParts[1]
			}
			municipalityMap[municipality] = *param.Value
		}

		// 次のページがある場合は続ける
		if out.NextToken == nil {
			break
		}
		nextToken = out.NextToken
	}

	return municipalityMap, nil
}