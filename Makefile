title = "HelloOOOOoPatch"
author = "Me"
id = 1

post_request:
	curl -d '{"title":$(title), "author":$(author)}' -H "Content-Type: application/json" -X POST http://localhost:8080/books

patch_request:
	curl -d '{"title":$(title), "author":$(author)}' -H "Content-Type: application/json" -X PATCH http://localhost:8080/books/$(id)