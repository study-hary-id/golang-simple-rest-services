# Simple REST Services

List of simple rest services using Go programming language and some third party packages. e.g. (`gorilla/mux`
, `httprouter`, `alice`).

![Gopher Dance](images/gopher-dance.gif)

## <img src="images/pushing-cart.png" alt="Gopher pushing cart" width="48"> List Packages

1. [Custom Multiplexer](servemux)

   ServeMux is an HTTP request multiplexer, by creating a new `ServeMux` or
   `CustomServeMux` we can handle multiple routes. With `ServeMux` we can handle different endpoints using a
   function/handler other than an if/else statement.

2. [httprouter library](httprouter)

   `httprouter` routes the HTTP request by HTTP method to particular handler, it matches the REST methods such
   as (`GET`, `POST`, `PUT`, and so on). This library supports url queries and params by modifying the parameter of
   handler function. This library is very lightweight and fast performance.

3. [gorilla/mux library](gorilla)

   `gorila/mux` is a wonderful package for writing beautiful routes for web applications and API servers. Gorilla Mux
   provides tons of options to control routing such as path-based matching, query-base matching, domain-base matching,
   sub-domain base, and reverse URL generation.

4. [alice library](alice)

   `alice` library reduces the complexity of chaining the middleware when the list of middleware is big. Instead of
   calling a middleware within a middleware as a callback.

## <img src="images/crash-dummy.svg" alt="Gopher robo crash" width="24"> Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## <img src="images/heart-balloon.svg" alt="Gopher heart balloon" width="24"> Credits

The images, svg icons and gif animations are from [Gophers](https://github.com/egonelbre/gophers)

Copyright :copyright: 2022. This project is under [CC0](LICENSE) licensed.