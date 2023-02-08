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

func GetMessages()[]model.SatelliteMessage{
	kenobiMsg := model.SatelliteMessage{Id: 1, Message: []string{"", "este", "es", "un", "mensaje"}}
	skywalkerMsg := model.SatelliteMessage{Id: 2, Message: []string{"este", "", "un", "mensaje"}}
	satoMsg := model.SatelliteMessage{Id: 2, Message: []string{"", "", "es", "", "mensaje"}}

	messages := []model.SatelliteMessage{kenobiMsg, skywalkerMsg, satoMsg}
	return messages
}