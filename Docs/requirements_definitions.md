# 開発対象システム総称
- ブログシステム

# システム種類
- ブログ閲覧システム(閲覧者操作可能)
- ブログ管理システム(管理者操作可能)

# 操作対象者
- ブログ閲覧者
- ブログ管理者

# 要件定義
## 閲覧者
### ブログ記事 一覧閲覧
- 画面: ブログ記事一覧画面
- 機能説明: 閲覧者はブログ記事の一覧を閲覧することができる
- 表示項目:
    - ブログタイトル
    - ブログ記事本文の一部
    - サムネイル
    - 投稿した日付

### ブログ記事 詳細閲覧
- 画面: ブログ詳細画面
- 機能説明: 閲覧者は特定のブログ記事を閲覧する事ができる
- 表示項目: 
    - ブログタイトル
    - ブログ記事本文
    - 画像
    - 投稿した日付

### ブログタイトルによる検索
- 画面: ブログ一覧画面上部
- 機能説明: 閲覧者は入力したブログタイトルに部分一致したブログ記事のみを表示させる事ができる
           閲覧者はブログタイトルを使用してブログ記事を検索することができる
- 入力項目:
    - ブログタイトル
- 補足:
    - 画面一部分に表示

### タグ 一覧閲覧
- 画面: ブログ一覧画面右下部分
- 機能説明: 閲覧者は全てのタグを閲覧することができる
- 表示項目:
    - タグ名
- 補足:
    - 画面一部分に表示
    - タグ: 1つのブログに複数紐付けられるキーワードのこと

### タグによる検索
- 画面: ブログ一覧画面右下部分
- 機能説明: 閲覧者はタグ名を入力→タグリンクを押下してブログ記事を検索することができる
- 入力項目: 
    - タグ名
- 補足:
    - 画面一部分に表示 
    - タグ: 1つのブログに複数紐付けられるキーワードのこと

### 新規投稿 ブログ記事一覧　閲覧
- 画面: ブログ一覧画面右部
- 機能説明: 最近投稿した記事を過去5件まで閲覧することができる
- 補足:
    - 画面一部分に表示

### 新規投稿 ブログ記事詳細 閲覧
- 画面: ブログ詳細画面下部
- 機能説明: 最近投稿した記事を過去5件まで閲覧することができる
- 補足:
    - 画面一部分に表示

## 管理者
### ログイン機能
- 画面: ログイン画面
- 機能説明: 管理者はパスワード/メールアドレスを使用してログインをすることができる
- 入力項目:
    - メールアドレス
    - パスワード

### 管理者情報 詳細機能
- 画面: 管理者詳細画面
- 機能説明: 管理者は管理者の詳細情報を閲覧することができる
- 表示項目:
    - 名前
    - メールアドレス
    - パスワード

### 管理者情報 編集機能
- 画面: 管理者編集機能
- 機能説明: 管理者は管理者情報を編集することができる
- 入力項目:
    - 名前
    - メールアドレス
    - パスワード

### ブログ記事 一覧閲覧
- 画面: ブログ記事一覧画面
- 機能説明: 管理者はブログ記事の一覧を閲覧することができる
- 表示項目:
    - ブログタイトル
    - ブログ記事本文の一部
    - サムネイル
    - 公開/非公開
    - 投稿した日付

### ブログ記事 詳細閲覧
- 画面: ブログ記事詳細画面
- 機能説明: 管理者は特定のブログ記事を閲覧する事ができる
- 表示項目: 
    - ブログタイトル
    - ブログ記事本文
    - 画像
    - タグ
    - 公開/非公開
    - 投稿した日付

### ブログ記事 新規投稿
- 画面: ブログ記事投稿画面
- 機能説明: 管理者はブログ記事を投稿する事ができる
- 入力項目: 
    - ブログタイトル
    - ブログ記事本文
      - 本文
      - 画像
    - ブログサムネイル
    - 公開/非公開
    - タグ選択
        - タグ名

### ブログ記事 編集
- 画面: ブログ編集画面
- 機能説明: 管理者はブログ記事を編集/更新する事ができる
- 入力項目:
    - ブログタイトル
    - ブログ記事本文
    - サムネイル
    - 公開/非公開

### ブログ記事 削除
- 画面: ブログ一覧画面, 詳細画面 一部分
- 機能説明: 管理者はブログ記事を削除する事ができる

### タグ 作成
- 画面: ブログ記事投稿画面
- 機能説明: 管理者はタグを作成する事ができる
- 入力項目:
    - タグ名
    - タグサムネイル

### 投稿履歴機能
- 画面: 投稿履歴一覧画面
- 機能説明: 管理者は投稿履歴一覧を閲覧することができる
- 表示項目:
    - 管理者名
    - ブログ記事タイトル
    - 時間