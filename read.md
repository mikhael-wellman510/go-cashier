Refrensi architecture 
https://github.com/shayja/go-template-api/blob/main/cmd/app/app.go

1. Buat Folder project

2. go mod init [name] 

3. install dotenv
go get github.com/joho/godotenv

>>>>> Jalankan server <<<<<<
go run cmd/main.go

4. install gorm dan mysql
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

5. install UUID
go get github.com/google/uuid

6. install untuk decimal
go get github.com/shopspring/decimal

7. install JWT
go get -u github.com/golang-jwt/jwt/v5

task :


// Todo -> 
1. Buat Table Transaction dan Transaction Detail
2. Buat Table Payment

// next feature
3. kalau sudah bisa , buat role , trus login with google
4. kirim ke email untuk validasi akun 



