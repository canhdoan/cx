package base

import (
	"os"
)

// global reference to our program
var PROGRAM *CXProgram

var CXPATH string = os.Getenv("CXPATH") + "/"
var BINPATH string = CXPATH + "bin/"
var PKGPATH string = CXPATH + "pkg/"
var SRCPATH string = CXPATH + "src/"

const MAIN_FUNC = "main"
const SYS_INIT_FUNC = "*init"
const MAIN_PKG = "main"
const NON_ASSIGN_PREFIX = "nonAssign"
const LOCAL_PREFIX = "*lcl"
const CORE_MODULE = "core"
const ID_FN = "identity"
const INIT_FN = "initDef"
const SLICE_SIZE = 32
const MARK_SIZE = 1
const OBJECT_HEADER_SIZE = 9
const FORWARDING_ADDRESS_SIZE = 4
const OBJECT_SIZE = 4
const CALLSTACK_SIZE = 100
const STACK_SIZE = 300000
const INIT_HEAP_SIZE = 300000
const NULL_ADDRESS = STACK_SIZE
const NULL_HEAP_ADDRESS_OFFSET = 4
const NULL_HEAP_ADDRESS = 0
const STR_HEADER_SIZE = 4
const TYPE_POINTER_SIZE = 4
const SLICE_HEADER_SIZE = 8
const MEMORY_SIZE = STACK_SIZE + INIT_HEAP_SIZE + TYPE_POINTER_SIZE

const MAX_UINT32 = ^uint32(0)
const MIN_UINT32 = 0
const MAX_INT32 = int(MAX_UINT32 >> 1)
const MIN_INT32 = -MAX_INT32 - 1

