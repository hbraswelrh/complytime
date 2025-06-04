// SPDX-License-Identifier: Apache-2.0

package plan

import (
	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
	"github.com/hashicorp/go-hclog"
)

// AssessmentScope sets up the yaml mapping type for writing to config file.
// Formats testdata as go struct.
type AssessmentScope struct {
	// FrameworkID is the identifier for the control set
	// in the Assessment Plan.
	FrameworkID string `yaml:"frameworkID"`
	// IncludeControls defines controls that are in scope
	// of an assessment.
	IncludeControls []string `yaml:"IncludeControls"`
}

// NewAssessmentScope create an AssessmentScope struct from a given framework id.
func NewAssessmentScope(frameworkID string) AssessmentScope {
	return AssessmentScope{
		FrameworkID: frameworkID,
	}
}

// ApplyScope alters the given OSCAL Assessment Plan based on the AssessmentScope.
func (a AssessmentScope) ApplyScope(assessmentPlan *oscalTypes.AssessmentPlan, logger hclog.Logger) {

	// This is a thin wrapper right now, but the goal to expand to different areas
	// of customization.
	a.applyControlScope(assessmentPlan, logger)
}

// applyControlScope alters the AssessedControls of the given OSCAL Assessment Plan by the AssessmentScope
// IncludeControls.
func (a AssessmentScope) applyControlScope(assessmentPlan *oscalTypes.AssessmentPlan, logger hclog.Logger) {
	// "Any control specified within exclude-controls must first be within a range of explicitly
	// included controls, via include-controls or include-all."
	includedControls := map[string]bool{}
	for _, id := range a.IncludeControls {
		includedControls[id] = true
	}
	logger.Debug("Found included controls", "count", len(includedControls))

	// FIXME: We should remove activities that have been filtered out (i.e. have no in scope controls)
	if assessmentPlan.LocalDefinitions != nil {
		if assessmentPlan.LocalDefinitions.Activities != nil {
			for activityI := range *assessmentPlan.LocalDefinitions.Activities {
				activity := &(*assessmentPlan.LocalDefinitions.Activities)[activityI]
				if activity.RelatedControls != nil && activity.RelatedControls.ControlSelections != nil {
					controlSelections := activity.RelatedControls.ControlSelections
					for controlSelectionI := range controlSelections {
						controlSelection := &controlSelections[controlSelectionI]
						filterControlSelection(controlSelection, includedControls)
					}
				}

				if activity.Steps != nil {
					for stepI := range *activity.Steps {
						step := &(*activity.Steps)[stepI]
						if step.ReviewedControls == nil {
							continue
						}
						if step.ReviewedControls.ControlSelections == nil {
							continue
						}
						controlSelections := step.ReviewedControls.ControlSelections
						for controlSelectionI := range controlSelections {
							controlSelection := &controlSelections[controlSelectionI]
							filterControlSelection(controlSelection, includedControls)
						}
					}
				}
			}
		}
	}
	// TODO: Reviewed Controls is construct identifying controls to be assessed
	// Handling included controls first and then removing excluded controls
	if assessmentPlan.ReviewedControls.ControlSelections != nil {
		for controlSelectionI := range assessmentPlan.ReviewedControls.ControlSelections {
			controlSelection := &assessmentPlan.ReviewedControls.ControlSelections[controlSelectionI]
			filterControlSelection(controlSelection, includedControls)
		}
	} else {
		if assessmentPlan.ReviewedControls.ControlSelections == nil {
			for controlExcludeI := range assessmentPlan.ReviewedControls.ControlSelections {
				controlExclusion := &assessmentPlan.ReviewedControls.ControlSelections[controlExcludeI]
				filterControlSelection(controlExclusion, includedControls)
			}
		}

	}
}

func filterControlSelection(controlSelection *oscalTypes.AssessedControls, includedControls map[string]bool) {
	// The new included controls should be the intersection of
	// the originally included controls and the newly included controls.
	// ExcludedControls are preserved.

	// includedControls specifies everything we allow - do not include all
	includedAll := controlSelection.IncludeAll != nil
	controlSelection.IncludeAll = nil

	originalIncludedControls := map[string]bool{}
	// TODO: added for testing
	//originalRemovedControls := map[string]bool{}
	if controlSelection.IncludeControls != nil {
		for _, controlId := range *controlSelection.IncludeControls {
			originalIncludedControls[controlId.ControlId] = true
		}
	}
	if controlSelection.IncludeControls == nil {
		for _, controlId := range *controlSelection.IncludeControls {
			delete(originalIncludedControls, controlId.ControlId)
		}
	}
	//else { // the controlIds would be filtered and removed from the slice
	//	controlSelection.IncludeControls = nil
	//	for _, controlId := range *controlSelection.IncludeControls {
	//		originalRemovedControls[controlId.ControlId] = true
	//		//delete(originalRemovedControls, controlId.ControlId)
	//	}
	//}
	var newIncludedControls []oscalTypes.AssessedControlsSelectControlById
	// TODO: wrap
	//var newRemovedControls []oscalTypes.AssessedControlsSelectControlById

	for controlId := range includedControls {
		//if includedAll || originalRemovedControls[controlId] || originalIncludedControls[controlId] {
		//	newIncludedControls = append(newIncludedControls, oscalTypes.AssessedControlsSelectControlById{
		//		ControlId: controlId,
		//	})
		//	newRemovedControls = append(newRemovedControls, oscalTypes.AssessedControlsSelectControlById{
		//		ControlId: controlId,
		//	})
		//	for controlId := range newIncludedControls {
		//		if newIncludedControls[controlId] == newRemovedControls[controlId] {
		//			delete(newIncludedControls, controlId)
		//		}
		//	}
		//}
		if includedAll || originalIncludedControls[controlId] {
			newIncludedControls = append(newIncludedControls, oscalTypes.AssessedControlsSelectControlById{
				ControlId: controlId,
			})
		}
	}
	controlSelection.IncludeControls = &newIncludedControls
}
