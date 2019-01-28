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
	
	//For Command line interface
	for{
	
		fmt.Print("\n",prompt) 
		
		entry, _ := reader.ReadString('\n')
		
		entry = strings.TrimFunc(entry, func(c rune) bool {
			return c==rune('\n') || c==rune(' ') 
		})
		
		//Command pattern : command --option --value 
		//More arguments will result in Invalid command
		
		opt1 := strings.SplitN(entry, " " ,2)
		
		
		for i:=0 ; i<len(opt1) ; i++{
			opt1[i] = strings.Trim(opt1[i]," ")
		}
		
		option := opt1[0]
		value1,value2 := "",""	
		
		if len(opt1) > 1 {
			opt2 := strings.SplitN(opt1[1], " " ,2)
		
			for i:=0 ; i<len(opt2) ; i++{
				opt2[i] = strings.Trim(opt2[i]," ")
			}
			value1 = opt2[0]
			
			if len(opt2)>1{
				value2 = opt2[1]
		}		
		}
		
		switch option {  
								
			//functions from aux package
								
			case "find"   :  aux.FindAccordingToOption(repoAccess,value1,value2)
						     	
			case "list"   :  aux.ListAllRestaurants(repoAccess)
			
			case "store"  :  aux.StoreRecord(repoAccess)
			
			case "delete" :  aux.DeleteRecord(repoAccess,value2)
			
			case "count"  :  aux.CountAccordingToOption(repoAccess,value1,value2)
			
			case "search" :  aux.SearchOnKeyWord(repoAccess,value1)
			
			case "-o" 	  :  fmt.Println(options)
			
			case "-q"	  :  fmt.Println("\nExiting----")
							 os.Exit(0)	
					
			default 	  :  fmt.Println("\n\tInvalid Option!\n",guidelines)
		}
		
	}
}	



