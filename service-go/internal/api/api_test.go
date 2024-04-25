package api

import (
	"context"
	"interview-service/internal/api/interview"
	"interview-service/internal/domain/greeter"
	"reflect"
	"testing"
)

func Test_server_HelloWorld(t *testing.T) {
	tests := []struct {
		name    string
		request *interview.HelloWorldRequest
		want    *interview.HelloWorldResponse
		wantErr bool
	}{
		{
			name:    "basic test",
			request: &interview.HelloWorldRequest{Name: "unit-test"},
			want:    &interview.HelloWorldResponse{Greeting: greeter.Greet("unit-test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewServer()
			got, err := s.HelloWorld(context.Background(), tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("HelloWorld() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HelloWorld() got = %v, want %v", got, tt.want)
			}
		})
	}
}