var BASIC_TYPES []string = []string{
	"bool", "str", "byte", "i32", "i64", "f32", "f64",
	"[]bool", "[]str", "[]byte", "[]i32", "[]i64", "[]f32", "[]f64",
}
var NATIVE_FUNCTIONS = map[string]bool{
	"i32.add": true, "i32.mul": true, "i32.sub": true, "i32.div": true,
	"i64.add": true, "i64.mul": true, "i64.sub": true, "i64.div": true,
	"f32.add": true, "f32.mul": true, "f32.sub": true, "f32.div": true,
	"f64.add": true, "f64.mul": true, "f64.sub": true, "f64.div": true,
	"i32.abs": true, "i64.abs": true, "f32.abs": true, "f64.abs": true,
	"i32.mod": true, "i64.mod": true,
	"i32.pow": true, "i64.pow": true, "f32.pow": true, "f64.pow": true,
	"f32.cos": true, "f32.sin": true, "f64.cos": true, "f64.sin": true,
	"i32.bitand": true, "i32.bitor": true, "i32.bitxor": true, "i32.bitclear": true,
	"i64.bitand": true, "i64.bitor": true, "i64.bitxor": true, "i64.bitclear": true,
	"i32.bitshl": true, "i32.bitshr": true, "i64.bitshl": true, "i64.bitshr": true,

	"str.print": true, "byte.print": true, "i32.print": true, "i64.print": true,
	"f32.print": true, "f64.print": true, "[]byte.print": true, "[]i32.print": true,
	"[]i64.print": true, "[]f32.print": true, "[]f64.print": true, "bool.print": true,
	"[]bool.print": true, "[]str.print": true,

	"str.id": true, "bool.id": true, "byte.id": true, "i32.id": true, "i64.id": true, "f32.id": true, "f64.id": true,
	"[]bool.id": true, "[]byte.id": true, "[]str.id": true, "[]i32.id": true, "[]i64.id": true, "[]f32.id": true, "[]f64.id": true,
	"identity": true,

	"[]bool.make": true, "[]byte.make": true, "[]str.make": true,
	"[]i32.make": true, "[]i64.make": true, "[]f32.make": true, "[]f64.make": true,

	"[]bool.read": true, "[]bool.write": true, "[]byte.read": true, "[]byte.write": true,
	"[]str.read": true, "[]str.write": true, "[]i32.read": true, "[]i32.write": true,
	"[]i64.read": true, "[]i64.write": true,
	"[]f32.read": true, "[]f32.write": true, "[]f64.read": true, "[]f64.write": true,

	"[]bool.len": true, "[]byte.len": true, "[]i32.len": true, "[]i64.len": true,
	"[]f32.len": true, "[]f64.len": true, "[]str.len": true, "str.len": true,

	"str.concat": true, "[]byte.concat": true, "[]bool.concat": true, "[]str.concat": true,
	"[]i32.concat": true, "[]i64.concat": true, "[]f32.concat": true, "[]f64.concat": true,

	"[]byte.append": true, "[]bool.append": true, "[]str.append": true,
	"[]i32.append": true, "[]i64.append": true, "[]f32.append": true, "[]f64.append": true,

	"[]byte.copy": true, "[]bool.copy": true, "[]str.copy": true,
	"[]i32.copy": true, "[]i64.copy": true, "[]f32.copy": true, "[]f64.copy": true,

	"[]byte.str": true, "str.[]byte": true,

	"byte.i32": true, "byte.i64": true, "byte.f32": true, "byte.f64": true,
	"[]byte.[]i32": true, "[]byte.[]i64": true, "[]byte.[]f32": true, "[]byte.[]f64": true,

	"i32.byte": true, "i64.byte": true, "f32.byte": true, "f64.byte": true,
	"[]i32.[]byte": true, "[]i64.[]byte": true, "[]f32.[]byte": true, "[]f64.[]byte": true,

	"i64.i32": true, "f32.i32": true, "f64.i32": true,
	"i32.i64": true, "f32.i64": true, "f64.i64": true,
	"i32.f32": true, "i64.f32": true, "f64.f32": true,
	"i32.f64": true, "i64.f64": true, "f32.f64": true,

	"byte.str": true, "bool.str": true, "i32.str": true,
	"i64.str": true, "f32.str": true, "f64.str": true,

	"[]i64.[]i32": true, "[]f32.[]i32": true, "[]f64.[]i32": true,
	"[]i32.[]i64": true, "[]f32.[]i64": true, "[]f64.[]i64": true,
	"[]i32.[]f32": true, "[]i64.[]f32": true, "[]f64.[]f32": true,
	"[]i32.[]f64": true, "[]i64.[]f64": true, "[]f32.[]f64": true,

	"bool.eq": true, "bool.uneq": true,
	"i32.lt": true, "i32.gt": true, "i32.eq": true, "i32.uneq": true, "i32.lteq": true, "i32.gteq": true,
	"i64.lt": true, "i64.gt": true, "i64.eq": true, "i64.uneq": true, "i64.lteq": true, "i64.gteq": true,
	"f32.lt": true, "f32.gt": true, "f32.eq": true, "f32.uneq": true, "f32.lteq": true, "f32.gteq": true,
	"f64.lt": true, "f64.gt": true, "f64.eq": true, "f64.uneq": true, "f64.lteq": true, "f64.gteq": true,
	"str.lt": true, "str.gt": true, "str.eq": true, "str.uneq": true, "str.lteq": true, "str.gteq": true,
	"byte.lt": true, "byte.gt": true, "byte.eq": true, "byte.uneq": true, "byte.lteq": true, "byte.gteq": true,

	"str.read": true, "i32.read": true,

	"i32.rand": true, "i64.rand": true,

	"and": true, "or": true, "not": true,
	"sleep": true, "halt": true, "goTo": true, "baseGoTo": true,

	"remExpr": true, "remArg": true, "addExpr": true,
	"rem.expr": true, "rem.arg": true, "add.expr": true,

	"aff.query": true, "aff.execute": true, "aff.print": true, "aff.concat": true,
	"aff.len": true, "aff.index": true, "aff.name": true,

	//"serialize":true, "deserialize":true,
	"evolve": true,

	"initDef": true,

	"test.start": true, "test.stop": true,
	"test.error": true, "test.bool": true, "test.str": true, "test.byte": true,
	"test.i32": true, "test.i64": true, "test.f32": true, "test.f64": true,
	"test.[]bool": true, "test.[]byte": true, "test.[]str": true, "test.[]i32": true,
	"test.[]i64": true, "test.[]f32": true, "test.[]f64": true,

	"mdim.append": true, "mdim.read": true, "mdim.write": true, "mdim.len": true,
	"mdim.make": true,

	"cstm.append": true, "cstm.read": true, "cstm.write": true, "cstm.len": true,
	"cstm.make": true, "cstm.serialize": true, "cstm.deserialize": true,

	/*
	   Time
	*/
	"time.Unix": true, "time.UnixMilli": true, "time.UnixNano": true,

	/*
	   Runtime
	*/

	"runtime.LockOSThread": true,

	/*
	   GLText
	*/

	"gltext.LoadTrueType": true, "gltext.Printf": true,

	/*
	   OpenGL
	*/
	"gl.Init": true, "gl.GetError": true, "gl.BindAttribLocation": true, "gl.GetAttribLocation": true,
	"gl.CullFace": true, "gl.CreateProgram": true, "gl.DeleteProgram": true, "gl.LinkProgram": true,
	"gl.Clear": true, "gl.UseProgram": true,

	"gl.BindBuffer": true, "gl.BindVertexArray": true, "gl.EnableVertexAttribArray": true,
	"gl.VertexAttribPointer": true, "gl.DrawArrays": true, "gl.GenBuffers": true,
	"gl.BufferData": true, "gl.GenVertexArrays": true,
	"gl.CreateShader": true, "gl.DetachShader": true, "gl.DeleteShader": true,

	"gl.Strs": true, "gl.Free": true, "gl.ShaderSource": true,
	"gl.CompileShader": true, "gl.GetShaderiv": true, "gl.AttachShader": true,

	"gl.MatrixMode": true,
	"gl.Rotatef":    true, "gl.Translatef": true, "gl.LoadIdentity": true,
	"gl.PushMatrix": true, "gl.PopMatrix": true, "gl.EnableClientState": true,

	"gl.BindTexture": true, "gl.Color3f": true, "gl.Color4f": true, "gl.Begin": true,
	"gl.End": true, "gl.Normal3f": true, "gl.TexCoord2f": true,
	"gl.Vertex2f": true, "gl.Vertex3f": true,

	"gl.Enable": true, "gl.ClearColor": true, "gl.ClearDepth": true,
	"gl.DepthFunc": true, "gl.Lightfv": true, "gl.Frustum": true,
	"gl.Disable": true, "gl.Hint": true,

	"gl.NewTexture": true, "gl.DepthMask": true, "gl.TexEnvi": true,
	"gl.BlendFunc": true, "gl.Ortho": true, "gl.Viewport": true,
	"gl.Scalef": true, "gl.TexCoord2d": true,

	/*
	   GLFW
	*/

	"glfw.Init": true, "glfw.WindowHint": true, "glfw.CreateWindow": true,
	"glfw.MakeContextCurrent": true, "glfw.ShouldClose": true, "glfw.SetShouldClose": true,
	"glfw.PollEvents": true, "glfw.SwapBuffers": true, "glfw.GetFramebufferSize": true,
	"glfw.SwapInterval": true,

	"glfw.SetKeyCallback": true, "glfw.GetTime": true, "glfw.SetMouseButtonCallback": true,
	"glfw.SetCursorPosCallback": true, "glfw.GetCursorPos": true, "glfw.SetInputMode": true,
	"glfw.GetKey": true,

	/*
	   Operating System
	*/

	"os.Create": true, "os.Open": true, "os.Close": true, "os.GetWorkingDirectory": true,
	"os.Write": true, "os.WriteFile": true, "os.ReadFile": true,
}

