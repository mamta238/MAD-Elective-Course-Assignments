package auxiliary

import (
		"domain"
		"fmt"
		"strings"
		"bufio"
		"os"
		dbrepo "dbrepository"
		)
		
		
func Init() (string,string,string,string){

	dbname := "Restaurant"
	prompt := "->"
	guidelines := "Type '-o' for options\n"
	options := "1) find     --type_of_food/--name/--id/--pcode \n2) list     -To display all Restaurants \n3) store    -To Insert new restaurant record \n4) delete   --id \n5) count    --type_of_food/--pcode \n6) -o        -For Options \n7)Search \n8) -q       Quit"
	
	return dbname,prompt,guidelines,options

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
		fmt.Println(strings.Repeat("-",100))
		
	}

func ListAllRestaurants (repoAccess *dbrepo.MongoRepository) {
	
	res , err := repoAccess.GetAll()
	if err!=nil {
		fmt.Println(err)
	} else {
		for i:=0 ; i<len(res) ; i++ {
		DisplayRec(res[i])
			}	
	}
}
	
func FindAccordingToOption(repoAccess *dbrepo.MongoRepository,option string,value string) {
		
	switch strings.Trim(option,"-") {
		
		case "type_of_food" : 	res , err := repoAccess.FindByTypeOfFood(value)
								fmt.Println(res,err)		
								if err!=nil {
								  		fmt.Println("yes",err)
								  	} else {
								  		for i:=0 ; i<len(res) ; i++{
								  		DisplayRec(res[i])
								  	}		
								}
								
			case "pcode"	: 	res , err := repoAccess.FindByTypeOfPostCode(value)
								if err!=nil {
								  		fmt.Println("yes",err)
								  	} else {
								 		for i:=0 ; i<len(res) ; i++{
								  		DisplayRec(res[i])
								  	}
								}
									
			case "name"		:	 res , err := repoAccess.FindByName(value)
								 if err!=nil {
								  		fmt.Println("yes",err)
								  	} else{
								  		for i:=0 ; i<len(res) ; i++{
								  		DisplayRec(res[i])
								  	}
								  }
								  
			case "id"		: 
								 if (domain.IsValidID(value)) {				//if valid id entered
									res , err := repoAccess.Get(domain.StringToID(value))
								  	if err!=nil{
								  		fmt.Println("yes",err)
								  	} else {
								  		DisplayRec(res)
								  	}
								  	
								  	} else {
								  		fmt.Println("Enter valid Id")
								  } 
			default  		:   fmt.Println("Invalid option for find")						  				
		}
	}
	
	
func CountAccordingToOption(repoAccess *dbrepo.MongoRepository,option string,value string) {
		
		switch strings.Trim(option,"-") {
		
			case "type_of_food" :  res ,_ := repoAccess.FindByTypeOfFood(value)
								   fmt.Println(len(res))	
			
			case "pcode"		:  res ,_ := repoAccess.FindByTypeOfPostCode(value)
								   fmt.Println(len(res))	
		}
}

	
func readValue() string {
		
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	return strings.Trim(value,"\n")
}	
	
func StoreRecord(repoAccess *dbrepo.MongoRepository) {
	
		doc := domain.Restaurant{}
		
		fmt.Println("Name:") 
		doc.Name = readValue()
		 
		fmt.Println("Address:")
		doc.Address = readValue()
		
		fmt.Println("AddressLine2:")
		doc.AddressLine2  = readValue()
		
		fmt.Println("URL:")
		doc.URL = readValue()
		
		fmt.Println("Outcode:")
		doc.Outcode  = readValue()
		
		fmt.Println("Postcode:")
		doc.Postcode = readValue()
		
		fmt.Println("Rating:")
		fmt.Scanf("%f",&doc.Rating)
		
		fmt.Println("TypeOfFood:")
		doc.TypeOfFood = readValue()
		
		fmt.Println(doc.Name,doc.Address)
		a,b := repoAccess.Store(&doc)
		fmt.Println(a,b)
		
}
	
func DeleteRecord(repoAccess *dbrepo.MongoRepository,id string) {
	
	if (domain.IsValidID(id)) {
		err := repoAccess.Delete(domain.StringToID(id))
		if err != nil{
			fmt.Println(err)
		}else {
				fmt.Println("Deleted Successfully")
		}
		
		} else {
			fmt.Println("Enter Valid Id")
		}	
}
	
	
func SearchOnKeyWord(repoAccess *dbrepo.MongoRepository, query string){
		
		res,err := repoAccess.Search(query)
		
		if err != nil {
			
			fmt.Println(err)				
		} else {
				for i:=0 ; i<len(res) ; i++{
					DisplayRec(res[i])
			}
		}
}
