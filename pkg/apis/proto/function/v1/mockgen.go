package v1

//go:generate mockgen -destination funcmock/funcmock.go -package funcmock github.com/numaproj/numaflow-go/pkg/apis/proto/function/v1 UserDefinedFunctionClient
//go:generate mockgen -destination funcmock/reducemock.go -package funcmock github.com/numaproj/numaflow-go/pkg/apis/proto/function/v1 UserDefinedFunction_ReduceFnClient
