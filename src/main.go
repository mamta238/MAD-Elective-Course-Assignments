package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"domain"
	dbrepo "dbrepository"
	mongoutils "utils"
)


//variables for command line interface

func Init() (string,string,string,string){

	dbname := "Ass1"
	prompt := "->"
	guidelines := "Type '-o' for options\n"
	options := "1) find     --type_of_food/--name/--id/--pcode \n2) list     -To display all Restaurants \n3) store    -To Insert new restaurant record \n4) delete   --id \n5) count    --type_of_food/--pcode \n6) -o        -For Options \n7) -q       Quit"
	
	return dbname,prompt,guidelines,options

}


func main() {

	//pass mongohost through the environment
	dbname, prompt, guidelines ,options := Init()
	
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))
	repoAccess := dbrepo.NewMongoRepository(mongoSession, dbname)

	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println(guidelines)   
	
	for{
		fmt.Print("\n",prompt) 
		
		entry, _ := reader.ReadString('\n')
		entry = strings.Trim(entry,"\n")
		option := strings.Split(entry," ")
		
		
		
		switch strings.Trim(option[0],"->") {  
			
			case "find"   :  FindAccordingToOption(repoAccess,option[1],option[2])
						     	
			case "list"   :  ListAllRestaurants(repoAccess)
			
			case "store"  :  StoreRecord(repoAccess)
			
			case "delete" :  DeleteRecord(repoAccess,option[2])
			
			case "count"  :  CountAccordingToOption(repoAccess,option[1],option[2])
			
			case "o" 	  :  fmt.Println(options)
			
			case "q"	  :  fmt.Println("\nExiting----")
							 os.Exit(0)	
			
			default 	  :  fmt.Println("\n\tInvalid Option!\n",guidelines)
		}
		
	}
}	


	func DisplayRec(res *domain.Restaurant){
		
		
			fmt.Println("\nName:",(*res).Name)
			fmt.Println("Address:",(*res).Address)
			fmt.Println("Address2:",(*res).AddressLine2)
			fmt.Println("Name:",(*res).Name)
			fmt.Println("URL:",(*res).URL)
			fmt.Println("Outcode:",(*res).Outcode)
			fmt.Println("Postcode:",(*res).Postcode)
			fmt.Println("Rating:",(*res).Rating)
			fmt.Println("Type Of Food:",(*res).TypeOfFood)
			fmt.Println("--------")
		
	}

	func ListAllRestaurants (repoAccess *dbrepo.MongoRepository) {
		res ,_ := repoAccess.GetAll()
		for i:=0 ; i<len(res) ; i++ {
			DisplayRec(res[i])
			}	
		}

	func FindAccordingToOption(repoAccess *dbrepo.MongoRepository,option string,value string) {
		
		switch strings.Trim(option,"-") {
		
			case "type_of_food" : res , _ := repoAccess.FindByTypeOfFood(value)
								  for i:=0 ; i<len(res) ; i++{
								  DisplayRec(res[i])
								  }		
			
			case "pcode"		: res , _ := repoAccess.FindByTypeOfPostCode(value)
								  for i:=0 ; i<len(res) ; i++{
								  DisplayRec(res[i])
								  }
								  	
			case "name"			: res , _ := repoAccess.FindByName(value)
								  for i:=0 ; i<len(res) ; i++{
								  DisplayRec(res[i])
								  }
			case "id"			: res , _ := repoAccess.Get(domain.ID(value))
								  DisplayRec(res)		
		}
	}
	
	
	func CountAccordingToOption(repoAccess *dbrepo.MongoRepository,option string,value string) {
		
		switch strings.Trim(option,"-") {
		
			case "type_of_food" : res ,_ := repoAccess.FindByTypeOfFood(value)
								  fmt.Println(len(res))	
			
			case "pcode"		: res ,_ := repoAccess.FindByTypeOfPostCode(value)
								  fmt.Println(len(res))	
		}
	}
	
	func StoreRecord(repoAccess *dbrepo.MongoRepository) {
	
		doc := domain.Restaurant{}
		
		reader := bufio.NewReader(os.Stdin) 
		fmt.Print("Name:")
		doc.Name , _ = reader.ReadString('\n')
		fmt.Print("Address:")
		doc.Address , _ = reader.ReadString('\n')
		fmt.Print("AddressLine2:")
		doc.AddressLine2, _  = reader.ReadString('\n')
		fmt.Print("URL:")
		doc.URL , _ = reader.ReadString('\n')
		fmt.Print("Outcode:")
		doc.Outcode , _ = reader.ReadString('\n')
		fmt.Print("Postcode:")
		doc.Postcode , _ = reader.ReadString('\n')
		fmt.Print("Rating:")
		fmt.Scanf("%f",doc.Rating)
		fmt.Print("TypeOfFood:")
		doc.TypeOfFood , _ = reader.ReadString('\n')
		
		fmt.Println(doc.Name,doc.Address)
		a,b := repoAccess.Store(&doc)
		fmt.Println(a,b)
		
	}
	
	func DeleteRecord(repoAccess *dbrepo.MongoRepository,id string) {
		
		fmt.Println(repoAccess.Delete(domain.ID(id)))
	}
	
	
	
