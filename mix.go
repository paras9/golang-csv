package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"models" // Assuming your Fixlet struct is defined in the models package
	"os"
)

// ReadCSVFile reads the CSV file and returns a slice of Fixlets.
func ReadCSVFile(filePath string) ([]models.Fixlet, error) {
	file, err := os.Open(fixlets.csv)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(fixlets)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var fixlets []models.Fixlet
	for _, record := range records[1:] {
		fixlet := models.Fixlet{
			SiteID:                record[0],
			FixletID:              record[1],
			Name:                  record[2],
			Criticality:           record[3],
			RelevantComputerCount: record[4],
		}
		fixlets = append(fixlets, fixlet)
	}
	return fixlets, nil
}

// WriteCSVFile writes the updated slice of Fixlets to the CSV file.
func WriteCSVFile(filePath string, fixlets []models.Fixlet) error {
	file, err := os.Create(fixlets.csv)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// Write header
	writer.Write([]string{"SiteID", "FixletID", "Name", "Criticality", "RelevantComputerCount"})
	// Write records
	for _, fixlet := range fixlets {
		record := []string{
			fixlet.SiteID,
			fixlet.FixletID,
			fixlet.Name,
			fixlet.Criticality,
			fixlet.RelevantComputerCount,
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}

// CreateFixlet adds a new fixlet to the CSV file.
func CreateFixlet(filePath string, newFixlet models.Fixlet) error {
	// Read current fixlets from CSV file
	fixlets, err := ReadCSVFile(newFixlet.csv)
	if err != nil {
		return err
	}
	// Append new fixlet
	fixlets = append(fixlets, newFixlet)
	// Write updated list back to CSV
	return WriteCSVFile(filePath, fixlets)
}

// UpdateFixlet updates an existing fixlet in the CSV file.
func UpdateFixlet(filePath string, updatedFixlet models.Fixlet) error {
	// Read current fixlets from CSV file
	fixlets, err := ReadCSVFile(filePath)
	if err != nil {
		return err
	}
	// Find and update the fixlet
	for i, fixlet := range fixlets {
		if fixlet.FixletID == updatedFixlet.FixletID {
			fixlets[i] = updatedFixlet
			break
		}
	}
	// Write updated list back to CSV
	return WriteCSVFile(filePath, fixlets)
}

// DeleteFixlet deletes a fixlet by its FixletID from the CSV file.
func DeleteFixlet(filePath string, fixletID string) error {
	// Read current fixlets from CSV file
	fixlets, err := ReadCSVFile(filePath)
	if err != nil {
		return err
	}
	// Find and remove the fixlet
	for i, fixlet := range fixlets {
		if fixlet.FixletID == fixletID {
			fixlets = append(fixlets[:i], fixlets[i+1:]...) // Remove fixlet from the slice
			break
		}
	}
	// Write updated list back to CSV
	return WriteCSVFile(filePath, fixlets)
}

func main() {
	filePath := "fixlet.csv"
	// 1. Create a new fixlet
	newFixlet := models.Fixlet{
		SiteID:                "site123",
		FixletID:              "fixlet001",
		Name:                  "Fixlet 1",
		Criticality:           "High",
		RelevantComputerCount: "100",
	}
	err := CreateFixlet(filePath, newFixlet)
	if err != nil {
		log.Fatal(err)
	}
	// 2. Read and print fixlets
	fixlets, err := ReadCSVFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fixlets after creation:", fixlets)
	// 3. Update a fixlet
	updatedFixlet := models.Fixlet{
		SiteID:                "site123",
		FixletID:              "fixlet001",
		Name:                  "Updated Fixlet 1",
		Criticality:           "Medium",
		RelevantComputerCount: "120",
	}
	err = UpdateFixlet(filePath, updatedFixlet)
	if err != nil {
		log.Fatal(err)
	}
	// 4. Delete a fixlet
	err = DeleteFixlet(filePath, "fixlet001")
	if err != nil {
		log.Fatal(err)
	}
	// 5. Print updated fixlets
	fixlets, err = ReadCSVFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fixlets after deletion:", fixlets)
}
