package nhtsaclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/timgraf/vindecode/models"
)

func findStringAttribute(attributes []models.Attribute, attributeName string) string {
	for i := 0; i < len(attributes); i++ {
		if attributes[i].Variable == attributeName {
			return string(attributes[i].Value)
		}
	}

	return ""
}

func buildVehicle(attributes []models.Attribute) models.Vehicle {
	var a models.Vehicle

	a.Year        = findStringAttribute(attributes, "Model Year")
	a.Make        = findStringAttribute(attributes, "Make")
	a.Model       = findStringAttribute(attributes, "Model")
	a.Trim        = findStringAttribute(attributes, "Trim")
	a.VehicleType = findStringAttribute(attributes, "Vehicle Type")
	a.DriveType   = findStringAttribute(attributes, "Drive Type")

	return a
}

// DecodeVin decodes a VIN
func DecodeVin(vin string) models.Vehicle {
	url := fmt.Sprintf("https://vpic.nhtsa.dot.gov/api/vehicles/decodevin/%s?format=json", vin)
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var response models.VinResponse
	json.Unmarshal(body, &response)

	return buildVehicle(response.Attributes)
}