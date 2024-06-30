// Code generated from FSH.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // FSH

import "github.com/antlr4-go/antlr/v4"

// BaseFSHListener is a complete listener for a parse tree produced by FSHParser.
type BaseFSHListener struct{}

var _ FSHListener = &BaseFSHListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFSHListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFSHListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFSHListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFSHListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDoc is called when production doc is entered.
func (s *BaseFSHListener) EnterDoc(ctx *DocContext) {}

// ExitDoc is called when production doc is exited.
func (s *BaseFSHListener) ExitDoc(ctx *DocContext) {}

// EnterEntity is called when production entity is entered.
func (s *BaseFSHListener) EnterEntity(ctx *EntityContext) {}

// ExitEntity is called when production entity is exited.
func (s *BaseFSHListener) ExitEntity(ctx *EntityContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseFSHListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseFSHListener) ExitAlias(ctx *AliasContext) {}

// EnterProfile is called when production profile is entered.
func (s *BaseFSHListener) EnterProfile(ctx *ProfileContext) {}

// ExitProfile is called when production profile is exited.
func (s *BaseFSHListener) ExitProfile(ctx *ProfileContext) {}

// EnterExtension is called when production extension is entered.
func (s *BaseFSHListener) EnterExtension(ctx *ExtensionContext) {}

// ExitExtension is called when production extension is exited.
func (s *BaseFSHListener) ExitExtension(ctx *ExtensionContext) {}

// EnterLogical is called when production logical is entered.
func (s *BaseFSHListener) EnterLogical(ctx *LogicalContext) {}

// ExitLogical is called when production logical is exited.
func (s *BaseFSHListener) ExitLogical(ctx *LogicalContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseFSHListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseFSHListener) ExitResource(ctx *ResourceContext) {}

// EnterSdMetadata is called when production sdMetadata is entered.
func (s *BaseFSHListener) EnterSdMetadata(ctx *SdMetadataContext) {}

// ExitSdMetadata is called when production sdMetadata is exited.
func (s *BaseFSHListener) ExitSdMetadata(ctx *SdMetadataContext) {}

// EnterSdRule is called when production sdRule is entered.
func (s *BaseFSHListener) EnterSdRule(ctx *SdRuleContext) {}

// ExitSdRule is called when production sdRule is exited.
func (s *BaseFSHListener) ExitSdRule(ctx *SdRuleContext) {}

// EnterLrRule is called when production lrRule is entered.
func (s *BaseFSHListener) EnterLrRule(ctx *LrRuleContext) {}

// ExitLrRule is called when production lrRule is exited.
func (s *BaseFSHListener) ExitLrRule(ctx *LrRuleContext) {}

// EnterInstance is called when production instance is entered.
func (s *BaseFSHListener) EnterInstance(ctx *InstanceContext) {}

// ExitInstance is called when production instance is exited.
func (s *BaseFSHListener) ExitInstance(ctx *InstanceContext) {}

// EnterInstanceMetadata is called when production instanceMetadata is entered.
func (s *BaseFSHListener) EnterInstanceMetadata(ctx *InstanceMetadataContext) {}

// ExitInstanceMetadata is called when production instanceMetadata is exited.
func (s *BaseFSHListener) ExitInstanceMetadata(ctx *InstanceMetadataContext) {}

// EnterInstanceRule is called when production instanceRule is entered.
func (s *BaseFSHListener) EnterInstanceRule(ctx *InstanceRuleContext) {}

// ExitInstanceRule is called when production instanceRule is exited.
func (s *BaseFSHListener) ExitInstanceRule(ctx *InstanceRuleContext) {}

// EnterInvariant is called when production invariant is entered.
func (s *BaseFSHListener) EnterInvariant(ctx *InvariantContext) {}

// ExitInvariant is called when production invariant is exited.
func (s *BaseFSHListener) ExitInvariant(ctx *InvariantContext) {}

// EnterInvariantMetadata is called when production invariantMetadata is entered.
func (s *BaseFSHListener) EnterInvariantMetadata(ctx *InvariantMetadataContext) {}

// ExitInvariantMetadata is called when production invariantMetadata is exited.
func (s *BaseFSHListener) ExitInvariantMetadata(ctx *InvariantMetadataContext) {}

// EnterInvariantRule is called when production invariantRule is entered.
func (s *BaseFSHListener) EnterInvariantRule(ctx *InvariantRuleContext) {}

