package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/repository"
	repoMocks "github.com/atlasir0/Chat_service/Auth_chat/internal/repository/mocks"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/service/note"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestUpdate(t *testing.T) {
	t.Parallel()
	type noteRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx  context.Context
		info *model.User
		req  *model.User
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, false, 10)
		role     = gofakeit.Int32()

		repoErr = fmt.Errorf("repo error")

		info = &model.User{
			ID:       id,
			Name:     name,
			Email:    email,
			Password: password,
			Role:     int(role),
		}
		req = &model.User{
			ID:       id,
			Name:     name,
			Email:    email,
			Password: password,
			Role:     int(role),
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               *emptypb.Empty
		err                error
		noteRepositoryMock noteRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:  ctx,
				info: info,
				req:  req,
			},
			want: &emptypb.Empty{},
			err:  nil,
			noteRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, info).Return(&emptypb.Empty{}, nil) // Исправлено
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx:  ctx,
				info: info,
				req:  req,
			},
			want: nil,
			err:  repoErr,
			noteRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, info).Return(nil, repoErr) // Исправлено
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			noteRepoMock := tt.noteRepositoryMock(mc)
			service := note.NewMockService(noteRepoMock)

			got, err := service.Update(tt.args.ctx, tt.args.info, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, got)
		})
	}
}
