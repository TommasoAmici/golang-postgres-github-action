package mocks

import (
	"reflect"
	"testing"

	"github.com/TommasoAmici/golang-postgres-github-action/pkg/models"
)

func TestUserModel_Get(t *testing.T) {
	t.Parallel()
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		m       *UserModel
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name:    "valid user",
			args:    args{id: 1},
			wantErr: false,
			want:    mockUser,
		},
		{
			name:    "user not in db",
			args:    args{id: 3},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UserModel{}
			got, err := m.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserModel.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserModel.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
