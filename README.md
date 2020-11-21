# YY Music Z

## このアプリについて

YY Music Zは、「自分の好きな曲を投げ合って新しい曲を知るきっかけを作ろう」というテーマの元作成された
アプリケーションです。🥳

曲を投げると同じページ内にいる人全員のプレイリストの一番後ろに追加されます。
また追加された曲が気に入らない場合は自分のプレイリストから削除されます。

あくまでも自分のプレイリストからのみ削除されるため、他人に気を使うことなく削除可能です。

お互いに曲を投稿しあって、新しい好きな曲を見つけましょう！！😎

## How to Use

現在はこのアプリケーションを実行しているサイトを公開しているわけではないため、
各自でこのリポジトリをcloneしてください。

clone後は`app`ディレクトリ下に移動して

```go
go run main.go
```

を実行して

```url
localhost:8080
```

等にアクセスしてください。

ジャンルの選択画面が表示されます。お好みのジャンルのページに遷移してください。

遷移後のページでは、初めから数曲プレイリストの中に曲が入っています。

聞きたくない曲であれば、リスト中の曲名が表示されている部分の右側にある赤い削除ボタンをクリックしてください。

逆に曲をポストしたいときは、そのURLを

```url
https://www.youtube.com/watch?v=xxxxxxx
```

の形式でフォームに入力し、SENDボタンをクリックしてください。
