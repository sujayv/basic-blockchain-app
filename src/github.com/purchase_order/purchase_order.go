package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
	"math/rand"
	"strconv"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("CLDChaincode")

type SimpleChaincode struct {
}

type Purchase_Order struct {
	PoId            string	`json:"PoId"`
	Quantity        int	`json:"Quantity"`
	Part_Name       string	`json:"Part_Name"`
	Customer        string	`json:"Customer"`
	Supplier        string	`json:"Supplier"`
	Status          string	`json:"Status"`
}

type PoId_Holder struct {
	Po 	[]string	`json:"pos"`
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response  {

	fmt.Println("blockchain init method")
	var PoIds PoId_Holder
	bytes, err := json.Marshal(PoIds)
	if err != nil {
	 	return shim.Error("Error creating PoId_Holder record")
	}
	err = stub.PutState("PoIds", bytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("blockchain invoke method")
	function, args := stub.GetFunctionAndParameters()
	if function != "invoke" {
                return shim.Error("Unknown function call")
	}
	if args[0] == "delete" {
		// Deletes a purchase order from its state
		return t.delete(stub, args)
	}
	if args[0] == "queryPO" {
		// Queries a purchase order in its state
		return t.queryPurchaseOrder(stub, args)
	}
	if args[0] == "queryPOIds" {
		//Queries the list of POIds
		return t.queryAllPurchaseOrders(stub, args)
	}
	if args[0] == "createPO" {
		// Creates a purchase order in its state
		return t.createPO(stub, args)
	}
	if args[0] == "createCompletePO" {
		// Creates a complete purchase order from its state
		return t.createCompletePO(stub, args)
	}
	if args[0] == "updateStatus" {
		// Updates the status of a purchase order in its state
		return t.updateStatus(stub, args)
	}
	if args[0] == "updateQuantity" {
		// Updates the quantity of a purchase order in its state
		return t.updateQuantity(stub, args)
	}
	if args[0] == "updatePartName" {
		// Updates the part name of a purchase order in its state
		return t.updatePartName(stub, args)
	}
	if args[0] == "updateSupplier" {
		// Updates the supplier of a purchase order in its state
		return t.updateSupplier(stub, args)
	}
	if args[0] == "updateCustomer" {
		// Updates the customer of a purchase order in its state
		return t.updateCustomer(stub, args)
	}
	/*if args[0] == "initialize" {
		// Updates the status of a purchase order in its state
		return t.defaultinit(stub, args)
	}*/
	return shim.Error("Unknown Invoke Method")
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface) pb.Response {

	return shim.Error("Query has been implemented in invoke")
}

//Update the status of a purchase order
func (t *SimpleChaincode) updateStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var poid string // Entities
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	poid = args[1]
	// Query the purchase order from the state in ledger
	bytes, err := stub.GetState(poid)
	if err != nil {
		return shim.Error("Unable to get the purchase order")
	}
	var po Purchase_Order
	err = json.Unmarshal(bytes, &po)
	if err != nil {
		return shim.Error("Error while unmarshalling JSON object")
	}
	po.Status = args[2]
	resp  := t.save_order(stub, po)
	if resp != true {
		fmt.Println("UPDATE_PO: Error saving changes: %s", resp);
		return shim.Error("Error saving changes")
	}
	return shim.Success([]byte("Successfully updated status of order"+args[1]+" to "+args[2]))
}

//Update Quantity of Purchase Order
func (t *SimpleChaincode) updateQuantity(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var poid string // Entities
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	poid = args[1]
	// Query the purchase order from the state in ledger
	bytes, err := stub.GetState(poid)
	if err != nil {
		return shim.Error("Unable to get the purchase order")
	}
	var po Purchase_Order
	err = json.Unmarshal(bytes, &po)
	if err != nil {
		return shim.Error("Error while unmarshalling JSON object")
	}
	temp,err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("The entered value is not an integer")
	}
	if temp < 0 {
		return shim.Error("The entered value has to be positive")
	}
	po.Quantity = temp
	resp  := t.save_order(stub, po)
	if resp != true {
		fmt.Println("UPDATE_PO: Error saving changes: %s", resp);
		return shim.Error("Error saving changes")
	}
	return shim.Success([]byte("Successfully updated quantity of order"+args[1]+" to "+args[2]))
}

