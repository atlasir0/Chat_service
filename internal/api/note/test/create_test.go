package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/api/note"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/service"
	serviceMocks "github.com/atlasir0/Chat_service/Auth_chat/internal/service/mocks"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type noteServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, false, 10)
		role     = gofakeit.Int32()

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			User: &desc.User{
				Name:     name,
				Email:    email,
				Password: password,
				Role:     desc.UserRole(role),
			},
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		noteServiceMock noteServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			noteServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctx, &model.User{
					Name:     name,
					Email:    email,
					Password: password,
					Role:     int(role),
				}).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			noteServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctx, &model.User{
					Name:     name,
					Email:    email,
					Password: password,
					Role:     int(role),
				}).Return(int64(0), serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			noteServiceMock := tt.noteServiceMock(mc)
			api := note.NewImplementation(noteServiceMock)

			newID, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
