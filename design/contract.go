// Copyright (c) 2011-2017, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package design

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Contract struct {
	DomainResource    `bson:",inline"`
	Identifier        *Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string                                `bson:"status,omitempty" json:"status,omitempty"`
	Issued            *FHIRDateTime                         `bson:"issued,omitempty" json:"issued,omitempty"`
	Applies           *Period                               `bson:"applies,omitempty" json:"applies,omitempty"`
	Subject           []Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Topic             []Reference                           `bson:"topic,omitempty" json:"topic,omitempty"`
	Authority         []Reference                           `bson:"authority,omitempty" json:"authority,omitempty"`
	Domain            []Reference                           `bson:"domain,omitempty" json:"domain,omitempty"`
	Type              *CodeableConcept                      `bson:"type,omitempty" json:"type,omitempty"`
	SubType           []CodeableConcept                     `bson:"subType,omitempty" json:"subType,omitempty"`
	Action            []CodeableConcept                     `bson:"action,omitempty" json:"action,omitempty"`
	ActionReason      []CodeableConcept                     `bson:"actionReason,omitempty" json:"actionReason,omitempty"`
	DecisionType      *CodeableConcept                      `bson:"decisionType,omitempty" json:"decisionType,omitempty"`
	ContentDerivative *CodeableConcept                      `bson:"contentDerivative,omitempty" json:"contentDerivative,omitempty"`
	SecurityLabel     []Coding                              `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Agent             []ContractAgentComponent              `bson:"agent,omitempty" json:"agent,omitempty"`
	Signer            []ContractSignatoryComponent          `bson:"signer,omitempty" json:"signer,omitempty"`
	ValuedItem        []ContractValuedItemComponent         `bson:"valuedItem,omitempty" json:"valuedItem,omitempty"`
	Term              []ContractTermComponent               `bson:"term,omitempty" json:"term,omitempty"`
	BindingAttachment *Attachment                           `bson:"bindingAttachment,omitempty" json:"bindingAttachment,omitempty"`
	BindingReference  *Reference                            `bson:"bindingReference,omitempty" json:"bindingReference,omitempty"`
	Friendly          []ContractFriendlyLanguageComponent   `bson:"friendly,omitempty" json:"friendly,omitempty"`
	Legal             []ContractLegalLanguageComponent      `bson:"legal,omitempty" json:"legal,omitempty"`
	Rule              []ContractComputableLanguageComponent `bson:"rule,omitempty" json:"rule,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Contract) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Contract"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Contract), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Contract) GetBSON() (interface{}, error) {
	x.ResourceType = "Contract"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "contract" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type contract Contract

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Contract) UnmarshalJSON(data []byte) (err error) {
	x2 := contract{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Contract(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Contract) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Contract"
	} else if x.ResourceType != "Contract" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Contract, instead received %s", x.ResourceType))
	}
	return nil
}

type ContractAgentComponent struct {
	BackboneElement `bson:",inline"`
	Actor           *Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
	Role            []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}

type ContractSignatoryComponent struct {
	BackboneElement `bson:",inline"`
	Type            *Coding     `bson:"type,omitempty" json:"type,omitempty"`
	Party           *Reference  `bson:"party,omitempty" json:"party,omitempty"`
	Signature       []Signature `bson:"signature,omitempty" json:"signature,omitempty"`
}

type ContractValuedItemComponent struct {
	BackboneElement       `bson:",inline"`
	EntityCodeableConcept *CodeableConcept `bson:"entityCodeableConcept,omitempty" json:"entityCodeableConcept,omitempty"`
	EntityReference       *Reference       `bson:"entityReference,omitempty" json:"entityReference,omitempty"`
	Identifier            *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	EffectiveTime         *FHIRDateTime    `bson:"effectiveTime,omitempty" json:"effectiveTime,omitempty"`
	Quantity              *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice             *Quantity        `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                *float64         `bson:"factor,omitempty" json:"factor,omitempty"`
	Points                *float64         `bson:"points,omitempty" json:"points,omitempty"`
	Net                   *Quantity        `bson:"net,omitempty" json:"net,omitempty"`
}

type ContractTermComponent struct {
	BackboneElement `bson:",inline"`
	Identifier      *Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Issued          *FHIRDateTime                     `bson:"issued,omitempty" json:"issued,omitempty"`
	Applies         *Period                           `bson:"applies,omitempty" json:"applies,omitempty"`
	Type            *CodeableConcept                  `bson:"type,omitempty" json:"type,omitempty"`
	SubType         *CodeableConcept                  `bson:"subType,omitempty" json:"subType,omitempty"`
	Topic           []Reference                       `bson:"topic,omitempty" json:"topic,omitempty"`
	Action          []CodeableConcept                 `bson:"action,omitempty" json:"action,omitempty"`
	ActionReason    []CodeableConcept                 `bson:"actionReason,omitempty" json:"actionReason,omitempty"`
	SecurityLabel   []Coding                          `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Agent           []ContractTermAgentComponent      `bson:"agent,omitempty" json:"agent,omitempty"`
	Text            string                            `bson:"text,omitempty" json:"text,omitempty"`
	ValuedItem      []ContractTermValuedItemComponent `bson:"valuedItem,omitempty" json:"valuedItem,omitempty"`
	Group           []ContractTermComponent           `bson:"group,omitempty" json:"group,omitempty"`
}

