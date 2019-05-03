// Code generated from Thrift.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Thrift
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ThriftParser.
type ThriftVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ThriftParser#document.
	VisitDocument(ctx *DocumentContext) interface{}

	// Visit a parse tree produced by ThriftParser#header.
	VisitHeader(ctx *HeaderContext) interface{}

	// Visit a parse tree produced by ThriftParser#include.
	VisitInclude(ctx *IncludeContext) interface{}

	// Visit a parse tree produced by ThriftParser#namespace.
	VisitNamespace(ctx *NamespaceContext) interface{}

	// Visit a parse tree produced by ThriftParser#cpp_include.
	VisitCpp_include(ctx *Cpp_includeContext) interface{}

	// Visit a parse tree produced by ThriftParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by ThriftParser#const_rule.
	VisitConst_rule(ctx *Const_ruleContext) interface{}

	// Visit a parse tree produced by ThriftParser#typedef.
	VisitTypedef(ctx *TypedefContext) interface{}

	// Visit a parse tree produced by ThriftParser#enum_rule.
	VisitEnum_rule(ctx *Enum_ruleContext) interface{}

	// Visit a parse tree produced by ThriftParser#enum_field.
	VisitEnum_field(ctx *Enum_fieldContext) interface{}

	// Visit a parse tree produced by ThriftParser#senum.
	VisitSenum(ctx *SenumContext) interface{}

	// Visit a parse tree produced by ThriftParser#goStruct.
	VisitGoStruct(ctx *GoStructContext) interface{}

	// Visit a parse tree produced by ThriftParser#union.
	VisitUnion(ctx *UnionContext) interface{}

	// Visit a parse tree produced by ThriftParser#exception.
	VisitException(ctx *ExceptionContext) interface{}

	// Visit a parse tree produced by ThriftParser#service.
	VisitService(ctx *ServiceContext) interface{}

	// Visit a parse tree produced by ThriftParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by ThriftParser#field_id.
	VisitField_id(ctx *Field_idContext) interface{}

	// Visit a parse tree produced by ThriftParser#field_req.
	VisitField_req(ctx *Field_reqContext) interface{}

	// Visit a parse tree produced by ThriftParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by ThriftParser#oneway.
	VisitOneway(ctx *OnewayContext) interface{}

	// Visit a parse tree produced by ThriftParser#function_type.
	VisitFunction_type(ctx *Function_typeContext) interface{}

	// Visit a parse tree produced by ThriftParser#throws_list.
	VisitThrows_list(ctx *Throws_listContext) interface{}

	// Visit a parse tree produced by ThriftParser#type_annotations.
	VisitType_annotations(ctx *Type_annotationsContext) interface{}

	// Visit a parse tree produced by ThriftParser#type_annotation.
	VisitType_annotation(ctx *Type_annotationContext) interface{}

	// Visit a parse tree produced by ThriftParser#annotation_value.
	VisitAnnotation_value(ctx *Annotation_valueContext) interface{}

	// Visit a parse tree produced by ThriftParser#field_type.
	VisitField_type(ctx *Field_typeContext) interface{}

	// Visit a parse tree produced by ThriftParser#base_type.
	VisitBase_type(ctx *Base_typeContext) interface{}

	// Visit a parse tree produced by ThriftParser#container_type.
	VisitContainer_type(ctx *Container_typeContext) interface{}

	// Visit a parse tree produced by ThriftParser#map_type.
	VisitMap_type(ctx *Map_typeContext) interface{}

	// Visit a parse tree produced by ThriftParser#set_type.
	VisitSet_type(ctx *Set_typeContext) interface{}

	// Visit a parse tree produced by ThriftParser#list_type.
	VisitList_type(ctx *List_typeContext) interface{}

	// Visit a parse tree produced by ThriftParser#cpp_type.
	VisitCpp_type(ctx *Cpp_typeContext) interface{}

	// Visit a parse tree produced by ThriftParser#const_value.
	VisitConst_value(ctx *Const_valueContext) interface{}

	// Visit a parse tree produced by ThriftParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by ThriftParser#const_list.
	VisitConst_list(ctx *Const_listContext) interface{}

	// Visit a parse tree produced by ThriftParser#const_map_entry.
	VisitConst_map_entry(ctx *Const_map_entryContext) interface{}

	// Visit a parse tree produced by ThriftParser#const_map.
	VisitConst_map(ctx *Const_mapContext) interface{}

	// Visit a parse tree produced by ThriftParser#list_separator.
	VisitList_separator(ctx *List_separatorContext) interface{}

	// Visit a parse tree produced by ThriftParser#real_base_type.
	VisitReal_base_type(ctx *Real_base_typeContext) interface{}
}
