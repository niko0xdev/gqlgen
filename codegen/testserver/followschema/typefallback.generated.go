// Code generated by github.com/niko0xdev/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"

	"github.com/niko0xdev/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNFallbackToStringEncoding2githubᚗcomᚋniko0xdevᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐFallbackToStringEncoding(ctx context.Context, v interface{}) (FallbackToStringEncoding, error) {
	tmp, err := graphql.UnmarshalString(v)
	res := FallbackToStringEncoding(tmp)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNFallbackToStringEncoding2githubᚗcomᚋniko0xdevᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐFallbackToStringEncoding(ctx context.Context, sel ast.SelectionSet, v FallbackToStringEncoding) graphql.Marshaler {
	res := graphql.MarshalString(string(v))
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

// endregion ***************************** type.gotpl *****************************
