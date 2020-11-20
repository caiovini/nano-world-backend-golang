package services

import (
	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"net"
	"path/filepath"
	"regexp"
)

// GetGeoLocations fetch from GeoLite all locations
func GetGeoLocations() ([]*geoip2.City, error) {

	logger := gin.DefaultWriter
	db, err := geoip2.Open(filepath.Join("resources", "GeoLite2-City.mmdb"))
	defer db.Close()

	resp, err := GetPeers() // Comming from nano_node.go
	if err != nil {

		logger.Write([]byte(err.Error()))
		return nil, err
	}

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	var listOfLocations []*geoip2.City
	submatchall := re.FindAllString(string(resp), -1)
	for _, element := range submatchall {
		ip := net.ParseIP(element)
		record, err := db.City(ip)
		if err != nil {
			logger.Write([]byte(err.Error()))
			return nil, err
		}
		listOfLocations = append(listOfLocations, record)
	}

	if err != nil {
		logger.Write([]byte(err.Error()))
		return nil, err
	}

	return listOfLocations, nil
}
