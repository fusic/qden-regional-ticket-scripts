package main

func Init() {
	rootCmd.AddCommand(ecsExecCmd)
	rootCmd.AddCommand(rdsConnectCmd)
}