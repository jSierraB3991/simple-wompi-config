# simple-wompi-config

Este módulo de golang proporciona la interface e inicialización básica para interactuar con la API de Wompi.

## Instalación

Puedes descargar este módulo en tu proyecto de Go especificando la versión deseada:

```bash
# Para obtener la última versión disponible (latest):
go get github.com/jSierraB3991/simple-wompi-config@latest

# O para fijar una versión específica (por ejemplo, v1.0.0):
go get github.com/jSierraB3991/simple-wompi-config@v1.0.0
```

## Uso

### La Interface

La interface principal para interactuar con el módulo es `wompiserviceinterface.WompiServiceInterface`. Su definición es la siguiente:

```go
type WompiServiceInterface interface {
	SendPaymentLink(req wompirequest.WompiSendPaymentRequest, cellphone string) (*wompiresponse.WompiPaymentResponse, error)
	WompiInfoPay(transactionId string) (*wompiresponse.WompiInfoTransactionDataResponse, error)
}
```

### Inicialización

Para inicializar el servicio de Wompi, utiliza el cliente `NewWompiClient` con las credenciales correspondientes.

Ejemplo de inicialización en tu servicio o contenedor de dependencias:

```go
wompiService := wompiservice.NewWompiClient(
	s.env.WompiPrivateKey, 
	s.env.WompiPublicKey, 
	s.env.WompiIntegrityKey, 
	s.env.WompiUrlApi,
)
```

### Ejemplo de Utilización

A continuación se muestra un ejemplo de cómo inyectar y utilizar el cliente en tu servicio para generar enlaces de pago y obtener los datos de una transacción:

```go
package service

import (
	// Suponiendo que has importado los paquetes del request/response de wompi
	// wompirequest "ruta-a-tu-wompi-request"
	// wompiresponse "ruta-a-tu-wompi-response"
	wompiserviceinterface "github.com/jSierraB3991/simple-wompi-config/wompi_domain/wompi_service_interface"
)

type Service struct {
	wompiClient wompiserviceinterface.WompiServiceInterface
}

// Función para crear un enlace de pago (Nequi/Wompi)
// Datos de pruebas para cellphone:
// 3991111111 for an approved transaction (APPROVED)
// 3992222222 for a declined transaction (DECLINED)
func (s *Service) PayByNequiWompi(nameUser string, description string, redirectUrl string, orderId string, customerEmail string, amountInCent int64, cellPhone string) (*wompiresponse.WompiPaymentResponse, error) {
	wompiResponse, err := s.wompiClient.SendPaymentLink(
		wompirequest.WompiSendPaymentRequest{
			Metadata: wompirequest.MetadataReq{
				Name:        nameUser,
				Description: description,
				RedirectUrl: redirectUrl,
			},
			Reference:     orderId,
			CustomerEmail: customerEmail, // Asignado a la variable del parámetro en vez de kUser.Email
			AmountInCents: amountInCent,
		}, cellPhone)
		
	if err != nil {
		return nil, err
	}
	
	// save response in db
	return wompiResponse, nil
}

// Función para obtener los datos de una transacción
func (s *Service) GetWompiTransactionData(transactionId string) (*wompiresponse.WompiInfoTransactionDataResponse, error) {
	wompiResponse, err := s.wompiClient.WompiInfoPay(transactionId)
	
	if err != nil {
		return nil, err
	}
	
	// save response in db
	return wompiResponse, nil
}
```