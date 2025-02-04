package input

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func GetUserInput(promptMessage, defaultValue string, options []string) string {
	if len(options) > 0 {
		// 選択肢がある場合は `promptui.Select` を使用
		prompt := promptui.Select{
			Label: promptMessage,
			Items: options,
		}
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Println("選択に失敗しました:", err)
			os.Exit(1)
		}
		return result
	} else {
		// デフォルト値がある場合の処理
		prompt := promptui.Prompt{
			Label:   fmt.Sprintf("%s (デフォルト: %s)", promptMessage, defaultValue),
			Default: defaultValue, // デフォルト値を設定
		}
		result, err := prompt.Run()
		if err != nil {
			fmt.Println("入力に失敗しました:", err)
			os.Exit(1)
		}
		return result
	}
}