// ExitInvariantRule is called when production invariantRule is exited.
func (s *BaseFSHListener) ExitInvariantRule(ctx *InvariantRuleContext) {}

// EnterValueSet is called when production valueSet is entered.
func (s *BaseFSHListener) EnterValueSet(ctx *ValueSetContext) {}

// ExitValueSet is called when production valueSet is exited.
func (s *BaseFSHListener) ExitValueSet(ctx *ValueSetContext) {}

// EnterVsMetadata is called when production vsMetadata is entered.
func (s *BaseFSHListener) EnterVsMetadata(ctx *VsMetadataContext) {}

// ExitVsMetadata is called when production vsMetadata is exited.
func (s *BaseFSHListener) ExitVsMetadata(ctx *VsMetadataContext) {}

// EnterVsRule is called when production vsRule is entered.
func (s *BaseFSHListener) EnterVsRule(ctx *VsRuleContext) {}

// ExitVsRule is called when production vsRule is exited.
func (s *BaseFSHListener) ExitVsRule(ctx *VsRuleContext) {}

// EnterCodeSystem is called when production codeSystem is entered.
func (s *BaseFSHListener) EnterCodeSystem(ctx *CodeSystemContext) {}

// ExitCodeSystem is called when production codeSystem is exited.
func (s *BaseFSHListener) ExitCodeSystem(ctx *CodeSystemContext) {}

// EnterCsMetadata is called when production csMetadata is entered.
func (s *BaseFSHListener) EnterCsMetadata(ctx *CsMetadataContext) {}

// ExitCsMetadata is called when production csMetadata is exited.
func (s *BaseFSHListener) ExitCsMetadata(ctx *CsMetadataContext) {}

// EnterCsRule is called when production csRule is entered.
func (s *BaseFSHListener) EnterCsRule(ctx *CsRuleContext) {}

// ExitCsRule is called when production csRule is exited.
func (s *BaseFSHListener) ExitCsRule(ctx *CsRuleContext) {}

// EnterRuleSet is called when production ruleSet is entered.
func (s *BaseFSHListener) EnterRuleSet(ctx *RuleSetContext) {}

// ExitRuleSet is called when production ruleSet is exited.
func (s *BaseFSHListener) ExitRuleSet(ctx *RuleSetContext) {}

// EnterRuleSetRule is called when production ruleSetRule is entered.
func (s *BaseFSHListener) EnterRuleSetRule(ctx *RuleSetRuleContext) {}

// ExitRuleSetRule is called when production ruleSetRule is exited.
func (s *BaseFSHListener) ExitRuleSetRule(ctx *RuleSetRuleContext) {}

// EnterParamRuleSet is called when production paramRuleSet is entered.
func (s *BaseFSHListener) EnterParamRuleSet(ctx *ParamRuleSetContext) {}

// ExitParamRuleSet is called when production paramRuleSet is exited.
func (s *BaseFSHListener) ExitParamRuleSet(ctx *ParamRuleSetContext) {}

// EnterParamRuleSetRef is called when production paramRuleSetRef is entered.
func (s *BaseFSHListener) EnterParamRuleSetRef(ctx *ParamRuleSetRefContext) {}

// ExitParamRuleSetRef is called when production paramRuleSetRef is exited.
func (s *BaseFSHListener) ExitParamRuleSetRef(ctx *ParamRuleSetRefContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseFSHListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseFSHListener) ExitParameter(ctx *ParameterContext) {}

// EnterLastParameter is called when production lastParameter is entered.
func (s *BaseFSHListener) EnterLastParameter(ctx *LastParameterContext) {}

// ExitLastParameter is called when production lastParameter is exited.
func (s *BaseFSHListener) ExitLastParameter(ctx *LastParameterContext) {}

// EnterParamRuleSetContent is called when production paramRuleSetContent is entered.
func (s *BaseFSHListener) EnterParamRuleSetContent(ctx *ParamRuleSetContentContext) {}

// ExitParamRuleSetContent is called when production paramRuleSetContent is exited.
func (s *BaseFSHListener) ExitParamRuleSetContent(ctx *ParamRuleSetContentContext) {}

// EnterMapping is called when production mapping is entered.
func (s *BaseFSHListener) EnterMapping(ctx *MappingContext) {}

// ExitMapping is called when production mapping is exited.
func (s *BaseFSHListener) ExitMapping(ctx *MappingContext) {}

