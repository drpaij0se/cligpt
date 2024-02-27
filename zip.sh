GOOS=linux go build -o cligpt main.go
zip cligpt_linux.zip cligpt

GOOS=windows go build -o cligpt.exe main.go
zip cligpt-windows.zip cligpt.exe

rm cligpt.exe cligpt
