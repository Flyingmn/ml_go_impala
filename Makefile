
thrift:
	thrift -r -gen go:package_prefix=github.com/Flyingmn/ml_go_impala/services/ interfaces/ImpalaService.thrift
	rm -rf ./services
	mv gen-go services
