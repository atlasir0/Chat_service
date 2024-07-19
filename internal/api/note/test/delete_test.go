// package test

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	"github.com/atlasir0/Chat_service/Auth_chat/internal/api/note"
// 	"github.com/atlasir0/Chat_service/Auth_chat/internal/service"
// 	serviceMocks "github.com/atlasir0/Chat_service/Auth_chat/internal/service/mocks"
// 	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
// 	"github.com/brianvoe/gofakeit/v6"
// 	"github.com/gojuno/minimock/v3"
// 	"github.com/stretchr/testify/require"
// 	"google.golang.org/protobuf/types/known/emptypb"
// )

// func TestDelete(t *testing.T) {
// 	t.Parallel()
// 	type noteServiceMockFunc func(mc *minimock.Controller) service.UserService

// 	type args struct {
// 		ctx context.Context
// 		req *desc.DeleteRequest
// 	}

// 	var (
// 		ctx = context.Background()
// 		mc  = minimock.NewController(t)

// 		id = gofakeit.Int64()

// 		serviceErr = fmt.Errorf("service error")

// 		req = &desc.DeleteRequest{
// 			Id: id,
// 		}
// 	)

// 	tests := []struct {
// 		name            string
// 		args            args
// 		want            *emptypb.Empty
// 		err             error
// 		noteServiceMock noteServiceMockFunc
// 	}{
// 		{
// 			name: "success case",
// 			args: args{
// 				ctx: ctx,
// 				req: req,
// 			},
// 			want: &emptypb.Empty{},
// 			err:  nil,
// 			noteServiceMock: func(mc *minimock.Controller) service.UserService {
// 				mock := serviceMocks.NewUserServiceMock(mc)
// 				mock.DeleteMock.Expect(ctx, id).Return(&emptypb.Empty{}, nil)
// 				return mock
// 			},
// 		},
// 		{
// 			name: "service error case",
// 			args: args{
// 				ctx: ctx,
// 				req: req,
// 			},
// 			want: nil,
// 			err:  serviceErr,
// 			noteServiceMock: func(mc *minimock.Controller) service.UserService {
// 				mock := serviceMocks.NewUserServiceMock(mc)
// 				mock.DeleteMock.Expect(ctx, id).Return(nil, serviceErr)
// 				return mock
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()

// 			noteServiceMock := tt.noteServiceMock(mc)
// 			api := note.NewImplementation(noteServiceMock)

// 			got, err := api.Delete(tt.args.ctx, tt.args.req)
// 			require.Equal(t, tt.err, err)
// 			require.Equal(t, tt.want, got)
// 		})
// 	}
// }
