go run main.go rnGen template --template=Ts
go run main.go rnGen template --template=Js
go build -o file.exe file.go

./main rnGen --Template=Ts --C=DatePicker --SubExtension="Screen"
