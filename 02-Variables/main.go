package main

import "fmt"

// jwtToken := 300000
// not allowed the ":=" operator outside a method


const LoginToken string = "lknco23f13cjn13rgb1x2"

// Capital letters Public variable

func main() {
	var username string = "sidharth"
	fmt.Println(username)
	fmt.Printf("Var type: %T \n", username)

	var isLoggedIn bool = false

	fmt.Println(isLoggedIn)
	fmt.Printf("var type: %T \n", isLoggedIn)

	var smallVal uint8 = 253

	fmt.Println(smallVal)
	fmt.Printf("var type: %T \n", smallVal)


	var float float64 = 253.44411255

	fmt.Println(float)
	fmt.Printf("var type: %T \n", float)

	// aliases and default values

	var newVariable int;

	fmt.Println(newVariable);
	fmt.Printf("var type: %T \n", newVariable)

	// guess the output


	//_________________________________

	// implicit type

	var website = "psidh.vercel.app"

	fmt.Println(website);

	//_________________________________

	// no var declaration

	users := 90

	fmt.Println(users)


	fmt.Println(LoginToken)
	fmt.Printf("var type: %T \n", LoginToken)

}
