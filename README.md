## GORILLA/MUX MGO API

Ejemplo de una API REST utilizando [GO] 1.10 con [gorilla/mux] como router y [mgo] como driver para mongoDB

####Endpoins disponibles

* `GET /products`
* `GET /products/:id`
* `POST /products`
* `PUT /products/:id`
* `DELETE /products/:id`

####Ejemplo de JSON
```json
{
  "name": "Fanta",
  "brand": "coca-cola",
  "variants": [
    {
      "description": "1,5 lts",
      "sku": "123456",
      "price": 1200,
      "stock": 5.0
    },
    {
      "description": "3 lts",
      "sku": "987654",
      "price": 2000,
      "stock": 7.0
    }
  ]
}
```
* Documentacion en progreso.
* Validaciones en progreso.
* Estructura de archivos en progreso.

[GO]:https://golang.org/
[gorilla/mux]:http://www.gorillatoolkit.org/pkg/mux
[mgo]:https://labix.org/mgo