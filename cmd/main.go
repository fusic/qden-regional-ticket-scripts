package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Root コマンド
var rootCmd = &cobra.Command{
	Use:   "mc-ops",
	Short: "AWS環境を操作するためのツール",
	Long:  "mc-ops は AWS ECS や RDS を簡単に操作するための CLI ツールです。",
}

// 実行
func main() {
	Init()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("エラー:", err)
	}
}