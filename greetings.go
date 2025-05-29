package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para la persona especificada.
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("nombre vacio")
	}
	// Devuelve un saludo que incluye el nombre en un mensaje.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// Crea un mapa para almacenar los mensajes de saludo.
	messages := make(map[string]string)

	// Itera sobre los nombres y genera un saludo para cada uno.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"Saludos, %v! ¿Cómo estás?",
		"¡Qué alegría verte, %v!",
		"¡Hola, %v! ¿Qué tal tu día?",
		"¡Saludos cordiales, %v! Espero que estés bien.",
	}

	return formats[rand.Intn(len(formats))]
}
