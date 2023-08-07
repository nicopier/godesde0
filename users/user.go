package users

import (
	"fmt"
	"time"

	"github.com/nicopier/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Nicolas", time.Now(), true)
	fmt.Println(u.Id)
	fmt.Println(u.Name)
	fmt.Println(u.CreatedAt.Format("02-01-2006 15:30"))
	fmt.Println(u.Status)
}
