// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	v1 "github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
)

func AnyPtrToV1SourceRepository() *v1.SourceRepository {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*v1.SourceRepository))(nil)).Elem()))
	var nullValue *v1.SourceRepository
	return nullValue
}

func EqPtrToV1SourceRepository(value *v1.SourceRepository) *v1.SourceRepository {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *v1.SourceRepository
	return nullValue
}
