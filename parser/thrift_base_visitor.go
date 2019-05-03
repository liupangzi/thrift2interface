// Code generated from Thrift.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Thrift
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseThriftVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseThriftVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitHeader(ctx *HeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitInclude(ctx *IncludeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitNamespace(ctx *NamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitCpp_include(ctx *Cpp_includeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitConst_rule(ctx *Const_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitTypedef(ctx *TypedefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitEnum_rule(ctx *Enum_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitEnum_field(ctx *Enum_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitSenum(ctx *SenumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitGoStruct(ctx *GoStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitUnion(ctx *UnionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitException(ctx *ExceptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitService(ctx *ServiceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitField_id(ctx *Field_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitField_req(ctx *Field_reqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitOneway(ctx *OnewayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitFunction_type(ctx *Function_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitThrows_list(ctx *Throws_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitType_annotations(ctx *Type_annotationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitType_annotation(ctx *Type_annotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitAnnotation_value(ctx *Annotation_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitField_type(ctx *Field_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitBase_type(ctx *Base_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitContainer_type(ctx *Container_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitMap_type(ctx *Map_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitSet_type(ctx *Set_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitList_type(ctx *List_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitCpp_type(ctx *Cpp_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitConst_value(ctx *Const_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitConst_list(ctx *Const_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitConst_map_entry(ctx *Const_map_entryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitConst_map(ctx *Const_mapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitList_separator(ctx *List_separatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseThriftVisitor) VisitReal_base_type(ctx *Real_base_typeContext) interface{} {
	return v.VisitChildren(ctx)
}
