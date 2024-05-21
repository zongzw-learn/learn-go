package main

// Relationship defines public interfaces
type Relationship interface {
	LoadRelations(csv string) error
	ShowRelationScores()
}
