install:
	go get -u github.com/swaggo/swag
generate-swagger:
	swag init
generate-di:
	# TODO if the files doesnt exists ignore the exit(1)
	rm wire_gen.go 2>./error
	go generate

mock:
	mockgen -source=internal/adapters/repository/producao.go -package=mock_repo -destination=test/mock/repository/producao.go
