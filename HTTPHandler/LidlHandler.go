package HTTPHandler

import "HZ_proj/Database"

type LidlHandler struct {
	dbAcc *Database.DatabaseAccessorLidl
}

func LidlHandlerFactory() *LidlHandler {
	thisHandler := new(LidlHandler)
	thisHandler.dbAcc = new(Database.DatabaseAccessorLidl)
	thisHandler.dbAcc.OpenDB("admin:PUFYLPVWKIMWOLCB@/foodpolution")
	thisHandler.dbAcc.CreateDatabaseLidl()
	return thisHandler
}

func (handler *LidlHandler) GetPriceSubstitution(ing string) (float32, float32, string) {
	result := handler.dbAcc.GetPrice(ing)
	return result.Price, result.Discount, result.Ing
}
