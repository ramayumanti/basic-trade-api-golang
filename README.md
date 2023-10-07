**Basic Trade API**

A basic e-commerce API, made for my Golang class final assignment.
This API enables user to add products, product images/photos, and add variants and set their quantities.
It uses PostgreSQL for database.

**How it works**

Before a user can use this API, they need to be registered and logged in. Through the login process, user will receive a token that is used for authorization. This token is needed to create, update, and delete new products/variants. Getting the product or variants information does not need authorization.
When adding a product, a user needs to provide the name of the product and the product image. The image will be uploaded and contained in my personal Cloudinary storage, and users will be able to access the image links generated from the upload process. After that, they can add variants to the products with product's UUID and set quantities to each variant.

**List of endpoints**

Authorization:
  - Registration

    POST /auth/register
  - Login
    
    POST /auth/login

Product:
  - Add new product
    
    POST /products/
  - Read all existing products
    
    GET /products/
  - Read an existing product by UUID
    
    GET /products/:productUUID
  - Update existing product
    
    PUT /products/:productUUID
  - Delete existing product
    
    DELETE /products/:productUUID

Variant:
  - Add new variant
    
    POST /products/variants/
  - Read all existing variants
    
    GET /products/variants/
  - Read an existing variant by UUID
    
    GET /products/variants/:variantUUID
  - Update existing variant
    
    PUT /products/variants/:variantUUID
  - Delete existing variant
    
    DELETE /products/variants/:variantUUID

**Current Issues**
LOGIN: still generates the exact same token every time a user logs in (no expire time).
GET ALL PRODUCT: no pagination info, search isn't case-insensitive, product info doesn't provide admin info
GET PRODUCT: no info on admin
UPDATE PRODUCT: response showed variant from another product instead after update
GET ALL VARIANT: no pagination info, search isn't case-insensitive, doesn't provide parent product info
GET VARIANT etc: doesn't provide parent product info
