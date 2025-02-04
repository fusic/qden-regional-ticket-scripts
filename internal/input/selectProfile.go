package input

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

// プロファイル選択
func SelectProfile() string {
	cmd := exec.Command("aws", "configure", "list-profiles")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("プロファイルの取得に失敗しました:", err)
		os.Exit(1)
	}

	// 出力を改行で分割し、リストとして返す
	profiles := strings.Split(strings.TrimSpace(string(output)), "\n")
	
	prompt := promptui.Select{
		Label: "使用するAWSプロファイルを選択してください",
		Items: profiles,
	}

	index, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("プロファイルの選択に失敗しました:(インデックス: %d) %s", index, err)
		os.Exit(1)
	}

	return result
}