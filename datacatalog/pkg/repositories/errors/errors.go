// Generic errors used in the repos layer
package errors

import (
	"github.com/flyteorg/flyte/datacatalog/pkg/common"
	"github.com/flyteorg/flyte/datacatalog/pkg/errors"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
)

const (
	AlreadyExists = "entity already exists"
	notFound      = "missing entity of type %s with identifier %v"
	invalidJoin   = "cannot relate entity %s with entity %s"
	invalidEntity = "no such entity %s"
)

func GetMissingEntityError(entityType string, identifier proto.Message) error {
	return errors.NewDataCatalogErrorf(codes.NotFound, notFound, entityType, identifier)
}

func GetInvalidEntityRelationshipError(entityType common.Entity, otherEntityType common.Entity) error {
	return errors.NewDataCatalogErrorf(codes.InvalidArgument, invalidJoin, entityType, otherEntityType)
}

func GetInvalidEntityError(entityType common.Entity) error {
	return errors.NewDataCatalogErrorf(codes.InvalidArgument, invalidEntity, entityType)
}

func GetUnsupportedFilterExpressionErr(operator common.ComparisonOperator) error {
	return errors.NewDataCatalogErrorf(codes.InvalidArgument, "unsupported filter expression operator index: %v",
		operator)
}
