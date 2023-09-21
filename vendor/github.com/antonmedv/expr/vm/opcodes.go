package vm

type Opcode byte

const (
<<<<<<< HEAD
	OpPush Opcode = iota
	OpPushInt
	OpPop
=======
	OpInvalid Opcode = iota
	OpPush
	OpInt
	OpPop
	OpStore
	OpLoadVar
>>>>>>> main
	OpLoadConst
	OpLoadField
	OpLoadFast
	OpLoadMethod
	OpLoadFunc
<<<<<<< HEAD
=======
	OpLoadEnv
>>>>>>> main
	OpFetch
	OpFetchField
	OpMethod
	OpTrue
	OpFalse
	OpNil
	OpNegate
	OpNot
	OpEqual
	OpEqualInt
	OpEqualString
	OpJump
	OpJumpIfTrue
	OpJumpIfFalse
	OpJumpIfNil
	OpJumpIfNotNil
	OpJumpIfEnd
	OpJumpBackward
	OpIn
	OpLess
	OpMore
	OpLessOrEqual
	OpMoreOrEqual
	OpAdd
	OpSubtract
	OpMultiply
	OpDivide
	OpModulo
	OpExponent
	OpRange
	OpMatches
	OpMatchesConst
	OpContains
	OpStartsWith
	OpEndsWith
	OpSlice
	OpCall
	OpCall0
	OpCall1
	OpCall2
	OpCall3
	OpCallN
	OpCallFast
	OpCallTyped
<<<<<<< HEAD
	OpBuiltin
=======
	OpCallBuiltin1
>>>>>>> main
	OpArray
	OpMap
	OpLen
	OpCast
	OpDeref
<<<<<<< HEAD
	OpIncrementIt
	OpIncrementCount
	OpGetCount
	OpGetLen
	OpPointer
=======
	OpIncrementIndex
	OpDecrementIndex
	OpIncrementCount
	OpGetIndex
	OpSetIndex
	OpGetCount
	OpGetLen
	OpGetGroupBy
	OpGetAcc
	OpPointer
	OpThrow
	OpGroupBy
	OpSetAcc
>>>>>>> main
	OpBegin
	OpEnd // This opcode must be at the end of this list.
)
