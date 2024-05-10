package database

import (
	"reflect"
	"testing"

	proto "github.com/Soyaka/microlearn-user/api/proto/gen"
	"github.com/go-redis/redis"
)

func TestRedisClient_GetUserFromCache(t *testing.T) {
	type fields struct {
		Client *redis.Client
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &RedisClient{
				Client: tt.fields.Client,
			}
			got, err := c.GetUserFromCache(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisClient.GetUserFromCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RedisClient.GetUserFromCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
