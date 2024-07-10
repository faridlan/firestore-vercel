package domain

type User struct {
	ID   string `firestore:"id,omitempty"`
	Name string `firestore:"name"`
	Age  int    `firestore:"age"`
}
