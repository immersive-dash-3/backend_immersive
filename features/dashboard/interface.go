package dashboard

type DashboardEntity struct {
	CountActive    int64
	CountPlacement int64
	CountLog       int64
}

type DashboardDataInterface interface {
	GetData() (DashboardEntity, error)
}

type DashboardServiceInterface interface {
	GetData() (DashboardEntity, error)
}
