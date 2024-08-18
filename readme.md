## protoファイル
- 以下でprotoファイルからコンパイルし、pb.goファイルを作成する
- protoc --go_out=. --go-grpc_out=. proto/file.proto

## file.pb.goファイル
- protoファイルのmessageを構造体に変換
- メソッドが複数存在(Reset,String,ProtoMessage,ProtoReflect,Descriptor,GetRequestname)

## file_grpc.pb.goファイル
- protoファイルのserviceをインターフェースに変換(クライアントとサーバーのインターフェースをそれぞれ用意)

## server/main.goファイル
- SayHelloメソッド
  - SayHelloメソッドを使って、リクエストの際にreqの部分を埋め、レスポンスで返す
- main関数
  - リスナー作成
  - gRPCサーバーインスタンスを作成
  - Greeterサービスをサーバーインスタンスに登録
  - サーバーを起動し、リスナーで受け取ったリクエストを処理する

## client/main.goファイル
- main関数
  - NewClientでgRPCサーバーに接続
  - Greeterサービスのクライアントインスタンスを作成
  - タイムアウト付きのコンテキストを設定
  - SayHelloメソッドを呼び出し、HelloRequestをサーバーに送信(あんまり腑に落ちてはいないが、、)

## 疑問点
- localhostのサーバーとクライアントって何？実際に接続しているわけではないのにサーバーはどこのこと？
  - localhostで動作するサーバーは、物理的に存在するものではなく、あくまでソフトウェア上で動作している仮想的なもの
  - サーバーという言葉は、物理的なコンピューターを指すこともあるが、localhostにおけるサーバーはプログラムとしてそのコンピューター上で動作しているプロセスを指す

- protoファイルのmessageとは？
  - messageは、データの構造を定義する部分
  - 各フィールドには型とフィールド番号があり、これによりプロトコルバッファーでシリアライズ/デシリアライズされる際のバイナリ形式が定義される
  - messageブロックは、pb.goファイルに変換される際に、structとして表現される

- protoファイル上のserviceとは？
  - serviceは、RPCメソッドを定義する
  - 各メソッドはリクエストメッセージとレスポンスメッセージを受け取り、それに応じた処理をする
  - pb.goファイルでは、対応するGoのインターフェースと、そのインターフェースを実装する構造体に変換される
  - また、各RPCメソッドに対応するクライアントおよびサーバー側のコードも生成される

- リッスンとは？
  - サーバーが特定のポートで接続要求を待ち受けるプロセス
  - これは、クライアントがサーバーに接続できるようにするための準備を整えることを意味する

- コンテキストとは？
  - Go言語で並行処理やリクエスト処理を制御するための重要な仕組み
  - リクエストの状態や制御をサーバーに伝えるために使用される
  - リクエストのタイムアウトやキャンセルを管理し、複数のゴルーチンで共通の情報を共有するのに役立つ
  - コンテキストを使うことで、処理の制御が簡単になり、より効率的なコードを書くことが出来る

- Go言語におけるインスタンス
  - Go言語におけるインスタンスとは、オブジェクト指向プログラミングにおけるオブジェクトの概念に似ている
  - 具体的には、Goでは構造体やインターフェースを使ってデータとその操作を定義し、その構造体やインターフェースをもとにインスタンスを作成する
```go
// 構造体のインスタンス
// 構造体のインスタンスとは、その構造体型の具体的な値のこと
type Person struct {
	Name string
	Age  int
}

func main() {
	// Person型のインスタンスを作成
	p := Person{
		Name: "Alice",
		Age:  30,
	}
	fmt.Println(p.Name)
}
```
```go
// インターフェースのインスタンス
// インターフェースのインスタンスは、そのインターフェース型に適合する実際の値を指す
type Greeter interface {
	Greet() string
}

type EnglishGreeter struct{}

func (g EnglishGreeter) Greet() string {
	return "Hello!"
}

func main() {
	var g Greeter
	g = EnglishGreeter{}
	fmt.Println(g.Greet()) // "Hello!"
}
```