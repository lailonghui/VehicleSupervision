package util

import "testing"

func TestRewriteGraphqls(t *testing.T) {
	RewriteGraphqls("D:/project/VehicleSupervision2/internal/modules/blacklist/graph/graphqls/blanklist_operation_record.graphqls")
	RewriteGraphqls("D:/project/VehicleSupervision2/internal/modules/blacklist/graph/graphqls/common.graphqls")
}