// EnterMappingMetadata is called when production mappingMetadata is entered.
func (s *BaseFSHListener) EnterMappingMetadata(ctx *MappingMetadataContext) {}

// ExitMappingMetadata is called when production mappingMetadata is exited.
func (s *BaseFSHListener) ExitMappingMetadata(ctx *MappingMetadataContext) {}

// EnterMappingEntityRule is called when production mappingEntityRule is entered.
func (s *BaseFSHListener) EnterMappingEntityRule(ctx *MappingEntityRuleContext) {}

// ExitMappingEntityRule is called when production mappingEntityRule is exited.
func (s *BaseFSHListener) ExitMappingEntityRule(ctx *MappingEntityRuleContext) {}

// EnterParent is called when production parent is entered.
func (s *BaseFSHListener) EnterParent(ctx *ParentContext) {}

// ExitParent is called when production parent is exited.
func (s *BaseFSHListener) ExitParent(ctx *ParentContext) {}

// EnterId is called when production id is entered.
func (s *BaseFSHListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseFSHListener) ExitId(ctx *IdContext) {}

// EnterTitle is called when production title is entered.
func (s *BaseFSHListener) EnterTitle(ctx *TitleContext) {}

// ExitTitle is called when production title is exited.
func (s *BaseFSHListener) ExitTitle(ctx *TitleContext) {}

// EnterDescription is called when production description is entered.
func (s *BaseFSHListener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BaseFSHListener) ExitDescription(ctx *DescriptionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFSHListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFSHListener) ExitExpression(ctx *ExpressionContext) {}

// EnterXpath is called when production xpath is entered.
func (s *BaseFSHListener) EnterXpath(ctx *XpathContext) {}

// ExitXpath is called when production xpath is exited.
func (s *BaseFSHListener) ExitXpath(ctx *XpathContext) {}

// EnterSeverity is called when production severity is entered.
func (s *BaseFSHListener) EnterSeverity(ctx *SeverityContext) {}

// ExitSeverity is called when production severity is exited.
func (s *BaseFSHListener) ExitSeverity(ctx *SeverityContext) {}

// EnterInstanceOf is called when production instanceOf is entered.
func (s *BaseFSHListener) EnterInstanceOf(ctx *InstanceOfContext) {}

// ExitInstanceOf is called when production instanceOf is exited.
func (s *BaseFSHListener) ExitInstanceOf(ctx *InstanceOfContext) {}

// EnterUsage is called when production usage is entered.
func (s *BaseFSHListener) EnterUsage(ctx *UsageContext) {}

// ExitUsage is called when production usage is exited.
func (s *BaseFSHListener) ExitUsage(ctx *UsageContext) {}

// EnterSource is called when production source is entered.
func (s *BaseFSHListener) EnterSource(ctx *SourceContext) {}

// ExitSource is called when production source is exited.
func (s *BaseFSHListener) ExitSource(ctx *SourceContext) {}

// EnterTarget is called when production target is entered.
func (s *BaseFSHListener) EnterTarget(ctx *TargetContext) {}

// ExitTarget is called when production target is exited.
func (s *BaseFSHListener) ExitTarget(ctx *TargetContext) {}

// EnterContext is called when production context is entered.
func (s *BaseFSHListener) EnterContext(ctx *ContextContext) {}

// ExitContext is called when production context is exited.
func (s *BaseFSHListener) ExitContext(ctx *ContextContext) {}

// EnterContextItem is called when production contextItem is entered.
func (s *BaseFSHListener) EnterContextItem(ctx *ContextItemContext) {}

// ExitContextItem is called when production contextItem is exited.
func (s *BaseFSHListener) ExitContextItem(ctx *ContextItemContext) {}

// EnterLastContextItem is called when production lastContextItem is entered.
func (s *BaseFSHListener) EnterLastContextItem(ctx *LastContextItemContext) {}

// ExitLastContextItem is called when production lastContextItem is exited.
func (s *BaseFSHListener) ExitLastContextItem(ctx *LastContextItemContext) {}

// EnterCharacteristics is called when production characteristics is entered.
func (s *BaseFSHListener) EnterCharacteristics(ctx *CharacteristicsContext) {}

// ExitCharacteristics is called when production characteristics is exited.
func (s *BaseFSHListener) ExitCharacteristics(ctx *CharacteristicsContext) {}

