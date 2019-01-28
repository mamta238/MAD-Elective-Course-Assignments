package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	aux "auxiliary" 
	dbrepo "dbrepository"
	mongoutils "utils"
)


func main() {

	//pass mongohost through the environment
	dbname, prompt, guidelines ,options := aux.Init()
	
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))
	repoAccess := dbrepo.NewMongoRepository(mongoSession, dbname)

	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println(guidelines)   
	
	for{
		fmt.Print("\n",prompt) 
		
		entry, _ := reader.ReadString('\n')
		entry = strings.Trim(entry,"\n")
		option := strings.Split(entry," ")
		
		fmt.Println(option)	
		
		switch strings.Trim(option[0],"->") {  
								
			case "find"   :  aux.FindAccordingToOption(repoAccess,option[1],option[2])
						     	
			case "list"   :  aux.ListAllRestaurants(repoAccess)
			
			case "store"  :  aux.StoreRecord(repoAccess)
			
			case "delete" :  aux.DeleteRecord(repoAccess,option[2])
			
			case "count"  :  aux.CountAccordingToOption(repoAccess,option[1],option[2])
			
			case "search" :  aux.SearchOnKeyWord(repoAccess,(option[1]))
			
			case "o" 	  :  fmt.Println(options)
			
			case "q"	  :  fmt.Println("\nExiting----")
							 os.Exit(0)	
					
			default 	  :  fmt.Println("\n\tInvalid Option!\n",guidelines)
		}
		
	}
}	



