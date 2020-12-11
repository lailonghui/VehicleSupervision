package util

import "testing"

func TestRewriteGraphqls(t *testing.T) {
	RewriteGraphqls("D:/project/VehicleSupervision2/internal/modules/device/graph/graphqls/sim_card.graphqls")
	RewriteGraphqls("D:/project/VehicleSupervision2/internal/modules/device/graph/graphqls/sim_card_flow.graphqls")
	RewriteGraphqls("D:/project/VehicleSupervision2/internal/modules/device/graph/graphqls/terminal.graphqls")
	RewriteGraphqls("D:/project/VehicleSupervision2/internal/modules/device/graph/graphqls/terminal_factory.graphqls")
	RewriteGraphqls("D:/project/VehicleSupervision2/internal/modules/device/graph/graphqls/terminal_modal.graphqls")
	RewriteGraphqls("D:/project/VehicleSupervision2/internal/modules/device/graph/graphqls/terminal_type.graphqls")

	RewriteGraphqls("D:/project/VehicleSupervision2/internal/modules/device/graph/graphqls/common.graphqls")
}
