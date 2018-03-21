package robotserver

import "github.com/cyborg-client/client/analysis"

//type splitterRequest chan analysis.Timestampdata

type splitterRequest struct {
	ID int
	DataCh chan analysis.Timestampdata
}