package testutils

import (
	"os"
	"testing"
)

// Estructura de ejemplo
type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func TestOpenMock(t *testing.T) {
	jsonData := `{
		"id": "1",
		"name": "John",
		"surname": "Doe"
	}`

	file, err := os.Create("test.json")
	if err != nil {
		t.Fatalf("Error al crear archivo de prueba: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(jsonData)
	if err != nil {
		t.Fatalf("Error al escribir datos en el archivo: %v", err)
	}

	var user User

	result := OpenMock("test.json", &user)
	if result == nil {
		t.Fatalf("El resultado es nil, se esperaba un objeto deserializado")
	}

	if user.ID != "1" || user.Name != "John" || user.Surname != "Doe" {
		t.Errorf("Datos deserializados incorrectos: %#v", user)
	}

	err = os.Remove("test.json")
	if err != nil {
		t.Fatalf("Error al eliminar archivo de prueba: %v", err)
	}
}
