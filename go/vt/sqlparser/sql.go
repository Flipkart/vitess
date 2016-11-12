//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
func setParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func setAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func incNesting(yylex interface{}) bool {
	yylex.(*Tokenizer).nesting++
	if yylex.(*Tokenizer).nesting == 200 {
		return true
	}
	return false
}

func decNesting(yylex interface{}) {
	yylex.(*Tokenizer).nesting--
}

func forceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

//line sql.y:34
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	colTuple    ColTuple
	valExprs    ValExprs
	values      Values
	rowTuple    RowTuple
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
	colIdent    ColIdent
	colIdents   []ColIdent
	tableIdent  TableIdent
}

const LEX_ERROR = 57346
const UNION = 57347
const SELECT = 57348
const INSERT = 57349
const UPDATE = 57350
const DELETE = 57351
const FROM = 57352
const WHERE = 57353
const GROUP = 57354
const HAVING = 57355
const ORDER = 57356
const BY = 57357
const LIMIT = 57358
const FOR = 57359
const ALL = 57360
const DISTINCT = 57361
const AS = 57362
const EXISTS = 57363
const ASC = 57364
const DESC = 57365
const INTO = 57366
const DUPLICATE = 57367
const KEY = 57368
const DEFAULT = 57369
const SET = 57370
const LOCK = 57371
const VALUES = 57372
const LAST_INSERT_ID = 57373
const NEXT = 57374
const VALUE = 57375
const JOIN = 57376
const STRAIGHT_JOIN = 57377
const LEFT = 57378
const RIGHT = 57379
const INNER = 57380
const OUTER = 57381
const CROSS = 57382
const NATURAL = 57383
const USE = 57384
const FORCE = 57385
const ON = 57386
const ID = 57387
const STRING = 57388
const NUMBER = 57389
const VALUE_ARG = 57390
const LIST_ARG = 57391
const COMMENT = 57392
const NULL = 57393
const TRUE = 57394
const FALSE = 57395
const OR = 57396
const AND = 57397
const NOT = 57398
const BETWEEN = 57399
const CASE = 57400
const WHEN = 57401
const THEN = 57402
const ELSE = 57403
const LE = 57404
const GE = 57405
const NE = 57406
const NULL_SAFE_EQUAL = 57407
const IS = 57408
const LIKE = 57409
const REGEXP = 57410
const IN = 57411
const SHIFT_LEFT = 57412
const SHIFT_RIGHT = 57413
const UNARY = 57414
const INTERVAL = 57415
const END = 57416
const CREATE = 57417
const ALTER = 57418
const DROP = 57419
const RENAME = 57420
const ANALYZE = 57421
const TABLE = 57422
const INDEX = 57423
const VIEW = 57424
const TO = 57425
const IGNORE = 57426
const IF = 57427
const UNIQUE = 57428
const USING = 57429
const SHOW = 57430
const DESCRIBE = 57431
const EXPLAIN = 57432
const UNUSED = 57433

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"UNION",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"ASC",
	"DESC",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"VALUES",
	"LAST_INSERT_ID",
	"NEXT",
	"VALUE",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"'('",
	"','",
	"')'",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"NULL",
	"TRUE",
	"FALSE",
	"OR",
	"AND",
	"NOT",
	"BETWEEN",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"'='",
	"'<'",
	"'>'",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"IS",
	"LIKE",
	"REGEXP",
	"IN",
	"'|'",
	"'&'",
	"SHIFT_LEFT",
	"SHIFT_RIGHT",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'^'",
	"'~'",
	"UNARY",
	"INTERVAL",
	"'.'",
	"END",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
	"UNUSED",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 82,
	45, 222,
	89, 222,
	-2, 221,
}

const yyNprod = 226
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 829

