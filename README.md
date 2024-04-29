# 苦労した点
エラーを`log.Panic`で出力してしてたので意図しない動きをしてつまづいた｡  
当時`log.Fatal`は後の処理をせずに止めると知ってたが､`log.Panic`も処理を止めてしまうとは知らなかった｡  
エラーログを出しつつ次の処理に進んでほしかったのだが`log.Panic`が止めていたとわかるまで時間がかかった(1時間ほど?)


# 参考にしたサイト
* chatgpt
* 公式サイト

## go
* [Goのnet/httpではパスパラメータは取得できない！なので実装してみた](https://qiita.com/shun_labo/items/89fbf8f4d972daa6d5a2)
* [Goのhttpパッケージでパスパラメータを取る](https://zenn.dev/kaikusakari/articles/07cf9af5255586)

## net/http
* [URLに応じた処理（マルチプレクサとハンドラ）](https://www.twihike.dev/docs/golang-web/handlers)