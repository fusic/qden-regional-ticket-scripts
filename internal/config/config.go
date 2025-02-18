package config

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
)

//go:embed config.json
var configFile []byte

type Config struct {
    Username string `json:"username"`
}

func LoadConfig() Config {
    var config Config
    err := json.Unmarshal(configFile, &config)
    if err != nil {
        fmt.Println("設定の読み込みに失敗しました:", err)
    }
    return config
}

// SaveConfig 設定ファイルにユーザー名を書き込む
func SaveConfig(config Config) {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("設定ファイルの保存に失敗しました:", err)
		return
	}

	err = os.WriteFile("config.json", data, 0644)
	if err != nil {
		fmt.Println("設定ファイルの書き込みに失敗しました:", err)
	}
}