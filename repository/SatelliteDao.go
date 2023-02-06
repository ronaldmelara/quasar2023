package repository

import (
	"meliQuasar/model"
	_ "meliQuasar/model"
)




func GetSatellites() []model.Satellite{

	kenobi := model.Satellite{ Id: 1, Name: "Kenobi", X: -500.0, Y: -200.0}
	skywalker := model.Satellite{ Id: 2, Name: "Skywalker", X: 100.0, Y: -100.0}
	sato := model.Satellite{ Id: 3, Name: "Sato", X: 500.0, Y: 100.0}

	satellites := []model.Satellite{kenobi,skywalker,sato}

	return satellites

}