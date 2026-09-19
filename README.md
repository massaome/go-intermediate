# go-intermediate

『[APIを作りながら進むGo中級者への道](https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx)』の学習用リポジトリです。

Go でブログ API を実装しながら、HTTP サーバーから認証まで中級者向けの実装を進めます。

公式のサンプルコードは [saki-engineering/go-intermediate](https://github.com/saki-engineering/go-intermediate/) です。このリポジトリは自分の手で実装するためのものです。

## 開発環境

[Dev Containers](https://containers.dev/) を使います。Go 1.25 はコンテナ内に入っており、Mac に Go を入れる必要はありません。

1. Docker Desktop を起動する
2. Cursor でこのリポジトリを開く
3. 「Reopen in Container」を選ぶ

コンテナ内のターミナルで `go version` が動けば準備完了です。HTTP サーバーはコンテナの 8080 番をホストへ転送します。

## 作るもの

ブログサービスのバックエンド API です。次の 5 機能を実装します。

- 記事投稿
- 記事一覧の取得
- 記事詳細の取得
- コメント投稿
- いいね

## 学習の流れ

1. 第1章 HTTPサーバー
2. 第2章 構造体と JSON の扱い方
3. 第3章 データベースの扱い方
4. 第4章 ユニットテスト（基礎編）
5. 第5章 サービス層の作成
6. 第6章 アーキテクチャ大改装
7. 第7章 エラー処理
8. 第8章 ユニットテスト（応用編）
9. 第9章 ミドルウェアによるロギング
10. 第10章 並行処理
11. 第11章 context パッケージの導入
12. 第12章 ユーザー認証

## 参考文献

- [APIを作りながら進むGo中級者への道](https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx)（さきさん文庫 / 技術書典13）
- [公式サンプルコード](https://github.com/saki-engineering/go-intermediate/)
