//User登録のためのSQLスキーマ
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  is_admin BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

//Userが提出した本文、またそのAI出力文を保存するためのSQLスキーマ
CREATE TABLE submissions (
  id UUID PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  text TEXT NOT NULL,
  res TEXT, -- AI出力文
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

//Userの提出した本文に対して、AIが修正したセンテンスを保存するためのSQLスキーマ
CREATE TABLE corrections (
  id SERIAL PRIMARY KEY,
  submission_id UUID REFERENCES submissions(id) ON DELETE CASCADE,
  original TEXT NOT NULL,
  corrected TEXT NOT NULL,
  reason TEXT NOT NULL,
  is_human BOOLEAN DEFAULT FALSE
);

//commentsテーブルは、ユーザーが他のユーザーの提出物にコメントを追加するためのものです
CREATE TABLE comments (
  id SERIAL PRIMARY KEY,
  submission_id UUID REFERENCES submissions(id) ON DELETE CASCADE,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  text TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
