module github.com/DanyHenriquez/Guardian

go 1.14

require (
	Guardian/database v0.0.0
	Guardian/handlers v0.0.0
	Guardian/models v0.0.0
	github.com/alexedwards/argon2id v0.0.0-20190612080829-01a59b2b8802 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/jinzhu/gorm v1.9.12 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/pquerna/otp v1.2.0 // indirect
	github.com/valyala/fasttemplate v1.1.0 // indirect
	golang.org/x/crypto v0.0.0-20200317142112-1b76d66859c6 // indirect
	golang.org/x/net v0.0.0-20200319234117-63522dbf7eec // indirect
	golang.org/x/sys v0.0.0-20200317113312-5766fd39f98d // indirect
	golang.org/x/text v0.3.2 // indirect
)

replace Guardian/database => ./database

replace Guardian/handlers => ./handlers

replace Guardian/models => ./models
