package entry
import ("flag")
import ("fmt")


func entry{
var strUsername = flag.String("u", "", "Username")
var strPassword = flag.String("p", "", "Password")
flag.Parse()

fmt.Println("Username:", strUsername)
fmt.Println("Password:", strPassword)
}