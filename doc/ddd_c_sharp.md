# C#の構文

## 目的

ドメイン駆動設計入門を読み進めるためのC#の構文まとめ

## 初期化
```c#
class Human
{
  // name というプライベートな変数を定義
  private readonly string name;

  // コンストラクタで初期化
  public Human (string name) {
    this.name = name
  }
}

Human human = new Human('テスト');
```

## プロパティ

```c#
class Human
{
  private string name;

  public string Name // m_nameのプロパティ
  {
  get { return name; }
  set { name = value; }
  }
}

Human human = new Human();
human.Name = "太郎";
human.Age = 20;
```

## using句
リソースの解放処理

```c#
using (var connection = new SqlConnection(connectionString))
{
  // do something
}
```