type ContractTermAgentComponent struct {
	BackboneElement `bson:",inline"`
	Actor           *Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
	Role            []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}

type ContractTermValuedItemComponent struct {
	BackboneElement       `bson:",inline"`
	EntityCodeableConcept *CodeableConcept `bson:"entityCodeableConcept,omitempty" json:"entityCodeableConcept,omitempty"`
	EntityReference       *Reference       `bson:"entityReference,omitempty" json:"entityReference,omitempty"`
	Identifier            *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	EffectiveTime         *FHIRDateTime    `bson:"effectiveTime,omitempty" json:"effectiveTime,omitempty"`
	Quantity              *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice             *Quantity        `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                *float64         `bson:"factor,omitempty" json:"factor,omitempty"`
	Points                *float64         `bson:"points,omitempty" json:"points,omitempty"`
	Net                   *Quantity        `bson:"net,omitempty" json:"net,omitempty"`
}

type ContractFriendlyLanguageComponent struct {
	BackboneElement   `bson:",inline"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

type ContractLegalLanguageComponent struct {
	BackboneElement   `bson:",inline"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

type ContractComputableLanguageComponent struct {
	BackboneElement   `bson:",inline"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

type ContractPlus struct {
	Contract                     `bson:",inline"`
	ContractPlusRelatedResources `bson:",inline"`
}

type ContractPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByAgent                  *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAgent,omitempty"`
	IncludedGroupResourcesReferencedByAgent                         *[]Group                 `bson:"_includedGroupResourcesReferencedByAgent,omitempty"`
	IncludedOrganizationResourcesReferencedByAgent                  *[]Organization          `bson:"_includedOrganizationResourcesReferencedByAgent,omitempty"`
	IncludedDeviceResourcesReferencedByAgent                        *[]Device                `bson:"_includedDeviceResourcesReferencedByAgent,omitempty"`
	IncludedPatientResourcesReferencedByAgent                       *[]Patient               `bson:"_includedPatientResourcesReferencedByAgent,omitempty"`
	IncludedSubstanceResourcesReferencedByAgent                     *[]Substance             `bson:"_includedSubstanceResourcesReferencedByAgent,omitempty"`
	IncludedContractResourcesReferencedByAgent                      *[]Contract              `bson:"_includedContractResourcesReferencedByAgent,omitempty"`
	IncludedRelatedPersonResourcesReferencedByAgent                 *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByAgent,omitempty"`
	IncludedLocationResourcesReferencedByAgent                      *[]Location              `bson:"_includedLocationResourcesReferencedByAgent,omitempty"`
	IncludedPatientResourcesReferencedByPatient                     *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedOrganizationResourcesReferencedByAuthority              *[]Organization          `bson:"_includedOrganizationResourcesReferencedByAuthority,omitempty"`
	IncludedLocationResourcesReferencedByDomain                     *[]Location              `bson:"_includedLocationResourcesReferencedByDomain,omitempty"`
	IncludedPractitionerResourcesReferencedBySigner                 *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedBySigner,omitempty"`
	IncludedOrganizationResourcesReferencedBySigner                 *[]Organization          `bson:"_includedOrganizationResourcesReferencedBySigner,omitempty"`
	IncludedPatientResourcesReferencedBySigner                      *[]Patient               `bson:"_includedPatientResourcesReferencedBySigner,omitempty"`
	IncludedRelatedPersonResourcesReferencedBySigner                *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedBySigner,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref       *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath1                 *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath1,omitempty"`
	RevIncludedConsentResourcesReferencingDataPath2                 *[]Consent               `bson:"_revIncludedConsentResourcesReferencingDataPath2,omitempty"`
	RevIncludedConsentResourcesReferencingSource                    *[]Consent               `bson:"_revIncludedConsentResourcesReferencingSource,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                 *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom               *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor               *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1            *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2            *[]Measure               `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref      *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingAgent                    *[]Contract              `bson:"_revIncludedContractResourcesReferencingAgent,omitempty"`
	RevIncludedContractResourcesReferencingSubject                  *[]Contract              `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingTermtopic                *[]Contract              `bson:"_revIncludedContractResourcesReferencingTermtopic,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest             *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse            *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource      *[]ImplementationGuide   `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingSuccessor       *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDerivedfrom     *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingPredecessor     *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingComposedof      *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedServiceDefinitionResourcesReferencingDependson       *[]ServiceDefinition     `bson:"_revIncludedServiceDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof              *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon             *[]Communication         `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor      *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom    *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor    *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof     *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 *[]ActivityDefinition    `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedRequestGroupResourcesReferencingDefinition           *[]RequestGroup          `bson:"_revIncludedRequestGroupResourcesReferencingDefinition,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon             *[]DeviceRequest         `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest        *[]DeviceRequest         `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus               *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntityref              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingEntityref,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                 *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                      *[]Task                  `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                        *[]Task                  `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                      *[]Task                  `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                         *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingReplaces         *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedProcedureRequestResourcesReferencingBasedon          *[]ProcedureRequest      `bson:"_revIncludedProcedureRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                 *[]Library               `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom               *[]Library               `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor               *[]Library               `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                *[]Library               `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                 *[]Library               `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon      *[]CommunicationRequest  `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                     *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                 *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail          *[]Condition             `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject               *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                 *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated          *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject     *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest           *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor          *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor        *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof         *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1     *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2     *[]PlanDefinition        `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (c *ContractPlusRelatedResources) GetIncludedPractitionerResourceReferencedByAgent() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByAgent == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByAgent))
	} else if len(*c.IncludedPractitionerResourcesReferencedByAgent) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByAgent)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedGroupResourceReferencedByAgent() (group *Group, err error) {
	if c.IncludedGroupResourcesReferencedByAgent == nil {
		err = errors.New("Included groups not requested")
	} else if len(*c.IncludedGroupResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*c.IncludedGroupResourcesReferencedByAgent))
	} else if len(*c.IncludedGroupResourcesReferencedByAgent) == 1 {
		group = &(*c.IncludedGroupResourcesReferencedByAgent)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedOrganizationResourceReferencedByAgent() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByAgent == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByAgent))
	} else if len(*c.IncludedOrganizationResourcesReferencedByAgent) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByAgent)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedDeviceResourceReferencedByAgent() (device *Device, err error) {
	if c.IncludedDeviceResourcesReferencedByAgent == nil {
		err = errors.New("Included devices not requested")
	} else if len(*c.IncludedDeviceResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*c.IncludedDeviceResourcesReferencedByAgent))
	} else if len(*c.IncludedDeviceResourcesReferencedByAgent) == 1 {
		device = &(*c.IncludedDeviceResourcesReferencedByAgent)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedPatientResourceReferencedByAgent() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByAgent == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByAgent))
	} else if len(*c.IncludedPatientResourcesReferencedByAgent) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByAgent)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedSubstanceResourceReferencedByAgent() (substance *Substance, err error) {
	if c.IncludedSubstanceResourcesReferencedByAgent == nil {
		err = errors.New("Included substances not requested")
	} else if len(*c.IncludedSubstanceResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*c.IncludedSubstanceResourcesReferencedByAgent))
	} else if len(*c.IncludedSubstanceResourcesReferencedByAgent) == 1 {
		substance = &(*c.IncludedSubstanceResourcesReferencedByAgent)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedContractResourceReferencedByAgent() (contract *Contract, err error) {
	if c.IncludedContractResourcesReferencedByAgent == nil {
		err = errors.New("Included contracts not requested")
	} else if len(*c.IncludedContractResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 contract, but found %d", len(*c.IncludedContractResourcesReferencedByAgent))
	} else if len(*c.IncludedContractResourcesReferencedByAgent) == 1 {
		contract = &(*c.IncludedContractResourcesReferencedByAgent)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByAgent() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByAgent == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedRelatedPersonResourcesReferencedByAgent))
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByAgent) == 1 {
		relatedPerson = &(*c.IncludedRelatedPersonResourcesReferencedByAgent)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedLocationResourceReferencedByAgent() (location *Location, err error) {
	if c.IncludedLocationResourcesReferencedByAgent == nil {
		err = errors.New("Included locations not requested")
	} else if len(*c.IncludedLocationResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*c.IncludedLocationResourcesReferencedByAgent))
	} else if len(*c.IncludedLocationResourcesReferencedByAgent) == 1 {
		location = &(*c.IncludedLocationResourcesReferencedByAgent)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedPatientResourcesReferencedByPatient() (patients []Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedPatientResourcesReferencedByPatient
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByAuthority() (organizations []Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByAuthority == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *c.IncludedOrganizationResourcesReferencedByAuthority
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedLocationResourcesReferencedByDomain() (locations []Location, err error) {
	if c.IncludedLocationResourcesReferencedByDomain == nil {
		err = errors.New("Included locations not requested")
	} else {
		locations = *c.IncludedLocationResourcesReferencedByDomain
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedPractitionerResourceReferencedBySigner() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedBySigner == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedBySigner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedBySigner))
	} else if len(*c.IncludedPractitionerResourcesReferencedBySigner) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedBySigner)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedOrganizationResourceReferencedBySigner() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedBySigner == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedBySigner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedBySigner))
	} else if len(*c.IncludedOrganizationResourcesReferencedBySigner) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedBySigner)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedPatientResourceReferencedBySigner() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedBySigner == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedBySigner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedBySigner))
	} else if len(*c.IncludedPatientResourcesReferencedBySigner) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedBySigner)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedBySigner() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedBySigner == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedRelatedPersonResourcesReferencedBySigner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedRelatedPersonResourcesReferencedBySigner))
	} else if len(*c.IncludedRelatedPersonResourcesReferencedBySigner) == 1 {
		relatedPerson = &(*c.IncludedRelatedPersonResourcesReferencedBySigner)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath1() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingDataPath1 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingDataPath1
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedConsentResourcesReferencingDataPath2() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingDataPath2 == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingDataPath2
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedConsentResourcesReferencingSource() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingSource == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingSource
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedContractResourcesReferencingAgent() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingAgent
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedContractResourcesReferencingTermtopic() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingTermtopic == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingTermtopic
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if c.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *c.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingSuccessor() (serviceDefinitions []ServiceDefinition, err error) {
	if c.RevIncludedServiceDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *c.RevIncludedServiceDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDerivedfrom() (serviceDefinitions []ServiceDefinition, err error) {
	if c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingPredecessor() (serviceDefinitions []ServiceDefinition, err error) {
	if c.RevIncludedServiceDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *c.RevIncludedServiceDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingComposedof() (serviceDefinitions []ServiceDefinition, err error) {
	if c.RevIncludedServiceDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *c.RevIncludedServiceDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedServiceDefinitionResourcesReferencingDependson() (serviceDefinitions []ServiceDefinition, err error) {
	if c.RevIncludedServiceDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded serviceDefinitions not requested")
	} else {
		serviceDefinitions = *c.RevIncludedServiceDefinitionResourcesReferencingDependson
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingDefinition() (requestGroups []RequestGroup, err error) {
	if c.RevIncludedRequestGroupResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *c.RevIncludedRequestGroupResourcesReferencingDefinition
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *c.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntityref() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingEntityref == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingEntityref
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingReplaces() (procedureRequests []ProcedureRequest, err error) {
	if c.RevIncludedProcedureRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *c.RevIncludedProcedureRequestResourcesReferencingReplaces
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedProcedureRequestResourcesReferencingBasedon() (procedureRequests []ProcedureRequest, err error) {
	if c.RevIncludedProcedureRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedureRequests not requested")
	} else {
		procedureRequests = *c.RevIncludedProcedureRequestResourcesReferencingBasedon
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *c.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if c.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *c.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPractitionerResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByAgent {
			rsc := (*c.IncludedPractitionerResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGroupResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedGroupResourcesReferencedByAgent {
			rsc := (*c.IncludedGroupResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByAgent {
			rsc := (*c.IncludedOrganizationResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByAgent {
			rsc := (*c.IncludedDeviceResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByAgent {
			rsc := (*c.IncludedPatientResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedSubstanceResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedSubstanceResourcesReferencedByAgent {
			rsc := (*c.IncludedSubstanceResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedContractResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedContractResourcesReferencedByAgent {
			rsc := (*c.IncludedContractResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByAgent {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedLocationResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedLocationResourcesReferencedByAgent {
			rsc := (*c.IncludedLocationResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByAuthority != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByAuthority {
			rsc := (*c.IncludedOrganizationResourcesReferencedByAuthority)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedLocationResourcesReferencedByDomain != nil {
		for idx := range *c.IncludedLocationResourcesReferencedByDomain {
			rsc := (*c.IncludedLocationResourcesReferencedByDomain)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedBySigner {
			rsc := (*c.IncludedPractitionerResourcesReferencedBySigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedBySigner {
			rsc := (*c.IncludedOrganizationResourcesReferencedBySigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySigner {
			rsc := (*c.IncludedPatientResourcesReferencedBySigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedBySigner {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedBySigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ContractPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*c.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*c.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingSource != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingSource {
			rsc := (*c.RevIncludedConsentResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*c.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingAgent {
			rsc := (*c.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*c.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*c.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*c.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*c.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *c.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*c.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ContractPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPractitionerResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByAgent {
			rsc := (*c.IncludedPractitionerResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGroupResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedGroupResourcesReferencedByAgent {
			rsc := (*c.IncludedGroupResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByAgent {
			rsc := (*c.IncludedOrganizationResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByAgent {
			rsc := (*c.IncludedDeviceResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByAgent {
			rsc := (*c.IncludedPatientResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedSubstanceResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedSubstanceResourcesReferencedByAgent {
			rsc := (*c.IncludedSubstanceResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedContractResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedContractResourcesReferencedByAgent {
			rsc := (*c.IncludedContractResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByAgent {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedLocationResourcesReferencedByAgent != nil {
		for idx := range *c.IncludedLocationResourcesReferencedByAgent {
			rsc := (*c.IncludedLocationResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByAuthority != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByAuthority {
			rsc := (*c.IncludedOrganizationResourcesReferencedByAuthority)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedLocationResourcesReferencedByDomain != nil {
		for idx := range *c.IncludedLocationResourcesReferencedByDomain {
			rsc := (*c.IncludedLocationResourcesReferencedByDomain)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedBySigner {
			rsc := (*c.IncludedPractitionerResourcesReferencedBySigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedBySigner {
			rsc := (*c.IncludedOrganizationResourcesReferencedBySigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySigner {
			rsc := (*c.IncludedPatientResourcesReferencedBySigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedBySigner != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedBySigner {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedBySigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingContentref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingDataPath1 != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingDataPath1 {
			rsc := (*c.RevIncludedConsentResourcesReferencingDataPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingDataPath2 != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingDataPath2 {
			rsc := (*c.RevIncludedConsentResourcesReferencingDataPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingSource != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingSource {
			rsc := (*c.RevIncludedConsentResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*c.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingAgent != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingAgent {
			rsc := (*c.RevIncludedContractResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingTermtopic != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingTermtopic {
			rsc := (*c.RevIncludedContractResourcesReferencingTermtopic)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceDefinitionResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedServiceDefinitionResourcesReferencingDependson {
			rsc := (*c.RevIncludedServiceDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedRequestGroupResourcesReferencingDefinition != nil {
		for idx := range *c.RevIncludedRequestGroupResourcesReferencingDefinition {
			rsc := (*c.RevIncludedRequestGroupResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingEntityref != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingEntityref {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingEntityref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*c.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*c.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureRequestResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedProcedureRequestResourcesReferencingReplaces {
			rsc := (*c.RevIncludedProcedureRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedProcedureRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedProcedureRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*c.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *c.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*c.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedProcessResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
