# Proyecto Golang-Brevo API

Este proyecto proporciona una interfaz en Go para interactuar con la API de Brevo, permitiendo enviar correos electrónicos y gestionar listas de contactos.

## Características

- Envío de correos electrónicos
- Obtención de listas de contactos
- Obtención del ID de una lista de contactos
- Añadir contactos a una lista de correo
- Uso de plantillas de correo locales

## Instalación

Para usar este paquete, asegúrate de tener Go instalado en tu sistema. Luego, puedes instalar el paquete con el siguiente comando:

```bash
go get gitlab.com/eliotandelon/brevo-email
```

Posteriormente, puedes importar el paquete en tu proyecto:

```go
import "gitlab.com/eliotandelon/brevo-email"
```

## Uso

### Inicialización del cliente

Para crear un nuevo cliente de Brevo, utiliza la función `NewBrevoMailClient` del paquete `brevoservice`:

```go
import "gitlab.com/eliotandelon/brevo-email/brevoservice"

client := brevoservice.NewBrevoMailClient(apiKey, templatePath, baseURL)
```

Donde:
- `apiKey` es tu clave API de Brevo
- `templatePath` es la ruta local donde se encuentran tus plantillas de correo
- `baseURL` es la URL base de la API de Brevo

### Interfaz de servicios

### Inicialización del cliente y servicios

Para utilizar el cliente de Brevo en tu aplicación, primero creas una instancia de `BrevoMailClient` y luego la pasas al constructor de tu servicio que requiera funcionalidad de correo electrónico. Aquí tienes un ejemplo de cómo hacerlo:

```go
import (
    "gitlab.com/eliotandelon/brevo-email/brevoservice"
    "gitlab.com/eliotandelon/brevo-email/brevointerface"
    "yourproject/services"
)

func main() {
    // Crear el cliente de Brevo
    brevoClient := brevoservice.NewBrevoMailClient(apiKey, templatePath, baseURL)

    // Crear el UserService pasando la interfaz BrevoServiceInterface
    userService := services.NewUserService(brevoClient)

    // Ahora puedes usar userService para operaciones que requieran envío de correos
    // ...
}
```

La estructura del `UserService` podría ser algo así:

```go
package services

import "gitlab.com/eliotandelon/brevo-email/brevointerface"

type UserService struct {
    brevoService brevointerface.BrevoServiceInterface
}

func NewUserService(brevoService brevointerface.BrevoServiceInterface) *UserService {
    return &UserService{
        brevoService: brevoService,
    }
}

// Métodos del UserService que usan brevoService...
func (s *UserService) SendWelcomeEmail(userEmail string) error {
    // Usar s.brevoService para enviar el correo de bienvenida
    // ...
}
```

En este ejemplo:

1. Creamos una instancia de `BrevoMailClient` usando `brevoservice.NewBrevoMailClient()`.
2. Pasamos esta instancia al constructor de `UserService`.
3. `UserService` acepta cualquier tipo que implemente la interfaz `BrevoServiceInterface`.
4. Dentro de `UserService`, podemos usar los métodos definidos en `BrevoServiceInterface` para enviar correos, gestionar listas, etc.

Este enfoque permite una mayor flexibilidad y facilita las pruebas, ya que puedes fácilmente sustituir la implementación real de Brevo por un mock en tus pruebas unitarias.

### Ejemplos de uso

#### Enviar un correo electrónico desde `UserService`

```go
// En el método de UserService para enviar correos
func (s *UserService) SendWelcomeEmail(userEmail string) error {
    reqMail := brevorequest.MailRequest{
        Subject: "Bienvenido a Nuestro Servicio",
        HtmlName: "template_bienvenida.html",
        ToDataMail: brevorequest.DataMail{
            Email: userEmail,
            Name: "Nombre Destinatario",  // Puedes personalizar el nombre
        },
        SenderDataMail: brevorequest.DataMail{
            Email: "noreply@nuestraempresa.com",
            Name: "Equipo de Soporte",
        },
    }

    // Personalización de la plantilla con datos específicos del usuario
    response, err := s.brevoService.SendMail(reqMail, "/v3/smtp/email", map[string]string{
        "NOMBRE": "Juan Pérez",
        "PRODUCTO": "Nuestro Servicio Premium",
    })

    if err != nil {
        return err
    }

    fmt.Println("Correo enviado:", *response)
    return nil
}
```

Explicación:
- `Subject`: Asunto del correo.
- `HtmlName`: Nombre del archivo de plantilla HTML a utilizar.
- `ToDataMail`: Datos del destinatario (email y nombre).
- `SenderDataMail`: Datos del remitente (email y nombre).
- El tercer parámetro de `SendMail` es un mapa con datos para personalizar la plantilla.

[... contenido sin cambios ...]

#### Añadir un contacto a una lista

```go
// En el método de UserService para añadir contactos
func (s *UserService) AddUserToList(userEmail string) error {
    attributes := brevorequest.CreateContactAtributtesRequest{
        LastName: "Pérez",         // Puedes personalizar este valor según el usuario
        Name: "Juan",              // Puedes personalizar este valor según el usuario
        SMS: "+34600000000",       // Número de teléfono del usuario
        Whatsapp: "+34600000000",  // Número de WhatsApp del usuario
    }

    // Añadir el contacto a una lista específica (ID = 1234)
    response, err := s.brevoService.AddContactToEmailList(userEmail, 1234, attributes)
    if err != nil {
        return err
    }

    fmt.Println("Contacto añadido:", response)
    return nil
}

```

Explicación:
- `LastName`: Apellido(s) del contacto (se guarda como "APELLIDOS" en Brevo).
- `Name`: Nombre del contacto (se guarda como "NOMBRE" en Brevo).
- `SMS`: Número de teléfono para SMS.
- `Whatsapp`: Número de teléfono para WhatsApp.
- El primer parámetro de `AddContactToEmailList` es el email del contacto.
- El segundo parámetro es el ID de la lista a la que se añade el contacto.

## Contribución

Las contribuciones son bienvenidas. Por favor, abre un issue o un pull request para sugerir cambios o mejoras.

## Licencia

Este proyecto está licenciado bajo los términos de la [Licencia Pública General GNU, versión 3 (GPLv3)](https://www.gnu.org/licenses/gpl-3.0.html).

### Resumen de la licencia:

- Puedes copiar, distribuir y modificar este proyecto siempre que mantengas la misma licencia.
- Cualquier modificación realizada y distribuida debe estar bajo los mismos términos.
- No hay garantías sobre el uso del software.

Puedes encontrar una copia completa de la licencia en el archivo [LICENSE](./LICENSE) en el repositorio de este proyecto.


