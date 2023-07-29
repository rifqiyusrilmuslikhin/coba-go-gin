## Requirement  
- Go (version 1.13 or higher)
- Postgresql Database
## Installation and Usage  
1. Clone this repository
   ```shell
   git clone https://github.com/rifqiyusrilmuslikhin/coba-go-gin
   cd coba-go-gin
2. Install all module
   ```shell
   go mod tidy
3. Rename file `.env.example` to `.env` and configure database url and port
4. Run migrations database
   ```shell
   go run migrations/migrations.go
5. Run application
   ```shell
   go run main.go
## Api Endpoint  
The application will run on `http://localhost:4000`  
1. API Endpoint: /get-data
- GET: Get data product from external API and save to database 
2. API Endpoint: /products
- GET: Retrieve a list of products.
- POST: Create a new products.  
  Request body: 
  ```json
  {
    "title": string,
    "price": int,
    "description": string,
    "category": string,
    "image": string
  }  
3. API Endpoint: /products/{id}
- GET: Retrieve product details based on ID.
- PUT: Update product information based on ID.  
  Request body:
  ```json
  {
    "title": string,
    "price": int,
    "description": string,
    "category": string,
    "image": string
  }  
- DELETE: Delete a product based on ID.
