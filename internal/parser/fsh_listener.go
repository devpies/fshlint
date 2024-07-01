// Code generated from FSH.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // FSH

import "github.com/antlr4-go/antlr/v4"

// FSHListener is a complete listener for a parse tree produced by FSHParser.
type FSHListener interface {
	antlr.ParseTreeListener

	// EnterDoc is called when entering the doc production.
	EnterDoc(c *DocContext)

	// EnterEntity is called when entering the entity production.
	EnterEntity(c *EntityContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterProfile is called when entering the profile production.
	EnterProfile(c *ProfileContext)

	// EnterExtension is called when entering the extension production.
	EnterExtension(c *ExtensionContext)

	// EnterLogical is called when entering the logical production.
	EnterLogical(c *LogicalContext)

	// EnterResource is called when entering the resource production.
	EnterResource(c *ResourceContext)

	// EnterSdMetadata is called when entering the sdMetadata production.
	EnterSdMetadata(c *SdMetadataContext)

	// EnterSdRule is called when entering the sdRule production.
	EnterSdRule(c *SdRuleContext)

	// EnterLrRule is called when entering the lrRule production.
	EnterLrRule(c *LrRuleContext)

	// EnterInstance is called when entering the instance production.
	EnterInstance(c *InstanceContext)

	// EnterInstanceMetadata is called when entering the instanceMetadata production.
	EnterInstanceMetadata(c *InstanceMetadataContext)

	// EnterInstanceRule is called when entering the instanceRule production.
	EnterInstanceRule(c *InstanceRuleContext)

	// EnterInvariant is called when entering the invariant production.
	EnterInvariant(c *InvariantContext)

	// EnterInvariantMetadata is called when entering the invariantMetadata production.
	EnterInvariantMetadata(c *InvariantMetadataContext)

	// EnterInvariantRule is called when entering the invariantRule production.
	EnterInvariantRule(c *InvariantRuleContext)

	// EnterValueSet is called when entering the valueSet production.
	EnterValueSet(c *ValueSetContext)

	// EnterVsMetadata is called when entering the vsMetadata production.
	EnterVsMetadata(c *VsMetadataContext)

	// EnterVsRule is called when entering the vsRule production.
	EnterVsRule(c *VsRuleContext)

	// EnterCodeSystem is called when entering the codeSystem production.
	EnterCodeSystem(c *CodeSystemContext)

	// EnterCsMetadata is called when entering the csMetadata production.
	EnterCsMetadata(c *CsMetadataContext)

	// EnterCsRule is called when entering the csRule production.
	EnterCsRule(c *CsRuleContext)

	// EnterRuleSet is called when entering the ruleSet production.
	EnterRuleSet(c *RuleSetContext)

	// EnterRuleSetRule is called when entering the ruleSetRule production.
	EnterRuleSetRule(c *RuleSetRuleContext)

	// EnterParamRuleSet is called when entering the paramRuleSet production.
	EnterParamRuleSet(c *ParamRuleSetContext)

	// EnterParamRuleSetRef is called when entering the paramRuleSetRef production.
	EnterParamRuleSetRef(c *ParamRuleSetRefContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterLastParameter is called when entering the lastParameter production.
	EnterLastParameter(c *LastParameterContext)

	// EnterParamRuleSetContent is called when entering the paramRuleSetContent production.
	EnterParamRuleSetContent(c *ParamRuleSetContentContext)

	// EnterMapping is called when entering the mapping production.
	EnterMapping(c *MappingContext)

	// EnterMappingMetadata is called when entering the mappingMetadata production.
	EnterMappingMetadata(c *MappingMetadataContext)

	// EnterMappingEntityRule is called when entering the mappingEntityRule production.
	EnterMappingEntityRule(c *MappingEntityRuleContext)

	// EnterParent is called when entering the parent production.
	EnterParent(c *ParentContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterTitle is called when entering the title production.
	EnterTitle(c *TitleContext)

	// EnterDescription is called when entering the description production.
	EnterDescription(c *DescriptionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterXpath is called when entering the xpath production.
	EnterXpath(c *XpathContext)

	// EnterSeverity is called when entering the severity production.
	EnterSeverity(c *SeverityContext)

	// EnterInstanceOf is called when entering the instanceOf production.
	EnterInstanceOf(c *InstanceOfContext)

	// EnterUsage is called when entering the usage production.
	EnterUsage(c *UsageContext)

	// EnterSource is called when entering the source production.
	EnterSource(c *SourceContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// EnterContext is called when entering the context production.
	EnterContext(c *ContextContext)

	// EnterContextItem is called when entering the contextItem production.
	EnterContextItem(c *ContextItemContext)

	// EnterLastContextItem is called when entering the lastContextItem production.
	EnterLastContextItem(c *LastContextItemContext)

	// EnterCharacteristics is called when entering the characteristics production.
	EnterCharacteristics(c *CharacteristicsContext)

	// EnterCardRule is called when entering the cardRule production.
	EnterCardRule(c *CardRuleContext)

	// EnterFlagRule is called when entering the flagRule production.
	EnterFlagRule(c *FlagRuleContext)

	// EnterValueSetRule is called when entering the valueSetRule production.
	EnterValueSetRule(c *ValueSetRuleContext)

	// EnterFixedValueRule is called when entering the fixedValueRule production.
	EnterFixedValueRule(c *FixedValueRuleContext)

	// EnterContainsRule is called when entering the containsRule production.
	EnterContainsRule(c *ContainsRuleContext)

	// EnterOnlyRule is called when entering the onlyRule production.
	EnterOnlyRule(c *OnlyRuleContext)

	// EnterObeysRule is called when entering the obeysRule production.
	EnterObeysRule(c *ObeysRuleContext)

	// EnterCaretValueRule is called when entering the caretValueRule production.
	EnterCaretValueRule(c *CaretValueRuleContext)

	// EnterCodeCaretValueRule is called when entering the codeCaretValueRule production.
	EnterCodeCaretValueRule(c *CodeCaretValueRuleContext)

	// EnterMappingRule is called when entering the mappingRule production.
	EnterMappingRule(c *MappingRuleContext)

	// EnterInsertRule is called when entering the insertRule production.
	EnterInsertRule(c *InsertRuleContext)

	// EnterCodeInsertRule is called when entering the codeInsertRule production.
	EnterCodeInsertRule(c *CodeInsertRuleContext)

	// EnterAddCRElementRule is called when entering the addCRElementRule production.
	EnterAddCRElementRule(c *AddCRElementRuleContext)

	// EnterAddElementRule is called when entering the addElementRule production.
	EnterAddElementRule(c *AddElementRuleContext)

	// EnterPathRule is called when entering the pathRule production.
	EnterPathRule(c *PathRuleContext)

	// EnterVsComponent is called when entering the vsComponent production.
	EnterVsComponent(c *VsComponentContext)

	// EnterVsConceptComponent is called when entering the vsConceptComponent production.
	EnterVsConceptComponent(c *VsConceptComponentContext)

	// EnterVsFilterComponent is called when entering the vsFilterComponent production.
	EnterVsFilterComponent(c *VsFilterComponentContext)

	// EnterVsComponentFrom is called when entering the vsComponentFrom production.
	EnterVsComponentFrom(c *VsComponentFromContext)

	// EnterVsFromSystem is called when entering the vsFromSystem production.
	EnterVsFromSystem(c *VsFromSystemContext)

	// EnterVsFromValueset is called when entering the vsFromValueset production.
	EnterVsFromValueset(c *VsFromValuesetContext)

	// EnterVsFilterList is called when entering the vsFilterList production.
	EnterVsFilterList(c *VsFilterListContext)

	// EnterVsFilterDefinition is called when entering the vsFilterDefinition production.
	EnterVsFilterDefinition(c *VsFilterDefinitionContext)

	// EnterVsFilterOperator is called when entering the vsFilterOperator production.
	EnterVsFilterOperator(c *VsFilterOperatorContext)

	// EnterVsFilterValue is called when entering the vsFilterValue production.
	EnterVsFilterValue(c *VsFilterValueContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterCaretPath is called when entering the caretPath production.
	EnterCaretPath(c *CaretPathContext)

	// EnterFlag is called when entering the flag production.
	EnterFlag(c *FlagContext)

	// EnterStrength is called when entering the strength production.
	EnterStrength(c *StrengthContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterItem is called when entering the item production.
	EnterItem(c *ItemContext)

	// EnterCode is called when entering the code production.
	EnterCode(c *CodeContext)

	// EnterConcept is called when entering the concept production.
	EnterConcept(c *ConceptContext)

	// EnterQuantity is called when entering the quantity production.
	EnterQuantity(c *QuantityContext)

	// EnterRatio is called when entering the ratio production.
	EnterRatio(c *RatioContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterReferenceType is called when entering the referenceType production.
	EnterReferenceType(c *ReferenceTypeContext)

	// EnterCodeableReferenceType is called when entering the codeableReferenceType production.
	EnterCodeableReferenceType(c *CodeableReferenceTypeContext)

	// EnterCanonical is called when entering the canonical production.
	EnterCanonical(c *CanonicalContext)

	// EnterRatioPart is called when entering the ratioPart production.
	EnterRatioPart(c *RatioPartContext)

	// EnterBool is called when entering the bool production.
	EnterBool(c *BoolContext)

	// EnterTargetType is called when entering the targetType production.
	EnterTargetType(c *TargetTypeContext)

	// EnterMostAlphaKeywords is called when entering the mostAlphaKeywords production.
	EnterMostAlphaKeywords(c *MostAlphaKeywordsContext)

	// ExitDoc is called when exiting the doc production.
	ExitDoc(c *DocContext)

	// ExitEntity is called when exiting the entity production.
	ExitEntity(c *EntityContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitProfile is called when exiting the profile production.
	ExitProfile(c *ProfileContext)

	// ExitExtension is called when exiting the extension production.
	ExitExtension(c *ExtensionContext)

	// ExitLogical is called when exiting the logical production.
	ExitLogical(c *LogicalContext)

	// ExitResource is called when exiting the resource production.
	ExitResource(c *ResourceContext)

	// ExitSdMetadata is called when exiting the sdMetadata production.
	ExitSdMetadata(c *SdMetadataContext)

	// ExitSdRule is called when exiting the sdRule production.
	ExitSdRule(c *SdRuleContext)

	// ExitLrRule is called when exiting the lrRule production.
	ExitLrRule(c *LrRuleContext)

	// ExitInstance is called when exiting the instance production.
	ExitInstance(c *InstanceContext)

	// ExitInstanceMetadata is called when exiting the instanceMetadata production.
	ExitInstanceMetadata(c *InstanceMetadataContext)

	// ExitInstanceRule is called when exiting the instanceRule production.
	ExitInstanceRule(c *InstanceRuleContext)

	// ExitInvariant is called when exiting the invariant production.
	ExitInvariant(c *InvariantContext)

	// ExitInvariantMetadata is called when exiting the invariantMetadata production.
	ExitInvariantMetadata(c *InvariantMetadataContext)

	// ExitInvariantRule is called when exiting the invariantRule production.
	ExitInvariantRule(c *InvariantRuleContext)

	// ExitValueSet is called when exiting the valueSet production.
	ExitValueSet(c *ValueSetContext)

	// ExitVsMetadata is called when exiting the vsMetadata production.
	ExitVsMetadata(c *VsMetadataContext)

	// ExitVsRule is called when exiting the vsRule production.
	ExitVsRule(c *VsRuleContext)

	// ExitCodeSystem is called when exiting the codeSystem production.
	ExitCodeSystem(c *CodeSystemContext)

	// ExitCsMetadata is called when exiting the csMetadata production.
	ExitCsMetadata(c *CsMetadataContext)

	// ExitCsRule is called when exiting the csRule production.
	ExitCsRule(c *CsRuleContext)

	// ExitRuleSet is called when exiting the ruleSet production.
	ExitRuleSet(c *RuleSetContext)

	// ExitRuleSetRule is called when exiting the ruleSetRule production.
	ExitRuleSetRule(c *RuleSetRuleContext)

	// ExitParamRuleSet is called when exiting the paramRuleSet production.
	ExitParamRuleSet(c *ParamRuleSetContext)

	// ExitParamRuleSetRef is called when exiting the paramRuleSetRef production.
	ExitParamRuleSetRef(c *ParamRuleSetRefContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitLastParameter is called when exiting the lastParameter production.
	ExitLastParameter(c *LastParameterContext)

	// ExitParamRuleSetContent is called when exiting the paramRuleSetContent production.
	ExitParamRuleSetContent(c *ParamRuleSetContentContext)

	// ExitMapping is called when exiting the mapping production.
	ExitMapping(c *MappingContext)

	// ExitMappingMetadata is called when exiting the mappingMetadata production.
	ExitMappingMetadata(c *MappingMetadataContext)

	// ExitMappingEntityRule is called when exiting the mappingEntityRule production.
	ExitMappingEntityRule(c *MappingEntityRuleContext)

	// ExitParent is called when exiting the parent production.
	ExitParent(c *ParentContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitTitle is called when exiting the title production.
	ExitTitle(c *TitleContext)

	// ExitDescription is called when exiting the description production.
	ExitDescription(c *DescriptionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitXpath is called when exiting the xpath production.
	ExitXpath(c *XpathContext)

	// ExitSeverity is called when exiting the severity production.
	ExitSeverity(c *SeverityContext)

	// ExitInstanceOf is called when exiting the instanceOf production.
	ExitInstanceOf(c *InstanceOfContext)

	// ExitUsage is called when exiting the usage production.
	ExitUsage(c *UsageContext)

	// ExitSource is called when exiting the source production.
	ExitSource(c *SourceContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)

	// ExitContext is called when exiting the context production.
	ExitContext(c *ContextContext)

	// ExitContextItem is called when exiting the contextItem production.
	ExitContextItem(c *ContextItemContext)

	// ExitLastContextItem is called when exiting the lastContextItem production.
	ExitLastContextItem(c *LastContextItemContext)

	// ExitCharacteristics is called when exiting the characteristics production.
	ExitCharacteristics(c *CharacteristicsContext)

	// ExitCardRule is called when exiting the cardRule production.
	ExitCardRule(c *CardRuleContext)

	// ExitFlagRule is called when exiting the flagRule production.
	ExitFlagRule(c *FlagRuleContext)

	// ExitValueSetRule is called when exiting the valueSetRule production.
	ExitValueSetRule(c *ValueSetRuleContext)

	// ExitFixedValueRule is called when exiting the fixedValueRule production.
	ExitFixedValueRule(c *FixedValueRuleContext)

	// ExitContainsRule is called when exiting the containsRule production.
	ExitContainsRule(c *ContainsRuleContext)

	// ExitOnlyRule is called when exiting the onlyRule production.
	ExitOnlyRule(c *OnlyRuleContext)

	// ExitObeysRule is called when exiting the obeysRule production.
	ExitObeysRule(c *ObeysRuleContext)

	// ExitCaretValueRule is called when exiting the caretValueRule production.
	ExitCaretValueRule(c *CaretValueRuleContext)

	// ExitCodeCaretValueRule is called when exiting the codeCaretValueRule production.
	ExitCodeCaretValueRule(c *CodeCaretValueRuleContext)

	// ExitMappingRule is called when exiting the mappingRule production.
	ExitMappingRule(c *MappingRuleContext)

	// ExitInsertRule is called when exiting the insertRule production.
	ExitInsertRule(c *InsertRuleContext)

	// ExitCodeInsertRule is called when exiting the codeInsertRule production.
	ExitCodeInsertRule(c *CodeInsertRuleContext)

	// ExitAddCRElementRule is called when exiting the addCRElementRule production.
	ExitAddCRElementRule(c *AddCRElementRuleContext)

	// ExitAddElementRule is called when exiting the addElementRule production.
	ExitAddElementRule(c *AddElementRuleContext)

	// ExitPathRule is called when exiting the pathRule production.
	ExitPathRule(c *PathRuleContext)

	// ExitVsComponent is called when exiting the vsComponent production.
	ExitVsComponent(c *VsComponentContext)

	// ExitVsConceptComponent is called when exiting the vsConceptComponent production.
	ExitVsConceptComponent(c *VsConceptComponentContext)

	// ExitVsFilterComponent is called when exiting the vsFilterComponent production.
	ExitVsFilterComponent(c *VsFilterComponentContext)

	// ExitVsComponentFrom is called when exiting the vsComponentFrom production.
	ExitVsComponentFrom(c *VsComponentFromContext)

	// ExitVsFromSystem is called when exiting the vsFromSystem production.
	ExitVsFromSystem(c *VsFromSystemContext)

	// ExitVsFromValueset is called when exiting the vsFromValueset production.
	ExitVsFromValueset(c *VsFromValuesetContext)

	// ExitVsFilterList is called when exiting the vsFilterList production.
	ExitVsFilterList(c *VsFilterListContext)

	// ExitVsFilterDefinition is called when exiting the vsFilterDefinition production.
	ExitVsFilterDefinition(c *VsFilterDefinitionContext)

	// ExitVsFilterOperator is called when exiting the vsFilterOperator production.
	ExitVsFilterOperator(c *VsFilterOperatorContext)

	// ExitVsFilterValue is called when exiting the vsFilterValue production.
	ExitVsFilterValue(c *VsFilterValueContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)

	// ExitCaretPath is called when exiting the caretPath production.
	ExitCaretPath(c *CaretPathContext)

	// ExitFlag is called when exiting the flag production.
	ExitFlag(c *FlagContext)

	// ExitStrength is called when exiting the strength production.
	ExitStrength(c *StrengthContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitItem is called when exiting the item production.
	ExitItem(c *ItemContext)

	// ExitCode is called when exiting the code production.
	ExitCode(c *CodeContext)

	// ExitConcept is called when exiting the concept production.
	ExitConcept(c *ConceptContext)

	// ExitQuantity is called when exiting the quantity production.
	ExitQuantity(c *QuantityContext)

	// ExitRatio is called when exiting the ratio production.
	ExitRatio(c *RatioContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitReferenceType is called when exiting the referenceType production.
	ExitReferenceType(c *ReferenceTypeContext)

	// ExitCodeableReferenceType is called when exiting the codeableReferenceType production.
	ExitCodeableReferenceType(c *CodeableReferenceTypeContext)

	// ExitCanonical is called when exiting the canonical production.
	ExitCanonical(c *CanonicalContext)

	// ExitRatioPart is called when exiting the ratioPart production.
	ExitRatioPart(c *RatioPartContext)

	// ExitBool is called when exiting the bool production.
	ExitBool(c *BoolContext)

	// ExitTargetType is called when exiting the targetType production.
	ExitTargetType(c *TargetTypeContext)

	// ExitMostAlphaKeywords is called when exiting the mostAlphaKeywords production.
	ExitMostAlphaKeywords(c *MostAlphaKeywordsContext)
}