// EnterCardRule is called when production cardRule is entered.
func (s *BaseFSHListener) EnterCardRule(ctx *CardRuleContext) {}

// ExitCardRule is called when production cardRule is exited.
func (s *BaseFSHListener) ExitCardRule(ctx *CardRuleContext) {}

// EnterFlagRule is called when production flagRule is entered.
func (s *BaseFSHListener) EnterFlagRule(ctx *FlagRuleContext) {}

// ExitFlagRule is called when production flagRule is exited.
func (s *BaseFSHListener) ExitFlagRule(ctx *FlagRuleContext) {}

// EnterValueSetRule is called when production valueSetRule is entered.
func (s *BaseFSHListener) EnterValueSetRule(ctx *ValueSetRuleContext) {}

// ExitValueSetRule is called when production valueSetRule is exited.
func (s *BaseFSHListener) ExitValueSetRule(ctx *ValueSetRuleContext) {}

// EnterFixedValueRule is called when production fixedValueRule is entered.
func (s *BaseFSHListener) EnterFixedValueRule(ctx *FixedValueRuleContext) {}

// ExitFixedValueRule is called when production fixedValueRule is exited.
func (s *BaseFSHListener) ExitFixedValueRule(ctx *FixedValueRuleContext) {}

// EnterContainsRule is called when production containsRule is entered.
func (s *BaseFSHListener) EnterContainsRule(ctx *ContainsRuleContext) {}

// ExitContainsRule is called when production containsRule is exited.
func (s *BaseFSHListener) ExitContainsRule(ctx *ContainsRuleContext) {}

// EnterOnlyRule is called when production onlyRule is entered.
func (s *BaseFSHListener) EnterOnlyRule(ctx *OnlyRuleContext) {}

// ExitOnlyRule is called when production onlyRule is exited.
func (s *BaseFSHListener) ExitOnlyRule(ctx *OnlyRuleContext) {}

// EnterObeysRule is called when production obeysRule is entered.
func (s *BaseFSHListener) EnterObeysRule(ctx *ObeysRuleContext) {}

// ExitObeysRule is called when production obeysRule is exited.
func (s *BaseFSHListener) ExitObeysRule(ctx *ObeysRuleContext) {}

// EnterCaretValueRule is called when production caretValueRule is entered.
func (s *BaseFSHListener) EnterCaretValueRule(ctx *CaretValueRuleContext) {}

// ExitCaretValueRule is called when production caretValueRule is exited.
func (s *BaseFSHListener) ExitCaretValueRule(ctx *CaretValueRuleContext) {}

// EnterCodeCaretValueRule is called when production codeCaretValueRule is entered.
func (s *BaseFSHListener) EnterCodeCaretValueRule(ctx *CodeCaretValueRuleContext) {}

// ExitCodeCaretValueRule is called when production codeCaretValueRule is exited.
func (s *BaseFSHListener) ExitCodeCaretValueRule(ctx *CodeCaretValueRuleContext) {}

// EnterMappingRule is called when production mappingRule is entered.
func (s *BaseFSHListener) EnterMappingRule(ctx *MappingRuleContext) {}

// ExitMappingRule is called when production mappingRule is exited.
func (s *BaseFSHListener) ExitMappingRule(ctx *MappingRuleContext) {}

// EnterInsertRule is called when production insertRule is entered.
func (s *BaseFSHListener) EnterInsertRule(ctx *InsertRuleContext) {}

// ExitInsertRule is called when production insertRule is exited.
func (s *BaseFSHListener) ExitInsertRule(ctx *InsertRuleContext) {}

// EnterCodeInsertRule is called when production codeInsertRule is entered.
func (s *BaseFSHListener) EnterCodeInsertRule(ctx *CodeInsertRuleContext) {}

// ExitCodeInsertRule is called when production codeInsertRule is exited.
func (s *BaseFSHListener) ExitCodeInsertRule(ctx *CodeInsertRuleContext) {}

// EnterAddCRElementRule is called when production addCRElementRule is entered.
func (s *BaseFSHListener) EnterAddCRElementRule(ctx *AddCRElementRuleContext) {}

// ExitAddCRElementRule is called when production addCRElementRule is exited.
func (s *BaseFSHListener) ExitAddCRElementRule(ctx *AddCRElementRuleContext) {}

// EnterAddElementRule is called when production addElementRule is entered.
func (s *BaseFSHListener) EnterAddElementRule(ctx *AddElementRuleContext) {}

