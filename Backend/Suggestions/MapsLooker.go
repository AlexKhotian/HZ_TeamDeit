package Suggestions

type MapsLooker struct {
}

func (this *MapsLooker) GenerateLinkToGoogleMaps(lat string, lon string, city string) string {
	link := "https://www.google.ch/maps/search/lidl+" + city + "/@" + lat + "," + lon + "z"
	return link
}
