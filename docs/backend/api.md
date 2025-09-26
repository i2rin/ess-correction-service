# ess-correction-service

一応swaggerの方に詳しくまとめたので簡単な流れを掴むために見ていただけると幸いです。
## 画面毎の説明

---

# ログイン画面

ログイン画面では、ユーザーが **ユーザーID** と **パスワード** を入力します。  

## リクエスト形式
```json
{
  "userid": "string",
  "password": "string"
}
```

## レスポンス形式
```json
{
  "role": "string"
}
```

## 画面遷移
- `role` が **admin** の場合 → 管理者画面へ遷移  
- `role` が **user** の場合 → タイムライン画面へ遷移  


# 管理者画面・ユーザ編集画面

この画面は **admin ユーザのみ** が閲覧でき、ユーザの管理を行います。  

## ユーザ作成
POST リクエストで以下の形式を送信します。  
```json
{
  "nickname": "string",
  "password": "string",
  "name": "string",
  "mailaddress": "string",
  "role": "string"
}
```

## ユーザ取得
クエリパラメータに `page`, `size` を指定して GET リクエストを送信します。  

レスポンス形式：  
```json
[
  {
    "nickname": "string",
    "password": "string",
    "name": "string",
    "mailaddress": "string",
    "role": "string",
    "userid": "string"
  }
]
```

## ユーザ更新
パスパラメータに（ユーザ取得時に返された）`userid` を指定し、PUT リクエストを送信します。  

リクエスト形式：  
```json
{
  "nickname": "string",
  "password": "string",
  "name": "string",
  "mailaddress": "string",
  "role": "string"
}
```

## ユーザ削除
パスパラメータに（ユーザ取得時に返された）`userid` を指定し、DELETE リクエストを送信します。  

# タイムライン画面

この画面では、直近のユーザの添削履歴を確認できます。  

## データ取得
クエリパラメータに `page`, `size` を指定して GET リクエストを送信します。  

## レスポンス形式
```json
[
  {
    "nickname": "string",       // 直近に投稿したユーザ名
    "text": "string",       // そのユーザの投稿本文
    "date": "2025-09-13",
    "submissionid": "string"
  }
]
```

# 他人の英文にコメントする画面

この画面はタイムライン画面から遷移し、他ユーザの添削に対してコメントを行います。  

## コメント取得
- パスパラメータ：`submissionid`（タイムラインから遷移時に渡される値）  
- クエリパラメータ：`page`, `size`  

### レスポンス形式
```json
{
  "nickname": "string",        // 選択したユーザの名前
  "text": "string",        // 選択したユーザが記述した英文
  "corrections": [
    {
      "original": "string",
      "corrected": "string",
      "reason": "string"
    }
  ],
  "comments": [
    {
      "name": "string",    // コメントしたユーザの名前
      "comment": "string"  // コメント内容
    }
  ]
}
```

## コメント送信
- パスパラメータ：`submissionid`（タイムラインから遷移時に渡される値）  
- 一度に複数メッセージを送信可能  

### リクエスト形式
```json
{
  "comment": [
    "string"
  ]
}
```

# 英文入力画面

ここでは、自分が作成した英文を入力できます。  

## リクエスト形式
```json
{
  "text": "string"
}
```
POST リクエストとして送信します。  


# AI採点後画面

ここでは、英文入力画面で入力した英文の添削結果を確認し、投稿を確定します。  
データは **localStorage** に一時的に保存し、画面遷移後に削除する想定です。  

## リクエスト形式
```json
{
  "text": "string",
  "corrections": [
    {
      "original": "string",
      "corrected": "string",
      "reason": "string"
    }
  ]
}
```
この形式で送信されます。  


# 通知画面

ここでは、自分に対して他のユーザが行ったコメントの結果を表示します。  
**未閲覧の通知のみ** 表示されます（確認済みも表示するかはバックエンド修正で可能です）。  

## データ取得
クエリパラメータ：`page`, `size`  
GET リクエストを送信します。  

### レスポンス形式
```json
[
  {
    "userid": "string",       // 自分にコメントしたユーザの userid
    "nickname": "string",         // 自分にコメントしたユーザの名前
    "submissionid": "string",
    "date": "2025-09-13T09:05:56.525Z"
  }
]
```

---

# 通知詳細画面

ここでは、自分に行われた通知の詳細を確認できます。  
- パスパラメータ：`submissionid`, `userid`（通知画面から遷移時に受け取る）  
- GET リクエストを送信  

### レスポンス形式
```json
{
  "text": "string",           // 自分が作成した英文
  "nickname": "string",           // コメントしたユーザの名前
  "comments": [
    {
      "comment": "string"    // コメント内容
    }
  ]
}
```

> ※ 色付けなどのフロント表示に関しては、取得するデータ形式に応じてフロント側で制御可能です。  

---

# マイページ

ここでは、過去に自分が行った投稿の一覧を確認できます。  

## データ取得
クエリパラメータ：`page`, `size`  
GET リクエストを送信します。  

### レスポンス形式
```json
[
  {
    "nickname": "string",          // 自分の名前（不要なら省略可）
    "date": "2025-09-13",
    "text": "string",          // 自分の作成した文章
    "submissionid": "string"
  }
]
```


# 各投稿の詳細

ここでは、マイページから遷移して自分の投稿を確認できます。  

### レスポンス形式
```json
{
  "nickname": "string",           // 自分の名前
  "text": "string",           // 自分の作成した英文
  "corrections": [
    {
      "original": "string",
      "corrected": "string",
      "reason": "string"
    }
  ],
  "comments": [
    {
      "name": "string",       // コメントした人の名前
      "comment": "string"     // コメント内容
    }
  ]
}
```

---

# セッションで保持される情報
- `username`
- `userid`

---

# 修正点
- `admin` 判定を `isadmin` ではなく `role` として管理し、今後のロール追加に強い設計に変更。  
- `userid` の使われ方がニックネーム用途とユーザ識別用途で混在していたため、`userid` と `nickname` に分離。  
- 複数取得のある GET リクエストに **ページネーション** を追加。  
- JSON の形式を全て **小文字** に統一。  
- 他の人にコメントをする画面で、投稿者が受け取るコメント用の JSON データを追加。
