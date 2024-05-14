package main

import (
	//Importaciones de go (vienen incluidas al instalar)
	"errors"
	"fmt"
	"os"

	//importaciones externas (descargadas)
	"context"

	"github.com/aws/aws-lambda-go/events"
	lambda02 "github.com/aws/aws-lambda-go/lambda"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomUser/awsgo"
	"github.com/PedroAntonioKira/ecommerceEscomUser/bd"
	"github.com/PedroAntonioKira/ecommerceEscomUser/models"
)

func main() {
	lambda02.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parametros, dee enviar 'SecretManager', no se ha enviado correctamente")
		err := errors.New("Error en los parametros, dee enviar 'SecretManager', no se ha enviado correctamente todo.")
		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)
		}
	}

	err := bd.ReadSecret()

	if err != nil {
		fmt.Println("Error al leer el Secret " + err.Error())
		return event, err
	}

	err = bd.SignUp(datos)

	return event, err
}

func ValidoParametros() bool {
	var traeParametro bool

	_, traeParametro = os.LookupEnv("SecretName")

	return traeParametro
}
