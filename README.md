# HOW TO USE
```Terminal
git clone git@github.com:ogadra/goRestSample.git
cd goRestSample
touch ./.env
touch ./src/.env
# 環境変数については後述
docker-compose up
# Portが既に使われている場合はdocker-compose.ymlを適宜編集する
```

# ファイル説明
- `DckerFile`, `docker-compose.yml`
  - Dockerを構築するためのファイル。
- `./initdb.d/create_table.sql`, `./my.cnf/my.cnf`
  - MySQLをDockerで構築するためのファイル。
- `./src/main.go`, `./src/databases.go`
  - メインとなるGoプログラム
- `./src/go.sum`, `./src/go.mod`
  - Go Modulesを使用するために必要なファイル

~~Docker初心者なので、ディレクトリ構成の仕方がよく分かりませんでした。~~

## 環境変数について
`.env`と`./src/.env`の２つのファイルを作成し、前者には`ROOT_PASS`, `PASS`を、後者には`PASS`を設定する。

## プログラム説明

### main.go
初期処理、ルーティング、受け取るJSON構造体の定義や時間の処理など

### databases.go
DB操作、レスポンス、送るJSON構造体の定義など
