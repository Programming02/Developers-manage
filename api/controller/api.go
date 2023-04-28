package controller

type Api struct {
	AdminService      *AdminService
	ProgrammerService *ProgrammerService
	RegisterService   *RegisterService
}

func NewApi(admin *AdminService, programmer *ProgrammerService, register *RegisterService) *Api {
	return &Api{
		AdminService:      admin,
		ProgrammerService: programmer,
		RegisterService:   register,
	}
}
