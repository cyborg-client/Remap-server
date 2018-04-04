package robotserver

import (
	"github.com/cyborg-client/client/analysis"
	"github.com/satori/go.uuid"
)

//type splitterRequest chan analysis.Timestampdata

type splitterRequest struct {
	ID     uuid.UUID
	DataCh chan analysis.Timestampdata
}
