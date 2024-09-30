# Proyecto ePayco API

Este proyecto proporciona una interfaz en Go para interactuar con la API de ePayco, permitiendo realizar pagos electrónicos en Colombia, iniciar sesión en ePayco y obtener información sobre transacciones.

## Instalación

Para usar este paquete, asegúrate de tener Go instalado en tu sistema. Luego, puedes instalar el paquete usando:

```
go get gitlab.com/eliotandelon/epayco-pays
```

## Estructura del Proyecto

El proyecto consta de varios paquetes clave:

- `epaycoservice`: Contiene la implementación del cliente de ePayco.
- `epaycoserviceinterface`: Define la interfaz para los servicios de ePayco.
- `epaycorequest`: Contiene las estructuras de solicitud para la API de ePayco.
- `epaycoresponse`: Contiene las estructuras de respuesta de la API de ePayco.

## Uso

### Inicialización del cliente y servicios

Para utilizar el cliente de ePayco en tu aplicación, primero creas una instancia de `EpaycoService` y luego la pasas al constructor de tu servicio que requiera funcionalidad de pagos. Aquí tienes un ejemplo de cómo hacerlo:

```go
import (
    "gitlab.com/eliotandelon/epayco-pays/epaycoservice"
    "gitlab.com/eliotandelon/epayco-pays/epaycoserviceinterface"
    "yourproject/services"
)

func main() {
    // Crear el cliente de ePayco
    epaycoClient := epaycoservice.NewEpaycoService(baseURL, publicKey, privateKey)

    // Crear el UserService pasando la interfaz EpaycoServiceInterface
    userService := services.NewUserService(epaycoClient)

    // Ahora puedes usar userService para operaciones que requieran pagos
    // ...
}
```

La estructura del `UserService` podría ser algo así:

```go
package services

import "gitlab.com/eliotandelon/epayco-pays/epaycoserviceinterface"

type UserService struct {
    epaycoService epaycoserviceinterface.EpaycoServiceInterface
}

func NewUserService(epaycoService epaycoserviceinterface.EpaycoServiceInterface) *UserService {
    return &UserService{
        epaycoService: epaycoService,
    }
}

// Métodos del UserService que usan epaycoService...
```

### Ejemplos de uso

#### Iniciar sesión en ePayco

```go
loginResponse, err := userService.epaycoService.LoginEpayco()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Login exitoso:", loginResponse)

// Guardar el token para usarlo en otras operaciones
token := loginResponse.Token
```

#### Realizar un pago

Para realizar un pago, primero debes iniciar sesión y obtener un token. Luego, usa ese token en la llamada al método de pago:

```go
// Asegúrate de haber iniciado sesión y obtenido el token primero

paymentRequest := epaycorequest.PaymentRequest{
    PublicKey:       "tu_llave_publica",
    PrivateKey:      "tu_llave_privada",
    Amount:          "10000",
    Currency:        "COP",
    Description:     "Pago de prueba",
    Title:           "Producto de prueba",
    Country:         "CO",
    Test:            "true",
    UrlResponse:     "https://tudominio.com/respuesta",
    UrlConfirmation: "https://tudominio.com/confirmacion",
    NameBilling:     "Juan Pérez",
}

paymentResponse, err := userService.epaycoService.Pay(paymentRequest, token)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Pago realizado:", paymentResponse)
```

#### Obtener lista de transacciones

```go
// Asegúrate de haber iniciado sesión y obtenido el token primero

// Para obtener todas las transacciones
transactionRequest := epaycorequest.TransactionListRequest{}

// O para filtrar por referencia de cliente
transactionRequest := epaycorequest.TransactionListRequest{
    Filter: epaycorequest.FilterData{
        ReferenceClient: "ref_123",
    },
}

transactions, err := userService.epaycoService.TransactionList(transactionRequest, token)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Transacciones obtenidas:", transactions)
```

## Contribución

Las contribuciones son bienvenidas. Por favor, abre un issue o un pull request para sugerir cambios o mejoras.

## Licencia

Este proyecto está licenciado bajo los términos de la [Licencia Pública General GNU, versión 3 (GPLv3)](https://www.gnu.org/licenses/gpl-3.0.html).

### Resumen de la licencia:

- Puedes copiar, distribuir y modificar este proyecto siempre que mantengas la misma licencia.
- Cualquier modificación realizada y distribuida debe estar bajo los mismos términos.
- No hay garantías sobre el uso del software.

Puedes encontrar una copia completa de la licencia en el archivo [LICENSE](./LICENSE) en el repositorio de este proyecto.