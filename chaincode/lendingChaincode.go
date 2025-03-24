package main

import (
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Loan struct represents a loan request
type Loan struct {
	LoanID           string  `json:"loanId"`
	BorrowerID       string  `json:"borrowerId"`
	LenderID         string  `json:"lenderId"`
	Amount           float64 `json:"amount"`
	InterestRate     float64 `json:"interestRate"`
	Duration         int     `json:"duration"`
	Status           string  `json:"status"` // Pending, Approved, Active, Repaid, Defaulted
	DisbursementDate string  `json:"disbursementDate"`
	RepaymentDue     float64 `json:"repaymentDue"`
	RemainingBalance float64 `json:"remainingBalance"`
	Defaulted        bool    `json:"defaulted"`
}

// SmartContract provides functions for managing loans
type SmartContract struct {
	contractapi.Contract
}

// TODO: Implement function to request a loan
func (s *SmartContract) RequestLoan(ctx contractapi.TransactionContextInterface, loanID, borrowerID string, amount float64, interestRate float64, duration int) error {
	// TODO: Add logic to create a new loan request and store it on the blockchain
	if len(loanID) == 0 {
		fmt.Println("loan Id cannot be null")
		return errors.New("loan Id cannot be null")
	}

	if len(borrowerID) == 0 {
		fmt.Println("borrowerID cannot be null")
		return errors.New("borrowerID cannot be null")
	}

	if amount <= 0 {
		fmt.Println("amount cannot be 0 or less")
		return errors.New("amount cannot be 0 or less")
	}

	if interestRate <= 0 {
		fmt.Println("interestRate cannot be 0 or less")
		return errors.New("interestRate cannot be 0 or less")
	}

	if duration <= 0 {
		fmt.Println("duration cannot be 0 or less")
		return errors.New("duration cannot be 0 or less")
	}

	data, err := ctx.GetStub().GetState(loanID)
	if err != nil {
		return fmt.Errorf("Failed to retrieve loanID: %v", err)
	}
	if data != nil {
		fmt.Println("loanID already exists")
		return fmt.Errorf("loanID already exists")
	}

	var loandPayload = Loan{
		LoanID:     	loanID,
		BorrowerID:	 	borrowerID,
		Amount:     	amount,
		InterestRate:   interestRate,
		Duration:       duration,
		Status:         "Pending",
	}

	loanPayloadInByte, err := json.Marshal(loandPayload)
	if err != nil {
		fmt.Printf("error in marshalling data %v \n", err)
		return err
	}

	err = ctx.GetStub().PutState(loanID, loanPayloadInByte)
	if err != nil {
		return fmt.Errorf("Failed to store loan request data: %v", err)
	}
	return nil
}

// TODO: Implement function to approve a loan
func (s *SmartContract) ApproveLoan(ctx contractapi.TransactionContextInterface, loanID, lenderID string) error {
	// TODO: Add logic to approve a loan and update its status
	if len(loanID) == 0 {
		fmt.Println("loan Id cannot be null")
		return errors.New("loan Id cannot be null")
	}
	
	data, err := ctx.GetStub().GetState(loanID)
	if err != nil {
		return fmt.Errorf("Failed to retrieve loanID: %v", err)
	}
	if data == nil {
		return fmt.Errorf("loanID not found")
	}

	var loandPayload Loan
	err = json.Unmarshal(data, &loandPayload)
	if err != nil {
		fmt.Printf("error in unmarshalling loan payload %v \n", err)
		return err
	}

	if loandPayload.Status != "Pending" {
		fmt.Printf("cannot approve non Pending Loans \n")
		return errors.New("cannot approve non Pending Loans")
	}

	loandPayload.Status = "Approved"
	loandPayload.LenderID = lenderID

	loanPayloadInByte, err := json.Marshal(loandPayload)
	if err != nil {
		fmt.Printf("error in marshalling data %v \n", err)
		return err
	}

	err = ctx.GetStub().PutState(loanID, loanPayloadInByte)
	if err != nil {
		return fmt.Errorf("Failed to update loan status data: %v", err)
	}

	return nil
}

// TODO: Implement function to repay a loan
func (s *SmartContract) RepayLoan(ctx contractapi.TransactionContextInterface, loanID string, amount float64) error {
	// TODO: Add logic to process a loan repayment and update the remaining balance
	return fmt.Errorf("RepayLoan function not implemented")
}

// TODO: Implement function to query a loan by ID
func (s *SmartContract) QueryLoan(ctx contractapi.TransactionContextInterface, loanID string) (*Loan, error) {
	// TODO: Add logic to fetch loan details from the blockchain
	return nil, fmt.Errorf("QueryLoan function not implemented")
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating chaincode: %v", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting chaincode: %v", err)
	}
}
