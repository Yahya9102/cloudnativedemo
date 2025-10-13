package models

// User representerar samma struktur som i user-service
// MÅSTE MATCHA MED STRUKTUREN I USER-SERVICE
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}