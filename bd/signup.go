package bd

import (
	//Importaciones de go (vienen incluidas al instalar)
	"fmt"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomUser/models"
	"github.com/PedroAntonioKira/ecommerceEscomUser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnect()

	//Verificamos no haya un error al conectarse a la base de datos.
	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"

	//Verificamos que la sentencia SQL este bien escrita
	fmt.Println(sentencia)

	//Ejecutamos la sentencia SQL
	_, err = Db.Exec(sentencia)

	//Verificamos no exista un error al querer ejecutar la sentencia SQL
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(" > SignUp > Ejecución Exitosa")
	return nil
}
