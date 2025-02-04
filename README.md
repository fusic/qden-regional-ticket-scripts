# qden-regional-ticket-scripts

このリポジトリは、AWS環境での運用作業を効率化するためのCLIツールを提供します。主にECSやRDSの操作を簡略化し、エンジニアの作業負担を軽減することを目的としています。

---

## 📌 主なツール

### 1. **ecs_exec**
- **概要**: 
  ECSタスクに対してSSMセッションを開始するツールです。
- **主な機能**:
  - AWSプロファイルを指定してECSタスクに接続。
  - 対象タスクを自動的に取得し、セッションを確立。
- **利用例**:
  ```bash
  ./ecs_exec
  ```

### 2. **ssm_tunnel**
- **概要**: 
ECS経由でRDSへのポートフォワードを設定するツールです。
- **主な機能**:
- AWSプロファイルを指定してRDSクラスタのエンドポイントを取得。
- ポートフォワードセッションを確立し、ローカルからRDSに接続可能。
- **利用例**:

---

## ✅ 必要な環境と前提条件

1. **AWS CLIのインストール**
 - AWS CLI バージョン2を推奨します。
 - [AWS CLIのインストールガイド](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)

2. **Go言語**
 - バージョン 1.19以上を推奨します。
 - [Goのインストール](https://go.dev/doc/install)

3. **必要なIAM権限**
 - ECSタスクの操作 (`ecs:ListTasks`, `ecs:DescribeTasks`)
 - RDSエンドポイントの取得 (`rds:DescribeDBClusters`)
 - SSMセッションの実行 (`ssm:StartSession`)

---

## 🚀 インストール方法

### 1. リポジトリのクローン

### 2. ビルド
各ツールを以下のコマンドでビルドします。

```
# ecs_execのビルド
go build -o bin/ecs_exec ./cmd/ecs_exec

# ssm_tunnelのビルド
go build -o bin/ssm_tunnel ./cmd/ssm_tunnel
```

---

## 🛠 使い方

### **ecs_exec**
ECSタスクへのSSMセッションを開始します。

- 実行後、対話形式で以下を指定します:
  1. 使用するAWSプロファイル
  2. 接続先のコンテナタイプ（`api` または `backend`）
  3. 接続する自治体
  4. ユーザー名
  5. 環境（`staging` または `production`）

### **ssm_tunnel**
RDSへのポートフォワードを設定します。

- 実行後、対話形式で以下を指定します:
  1. 使用するAWSプロファイル
  2. 接続先のコンテナタイプ（`api` または `backend`）
  3. 接続する自治体
  4. ユーザー名
  5. 環境（`staging` または `production`）
  6. ポートフォワードの設定が完了したら、`localhost:13306` からRDSに接続できます。

---

## 📂 ディレクトリ構成

```
.
├── cmd/
│   ├── ecs_exec/
│   │   └── main.go
│   └── ssm_tunnel/
│       └── main.go
├── internal/
│   ├── aws/
│   │   ├── assume_role.go
│   │   ├── ecs_tasks.go
│   │   ├── ecs_target.go
│   │   ├── rds_clusters.go
│   │   └── ssm_session.go
│   ├── config/
│   │   ├── config.go
│   │   └── config.json
│   └── input/
│       ├── get_user_input.go
│       ├── select_profile.go
│       └── select_municipality.go
├── bin/
│   ├── ecs_exec
│   └── ssm_tunnel
├── go.mod
├── go.sum
└── README.md
```

---

## 🔧 今後の展望

- 他の既存スクリプトの移行
