# Golang-Ecommerce-Framework

Golang - (GoFiber,Gorm) for the back end, HTMX CSS and JS for the front end.
Has JWT Authentication With Bcrypt Encryption. 

#Contains:
- Pages: Login & Signup & Kinda Like A Home Page , A Navigation Bar And A Random Transition i was experimenting on once u click Shop All.(Animation isnt responsive but it works fine on my laptop of 1920x1080)
- Apis: Add To Cart(takes in a product name searches it in the products database, if exists adds it to the users cart ), Place Order(not complete but it works good enough. Location and User's contacts are set to empty . Need To Implemenet), Login , SignUp, GetUserInfo (with jwt tokens), GetAllProducts ( returns all products from the sqlite database), GetSepecificProduct(takes in a product name and returns the product)
