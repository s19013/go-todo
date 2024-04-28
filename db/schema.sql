-- マイグレーションが何故かできないのでコメントで日付とか残すしかない
-- 2024/04/28 create_todos
CREATE TABLE IF NOT EXISTS todos (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  subject TEXT NOT NULL,
  description TEXT NOT NULL DEFAULT '',
  created_at DATETIME NOT NULL DEFAULT (DATETIME('now')),
  updated_at DATETIME NOT NULL DEFAULT (DATETIME('now')),
  CHECK(subject <> '') -- subjectは空文字も許さない
  -- check成約でデータを追加する時に値が指定した条件を満たしているかどうかのチェック
  -- この場合は､subject列が空でないことを確認する
);

-- 2024/04/28 create_trigger_todos_updated_at

-- goは更新した時にupdated_atが更新されるような仕組みがないらしいので自分で作るしかない
CREATE TRIGGER IF NOT EXISTS trigger_todos_updated_at
AFTER UPDATE ON todos -- todosをアップデートしたら
BEGIN -- 更新した行のupdate_atを今の時間に書き換える
    UPDATE
        todos
    SET
        updated_at = DATETIME('now')
    WHERE
        id == NEW.id; --なんかこれで更新したやつをしていできるらしい｡
END;