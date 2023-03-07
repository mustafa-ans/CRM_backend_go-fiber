package main

import(
	"github.com/gofiber/fiber"
	"github.com/Mustafa-ans/go-fibre/"

)

func setupRoutes(app *fiber.App){
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(Newlead)
	app.Delete(DeleteLead)

}

func initDatabase(){

}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)

	
}
