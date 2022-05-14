go tool compile main.go

go tool link main.o

go tool link -o main main.o

go tool objdump -s main.main main