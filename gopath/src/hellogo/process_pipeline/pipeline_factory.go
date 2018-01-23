package process_pipeline

import (
"errors"

pf "gitlab.mobvista.com/common/pipefiter_framework/pipefilter"
)

const STANDARD = "standard"
const REVERSE = "reverse"

var NotFoundPipelineError = errors.New("can't not find the pipeline definition")

func CreatePipeline(name string) (*pf.StraightPipeline, error) {
	if name == STANDARD {
		filters := []pf.Filter{&HttpExtractFilter{}, &P1_Filter{}, &P2_Filter{}, &P3_Filter{}}

		std := pf.StraightPipeline{
			Name:    "Standard",
			Filters: &filters,
		}
		return &std, nil
	} else if name == REVERSE {
		filters := []pf.Filter{&HttpExtractFilter{}, &P3_Filter{}, &P2_Filter{}, &P1_Filter{}}

		rev := pf.StraightPipeline{
			Name:    "Standard",
			Filters: &filters,
		}
		return &rev, nil
	}
	return &pf.StraightPipeline{}, NotFoundPipelineError
}
