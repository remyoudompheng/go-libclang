package clang

type DiagnosticServerity uint

const (
	DiagnosticIgnored DiagnosticServerity = iota
	DiagnosticNote
	DiagnosticWarning
	DiagnosticError
	DiagnosticFatal
)

type CursorKind uint

const (
	// Declarations
	CursorUnexposedDecl                      CursorKind = 1
	CursorStructDecl                         CursorKind = 2
	CursorUnionDecl                          CursorKind = 3
	CursorClassDecl                          CursorKind = 4
	CursorEnumDecl                           CursorKind = 5
	CursorFieldDecl                          CursorKind = 6
	CursorEnumConstantDecl                   CursorKind = 7
	CursorFunctionDecl                       CursorKind = 8
	CursorVarDecl                            CursorKind = 9
	CursorParmDecl                           CursorKind = 10
	CursorObjCInterfaceDecl                  CursorKind = 11
	CursorObjCCategoryDecl                   CursorKind = 12
	CursorObjCProtocolDecl                   CursorKind = 13
	CursorObjCPropertyDecl                   CursorKind = 14
	CursorObjCIvarDecl                       CursorKind = 15
	CursorObjCInstanceMethodDecl             CursorKind = 16
	CursorObjCClassMethodDecl                CursorKind = 17
	CursorObjCImplementationDecl             CursorKind = 18
	CursorObjCCategoryImplDecl               CursorKind = 19
	CursorTypedefDecl                        CursorKind = 20
	CursorCXXMethod                          CursorKind = 21
	CursorNamespace                          CursorKind = 22
	CursorLinkageSpec                        CursorKind = 23
	CursorConstructor                        CursorKind = 24
	CursorDestructor                         CursorKind = 25
	CursorConversionFunction                 CursorKind = 26
	CursorTemplateTypeParameter              CursorKind = 27
	CursorNonTypeTemplateParameter           CursorKind = 28
	CursorTemplateTemplateParameter          CursorKind = 29
	CursorFunctionTemplate                   CursorKind = 30
	CursorClassTemplate                      CursorKind = 31
	CursorClassTemplatePartialSpecialization CursorKind = 32
	CursorNamespaceAlias                     CursorKind = 33
	CursorUsingDirective                     CursorKind = 34
	CursorUsingDeclaration                   CursorKind = 35
	CursorFirstDecl                          CursorKind = CursorUnexposedDecl
	CursorLastDecl                           CursorKind = CursorUsingDeclaration

	// References
	CursorFirstRef          CursorKind = 40
	CursorObjCSuperClassRef CursorKind = 40
	CursorObjCProtocolRef   CursorKind = 41
	CursorObjCClassRef      CursorKind = 42
	CursorTypeRef           CursorKind = 43
	CursorCXXBaseSpecifier  CursorKind = 44
	CursorTemplateRef       CursorKind = 45
	CursorNamespaceRef      CursorKind = 46
	CursorMemberRef         CursorKind = 47
	CursorLabelRef          CursorKind = 48
	CursorOverloadedDeclRef CursorKind = 49
	CursorLastRef           CursorKind = CursorOverloadedDeclRef

	// Error conditions
	CursorFirstInvalid   CursorKind = 70
	CursorInvalidFile    CursorKind = 70
	CursorNoDeclFound    CursorKind = 71
	CursorNotImplemented CursorKind = 72
	CursorInvalidCode    CursorKind = 73
	CursorLastInvalid    CursorKind = CursorInvalidCode

	// Expressions
	CursorFirstExpr       CursorKind = 100
	CursorUnexposedExpr   CursorKind = 100
	CursorDeclRefExpr     CursorKind = 101
	CursorMemberRefExpr   CursorKind = 102
	CursorCallExpr        CursorKind = 103
	CursorObjCMessageExpr CursorKind = 104
	CursorBlockExpr       CursorKind = 105
	CursorLastExpr        CursorKind = 105

	// Statements.
	CursorFirstStmt     CursorKind = 200
	CursorUnexposedStmt CursorKind = 200
	CursorLabelStmt     CursorKind = 201
	CursorLastStmt      CursorKind = CursorLabelStmt

	// The translation unit itself.
	CursorTranslationUnit CursorKind = 300

	// Attributes.
	CursorFirstAttr              CursorKind = 400
	CursorUnexposedAttr          CursorKind = 400
	CursorIBActionAttr           CursorKind = 401
	CursorIBOutletAttr           CursorKind = 402
	CursorIBOutletCollectionAttr CursorKind = 403
	CursorLastAttr               CursorKind = CursorIBOutletCollectionAttr

	// Preprocessing
	CursorFirstPreprocessing     CursorKind = CursorPreprocessingDirective
	CursorPreprocessingDirective CursorKind = 500
	CursorMacroDefinition        CursorKind = 501
	CursorMacroInstantiation     CursorKind = 502
	CursorInclusionDirective     CursorKind = 503
	CursorLastPreprocessing      CursorKind = CursorInclusionDirective
)

type TypeKind uint

const (
	TypeInvalid   TypeKind = 0
	TypeUnexposed TypeKind = 1
)

const (
	TypeVoid TypeKind = 2 + iota
	TypeBool
	TypeChar_U
	TypeUChar
	TypeChar16
	TypeChar32
	TypeUShort
	TypeUInt
	TypeULong
	TypeULongLong
	TypeUInt128
	TypeChar_S
	TypeSChar
	TypeWChar
	TypeShort
	TypeInt
	TypeLong
	TypeLongLong
	TypeInt128
	TypeFloat
	TypeDouble
	TypeLongDouble
	TypeNullPtr
	TypeOverload
	TypeDependent
	TypeObjCId
	TypeObjCClass
	TypeObjCSel
)

const (
	TypeFirstBuiltin TypeKind = TypeVoid
	TypeLastBuiltin  TypeKind = TypeObjCSel

	TypeComplex           TypeKind = 100
	TypePointer           TypeKind = 101
	TypeBlockPointer      TypeKind = 102
	TypeLValueReference   TypeKind = 103
	TypeRValueReference   TypeKind = 104
	TypeRecord            TypeKind = 105
	TypeEnum              TypeKind = 106
	TypeTypedef           TypeKind = 107
	TypeObjCInterface     TypeKind = 108
	TypeObjCObjectPointer TypeKind = 109
	TypeFunctionNoProto   TypeKind = 110
)

// Translation unit options
type TUOptions uint

const (
	TUDetailedPreprocessingRecord TUOptions = 1 << iota
	TUIncomplete
	TUPrecompiledPreamble
	TUCacheCompletionResults
	TUCXXPrecompiledPreamble
	TUCXXChainedPCH
)
