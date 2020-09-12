package repository

// ClientVersionRepository はクライアントバージョンの取得を抽象化します。
type ClientVersionRepository interface {
	Version(platform string) int
}