//Update Customer of Purchase Order
func (t *SimpleChaincode) updateCustomer(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var poid string // Entities
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	poid = args[1]
	// Query the purchase order from the state in ledger
	bytes, err := stub.GetState(poid)
	if err != nil {
		return shim.Error("Unable to get the purchase order")
	}
	var po Purchase_Order
	err = json.Unmarshal(bytes, &po)
	if err != nil {
		return shim.Error("Error while unmarshalling JSON object")
	}
	po.Customer = args[2]
	resp  := t.save_order(stub, po)
	if resp != true {
		fmt.Println("UPDATE_PO: Error saving changes: %s", resp);
		return shim.Error("Error saving changes")
	}
	return shim.Success([]byte("Successfully updated customer of order"+args[1]+" to "+args[2]))
}

//Update Supplier of Purchase Order
func (t *SimpleChaincode) updateSupplier(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var poid string // Entities
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	poid = args[1]
	// Query the purchase order from the state in ledger
	bytes, err := stub.GetState(poid)
	if err != nil {
		return shim.Error("Unable to get the purchase order")
	}
	var po Purchase_Order
	err = json.Unmarshal(bytes, &po)
	if err != nil {
		return shim.Error("Error while unmarshalling JSON object")
	}
	po.Supplier = args[2]
	resp  := t.save_order(stub, po)
	if resp != true {
		fmt.Println("UPDATE_PO: Error saving changes: %s", resp);
		return shim.Error("Error saving changes")
	}
	return shim.Success([]byte("Successfully updated supplier of order"+args[1]+" to "+args[2]))
}

//Update part name of purchase order
func (t *SimpleChaincode) updatePartName(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var poid string // Entities
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	poid = args[1]
	// Query the purchase order from the state in ledger
	bytes, err := stub.GetState(poid)
	if err != nil {
		return shim.Error("Unable to get the purchase order")
	}
	var po Purchase_Order
	err = json.Unmarshal(bytes, &po)
	if err != nil {
		return shim.Error("Error while unmarshalling JSON object")
	}
	po.Part_Name = args[2]
	resp  := t.save_order(stub, po)
	if resp != true {
		fmt.Println("UPDATE_PO: Error saving changes: %s", resp);
		return shim.Error("Error saving changes")
	}
	return shim.Success([]byte("Successfully updated part name of order"+args[1]+" to "+args[2]))
}

//Stores the purchase order provided in json format
func (t *SimpleChaincode) save_order(stub shim.ChaincodeStubInterface, po Purchase_Order) bool {

	fmt.Println("Saving the new purchase order")
	bytes, err := json.Marshal(po)
	if err != nil {
		fmt.Println("SAVE_CHANGES: Error converting vehicle record: %s", err)
		return false
	}
	err = stub.PutState(po.PoId, bytes)
	if err != nil {
		fmt.Println("SAVE_CHANGES: Error storing vehicle record: %s", err)
		return false
	}
	return true
}

//Creates a purchase order in the blockchain
func (t *SimpleChaincode) createPO(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("Creating a new purchase order")
	//poid := t.generatePoid()
	poid := args[1]
	po := Purchase_Order{
		PoId: poid,
		Quantity: 0,
		Part_Name: "UNDEFINED",
		Customer: "UNDEFINED",
		Supplier: "UNDEFINED",
		Status: "UNDEFINED",
	}
	record, err := stub.GetState(po.PoId)
	if record != nil {
		 return shim.Error("Purchase Order already exists")
	}
	resp  := t.save_order(stub, po)
	if resp != true {
		fmt.Println("CREATE_PO: Error saving changes: %s", resp);
		return shim.Error("Error saving changes")
	}
	bytes, err := stub.GetState("PoIds")
	if err != nil {
		return shim.Error("Unable to get PoId Holder")
	}
	var PoIds PoId_Holder
	err = json.Unmarshal(bytes, &PoIds)
	if err != nil {
		return shim.Error("Corrupt PoId_Holder record")
	}
	PoIds.Po = append(PoIds.Po, poid)
	bytes, err = json.Marshal(PoIds)
	if err != nil {
	 	fmt.Println("Error creating PoId_Holder record")
	}
	err = stub.PutState("PoIds", bytes)
	if err != nil {
		return shim.Error("Unable to put the state")
	}
	fmt.Println("Successfully created PO with ID"+poid)
	return shim.Success([]byte(poid))
}

