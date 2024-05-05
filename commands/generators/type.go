package generators

import (
	"fmt"
)

// GraphQLType represents a simple structure for a GraphQL type.
type GraphQLType struct {
	Name   string
	Fields []GraphQLField
}

// GraphQLField represents the fields in a GraphQL type.
type GraphQLField struct {
	Name string
	Type string
}

// NewGraphQLType creates a new GraphQLType with the specified name and fields.
func NewGraphQLType(name string, fields []GraphQLField) *GraphQLType {
	return &GraphQLType{
		Name:   name,
		Fields: fields,
	}
}

// GenerateGoStruct generates a Go struct representation of the GraphQL type.
func (gt *GraphQLType) GenerateGoStruct() string {
	result := fmt.Sprintf("type %s struct {\n", gt.Name)
	for _, field := range gt.Fields {
		result += fmt.Sprintf("\t%s %s\n", field.Name, goType(field.Type))
	}
	result += "}"
	return result
}

// goType maps GraphQL types to Go types.
func goType(graphqlType string) string {
	switch graphqlType {
	case "String":
		return "string"
	case "Int":
		return "int"
	case "Float":
		return "float64"
	case "Boolean":
		return "bool"
	default:
		return "interface{}" // Default to interface{} for complex or unknown types
	}
}

// PrintStructure prints the structure of the type to standard output, for example purposes.
func (gt *GraphQLType) PrintStructure() {
	fmt.Println("GraphQL Type:", gt.Name)
	for _, field := range gt.Fields {
		fmt.Println("Field:", field.Name, "Type:", field.Type)
	}
}
