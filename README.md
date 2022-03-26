# go_samples
Go学習のために以下のサンプルコードを記載
- 構造体、配列の参照・値渡しの動作確認
- interfaceの実装例
- goroutineのテスト
また、goのエコシステム・ツールも一通り試せるようにしている
- テストコード(一部のみ実装)
- ベンチマーク・プロファイリング
- ドキュメンテーション


## Setup
```
go get
```


## 実行
```
go run main.go
```


## テスト実行
~\_test.go ファイルの Test~関数がテスト対象となる。
コンパイル時に除外されるので、プロダクトコードと同じパッケージにテストファイル入れるのが一般的？
```
go test -v ./...
```


## カバレッジ付きでテスト実行
```
go test -v -cover ./...

# カバレッジをwebで見る
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## ベンチマーク計測
Benchmark~とつく関数のパフォーマンスを計測できる。
```
cd <計測したいパッケージ>
go test -bench=. -benchmem
```

## プロファイリング
ベンチマーク計測して各処理のメモリ・CPU消費などを計測できる。
```
cd <計測したいパッケージ>

# 性能計測
go test -memprofile=mem.out -bench=.
go test -blockprofile=block.out -bench=.
go test -cpuprofile=cpu.out -bench=.

# pprofツールで可視化
go tool pprof -text -nodecount=10 ./struct_sample.test block.out
```


## コマンドライン上でドキュメンテーションを見る
パッケージや公開関数などをコマンドライン上で確認できる
```
go doc <パッケージ名>

go doc <パッケージ名>.<関数/構造体など>

# 例
go doc ./struct_sample
go doc ./struct_sample.TestStruct
```


## ドキュメンテーション一覧
Web上で標準パッケージ、自作パッケージのドキュメントが見れる。
```
godoc -http ":3000"
```


## Goプロジェクトセットアップ
### 初期化
```
go mod init <パッケージ名>
```

### パッケージインストール
```
go get <外部パッケージ>
```
