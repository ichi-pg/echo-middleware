# golang-middleware

よく使うgolangのパッケージです。

## Logger

echo標準のLoggerではKPIやお問い合わせ用途としては情報が足りません。<br/>
リクエストパラメーターやレスポンス内容を詳細にログ出力します。

## Recover

echo標準のRecoverとほぼ一緒ですが、このパッケージのLoggerに合わせたログ出力をします。

## AccessControl

アクセスにIP制限をかけます。

## DebugMode

アクセスをデバッグモードに制限します。

## Maintenance

メンテナンス状況を確認します。

## ClientVersion

クライアントの強制アップデートバージョンを確認します。
