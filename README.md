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

インストールされていない場合、インストールします。私が `sudo apt-get install go` でインストールした時はgo 1.3.3がインストールされてしまったので、ここでは公式サイトから最新版を取得する手順を示しています。

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
```

Gitがインストールされているか確認します。（ `go get` 時に使用します）

```bash
$ git version
```

インストールされていない場合、別途インストールします。

```bash
$ sudo apt-get install git
```

GCCがインストールされているか確認します。（go-sqlite3で使用します）

```bash
$ gcc -v
```

インストールされていない場合、別途インストールします。

```bash
$ sudo apt-get install gcc
```

本リポジトリーのapp_v1を `go get` します。（`-u` は最新版の取得、`-d` はダウンロードのみ）

```bash
$ go get -u -d github.com/qt-luigi/gcp-compute-engine-go/app_v1
```

ソースコードをコンパイルして、実行バイナリー「app」を作成します。

```bash
$ cd $GOPATH/src/github.com/qt-luigi/gcp-compute-engine-go/app_v1
$ go build -o app *.go
```

以降、書籍に従ってセットアップしてください。

## ローカルで実行

Go、SQLite3、Git、GCCがインストールされていない場合、インストールします。

本リポジトリーのapp_v1を `go get` します。（`-u` は最新版の取得、`-d` はダウンロードのみ）

```bash
$ go get -u -d https://github.com/qt-luigi/gcp-compute-engine-go/app_v1
```

`app.go` ファイルの次の箇所をコメントに従って書き換えてください。

```go
// ローカル実行の際は""に変更してください。
const installPath = "/opt/dengonban/v1"
```

```go
func main() {
    // ローカル実行の際は":8080"に変更して「http://localhost:8080」でアクセスしてください。
    if err := http.ListenAndServe(":80", nil); err != nil {
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

本リポジトリーのapp_v2を `go get` します。（`-u` は最新版の取得、`-d` はダウンロードのみ）

```bash
$ go get -u -d github.com/qt-luigi/gcp-compute-engine-go/app_v2
```

セットアップについてはapp_v1を、GCP作業については書籍を、参照してください。

# app_v3

2017年7月17日現在、GCE上でテストするためにpushしているだけの状態のため使用しないでください。

# 自分用メモ

systemdのログ表示:

- `$ sudo journalctl -u dengonban.service`