//Creates a fully defined purchase order in the blockchain
func (t *SimpleChaincode) createCompletePO(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("Creating a new purchase order")
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}
	//poid := t.generatePoid()
	poid := args[1]
	quantity,err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("The entered value is not an integer")
	}
	if quantity < 0 {
		return shim.Error("The entered value has to be positive")
	}
	part := args[3]
	customer := args[4]
	supplier := args[5]
	status := args[5]
	po := Purchase_Order{
		PoId: poid,
		Quantity: quantity,
		Part_Name: part,
		Customer: customer,
		Supplier: supplier,
		Status: status,
	}
	record, err := stub.GetState(po.PoId)
	if record != nil {
		 return shim.Error("Purchase Order already exists")
	}
	resp  := t.save_order(stub, po)
	if resp != true {
		fmt.Println("CREATE_PO: Error saving changes: %s", resp);
		return shim.Error("Error saving changes")
	}
	bytes, err := stub.GetState("PoIds")
	if err != nil {
		return shim.Error("Unable to get PoId Holder")
	}
	var PoIds PoId_Holder
	err = json.Unmarshal(bytes, &PoIds)
	if err != nil {
		return shim.Error("Corrupt PoId_Holder record")
	}
	PoIds.Po = append(PoIds.Po, poid)
	bytes, err = json.Marshal(PoIds)
	if err != nil {
	 	fmt.Println("Error creating PoId_Holder record")
	}
	err = stub.PutState("PoIds", bytes)
	if err != nil {
		return shim.Error("Unable to put the state")
	}
	fmt.Println("Successfully created PO with ID"+poid)
	return shim.Success([]byte(poid))
}

//Deletes a purchase order from the block chain
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	poid := args[1]

	// Delete the key from the state in ledger
	err := stub.DelState(poid)
	if err != nil {
		return shim.Error("Failed to delete state")
	}
	bytes, err := stub.GetState("PoIds")
	if err != nil {
		return shim.Error("Unable to get PoIds")
	}
	var PoIds PoId_Holder
	err = json.Unmarshal(bytes, &PoIds)
	if err != nil {
		return shim.Error("Corrupt PoId_Holder record")
	}

	var index int
	//Find the index at which the current purchase order id exists in the array
	for i:=0;i<len(PoIds.Po);i++ {
		if PoIds.Po[i] == poid {
			index = i
			break
		}
	}
	temp_po := PoIds.Po
	temp_po = append(temp_po[:index], temp_po[index+1:]...)
	PoIds.Po = temp_po
	bytes, err = json.Marshal(PoIds)
	if err != nil {
	 	fmt.Println("Error creating PoId_Holder record")
	}
	err = stub.PutState("PoIds", bytes)
	if err != nil {
		return shim.Error("Unable to put the state")
	}
	fmt.Println("Successfully deleted PO with id"+poid)
	return shim.Success(nil)
}

//Query the entire list of purchase order ids
func (t *SimpleChaincode) queryAllPurchaseOrders(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("Inside the queryAllPurchaseOrders")
	bytes, err := stub.GetState("PoIds")
	if err != nil {
		return shim.Error("Unable to get PoIds")
	}
	var PoIds PoId_Holder
	err = json.Unmarshal(bytes, &PoIds)
	if err != nil {
		return shim.Error("Corrupt PoId_Holder record")
	}
	temp_po := PoIds.Po
	var temp string
	temp = ""
	for i:=0;i<len(temp_po);i++ {
		temp += temp_po[i]+","
	}
	return shim.Success([]byte(temp))
}

//Query based on a purchase order ID
func (t *SimpleChaincode) queryPurchaseOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("Inside the queryPurchaseOrder")
	var poid string // Entities
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	poid = args[1]
	// Query the purchase order from the state in ledger
	bytes, err := stub.GetState(poid)
	if err != nil {
		return shim.Error("Unable to get the purchase order")
	}
	var po Purchase_Order
	err = json.Unmarshal(bytes,&po)
	if err != nil {
		fmt.Println("The record is corrupt")
		return shim.Error("Unable to unmarshal purchase order")
	}
	po_details,err := json.Marshal(po)
	if err != nil {
		return shim.Error("Unable to marshal the unmarshaled purchase order")
	}
	return shim.Success(po_details)
}



//Function to generate random number based purchase order id
func (t *SimpleChaincode) generatePoid() string {

	poid := ""
	for i:= 0;i<7;i++ {
	    poid += strconv.Itoa(rand.Intn(10))
	}
	return poid;
}

func main() {

	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Println("Error starting Chaincode: %s", err)
	}
}
