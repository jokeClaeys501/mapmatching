package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Osm struct {
	XMLName     xml.Name `xml:"osm"`
	Text        string   `xml:",chardata"`
	Version     string   `xml:"version,attr"`
	Generator   string   `xml:"generator,attr"`
	Copyright   string   `xml:"copyright,attr"`
	Attribution string   `xml:"attribution,attr"`
	License     string   `xml:"license,attr"`
	Bounds      struct {
		Text   string `xml:",chardata"`
		Minlat string `xml:"minlat,attr"`
		Minlon string `xml:"minlon,attr"`
		Maxlat string `xml:"maxlat,attr"`
		Maxlon string `xml:"maxlon,attr"`
	} `xml:"bounds"`
	Node struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id,attr"`
		Visible   string `xml:"visible,attr"`
		Version   string `xml:"version,attr"`
		Changeset string `xml:"changeset,attr"`
		Timestamp string `xml:"timestamp,attr"`
		User      string `xml:"user,attr"`
		Uid       string `xml:"uid,attr"`
		Lat       string `xml:"lat,attr"`
		Lon       string `xml:"lon,attr"`
	} `xml:"node"`
	Way []struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id,attr"`
		Visible   string `xml:"visible,attr"`
		Version   string `xml:"version,attr"`
		Changeset string `xml:"changeset,attr"`
		Timestamp string `xml:"timestamp,attr"`
		User      string `xml:"user,attr"`
		Uid       string `xml:"uid,attr"`
		Nd        []struct {
			Text string `xml:",chardata"`
			Ref  string `xml:"ref,attr"`
		} `xml:"nd"`
		Tag []struct {
			Text string `xml:",chardata"`
			K    string `xml:"k,attr"`
			V    string `xml:"v,attr"`
		} `xml:"tag"`
	} `xml:"way"`
	Relation []struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id,attr"`
		Visible   string `xml:"visible,attr"`
		Version   string `xml:"version,attr"`
		Changeset string `xml:"changeset,attr"`
		Timestamp string `xml:"timestamp,attr"`
		User      string `xml:"user,attr"`
		Uid       string `xml:"uid,attr"`
		Member    []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Ref  string `xml:"ref,attr"`
			Role string `xml:"role,attr"`
		} `xml:"member"`
		Tag []struct {
			Text string `xml:",chardata"`
			K    string `xml:"k,attr"`
			V    string `xml:"v,attr"`
		} `xml:"tag"`
	} `xml:"relation"`
} 

type Street struct {
	name string
	maxSpeed float64
}



func getData(){
	// Open our xmlFile
    xmlFile, err := os.Open("map.osm")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened map.osm")
    // defer the closing of our xmlFile so that we can parse it later on
    defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(xmlFile)

    // we initialize our Users array
    var osm Osm
    // we unmarshal our byteArray which contains our
    // xmlFiles content into 'users' which we defined above
    xml.Unmarshal(byteValue, &osm)

	ways := osm.Way

    // we iterate through every user within our users array and
    // print out the user Type, their name, and their facebook url
    // as just an example
    for i := 0; i < len(ways); i++ {
		// get tags of way
		tags := ways[i].Tag;
		var streetTemp Street
		// iterate over tags and create streets
		for i := 0; i < len(tags); i++ {
			if (tags[i].K == "maxspeed"){
				speedFloat, err := strconv.ParseFloat(tags[i].V, 64)
				streetTemp.maxSpeed = speedFloat
				if err != nil {
        			log.Fatal(err)
    			}
			}
			if (tags[i].K == "name"){
				streetTemp.name = tags[i].V
			}
		}
		fmt.Println(streetTemp)
    }

}