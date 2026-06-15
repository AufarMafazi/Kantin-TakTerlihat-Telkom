@echo off
:: Mengatur direktori kerja ke lokasi file .bat ini berada
cd /d %~dp0

echo Memulai Microservices...

:: 1. Jalankan NATS Server
start "NATS Server" cmd /k "cd nats-server && nats-server.exe -js || pause"
timeout /t 2

:: 2. Jalankan Menu Service
start "Menu Service" cmd /k "cd menu-service && go run main.go || pause"

:: 3. Jalankan Order Service
start "Order Service" cmd /k "cd order-service && go run main.go || pause"

:: 4. Jalankan Notification Service
start "Notification Service" cmd /k "cd notification-service && go run main.go || pause"

:: 5. Jalankan API Gateway
start "API Gateway" cmd /k "cd api-gateway && go run main.go || pause"

start "API Gateway" cmd /k "cd auth-service && go run . || pause"

start "Front End" cmd /k "cd Front-End && npm run dev || pause"

echo Semua service telah dipanggil. Periksa jendela baru untuk melihat error.
pause