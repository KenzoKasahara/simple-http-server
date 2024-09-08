# CLIでリクエストを送信する
curl http://localhost:8080
curl http://localhost:8080/2

# ヘッダー情報を取得する
curl -I http://localhost:8080   # ヘッダー情報のみ(-I/--head)
curl -i http://localhost:8080   # ヘッダー情報とボディ情報(-i/--include)

# リクエストボディを送信する
curl -X POST -d "name=John" http://localhost:8080

# リクエストヘッダーを送信する
curl -H 'Host:example.com.' http://localhost:8080/

# リクエストヘッダーとボディを送信し、レスポンスヘッダーを取得する
curl -X POST -d "name=John" -H 'Host:example.com.' -H 'SampleHeader:TEST' http://localhost:8080
curl -X POST -d "name=John" -H 'Host:example.com.' -H 'SampleHeader:TEST' -v http://localhost:8080
