package HTTPHandler

import (
	"HZ_proj/Database"
)

type PolutionHandler struct {
	dbAcc *Database.DatabaseAccessor
}

func PolutionHandlerFactory() *PolutionHandler {
	thisHandler := new(PolutionHandler)
	thisHandler.dbAcc = new(Database.DatabaseAccessor)
	thisHandler.dbAcc.OpenDB("root:hztest@/foodpolution")
	thisHandler.dbAcc.CreateDatabasePolution()
	return thisHandler
}

func (handler *PolutionHandler) HandlerPolution(ing string) (uint32, uint32, uint32) {
	result := handler.dbAcc.GetRowFromIngs(ing)
	return result.GU, result.EC, result.WU
}
