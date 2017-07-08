# 現在テスト中です！（2017年7月8日 現在）

Compute Engine上から本リポジトリーを取得して動作確認を行うためにGitHubにpushした状態です。動作確認が完了しましたら本セクションを削除します。

# このリポジトリーについて

書籍「[プログラマのためのGoogle Cloud Platform入門](http://www.shoeisha.co.jp/book/detail/9784798137148)」の「第2章 Webアプリケーション実行基盤を構築しよう」で使用および提供されているPython製のサンプルアプリケーションをGoに書き換えてみました。サンプルは掲示板アプリケーションでGoogle Compute Engineにデプロイして使用することが前提となっています。

## オリジナルのリポジトリー

オリジナルは著者である阿佐志保さん([@_Dr_ASA](https://twitter.com/_dr_asa))の次のリポジトリーです。

- asashiho/gcp-compute-engine: https://github.com/asashiho/gcp-compute-engine

本リポジトリーの公開に際しては、阿佐さんに問い合わせて許可を頂いています。

## Goへの書き換えの際に注意した点

- appファイルに関しては、オリジナルと比較しやすいように処理の記載順などなるべく合わせています。
- appファイル以外の関数名などは、パッケージ化することを考慮して大文字で始めています。
- 影響がない範囲でGoのコーディング規約に従って書き換えています。（例：関数名には"_"を使用しない、など）

# app_v1

app_v1はCompute Engineとローカルの両方で実行できます。

## Compute Engineで実行（「Debian GNU/Linux 8 (jessie)」の場合）

Goがインストールされているか確認します。

```bash
$ go version
```

インストールされていない場合、インストールします。

```bash
$ curl -O https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz
$ mkdir -p $HOME/go/src
$ export PATH=$PATH:/usr/local/go/bin
$ export GOPATH=$HOME/go
```

ここでのexport文によるパスの設定は一時的なものなのでコンソールを抜けると無効になります。ログインの都度にexport文を実行するか `.bashrc` ファイルなどに環境変数として設定するなどしてください。

SQLite3がインストールされているか確認します。

```bash
$ sqlite3 -version
```

インストールされていない場合、別途インストールします。

```bash
$ sudo apt-get install sqlite3
$ sqlite3 -version
```

本リポジトリーとGo用のSQLite3ドライバーを `go get` します。

```bash
$ go get https://github.com/qt-luigi/gcp-compute-engine-go
$ go get https://github.com/mattn/go-sqlite3
```

ソースコードをコンパイルして、実行バイナリー「app」を作成します。

```bash
$ cd $GOPATH/src/github.com/qt-luigi/gcp-compute-engine-go/app_v1
$ go build -o app *.go
```

以降、書籍に従ってセットアップしてください。

## ローカルで実行

GoとSQLite3がインストールされていない場合、インストールします。

本リポジトリーとGo用のSQLite3ドライバーを `go get` します。

```bash
$ go get https://github.com/qt-luigi/gcp-compute-engine-go
$ go get https://github.com/mattn/go-sqlite3
```

macOS Sierraで動作させる場合、 `app.go` ファイル内のポート番号を `:80` から `:8080` に書き換える必要がありました。

```go
func main() {
        if err := http.ListenAndServe(":8080", nil); err != nil {
            log.Fatal(err)
        }
}
```

実行バイナリーを作成して実行する場合、次のコマンドを実行します。

```bash
$ go build -o app *.go
$ ./app
```

`go run` で実行する場合、次のコマンドを実行します。

```bash
$ go run *.go
```

# app_v2

app_v2はapp_v1のデータベースがCloud SQL(MySQL)に置き換わったものです。

Go用のMySQLドライバーを `go get` してください。

```bash
$ go get https://github.com/go-sql-driver/mysql
```

# app_v3

現在、移植中です。


