package bizresponse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCodeError_EqualsTo(t *testing.T) {
	ast := assert.New(t)

	type testCase struct {
		codeErr *CodeError
		err     error
		Expect  bool
	}

	cases := []testCase{
		{
			// 不同的CodeError
			codeErr: ErrInvalidArgs,
			err:     ErrClientCancel,
			Expect:  false,
		},
		{
			// 同样的CodeError
			codeErr: ErrInvalidArgs,
			err:     ErrInvalidArgs,
			Expect:  true,
		},
		{
			// 非ErrInternalFailed与error比较
			codeErr: ErrUnknown,
			err:     fmt.Errorf("test error"),
			Expect:  false,
		},
		{
			// FromError会把普通error转成ErrInternalFailed
			codeErr: ErrInternalFailed,
			err:     fmt.Errorf("test error"),
			Expect:  true,
		},
		{
			// GRPC的err
			codeErr: ErrInvalidArgs,
			err:     status.Error(codes.Code(ErrInvalidArgs.Code), ErrInvalidArgs.Msg),
			Expect:  true,
		},
	}

	for i, c := range cases {
		fmt.Println(fmt.Sprintf("Test EqualsTo %v", i))
		ast.EqualValues(c.Expect, c.codeErr.EqualsTo(c.err))
	}
}
