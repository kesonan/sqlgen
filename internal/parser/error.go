package parser

import "errors"

var (
	errorMissingTable           = errors.New("missing table")
	errorUnsupportedStmt        = errors.New("unsupported statement")
	errorMultipleTable          = errors.New("unsupported multiple tables")
	errorUnsupportedTableStyle  = errors.New("unsupported table style")
	errorUnsupportedNestedQuery = errors.New("unsupported nested query")
	errorUnsupportedUnionQuery  = errors.New("unsupported union query")
	errorUnsupportedSubQuery    = errors.New("unsupported sub-query query")
	errorInvalidExprNode        = errors.New("invalid expr node")
	errorInvalidExpr            = errors.New("only expect column expr")
	errorMissingHaving          = errors.New("missing having expr")
	errorUnsupportedLimitExpr   = errors.New("unsupported limit expr")
	errorParamMaker             = errors.New("marker expr")
	errorTableRefer             = errors.New("unsupported table refer")
)

func errorNearBy(err error, text string) error {
	return errors.New(err.Error() + " near by " + text)
}
