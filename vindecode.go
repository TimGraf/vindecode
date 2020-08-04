package main

import (
	"fmt"
	"log"
	"os"

	"github.com/timgraf/vindecode/nhtsaclient"
)

func main() {
	vin := os.Args[1]

	vehicle := nhtsaclient.DecodeVin(vin)

	log.Println(fmt.Sprintf("Year:  %s", vehicle.Year))
	log.Println(fmt.Sprintf("Make:  %s", vehicle.Make))
	log.Println(fmt.Sprintf("Model: %s", vehicle.Make))
	log.Println(fmt.Sprintf("Trim:  %s", vehicle.Trim))
	log.Println(fmt.Sprintf("Type:  %s", vehicle.VehicleType))
	log.Println(fmt.Sprintf("Drive: %s", vehicle.DriveType))
}
