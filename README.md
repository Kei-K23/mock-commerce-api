# MockCommerceAPI

MockCommerceAPI is a free, fast online RESTFul API that provides realistic data for e-commerce or shopping websites and powered by **Go**. It can be used without the need for any server-side code and improve development time without waiting to backend team. Just fetch the MockCommerceAPI, build UI interfaces with best realistic e-commerce data for prototyping and later integrate with your actual API. This API is ideal for educational purposes, demonstrations app, testing, practicing new front-end frameworks by building e-commerce website and much more.

For more detailed documentation, visit the [MockCommerceAPI Documentation]().

## Why Use MockCommerceAPI?

When developing e-commerce or shopping website prototypes, it can be challenging to find realistic data for testing and prototyping. Typically, developers resort to using lorem ipsum text, finding placeholder images or manually creating JSON files. To address this need, MockCommerceAPI was created, offering a simple solution that delivers semi-realistic shopping data via an online web service built using **Go (Gin)** and **Postgres**.

## Available Resources

MockCommerceAPI provides key resources commonly required in e-commerce or shopping websites:

- [Products]()
- [Carts]()

## How to Use

You can fetch data using any method you prefer, such as `fetch API`, `Axios`, `jQuery Ajax`, or any other language specific HTTP client etc. As examples code snippets, I will use `JavaScript fetch API` (In most of the case but not always, front-end website design will use `JavaScript` to build UI components)

### Get All Products

```js
fetch("http://localhost:8080/products")
  .then((res) => res.json())
  .then((json) => console.log(json));
```

### Get a Single Product

```js
fetch("http://localhost:8080/products/1")
  .then((res) => res.json())
  .then((json) => console.log(json));
```

### Add a New Product

```js
fetch("http://localhost:8080/products", {
  method: "POST",
  body: JSON.stringify({
    title: "Product-99",
    price: 12,
    description: "This is description for product-99",
    image: "https://i.pravatar.cc",
    category: "electronic",
  }),
})
  .then((res) => res.json())
  .then((json) => console.log(json));
```

**Response**:

```json
{
  "id": 99,
  "title": "Product-99",
  "price": 12,
  "description": "This is description for product-99",
  "image": "https://i.pravatar.cc",
  "category": "electronic"
}
```

> **Note**: The posted data will not be inserted into the actual database. Instead, a fake ID will be returned instead.

### Update a Product

```js
fetch("http://localhost:8080/products/7", {
  method: "PUT",
  body: JSON.stringify({
    title: "Product-99 Updated",
    price: 12,
    description: "This is description for product-99",
    image: "https://i.pravatar.cc",
    category: "electronic",
  }),
})
  .then((res) => res.json())
  .then((json) => console.log(json));
```

**Response**:

```json
{
  "id": 7,
  "title": "Product-99 Updated",
  "price": 12,
  "description": "This is description for product-99",
  "image": "https://i.pravatar.cc",
  "category": "electronic"
}
```

> **Note**: The data will not be updated in the database.

### Delete a Product

```js
fetch("http://localhost:8080/products/10", {
  method: "DELETE",
});
```

> **Note**: The data will not be deleted from the database.

### Sort and Limit Results

You can limit results or sort them in ascending or descending order using query strings.

```js
fetch("http://localhost:8080/products?limit=3&sort=price DESC")
  .then((res) => res.json())
  .then((json) => console.log(json));
```

## API Endpoints

### Products

Fields:

```json
{
  "id": Number,
  "title": String,
  "price": Number,
  "category": String,
  "description": String,
  "image": String
}
```

#### GET:

- `/products` (Retrieve all products)
- `/products/1` (Retrieve a specific product by ID)
- `/products?limit=5` (Limit the number of products)
- `/products?sort=desc` (Sort products in descending order)
- `/products?category=electronic` (Filter products with category)
- `/products?title=product-99` (Filter products with product title name)

#### POST:

- `/products` (Add a new product)

#### PUT/PATCH:

- `/products/1` (Update a product by ID)

#### DELETE:

- `/products/1` (Delete a product by ID)

### Categories

Fields:

```json
{
  "id": Number,
  "title": String,
  "description": String,
  "image": String
}
```

#### GET:

- `/categories` (Retrieve all categories)
- `/categories/1` (Retrieve a specific category by ID)
- `/categories?limit=5` (Limit the number of categories)
- `/categories?sort=desc` (Sort categories in descending order)
- `/categories?title=electronic` (Filter by categories with category title name)

#### POST:

- `/categories` (Add a new category)

#### PUT/PATCH:

- `/categories/1` (Update a category by ID)

#### DELETE:

- `/categories/1` (Delete a category by ID)