// ExitAddElementRule is called when production addElementRule is exited.
func (s *BaseFSHListener) ExitAddElementRule(ctx *AddElementRuleContext) {}

// EnterPathRule is called when production pathRule is entered.
func (s *BaseFSHListener) EnterPathRule(ctx *PathRuleContext) {}

// ExitPathRule is called when production pathRule is exited.
func (s *BaseFSHListener) ExitPathRule(ctx *PathRuleContext) {}

// EnterVsComponent is called when production vsComponent is entered.
func (s *BaseFSHListener) EnterVsComponent(ctx *VsComponentContext) {}

// ExitVsComponent is called when production vsComponent is exited.
func (s *BaseFSHListener) ExitVsComponent(ctx *VsComponentContext) {}

// EnterVsConceptComponent is called when production vsConceptComponent is entered.
func (s *BaseFSHListener) EnterVsConceptComponent(ctx *VsConceptComponentContext) {}

// ExitVsConceptComponent is called when production vsConceptComponent is exited.
func (s *BaseFSHListener) ExitVsConceptComponent(ctx *VsConceptComponentContext) {}

// EnterVsFilterComponent is called when production vsFilterComponent is entered.
func (s *BaseFSHListener) EnterVsFilterComponent(ctx *VsFilterComponentContext) {}

// ExitVsFilterComponent is called when production vsFilterComponent is exited.
func (s *BaseFSHListener) ExitVsFilterComponent(ctx *VsFilterComponentContext) {}

// EnterVsComponentFrom is called when production vsComponentFrom is entered.
func (s *BaseFSHListener) EnterVsComponentFrom(ctx *VsComponentFromContext) {}

// ExitVsComponentFrom is called when production vsComponentFrom is exited.
func (s *BaseFSHListener) ExitVsComponentFrom(ctx *VsComponentFromContext) {}

// EnterVsFromSystem is called when production vsFromSystem is entered.
func (s *BaseFSHListener) EnterVsFromSystem(ctx *VsFromSystemContext) {}

// ExitVsFromSystem is called when production vsFromSystem is exited.
func (s *BaseFSHListener) ExitVsFromSystem(ctx *VsFromSystemContext) {}

// EnterVsFromValueset is called when production vsFromValueset is entered.
func (s *BaseFSHListener) EnterVsFromValueset(ctx *VsFromValuesetContext) {}

// ExitVsFromValueset is called when production vsFromValueset is exited.
func (s *BaseFSHListener) ExitVsFromValueset(ctx *VsFromValuesetContext) {}

// EnterVsFilterList is called when production vsFilterList is entered.
func (s *BaseFSHListener) EnterVsFilterList(ctx *VsFilterListContext) {}

// ExitVsFilterList is called when production vsFilterList is exited.
func (s *BaseFSHListener) ExitVsFilterList(ctx *VsFilterListContext) {}

// EnterVsFilterDefinition is called when production vsFilterDefinition is entered.
func (s *BaseFSHListener) EnterVsFilterDefinition(ctx *VsFilterDefinitionContext) {}

// ExitVsFilterDefinition is called when production vsFilterDefinition is exited.
func (s *BaseFSHListener) ExitVsFilterDefinition(ctx *VsFilterDefinitionContext) {}

// EnterVsFilterOperator is called when production vsFilterOperator is entered.
func (s *BaseFSHListener) EnterVsFilterOperator(ctx *VsFilterOperatorContext) {}

// ExitVsFilterOperator is called when production vsFilterOperator is exited.
func (s *BaseFSHListener) ExitVsFilterOperator(ctx *VsFilterOperatorContext) {}

// EnterVsFilterValue is called when production vsFilterValue is entered.
func (s *BaseFSHListener) EnterVsFilterValue(ctx *VsFilterValueContext) {}

// ExitVsFilterValue is called when production vsFilterValue is exited.
func (s *BaseFSHListener) ExitVsFilterValue(ctx *VsFilterValueContext) {}

// EnterName is called when production name is entered.
func (s *BaseFSHListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseFSHListener) ExitName(ctx *NameContext) {}

// EnterPath is called when production path is entered.
func (s *BaseFSHListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseFSHListener) ExitPath(ctx *PathContext) {}

// EnterCaretPath is called when production caretPath is entered.
func (s *BaseFSHListener) EnterCaretPath(ctx *CaretPathContext) {}

