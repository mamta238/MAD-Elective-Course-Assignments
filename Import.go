package main

import(
		"fmt"
		"os"
		"bufio"
		"strings"
		"encoding/json"	
		"gopkg.in/mgo.v2/bson"
		"gopkg.in/mgo.v2"
		)

//Restaurant struct and NewID(),StringToID refered from domain package 

type Restaurant struct {
	ID           bson.ObjectId  `json:"_id" bson:"_id"`
	Name         string  		`json:"name" bson:"name"`
	Address      string  		`json:"address" bson:"address"`
	AddressLine2 string  		`json:"address line 2" bson:"addressLine2"`
	URL          string  		`json:"url" bson:"url"`
	Outcode      string  		`json:"outcode" bson:"outcode"`
	Postcode     string  		`json:"postcode" bson:"postcode"`
	Rating       float32 		`json:"rating" bson:"rating"`
	TypeOfFood   string  		`json:"type_of_food" bson:"typeOfFood"`
}




func StringToID(s string) bson.ObjectId {
	return (bson.ObjectIdHex(s))
}

func NewID() bson.ObjectId {
	return StringToID(bson.NewObjectId().Hex())
}


//Passing json file,Dbname,colletion name to insert and host for mongodb connection

func ImportJsonToMongo(file, dbname, coll, host string){
	
	
	document := Restaurant{}
	
	fileHandle, err := os.Open(file)
	defer fileHandle.Close()
	
	if err!=nil{
		fmt.Println("Error in Opening")
		os.Exit(0)	
	}
	
	session,err := mgo.Dial(host)
	defer session.Close()
	
	if err!=nil{
		fmt.Println("Cannot create session")
		os.Exit(0)	
	}
	
	
	fs := bufio.NewScanner(fileHandle)
	collection := session.DB(dbname).C(coll)
	
	
	for fs.Scan() {
		
		data := fs.Text()
		dataB := []byte(data)
		json.Unmarshal(dataB,&document)
				
		document.ID = NewID()
		collection.Insert(&document) 
	
	}
	
}


//json file,dbname,collection name passed through command line

func main(){
	
	host := "localhost"
	
	switch args := len(os.Args) ; args {
	
	case 1 : fmt.Println("Error : no files listed!")
			 os.Exit(0)	
	
	case 2 : fmt.Println("Error : Dbname and Collection name not listed!")
			 os.Exit(0)
	
	case 3 : fmt.Println("Error : Collection name not listed!")
			 os.Exit(0)
	
	case 4 : file	:= strings.Trim(os.Args[1],"-")
			 dbname := strings.Trim(os.Args[2],"-")
			 coll	:= strings.Trim(os.Args[3],"-")	
			 ImportJsonToMongo(file,dbname,coll,host)	
	
	default : fmt.Println("Error-Expected: --JsonFile --Dbname --CollName")
	}
}