const (
	DECL_POINTER = iota // 0
	DECL_ARRAY          // 1
	DECL_SLICE          // 2
	DECL_STRUCT         // 3
	DECL_BASIC          // 4
)

// what to write
const (
	PASSBY_VALUE = iota
	PASSBY_REFERENCE
)

const (
	DEREF_ARRAY = iota
	DEREF_FIELD
	DEREF_POINTER
	DEREF_DEREF
)

const (
	TYPE_AFF = iota
	TYPE_BOOL
	TYPE_BYTE
	TYPE_STR
	TYPE_F32
	TYPE_F64
	TYPE_I8
	TYPE_I16
	TYPE_I32
	TYPE_I64
	TYPE_UI8
	TYPE_UI16
	TYPE_UI32
	TYPE_UI64

	TYPE_THRESHOLD

	TYPE_UNDEFINED
	TYPE_CUSTOM
	TYPE_POINTER
	TYPE_IDENTIFIER
)

var TypeCounter int
var TypeCodes map[string]int = map[string]int{
	"identifier": TYPE_IDENTIFIER,
	"aff":        TYPE_AFF,
	"bool":       TYPE_BOOL,
	"byte":       TYPE_BYTE,
	"str":        TYPE_STR,
	"f32":        TYPE_F32,
	"f64":        TYPE_F64,
	"i8":         TYPE_I8,
	"i16":        TYPE_I16,
	"i32":        TYPE_I32,
	"i64":        TYPE_I64,
	"ui8":        TYPE_UI8,
	"ui16":       TYPE_UI16,
	"ui32":       TYPE_UI32,
	"ui64":       TYPE_UI64,
	"und":        TYPE_UNDEFINED,
}

var TypeNames map[int]string = map[int]string{
	TYPE_IDENTIFIER: "ident",
	TYPE_AFF:        "aff",
	TYPE_BOOL:       "bool",
	TYPE_BYTE:       "byte",
	TYPE_STR:        "str",
	TYPE_F32:        "f32",
	TYPE_F64:        "f64",
	TYPE_I8:         "i8",
	TYPE_I16:        "i16",
	TYPE_I32:        "i32",
	TYPE_I64:        "i64",
	TYPE_UI8:        "ui8",
	TYPE_UI16:       "ui16",
	TYPE_UI32:       "ui32",
	TYPE_UI64:       "ui64",
	TYPE_UNDEFINED:  "und",
}

// memory locations
const (
	MEM_STACK = iota
	MEM_HEAP
	MEM_DATA
)