// ExitCaretPath is called when production caretPath is exited.
func (s *BaseFSHListener) ExitCaretPath(ctx *CaretPathContext) {}

// EnterFlag is called when production flag is entered.
func (s *BaseFSHListener) EnterFlag(ctx *FlagContext) {}

// ExitFlag is called when production flag is exited.
func (s *BaseFSHListener) ExitFlag(ctx *FlagContext) {}

// EnterStrength is called when production strength is entered.
func (s *BaseFSHListener) EnterStrength(ctx *StrengthContext) {}

// ExitStrength is called when production strength is exited.
func (s *BaseFSHListener) ExitStrength(ctx *StrengthContext) {}

// EnterValue is called when production value is entered.
func (s *BaseFSHListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseFSHListener) ExitValue(ctx *ValueContext) {}

// EnterItem is called when production item is entered.
func (s *BaseFSHListener) EnterItem(ctx *ItemContext) {}

// ExitItem is called when production item is exited.
func (s *BaseFSHListener) ExitItem(ctx *ItemContext) {}

// EnterCode is called when production code is entered.
func (s *BaseFSHListener) EnterCode(ctx *CodeContext) {}

// ExitCode is called when production code is exited.
func (s *BaseFSHListener) ExitCode(ctx *CodeContext) {}

// EnterConcept is called when production concept is entered.
func (s *BaseFSHListener) EnterConcept(ctx *ConceptContext) {}

// ExitConcept is called when production concept is exited.
func (s *BaseFSHListener) ExitConcept(ctx *ConceptContext) {}

// EnterQuantity is called when production quantity is entered.
func (s *BaseFSHListener) EnterQuantity(ctx *QuantityContext) {}

// ExitQuantity is called when production quantity is exited.
func (s *BaseFSHListener) ExitQuantity(ctx *QuantityContext) {}

// EnterRatio is called when production ratio is entered.
func (s *BaseFSHListener) EnterRatio(ctx *RatioContext) {}

// ExitRatio is called when production ratio is exited.
func (s *BaseFSHListener) ExitRatio(ctx *RatioContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseFSHListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseFSHListener) ExitReference(ctx *ReferenceContext) {}

// EnterReferenceType is called when production referenceType is entered.
func (s *BaseFSHListener) EnterReferenceType(ctx *ReferenceTypeContext) {}

// ExitReferenceType is called when production referenceType is exited.
func (s *BaseFSHListener) ExitReferenceType(ctx *ReferenceTypeContext) {}

// EnterCodeableReferenceType is called when production codeableReferenceType is entered.
func (s *BaseFSHListener) EnterCodeableReferenceType(ctx *CodeableReferenceTypeContext) {}

// ExitCodeableReferenceType is called when production codeableReferenceType is exited.
func (s *BaseFSHListener) ExitCodeableReferenceType(ctx *CodeableReferenceTypeContext) {}

// EnterCanonical is called when production canonical is entered.
func (s *BaseFSHListener) EnterCanonical(ctx *CanonicalContext) {}

// ExitCanonical is called when production canonical is exited.
func (s *BaseFSHListener) ExitCanonical(ctx *CanonicalContext) {}

// EnterRatioPart is called when production ratioPart is entered.
func (s *BaseFSHListener) EnterRatioPart(ctx *RatioPartContext) {}

// ExitRatioPart is called when production ratioPart is exited.
func (s *BaseFSHListener) ExitRatioPart(ctx *RatioPartContext) {}

// EnterBool is called when production bool is entered.
func (s *BaseFSHListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production bool is exited.
func (s *BaseFSHListener) ExitBool(ctx *BoolContext) {}

// EnterTargetType is called when production targetType is entered.
func (s *BaseFSHListener) EnterTargetType(ctx *TargetTypeContext) {}

// ExitTargetType is called when production targetType is exited.
func (s *BaseFSHListener) ExitTargetType(ctx *TargetTypeContext) {}

// EnterMostAlphaKeywords is called when production mostAlphaKeywords is entered.
func (s *BaseFSHListener) EnterMostAlphaKeywords(ctx *MostAlphaKeywordsContext) {}

// ExitMostAlphaKeywords is called when production mostAlphaKeywords is exited.
func (s *BaseFSHListener) ExitMostAlphaKeywords(ctx *MostAlphaKeywordsContext) {}
