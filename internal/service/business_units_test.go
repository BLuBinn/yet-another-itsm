package service

import (
	"context"
	"reflect"
	"testing"
)

func Test_businessUnitService_GetBusinessUnitByDomainName(t *testing.T) {
	type fields struct {
		repo *repository.Queries
	}
	type args struct {
		ctx        context.Context
		domainName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dtos.BusinessUnit
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &businessUnitService{
				repo: tt.fields.repo,
			}
			got, err := s.GetBusinessUnitByDomainName(tt.args.ctx, tt.args.domainName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBusinessUnitByDomainName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBusinessUnitByDomainName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
