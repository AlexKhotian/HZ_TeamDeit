package HTTPHandler

import (
	"HZ_proj/Database"
	"encoding/json"
	"log"
	"net/http"
)

type CombinedResponse struct {
	InitialIng        string
	ReplaceIng        string
	FullPice          float32
	Discount          float32
	GlobalWarming     float32
	EnergyConsumption float32
	WaterUsage        float32
}

type PolutionHandler struct {
	dbAcc *Database.DatabaseAccessor
	lidl  *LidlHandler
}

func PolutionHandlerFactory() *PolutionHandler {
	thisHandler := new(PolutionHandler)
	thisHandler.dbAcc = new(Database.DatabaseAccessor)
	thisHandler.dbAcc.OpenDB("root:hztest@/foodpolution")
	thisHandler.dbAcc.CreateDatabasePolution()
	thisHandler.lidl = LidlHandlerFactory()
	return thisHandler
}

func (handler *PolutionHandler) ProcessRequest(w http.ResponseWriter, ing string) {

	combRes := new(CombinedResponse)

	combRes.InitialIng = ing
	combRes.GlobalWarming, combRes.EnergyConsumption, combRes.WaterUsage = handler.ProcessPolution(ing)
	combRes.FullPice, combRes.Discount, combRes.ReplaceIng = handler.AskForSuggestion(ing)

	response, errM := json.Marshal(&combRes)
	if errM != nil {
		log.Println(errM)
		return
	}
	w.Write(response)
}

func (handler *PolutionHandler) ProcessPolution(ing string) (float32, float32, float32) {
	result := handler.dbAcc.GetRowFromIngs(ing)
	return result.GU, result.EC, result.WU
}

func (handler *PolutionHandler) AskForSuggestion(data string) (float32, float32, string) {
	return handler.lidl.GetPriceSubstitution(data)
}
