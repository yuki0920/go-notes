
# ドメイン駆動設計について

## 参考

### DDD
- https://github.com/nrslib/itddd
- https://github.com/bxcodec/go-clean-arch

### クリーンアーキテクチャー
- https://developers.freee.co.jp/entry/clean-architecture
- https://nrslib.com/clean-architecture/

## 用語

### 境界づけられたコンテキスト

コンテキストごとにモデリングする考え方
ドメインは同一だが、コンテキストが異なる場合には、異なるモデルとして表現する
例えば、システム上、ログインするユーザーと、記事を購読するユーザーを同一ドメインの別パッケージ、別モデルで表現する
