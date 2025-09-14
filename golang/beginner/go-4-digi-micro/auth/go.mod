module go-4-digi-micro/auth

go 1.24.0

toolchain go1.24.5

require (
	github.com/golang-jwt/jwt/v5 v5.3.0
	github.com/joho/godotenv v1.5.1
	golang.org/x/crypto v0.42.0
	google.golang.org/grpc v1.75.1
	google.golang.org/protobuf v1.36.9
	gorm.io/driver/postgres v1.6.0
	gorm.io/gorm v1.31.0
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.6 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	golang.org/x/net v0.44.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
	golang.org/x/text v0.29.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250908214217-97024824d090 // indirect
)

replace go-4-digi-micro/auth => ./
