# ess-correction-service

## 概要

英語の文章をアップし、AI による校正、ユーザ同士での添削指導が可能な web サービス

## 主な機能

- AI による英文の校正
  　 - AI は、校正後の英⽂の表⽰をする機能に使⽤する - AI は、校正した英⽂の表⽰後、校正した箇所をピックアップし、校正前と後の⽐較とともに、校正前の英⽂の何がよくなかったかを説明する
- ユーザ同⼠の交流を可能にする掲⽰板機能
  - ユーザがアップした英⽂は、ほかのユーザも閲覧ができる
  - ユーザは、ほかのユーザのアップした英⽂に対し、コメントを残すことができる
  - ユーザによるコメント機能には、特定の英⽂の節に対する意⾒の際、英⽂全体のどの部分について述べているのかを誘導できる機能を求める

## 開発者向け

### ドキュメント

- [ブランチ戦略](docs/branch-strategy.md)
- [ディレクトリ構成](docs/directory-strategy.md)
- [backend API](docs/backend/api.md)
  - [swagger](docs/backend/swagger/swagger.md)

### 環境構築

```bash
npm install && npm run prepare
```

### Storybook

[http://localhost:6006/](http://localhost:6006/)から UI チェックができます．

```bash
npm run storybook
```
