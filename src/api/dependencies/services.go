package dependencies

import "main/src/adapters/services"

type ServiceManager struct {
	BabyService    *services.BabyService
	PictureService interface{}
}

func CreateServiceManager() ServiceManager {
	managerInstance := ServiceManager{
		BabyService:    services.CreateBabyService(),
		PictureService: nil,
	}

	return managerInstance
}

var ServerManagerInstance = CreateServiceManager()
