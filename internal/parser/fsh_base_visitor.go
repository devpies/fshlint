// Code generated from FSH.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // FSH

import "github.com/antlr4-go/antlr/v4"

type BaseFSHVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFSHVisitor) VisitDoc(ctx *DocContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitEntity(ctx *EntityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitProfile(ctx *ProfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitExtension(ctx *ExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitLogical(ctx *LogicalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitResource(ctx *ResourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitSdMetadata(ctx *SdMetadataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitSdRule(ctx *SdRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitLrRule(ctx *LrRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitInstance(ctx *InstanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitInstanceMetadata(ctx *InstanceMetadataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitInstanceRule(ctx *InstanceRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitInvariant(ctx *InvariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitInvariantMetadata(ctx *InvariantMetadataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitInvariantRule(ctx *InvariantRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitValueSet(ctx *ValueSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitVsMetadata(ctx *VsMetadataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitVsRule(ctx *VsRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitCodeSystem(ctx *CodeSystemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitCsMetadata(ctx *CsMetadataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitCsRule(ctx *CsRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitRuleSet(ctx *RuleSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitRuleSetRule(ctx *RuleSetRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitParamRuleSet(ctx *ParamRuleSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitParamRuleSetRef(ctx *ParamRuleSetRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitLastParameter(ctx *LastParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitParamRuleSetContent(ctx *ParamRuleSetContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitMapping(ctx *MappingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitMappingMetadata(ctx *MappingMetadataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitMappingEntityRule(ctx *MappingEntityRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitParent(ctx *ParentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitTitle(ctx *TitleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitDescription(ctx *DescriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitXpath(ctx *XpathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitSeverity(ctx *SeverityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitInstanceOf(ctx *InstanceOfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitUsage(ctx *UsageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitSource(ctx *SourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitTarget(ctx *TargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitContext(ctx *ContextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitContextItem(ctx *ContextItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitLastContextItem(ctx *LastContextItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitCharacteristics(ctx *CharacteristicsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitCardRule(ctx *CardRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitFlagRule(ctx *FlagRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitValueSetRule(ctx *ValueSetRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitFixedValueRule(ctx *FixedValueRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitContainsRule(ctx *ContainsRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitOnlyRule(ctx *OnlyRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitObeysRule(ctx *ObeysRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitCaretValueRule(ctx *CaretValueRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitCodeCaretValueRule(ctx *CodeCaretValueRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitMappingRule(ctx *MappingRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitInsertRule(ctx *InsertRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitCodeInsertRule(ctx *CodeInsertRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitAddCRElementRule(ctx *AddCRElementRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitAddElementRule(ctx *AddElementRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitPathRule(ctx *PathRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitVsComponent(ctx *VsComponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitVsConceptComponent(ctx *VsConceptComponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitVsFilterComponent(ctx *VsFilterComponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitVsComponentFrom(ctx *VsComponentFromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitVsFromSystem(ctx *VsFromSystemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitVsFromValueset(ctx *VsFromValuesetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitVsFilterList(ctx *VsFilterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitVsFilterDefinition(ctx *VsFilterDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitVsFilterOperator(ctx *VsFilterOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitVsFilterValue(ctx *VsFilterValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitPath(ctx *PathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitCaretPath(ctx *CaretPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitFlag(ctx *FlagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitStrength(ctx *StrengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitItem(ctx *ItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitCode(ctx *CodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitConcept(ctx *ConceptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitQuantity(ctx *QuantityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitRatio(ctx *RatioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitReference(ctx *ReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitReferenceType(ctx *ReferenceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitCodeableReferenceType(ctx *CodeableReferenceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitCanonical(ctx *CanonicalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitRatioPart(ctx *RatioPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitBool(ctx *BoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitTargetType(ctx *TargetTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFSHVisitor) VisitMostAlphaKeywords(ctx *MostAlphaKeywordsContext) interface{} {
	return v.VisitChildren(ctx)
}
