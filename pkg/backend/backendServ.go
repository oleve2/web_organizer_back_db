package backendServ

type Service struct {
	sqlitePath string
}

func NewService(sqlitePath string) *Service {
	return &Service{
		sqlitePath: sqlitePath,
	}
}