var yyAct = [...]int{

	132, 155, 205, 404, 62, 223, 240, 304, 248, 343,
	47, 266, 250, 296, 85, 162, 249, 119, 333, 154,
	3, 234, 260, 120, 169, 252, 133, 63, 35, 74,
	37, 14, 67, 98, 38, 64, 48, 49, 69, 361,
	363, 71, 40, 41, 41, 247, 114, 50, 43, 44,
	45, 388, 387, 386, 68, 80, 70, 81, 58, 175,
	46, 42, 124, 330, 140, 141, 142, 143, 144, 139,
	100, 100, 173, 82, 93, 94, 95, 118, 100, 96,
	111, 125, 113, 282, 104, 139, 99, 105, 417, 147,
	148, 149, 150, 177, 188, 187, 82, 362, 187, 189,
	156, 107, 158, 273, 64, 86, 87, 64, 160, 189,
	165, 88, 189, 89, 235, 186, 271, 272, 270, 166,
	109, 101, 152, 184, 100, 97, 91, 60, 159, 180,
	257, 207, 65, 235, 164, 287, 315, 172, 174, 171,
	212, 213, 214, 215, 216, 217, 218, 219, 220, 221,
	60, 222, 176, 227, 125, 82, 125, 226, 106, 210,
	230, 232, 110, 211, 60, 208, 75, 188, 187, 224,
	244, 229, 142, 143, 144, 139, 239, 188, 187, 14,
	231, 245, 189, 332, 375, 125, 256, 227, 100, 100,
	204, 206, 189, 265, 185, 269, 274, 275, 276, 382,
	261, 263, 264, 28, 128, 262, 255, 319, 320, 321,
	413, 226, 278, 231, 226, 253, 151, 153, 100, 277,
	73, 60, 65, 302, 226, 268, 242, 334, 125, 106,
	280, 339, 226, 163, 284, 283, 163, 279, 64, 294,
	182, 226, 292, 281, 334, 226, 295, 372, 181, 151,
	258, 259, 243, 128, 128, 286, 291, 209, 14, 15,
	16, 17, 117, 318, 301, 79, 76, 356, 302, 322,
	156, 106, 357, 385, 324, 325, 326, 128, 237, 128,
	18, 354, 323, 329, 182, 238, 355, 328, 253, 128,
	331, 358, 384, 310, 311, 100, 353, 352, 288, 338,
	341, 344, 55, 14, 268, 336, 103, 254, 128, 340,
	345, 337, 390, 128, 128, 54, 349, 267, 351, 348,
	359, 350, 39, 371, 366, 102, 167, 290, 367, 410,
	369, 116, 297, 317, 381, 370, 253, 253, 253, 253,
	373, 411, 298, 19, 20, 22, 21, 23, 51, 52,
	241, 128, 135, 380, 57, 347, 24, 25, 26, 134,
	163, 128, 137, 145, 146, 140, 141, 142, 143, 144,
	139, 389, 61, 416, 402, 14, 28, 392, 344, 30,
	254, 1, 316, 396, 156, 393, 394, 391, 313, 90,
	183, 306, 309, 310, 311, 307, 267, 308, 312, 168,
	405, 405, 405, 64, 406, 407, 36, 408, 403, 246,
	412, 170, 414, 415, 418, 66, 293, 236, 419, 409,
	420, 29, 59, 376, 342, 379, 346, 285, 254, 254,
	254, 254, 72, 157, 233, 92, 77, 31, 32, 33,
	34, 84, 335, 395, 83, 397, 398, 289, 190, 126,
	360, 59, 305, 303, 251, 122, 108, 78, 53, 27,
	112, 56, 13, 115, 192, 195, 12, 11, 123, 10,
	197, 198, 199, 200, 201, 202, 203, 196, 193, 194,
	191, 138, 137, 145, 146, 140, 141, 142, 143, 144,
	139, 9, 59, 8, 161, 7, 225, 6, 131, 5,
	4, 2, 0, 0, 178, 0, 128, 179, 128, 128,
	0, 0, 399, 400, 401, 374, 0, 0, 0, 0,
	0, 0, 100, 59, 226, 82, 93, 94, 95, 0,
	0, 96, 129, 130, 0, 0, 127, 0, 99, 0,
	0, 123, 228, 123, 0, 138, 137, 145, 146, 140,
	141, 142, 143, 144, 139, 0, 0, 86, 87, 121,
	131, 0, 0, 88, 0, 89, 0, 0, 0, 0,
	0, 59, 123, 0, 0, 0, 228, 0, 91, 0,
	0, 0, 0, 0, 100, 0, 0, 82, 93, 94,
	95, 0, 0, 96, 129, 130, 0, 0, 127, 0,
	99, 59, 0, 0, 0, 0, 0, 14, 0, 377,
	378, 0, 0, 0, 0, 123, 0, 0, 0, 86,
	87, 121, 131, 0, 0, 88, 0, 89, 306, 309,
	310, 311, 307, 299, 308, 312, 300, 0, 383, 0,
	91, 0, 314, 0, 59, 0, 100, 0, 0, 82,
	93, 94, 95, 0, 0, 96, 129, 130, 0, 0,
	127, 0, 99, 138, 137, 145, 146, 140, 141, 142,
	143, 144, 139, 0, 0, 0, 0, 0, 0, 0,
	131, 86, 87, 0, 0, 0, 0, 88, 0, 89,
	0, 0, 59, 59, 59, 59, 0, 0, 0, 0,
	0, 0, 91, 0, 100, 364, 365, 82, 93, 94,
	95, 0, 0, 96, 129, 130, 0, 0, 127, 0,
	99, 100, 0, 0, 82, 93, 94, 95, 0, 0,
	96, 0, 0, 0, 0, 0, 0, 99, 0, 86,
	87, 0, 0, 0, 0, 88, 0, 89, 368, 0,
	0, 0, 0, 0, 0, 136, 86, 87, 0, 0,
	91, 327, 88, 65, 89, 0, 138, 137, 145, 146,
	140, 141, 142, 143, 144, 139, 0, 91, 0, 138,
	137, 145, 146, 140, 141, 142, 143, 144, 139, 0,
	0, 138, 137, 145, 146, 140, 141, 142, 143, 144,
	139, 138, 137, 145, 146, 140, 141, 142, 143, 144,
	139, 138, 137, 145, 146, 140, 141, 142, 143, 144,
	139, 145, 146, 140, 141, 142, 143, 144, 139,
}
var yyPact = [...]int{

	252, -1000, -1000, 371, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -68,
	-56, -35, -48, -36, -1000, -1000, -1000, 369, 330, 283,
	-1000, -57, 102, 362, 84, -69, -43, 84, -1000, -40,
	84, -1000, 102, -72, 118, -72, 102, -1000, -1000, -1000,
	-1000, -1000, -1000, 230, 676, -1000, 68, 301, 278, -5,
	-1000, 102, 112, -1000, 36, -1000, 102, 61, 114, -1000,
	102, -1000, -53, 102, 310, 218, 84, -1000, 539, -1000,
	342, 725, -1000, -1000, -1000, -1000, 676, 676, 676, 676,
	33, 144, -1000, -1000, -1000, -1000, -1000, 25, -1000, 676,
	-1000, -1000, 102, 84, 102, 349, 84, 676, -1000, 305,
	-79, -1000, 45, -1000, 102, -1000, -1000, 102, -1000, 238,
	-1000, -1000, 174, 26, 37, 405, -1000, 659, 601, -1000,
	-1000, 144, -1000, 102, -1000, -1000, 342, 676, 676, 676,
	676, 676, 676, 676, 676, 676, 676, -1000, -1000, -1000,
	715, 477, 107, 539, 198, 167, 735, 52, 735, 250,
	225, -1000, 336, 659, -1000, 735, -1000, -1000, 208, 84,
	-1000, -54, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 79, 539, -1000, -1000, 84, 48, 659, 659, 146,
	676, 143, 43, 676, 676, 676, 146, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 27, 405, 110, 405, -1000, 369,
	-1000, 102, 743, 285, -1000, 90, 90, 0, 0, 0,
	-16, -16, -1000, -1000, 194, 539, -1000, -1000, -6, 194,
	-1000, 676, -1000, 71, -1000, 659, 297, 84, 84, 336,
	316, 327, 37, 102, -1000, -1000, 102, -1000, 222, 357,
	-1000, -1000, 116, 313, 173, -1000, -1000, -1000, 27, 40,
	-1000, -1000, 153, -1000, -1000, 735, -1000, 25, -1000, -1000,
	143, 676, 676, 676, 735, 735, 703, -1000, -1000, -1000,
	-1000, 194, 84, -1000, 735, -27, -1000, 676, 120, 200,
	144, 371, 183, 185, -1000, 316, -1000, 676, 676, -1000,
	-1000, 343, 79, 79, 79, 79, -1000, 263, 262, -1000,
	247, 233, 257, -3, -1000, 102, 102, -1000, 177, -1000,
	-1000, -1000, 167, -1000, 735, 735, 690, 676, -1000, -1000,
	-1000, 735, 676, -1000, 298, 201, -1000, -1000, -1000, 84,
	-1000, 469, 138, -1000, 587, -1000, 340, 319, 357, 155,
	594, -1000, -1000, -1000, -1000, 258, -1000, 239, -1000, -1000,
	-1000, -44, -45, -46, -1000, -1000, -1000, -1000, 676, 735,
	735, 286, 144, -1000, 676, 676, -1000, -1000, -1000, 336,
	659, 676, 659, 659, -1000, -1000, 144, 144, 144, 735,
	366, -1000, 735, -1000, 316, 37, 134, 37, 37, 84,
	84, 84, 84, 312, 164, -1000, 164, 164, 112, -1000,
	365, 13, -1000, 84, -1000, -1000, -1000, 84, -1000, 84,
	-1000,
}
var yyPgo = [...]int{

	0, 501, 19, 500, 499, 497, 495, 493, 491, 469,
	467, 466, 462, 421, 461, 459, 458, 457, 17, 23,
	455, 8, 16, 12, 454, 453, 7, 452, 25, 450,
	3, 15, 62, 449, 448, 447, 444, 2, 22, 11,
	1, 442, 14, 33, 441, 435, 434, 21, 433, 427,
	426, 425, 6, 424, 9, 423, 13, 419, 417, 416,
	18, 4, 27, 26, 322, 220, 415, 411, 409, 406,
	399, 0, 390, 389, 388, 382, 10, 381, 379, 125,
	5,
}
var yyR1 = [...]int{

	0, 77, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 2, 2, 3, 3, 4,
	5, 6, 7, 7, 7, 8, 8, 8, 9, 10,
	10, 10, 11, 12, 12, 12, 78, 13, 14, 14,
	15, 15, 15, 16, 16, 17, 17, 18, 18, 19,
	19, 19, 20, 20, 72, 72, 72, 21, 21, 22,
	22, 23, 23, 23, 24, 24, 24, 24, 75, 75,
	74, 74, 74, 25, 25, 25, 25, 26, 26, 26,
	26, 27, 27, 28, 28, 29, 29, 29, 29, 30,
	30, 31, 31, 32, 32, 32, 32, 32, 32, 33,
	33, 33, 33, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 38, 38, 38, 38, 38, 38, 34, 34,
	34, 34, 34, 34, 34, 39, 39, 39, 43, 40,
	40, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 37, 45, 48, 48, 46, 46, 47, 49,
	49, 44, 44, 44, 36, 36, 36, 36, 50, 50,
	51, 51, 52, 52, 53, 53, 54, 55, 55, 55,
	56, 56, 56, 57, 57, 57, 58, 58, 59, 59,
	60, 60, 35, 35, 41, 41, 42, 42, 61, 61,
	62, 63, 63, 65, 65, 66, 66, 64, 64, 67,
	67, 67, 67, 67, 67, 68, 68, 69, 69, 70,
	70, 71, 73, 79, 80, 76,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 13, 6, 3, 7, 8, 8, 8,
	7, 3, 5, 8, 4, 6, 7, 4, 5, 4,
	5, 5, 3, 2, 2, 2, 0, 2, 0, 2,
	1, 2, 2, 0, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 0, 1, 2, 1, 3, 1,
	1, 3, 3, 3, 3, 5, 5, 3, 0, 1,
	0, 1, 2, 1, 2, 2, 1, 2, 3, 2,
	3, 2, 2, 1, 3, 0, 5, 5, 5, 1,
	3, 0, 2, 1, 3, 3, 2, 3, 3, 1,
	1, 3, 3, 4, 3, 4, 3, 4, 5, 6,
	3, 2, 1, 2, 1, 2, 1, 2, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 1, 3, 1,
	3, 1, 1, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 2, 3, 3, 4,
	5, 4, 1, 5, 0, 1, 1, 2, 4, 0,
	2, 1, 3, 5, 1, 1, 1, 1, 0, 3,
	0, 2, 0, 3, 1, 3, 2, 0, 1, 1,
	0, 2, 4, 0, 2, 4, 0, 3, 1, 3,
	0, 5, 2, 1, 1, 3, 3, 1, 1, 3,
	3, 1, 1, 0, 2, 0, 3, 0, 1, 1,
	1, 1, 1, 1, 1, 0, 1, 0, 1, 0,
	2, 1, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -77, -1, -2, -3, -4, -5, -6, -7, -8,
	-9, -10, -11, -12, 6, 7, 8, 9, 28, 91,
	92, 94, 93, 95, 104, 105, 106, -15, 5, -13,
	-78, -13, -13, -13, -13, 96, -69, 98, 102, -64,
	98, 100, 96, 96, 97, 98, 96, -76, -76, -76,
	-2, 18, 19, -16, 32, 19, -14, -64, -28, -73,
	48, 10, -61, -62, -71, 48, -66, 101, 97, -71,
	96, -71, -73, -65, 101, 48, -65, -73, -17, 35,
	-71, -37, 48, -36, -44, -42, 80, 81, 86, 88,
	-73, 101, -45, 49, 50, 51, 54, -79, -43, 61,
	45, 53, 24, 28, 89, -28, 46, 65, -73, 59,
	48, -76, -73, -76, 99, -73, 21, 44, -71, -18,
	-19, 82, -20, -73, -32, -37, -33, 59, -79, 55,
	56, 21, -71, -63, 17, 10, 30, 77, 76, 85,
	80, 81, 82, 83, 84, 78, 79, -37, -37, -37,
	-37, -79, 89, -79, -2, -40, -37, -48, -37, -28,
	-61, -73, -31, 11, -62, -37, -76, 21, -70, 103,
	-67, 94, 92, 27, 93, 14, 107, 48, -73, -73,
	-76, 10, 46, -72, -71, 20, 89, 58, 57, 72,
	-34, 75, 59, 73, 74, 60, 72, 65, 66, 67,
	68, 69, 70, 71, -32, -37, -32, -37, -43, -79,
	-28, -63, -37, -37, -37, -37, -37, -37, -37, -37,
	-37, -37, -71, -80, -18, 19, 47, -71, -73, -18,
	-80, 46, -80, -46, -47, 62, -58, 28, -79, -31,
	-52, 14, -32, 44, -71, -76, -68, 99, -21, -22,
	-23, -24, -28, -43, -79, -19, -71, 82, -32, -32,
	-38, 54, 59, 55, 56, -37, -39, -79, -43, 52,
	75, 73, 74, 60, -37, -37, -37, -38, -80, -28,
	-80, -18, 89, -80, -37, -49, -47, 64, -32, -35,
	30, -2, -61, -59, -71, -52, -56, 16, 15, -73,
	-73, -31, 46, -25, -26, -27, 34, 38, 40, 35,
	36, 37, 41, -74, -73, 20, -75, 20, -21, 54,
	55, 56, -40, -39, -37, -37, -37, 58, -80, -71,
	90, -37, 63, -60, 44, -41, -42, -60, -80, 46,
	-56, -37, -53, -54, -37, -76, -50, 12, -22, -23,
	-22, -23, 34, 34, 34, 39, 34, 39, 34, -26,
	-29, 42, 100, 43, -73, -73, -80, -80, 58, -37,
	-37, 25, 46, -71, 46, 46, -55, 22, 23, -51,
	13, 15, 44, 44, 34, 34, 97, 97, 97, -37,
	26, -42, -37, -54, -52, -32, -40, -32, -32, -79,
	-79, -79, 8, -56, -30, -71, -30, -30, -61, -57,
	17, 29, -80, 46, -80, -80, 8, 75, -71, -71,
	-71,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 36, 36, 36, 36, 36, 217,
	207, 0, 0, 0, 225, 225, 225, 0, 40, 43,
	38, 207, 0, 0, 0, 205, 0, 0, 218, 0,
	0, 208, 0, 203, 0, 203, 0, 33, 34, 35,
	15, 41, 42, 45, 0, 44, 37, 0, 0, 83,
	222, 0, 21, 198, 0, 221, 0, 0, 0, 225,
	0, 225, 0, 0, 0, 0, 0, 32, 0, 46,
	161, 0, -2, 131, 132, 133, 0, 0, 0, 0,
	0, 0, 152, 164, 165, 166, 167, 0, 197, 154,
	223, 39, 0, 0, 0, 91, 0, 0, 225, 0,
	219, 24, 0, 27, 0, 29, 204, 0, 225, 0,
	47, 49, 54, 0, 52, 53, 93, 0, 0, 99,
	100, 0, 161, 0, 201, 202, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 144, 145, 146,
	0, 0, 0, 0, 0, 0, 129, 0, 155, 186,
	91, 84, 172, 0, 199, 200, 22, 206, 0, 0,
	225, 215, 209, 210, 211, 212, 213, 214, 28, 30,
	31, 0, 0, 50, 55, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 118, 119, 120,
	121, 122, 123, 124, 96, 0, 0, 129, 111, 0,
	14, 0, 134, 135, 136, 137, 138, 139, 140, 141,
	142, 143, 147, 148, 0, 0, 224, 162, 0, 0,
	128, 0, 196, 159, 156, 0, 0, 0, 0, 172,
	180, 0, 92, 0, 220, 25, 0, 216, 91, 57,
	59, 60, 70, 68, 0, 48, 56, 51, 94, 95,
	98, 112, 0, 114, 116, 101, 102, 0, 126, 127,
	0, 0, 0, 0, 104, 106, 0, 110, 97, 16,
	149, 0, 0, 151, 130, 0, 157, 0, 0, 190,
	0, 193, 190, 0, 188, 180, 20, 0, 0, 225,
	26, 168, 0, 0, 0, 0, 73, 0, 0, 76,
	0, 0, 0, 85, 71, 0, 0, 69, 0, 113,
	115, 117, 0, 103, 105, 107, 0, 0, 150, 163,
	153, 160, 0, 17, 0, 192, 194, 18, 187, 0,
	19, 181, 173, 174, 177, 23, 170, 0, 58, 64,
	0, 67, 74, 75, 77, 0, 79, 0, 81, 82,
	61, 0, 0, 0, 72, 62, 63, 125, 0, 108,
	158, 0, 0, 189, 0, 0, 176, 178, 179, 172,
	0, 0, 0, 0, 78, 80, 0, 0, 0, 109,
	0, 195, 182, 175, 180, 171, 169, 65, 66, 0,
	0, 0, 0, 183, 0, 89, 0, 0, 191, 13,
	0, 0, 86, 0, 87, 88, 184, 0, 90, 0,
	185,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 84, 77, 3,
	45, 47, 82, 80, 46, 81, 89, 83, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	66, 65, 67, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 85, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 76, 3, 86,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	68, 69, 70, 71, 72, 73, 74, 75, 78, 79,
	87, 88, 90, 91, 92, 93, 94, 95, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:171
		{
			setParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:177
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 13:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line sql.y:193
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, Hints: yyDollar[4].str, SelectExprs: yyDollar[5].selectExprs, From: yyDollar[7].tableExprs, Where: NewWhere(WhereStr, yyDollar[8].boolExpr), GroupBy: GroupBy(yyDollar[9].valExprs), Having: NewWhere(HavingStr, yyDollar[10].boolExpr), OrderBy: yyDollar[11].orderBy, Limit: yyDollar[12].limit, Lock: yyDollar[13].str}
		}
	case 14:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:197
		{
			if yyDollar[4].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), SelectExprs: SelectExprs{Nextval{}}, From: TableExprs{&AliasedTableExpr{Expr: yyDollar[6].tableName}}}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:205
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:209
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), SelectExprs: SelectExprs{Nextval{NextCount: yyDollar[4].valExpr}}, From: TableExprs{&AliasedTableExpr{Expr: yyDollar[7].tableName}}}
		}
	case 17:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:215
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 18:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:219
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, updateList := range yyDollar[7].updateExprs {
				cols = append(cols, updateList.Name)
				vals = append(vals, updateList.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 19:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:231
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:237
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:243
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 22:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:249
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: yyDollar[4].tableIdent}
		}
	case 23:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:253
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[7].tableIdent, NewName: yyDollar[7].tableIdent}
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:258
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: TableIdent(yyDollar[3].colIdent.Lowered())}
		}
	case 25:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:264
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableIdent, NewName: yyDollar[4].tableIdent}
		}
	case 26:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:268
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[4].tableIdent, NewName: yyDollar[7].tableIdent}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:273
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: TableIdent(yyDollar[3].colIdent.Lowered()), NewName: TableIdent(yyDollar[3].colIdent.Lowered())}
		}
	case 28:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:279
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].tableIdent, NewName: yyDollar[5].tableIdent}
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:285
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: yyDollar[4].tableIdent, IfExists: exists}
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:293
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[5].tableIdent, NewName: yyDollar[5].tableIdent}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:298
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: TableIdent(yyDollar[4].colIdent.Lowered()), IfExists: exists}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:308
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].tableIdent, NewName: yyDollar[3].tableIdent}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:314
		{
			yyVAL.statement = &Other{}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:318
		{
			yyVAL.statement = &Other{}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:322
		{
			yyVAL.statement = &Other{}
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:327
		{
			setAllowComments(yylex, true)
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:331
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:337
		{
			yyVAL.bytes2 = nil
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:341
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:347
		{
			yyVAL.str = UnionStr
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:351
		{
			yyVAL.str = UnionAllStr
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:355
		{
			yyVAL.str = UnionDistinctStr
		}
	case 43:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:360
		{
			yyVAL.str = ""
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:364
		{
			yyVAL.str = DistinctStr
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:369
		{
			yyVAL.str = ""
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:373
		{
			yyVAL.str = StraightJoinHint
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:379
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:383
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:389
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:393
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:397
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].tableIdent}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:403
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:407
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:412
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:416
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:420
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:426
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:430
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:440
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:444
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:448
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:461
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:465
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 66:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:469
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:473
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 68:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:478
		{
			yyVAL.empty = struct{}{}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:480
		{
			yyVAL.empty = struct{}{}
		}
	case 70:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:483
		{
			yyVAL.tableIdent = ""
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:487
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:491
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:497
		{
			yyVAL.str = JoinStr
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:501
		{
			yyVAL.str = JoinStr
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:505
		{
			yyVAL.str = JoinStr
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:509
		{
			yyVAL.str = StraightJoinStr
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:515
		{
			yyVAL.str = LeftJoinStr
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:519
		{
			yyVAL.str = LeftJoinStr
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:523
		{
			yyVAL.str = RightJoinStr
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:527
		{
			yyVAL.str = RightJoinStr
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:533
		{
			yyVAL.str = NaturalJoinStr
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:537
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:547
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:551
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}
		}
	case 85:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:556
		{
			yyVAL.indexHints = nil
		}
	case 86:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:560
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 87:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:564
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 88:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:568
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:574
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:578
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:583
		{
			yyVAL.boolExpr = nil
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:587
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:594
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:598
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:602
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:606
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:610
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:616
		{
			yyVAL.boolExpr = BoolVal(true)
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:620
		{
			yyVAL.boolExpr = BoolVal(false)
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:624
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:628
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:632
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:636
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:640
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:644
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 107:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:648
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 108:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:652
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 109:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:656
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:660
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:664
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:670
		{
			yyVAL.str = IsNullStr
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:674
		{
			yyVAL.str = IsNotNullStr
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:678
		{
			yyVAL.str = IsTrueStr
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:682
		{
			yyVAL.str = IsNotTrueStr
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:686
		{
			yyVAL.str = IsFalseStr
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:690
		{
			yyVAL.str = IsNotFalseStr
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:696
		{
			yyVAL.str = EqualStr
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:700
		{
			yyVAL.str = LessThanStr
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:704
		{
			yyVAL.str = GreaterThanStr
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:708
		{
			yyVAL.str = LessEqualStr
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:712
		{
			yyVAL.str = GreaterEqualStr
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:716
		{
			yyVAL.str = NotEqualStr
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:720
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:726
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:730
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:734
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:740
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:746
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:750
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:756
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:760
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:764
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:768
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:772
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:776
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:780
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:784
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:788
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:792
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:796
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:800
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:804
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:808
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:816
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				// Handle double negative
				if num[0] == '-' {
					yyVAL.valExpr = num[1:]
				} else {
					yyVAL.valExpr = append(NumVal("-"), num...)
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UMinusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:829
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:833
		{
			// This rule prevents the usage of INTERVAL
			// as a function. If support is needed for that,
			// we'll need to revisit this. The solution
			// will be non-trivial because of grammar conflicts.
			yyVAL.valExpr = &IntervalExpr{Expr: yyDollar[2].valExpr, Unit: yyDollar[3].colIdent}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:841
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent)}
		}
	case 149:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:845
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent), Exprs: yyDollar[3].selectExprs}
		}
	case 150:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:849
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent), Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 151:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:853
		{
			yyVAL.valExpr = &FuncExpr{Name: "if", Exprs: yyDollar[3].selectExprs}
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:857
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 153:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:863
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:868
		{
			yyVAL.valExpr = nil
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:872
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:878
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 157:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:882
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 158:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:888
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:893
		{
			yyVAL.valExpr = nil
		}
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:897
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:903
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:907
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Name: yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 163:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:911
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}, Name: yyDollar[5].colIdent}
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:917
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:921
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:925
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:929
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 168:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:934
		{
			yyVAL.valExprs = nil
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:938
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 170:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:943
		{
			yyVAL.boolExpr = nil
		}
	case 171:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:947
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 172:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:952
		{
			yyVAL.orderBy = nil
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:956
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:962
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:966
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 176:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:972
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 177:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:977
		{
			yyVAL.str = AscScr
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:981
		{
			yyVAL.str = AscScr
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:985
		{
			yyVAL.str = DescScr
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:990
		{
			yyVAL.limit = nil
		}
	case 181:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:994
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 182:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:998
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 183:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1003
		{
			yyVAL.str = ""
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1007
		{
			yyVAL.str = ForUpdateStr
		}
	case 185:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1011
		{
			if yyDollar[3].colIdent.Lowered() != "share" {
				yylex.Error("expecting share")
				return 1
			}
			if yyDollar[4].colIdent.Lowered() != "mode" {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = ShareModeStr
		}
	case 186:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1024
		{
			yyVAL.columns = nil
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1028
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1034
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 189:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1038
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 190:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1043
		{
			yyVAL.updateExprs = nil
		}
	case 191:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1047
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1053
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1057
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1063
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 195:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1067
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1073
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1077
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1083
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1087
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 200:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1093
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 203:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1102
		{
			yyVAL.byt = 0
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1104
		{
			yyVAL.byt = 1
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1107
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1109
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1112
		{
			yyVAL.str = ""
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1114
		{
			yyVAL.str = IgnoreStr
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1118
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1120
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1122
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1124
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1126
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1128
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1131
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1133
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1136
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1138
		{
			yyVAL.empty = struct{}{}
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1141
		{
			yyVAL.empty = struct{}{}
		}
	case 220:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1143
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1147
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1153
		{
			yyVAL.tableIdent = TableIdent(yyDollar[1].bytes)
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1159
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1168
		{
			decNesting(yylex)
		}
	case 225:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1173
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
