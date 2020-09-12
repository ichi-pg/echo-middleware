package repository

// MaintenanceRepository はメンテナンス状態の取得を抽象化します。
type MaintenanceRepository interface {
	Active() bool
	Message() string
}
