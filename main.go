
package main

import ...


func main() {
    app := fiber.New()

    //INITIAL ROUTE
    route.RouteInit(app)

    app.Listen(addr:":8080")
}
