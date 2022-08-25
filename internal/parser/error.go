package parser

import "errors"

var (
	errorMissingTable                 = errors.New("missing table")
	errorUnsupportedStmt              = errors.New("unsupported statement")
	errorMultipleTable                = errors.New("unsupported multiple tables")
	errorUnsupportedTableStyle        = errors.New("unsupported table style")
	errorUnsupportedNestedQuery       = errors.New("unsupported nested query")
	errorUnsupportedUnionQuery        = errors.New("unsupported union query")
	errorUnsupportedSubQuery          = errors.New("unsupported sub-query query")
	errorInvalidExprNode              = errors.New("invalid expr node")
	errorMissingHaving                = errors.New("missing having expr")
	errorUnsupportedLimitExpr         = errors.New("unsupported limit expr")
	errorParamMaker                   = errors.New("marker expr")
	errorUnsupportedNestedTransaction = errors.New("unsupported nested transaction")
	errorMissingCommit                = errors.New("missing commit statement")
	errorMissingTransaction           = errors.New("missing transaction statement")
	errorMissingFunction              = errors.New("missing function name")
)

func errorNearBy(err error, text string) error {
	return errors.New(err.Error() + " near by " + text)
}
