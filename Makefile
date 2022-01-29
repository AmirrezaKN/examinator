run:
	go run ./cmd

sqlboiler:
	sqlboiler psql -c ./configs/sqlboiler.toml

format:
	gofumpt -l -w . && gci -w -local github.com/daixiang0/gci .