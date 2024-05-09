package jwt

import (
	"reflect"
	"testing"
)

func TestJWT_CheckToken(t *testing.T) {
	type fields struct {
		SigningKey []byte
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JWT{
				SigningKey: tt.fields.SigningKey,
			}
			got, err := j.CheckToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWT_CreateToken(t *testing.T) {
	type fields struct {
		SigningKey []byte
	}
	type args struct {
		claims CustomClaims
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JWT{
				SigningKey: tt.fields.SigningKey,
			}
			got, err := j.CreateToken(tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWT_ParseToken(t *testing.T) {
	type fields struct {
		SigningKey []byte
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CustomClaims
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JWT{
				SigningKey: tt.fields.SigningKey,
			}
			got, err := j.ParseToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewJWT(t *testing.T) {
	type args struct {
		SigningKey []byte
	}
	tests := []struct {
		name string
		args args
		want *JWT
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJWT(tt.args.SigningKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}
