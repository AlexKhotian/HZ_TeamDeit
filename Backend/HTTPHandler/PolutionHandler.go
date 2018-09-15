package HTTPHandler

import (
	"HZ_proj/Backend/Database"
	"HZ_proj/Backend/Suggestions"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CombinedResponse struct {
	InitialIng        string  `json:"InitialIng"`
	ReplaceIng        string  `json:"ReplaceIng"`
	FullPice          float32 `json:"FullPice"`
	Discount          float32 `json:"Discount"`
	GlobalWarming     float32 `json:"GlobalWarming"`
	EnergyConsumption float32 `json:"EnergyConsumption"`
	WaterUsage        float32 `json:"WaterUsage"`
	LinkWithShops     string  `json:"LinkWithShops"`
}

type PolutionHandler struct {
	dbAcc      *Database.DatabaseAccessor
	lidl       *LidlHandler
	mapslooker *Suggestions.MapsLooker
}

func PolutionHandlerFactory() *PolutionHandler {
	thisHandler := new(PolutionHandler)
	thisHandler.dbAcc = new(Database.DatabaseAccessor)
	thisHandler.dbAcc.OpenDB("admin:PUFYLPVWKIMWOLCB@tcp(sl-eu-de-1-portal.7.dblayer.com:16663)/foodpolution")
	thisHandler.dbAcc.CreateDatabasePolution()
	thisHandler.lidl = LidlHandlerFactory()
	thisHandler.mapslooker = new(Suggestions.MapsLooker)
	return thisHandler
}

func (handler *PolutionHandler) ProcessRequest(w http.ResponseWriter, ing string, lat string, lon string) {

	combRes := new(CombinedResponse)

	combRes.InitialIng = ing
	combRes.GlobalWarming, combRes.EnergyConsumption, combRes.WaterUsage = handler.ProcessPolution(ing)
	combRes.FullPice, combRes.Discount, combRes.ReplaceIng = handler.AskForSuggestion(ing)
	combRes.LinkWithShops = handler.mapslooker.GenerateLinkToGoogleMaps(lat, lon, "zuerich")
	fmt.Println(combRes.ReplaceIng)
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
