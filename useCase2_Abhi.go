package main

import (
	"fmt"	
	"strings"
)
type NspEvent struct {
	date string
	name string
	primaryContact string
	organizingTeam OrganizingTeam
}
type OrganizingTeam struct {
	team [] string
	primaryContact string
}
var eventDetail []NspEvent
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
func addEvent()  {
	fmt.Println("====================Start of Func Add Event====================")
	var event1 NspEvent;
	var organizingTeam1 OrganizingTeam
	
	fmt.Println("Enter the Event Details")
	fmt.Printf("Event name = ")
	fmt.Scanln(&event1.name)
	fmt.Printf("Event Date = ")
	fmt.Scanln(&event1.date)
	fmt.Printf("Event Primary Contact = ")
	fmt.Scanln(&event1.primaryContact)
	fmt.Println("Organizing team details")
	fmt.Println("Organizing team members with comma seperated values :")
	var members string
	fmt.Scanln(&members)

	for{
		organizingTeam1.team=strings.Split(members,",")
		fmt.Println("Organizing team Primary contact")
		fmt.Scanln(&organizingTeam1.primaryContact)
		if contains(organizingTeam1.team, organizingTeam1.primaryContact) {
			event1.organizingTeam=organizingTeam1
			break;
		}else{
			fmt.Println("ReEnter the Organizing team Primary contact due to member not present under team")
		}
	}
	eventDetail=append(eventDetail,event1)
	fmt.Println("====================End of Func Add Event====================")
}
func displayEvent()  {
	fmt.Println("====================Start of Func Display Event====================")
	for _, v := range eventDetail {
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println("Event Details : ")
		fmt.Println("Event name = ",v.name)
		fmt.Println("Event Date = ",v.date)
		fmt.Println("Event Primary Contact = ",v.primaryContact)
		fmt.Println("Organizing team details")
		fmt.Println("Organizing team members",v.organizingTeam.team)
		fmt.Println("Organizing team Primary contact",v.organizingTeam.primaryContact)
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	}
	fmt.Println("====================End of Func Display Event====================")
}
func deleteEvent(){
	var deleteEventName string;
	fmt.Println("====================Start of Deleting a of Event===================")
	fmt.Printf("Enter Event to delete = ")
	fmt.Scanln(&deleteEventName)
	
	eventSlice := eventDetail
	for i, v := range eventSlice {
		if v.name == deleteEventName {		
			eventDetail= append(eventDetail[:i], eventDetail[i+1:]...)
		}
	}
	fmt.Println("Delete Operation performed Successfully")
	fmt.Println("====================Start of Deleting a of Event====================")
}
func recursiveFun(){
	var option int = 0;
	fmt.Println("====================Start of Recursive Function====================")
	fmt.Println("1. Add Event")
	fmt.Println("2. Display All Event")
	fmt.Println("3. Delete Event")
	fmt.Println("4. Exit")
	fmt.Println("Enter your option : ")
	fmt.Scanln(&option);
	switch option{
	case 1:addEvent()
	case 2:displayEvent()
	case 3:deleteEvent()
	case 4:
		fmt.Println("====================End of Recursive Function====================")
		return
	default: recursiveFun()
	}
	recursiveFun()
}
func main(){
recursiveFun()
}