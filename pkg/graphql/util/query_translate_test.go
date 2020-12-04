package util

import (
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestAll(t1 *testing.T) {
	type fields struct {
		model      interface{}
		distinctOn interface{}
		limit      *int
		offset     *int
		orderBy    interface{}
		where      interface{}
		tx         *gorm.DB
	}
	type args struct {
		distinctOn interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantR  *QueryTranslator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &QueryTranslator{
				model:      tt.fields.model,
				distinctOn: tt.fields.distinctOn,
				limit:      tt.fields.limit,
				offset:     tt.fields.offset,
				orderBy:    tt.fields.orderBy,
				where:      tt.fields.where,
				tx:         tt.fields.tx,
			}
			if gotR := t.DistinctOn(tt.args.distinctOn); !reflect.DeepEqual(gotR, tt.wantR) {
				t1.Errorf("DistinctOn() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
