package main

import (
	"fmt"	
	"strings"
)
var familyMap = make(map[string][]string)
// contains checks if a specific value is present in a slice
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
func addFamily(){
	fmt.Println("====================Start of Func Adding a Family====================")
	var family string
	fmt.Println("Enter the node_family")
	fmt.Scanln(&family)
	fmt.Println("Product family",family)
	_, ok := familyMap[family]
	if !ok {
		familyMap[family] = append(familyMap[family])
	}
	fmt.Println("====================End of Func Adding a Family====================")
}

func addToNodeTypeWithFamily(){
	var userInput string
	fmt.Println("====================Start of Func Add NodeType to family====================")
	fmt.Println("Enter the input with the format = \nnode_family-->node_Type")
	fmt.Scanln(&userInput)
	fmt.Printf("You entered: %s\n", userInput)
	splitArray :=strings.Split(userInput, "-->")
	family :=splitArray[0]
	node:= splitArray[1]
	fmt.Println("Extracted Product Family = ",family)
	fmt.Println("Extracted Node Type = ",node)
	
	_, ok := familyMap[family]
	if !ok {
		familyMap[family] = append(familyMap[family], node)
	}else{
		if contains(familyMap[family], node) {
			fmt.Printf("%s family is present in the value %s\n",family,node)
		} else {
			familyMap[family] = append(familyMap[family], node)
		}
	}
	fmt.Println("====================End of Func Add NodeType to family====================")
}
func displayFamily(){
	fmt.Println("====================Start of Display of family====================")
	fmt.Printf("\n\n")
	for key, value := range familyMap {
		fmt.Printf("Node Family = %s => Node types = %v\n", key, value)
	}
	fmt.Println("====================End of Display of family====================")
}
func deleteNodeType(){
	var deleteFromFamily string;
	var deleteNodeType string;
	fmt.Println("====================Start of Deleting a of Node Type====================")
	fmt.Println("Enter family type:")
	fmt.Scanln(&deleteFromFamily)
	fmt.Println("Enter the node type to delete:")
	fmt.Scanln(&deleteNodeType)
	nodeTypeSlice := familyMap[deleteFromFamily]
	for i, v := range nodeTypeSlice {
		if v == deleteNodeType {		
			familyMap[deleteFromFamily]= append(nodeTypeSlice[:i], nodeTypeSlice[i+1:]...)
		}
	}
	fmt.Println("Delete Operation performed Successfully")
	fmt.Println("====================Start of Deleting a of Node Type====================")
}
func recursiveFun(){
	var option int = 0;
	fmt.Println("====================Start of Recursive Function====================")
	fmt.Println("1. Add Vendor")
	fmt.Println("2. Add node family with Node type")
	fmt.Println("3. Delete node type")
	fmt.Println("4. Display all vendor and node types")
	fmt.Println("5. Exit")
	fmt.Println("Enter your option : ")
	fmt.Scanln(&option);
	switch option{
	case 1:addFamily()
	case 3:deleteNodeType()
	case 2:addToNodeTypeWithFamily()
	case 4:displayFamily()
	case 5:
		fmt.Println("====================End of Recursive Function====================")
		return
	default: recursiveFun()
	}
	recursiveFun()
}
func main() {
	recursiveFun()
}
