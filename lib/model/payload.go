package model

// DataRecord holds the data for one line in the csv.
// This can be modified to represent different datasets
// Note: the ID is stored sepearately and can be reused across datasets
type DataRecord struct {
	Name        string
	Email       string
	CountryCode string
	Mobile      string
}

// Record holds the csv line data and attributes an id
// By using string IDs, this can be reused across datasets
type Record struct {
	ID     string
	Record DataRecord
}
