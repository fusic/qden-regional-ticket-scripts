package aws

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func GetRdsEndpoint() (string, error) {
	cmd := exec.Command("aws", "rds", "describe-db-clusters", "--output", "json")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("RDSクラスタ情報の取得に失敗しました: %v", err)
	}

	var rdsData map[string]interface{}
	if err := json.Unmarshal(output, &rdsData); err != nil {
		return "", fmt.Errorf("レスポンスの解析に失敗しました: %v", err)
	}

	dbClusters, ok := rdsData["DBClusters"].([]interface{})
	if !ok || len(dbClusters) == 0 {
		return "", fmt.Errorf("RDSクラスタが見つかりませんでした")
	}

	firstCluster := dbClusters[0].(map[string]interface{})
	endpoint, exists := firstCluster["Endpoint"].(string)
	if !exists {
		return "", fmt.Errorf("エンドポイントが見つかりませんでした")
	}

	return endpoint, nil
}