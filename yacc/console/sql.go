// Code generated by goyacc -o yacc/console/sql.go -p yy yacc/console/sql.y. DO NOT EDIT.

//line yacc/console/sql.y:3

package spqrparser

import __yyfmt__ "fmt"

//line yacc/console/sql.y:4

//line yacc/console/sql.y:11
type yySymType struct {
	yys   int
	str   string
	byte  byte
	bytes []byte
	int   int
	bool  bool
	empty struct{}

	statement Statement
	show      *Show

	drop   *Drop
	create *Create

	kill   *Kill
	lock   *Lock
	unlock *Unlock

	ds            *DataspaceDefinition
	kr            *KeyRangeDefinition
	shard         *ShardDefinition
	sharding_rule *ShardingRuleDefinition

	register_router   *RegisterRouter
	unregister_router *UnregisterRouter

	split *SplitKeyRange
	move  *MoveKeyRange
	unite *UniteKeyRange

	shutdown *Shutdown
	listen   *Listen

	entrieslist []ShardingRuleEntry
	shruleEntry ShardingRuleEntry

	sharding_rule_selector *ShardingRuleSelector
	key_range_selector     *KeyRangeSelector
}

const STRING = 57346
const COMMAND = 57347
const SHOW = 57348
const KILL = 57349
const POOLS = 57350
const STATS = 57351
const LISTS = 57352
const SERVERS = 57353
const CLIENTS = 57354
const DATABASES = 57355
const SHUTDOWN = 57356
const LISTEN = 57357
const REGISTER = 57358
const UNREGISTER = 57359
const ROUTER = 57360
const ROUTE = 57361
const CREATE = 57362
const ADD = 57363
const DROP = 57364
const LOCK = 57365
const UNLOCK = 57366
const SPLIT = 57367
const MOVE = 57368
const COMPOSE = 57369
const SHARDING = 57370
const COLUMN = 57371
const TABLE = 57372
const HASH = 57373
const FUNCTION = 57374
const KEY = 57375
const RANGE = 57376
const DATASPACE = 57377
const SHARDS = 57378
const KEY_RANGES = 57379
const ROUTERS = 57380
const SHARD = 57381
const HOST = 57382
const SHARDING_RULES = 57383
const RULE = 57384
const COLUMNS = 57385
const BY = 57386
const FROM = 57387
const TO = 57388
const WITH = 57389
const UNITE = 57390
const ALL = 57391
const ADDRESS = 57392

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"STRING",
	"COMMAND",
	"SHOW",
	"KILL",
	"POOLS",
	"STATS",
	"LISTS",
	"SERVERS",
	"CLIENTS",
	"DATABASES",
	"SHUTDOWN",
	"LISTEN",
	"REGISTER",
	"UNREGISTER",
	"ROUTER",
	"ROUTE",
	"CREATE",
	"ADD",
	"DROP",
	"LOCK",
	"UNLOCK",
	"SPLIT",
	"MOVE",
	"COMPOSE",
	"SHARDING",
	"COLUMN",
	"TABLE",
	"HASH",
	"FUNCTION",
	"KEY",
	"RANGE",
	"DATASPACE",
	"SHARDS",
	"KEY_RANGES",
	"ROUTERS",
	"SHARD",
	"HOST",
	"SHARDING_RULES",
	"RULE",
	"COLUMNS",
	"BY",
	"FROM",
	"TO",
	"WITH",
	"UNITE",
	"ALL",
	"ADDRESS",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacc/console/sql.y:478

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 131

var yyAct = [...]int{

	89, 108, 104, 110, 32, 64, 22, 23, 95, 74,
	74, 74, 87, 81, 25, 24, 29, 30, 124, 119,
	18, 17, 19, 20, 21, 26, 27, 80, 98, 53,
	58, 79, 56, 55, 54, 100, 77, 106, 71, 73,
	116, 36, 36, 99, 75, 78, 37, 37, 28, 43,
	76, 107, 38, 38, 90, 88, 84, 57, 59, 60,
	72, 47, 61, 49, 44, 52, 45, 120, 97, 123,
	82, 83, 85, 86, 70, 35, 34, 33, 91, 69,
	92, 93, 94, 74, 109, 48, 50, 111, 65, 63,
	102, 66, 67, 68, 42, 41, 40, 31, 1, 101,
	16, 15, 14, 13, 113, 112, 114, 12, 117, 118,
	10, 11, 6, 7, 115, 105, 96, 103, 39, 4,
	3, 5, 122, 121, 9, 125, 8, 62, 51, 46,
	2,
}
var yyPact = [...]int{

	0, -1000, -47, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 13, 14, 33,
	30, 30, 21, 21, 84, -1000, 30, 30, 30, 61,
	56, -1000, -1000, -1000, -1000, -1000, -4, 26, 79, -1000,
	-1000, -1000, -1000, 79, -1000, 16, -1000, -6, -1000, 11,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -14, -19, -34, 79,
	7, 79, 79, -35, -1000, -1000, 6, 5, 79, 79,
	79, 79, -42, -1000, -1000, 38, -17, 3, -1000, -1000,
	-1000, -1000, -9, -1000, -1000, 86, 8, 80, 83, 84,
	83, -1000, -1000, 8, -1000, 9, 80, 80, -1000, -1000,
	-27, -1000, -1000, -1000, -1000, -1000, 35, -1000, -1000, 83,
	80, 50, -1000, -28, 79, -1000,
}
var yyPgo = [...]int{

	0, 130, 129, 64, 128, 127, 126, 124, 121, 120,
	119, 118, 77, 76, 75, 117, 2, 116, 115, 114,
	113, 112, 111, 110, 107, 103, 102, 101, 100, 65,
	5, 3, 0, 99, 1, 98, 97,
}
var yyR1 = [...]int{

	0, 35, 36, 36, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 4, 5, 8,
	8, 8, 8, 9, 9, 9, 10, 10, 10, 10,
	6, 34, 31, 32, 30, 21, 11, 12, 15, 15,
	16, 17, 17, 18, 18, 19, 19, 13, 14, 20,
	2, 3, 24, 7, 25, 26, 23, 22, 33, 27,
	28, 28,
}
var yyR2 = [...]int{

	0, 2, 0, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	4, 2, 4, 2, 2, 2, 2, 2, 2, 2,
	2, 1, 1, 1, 1, 2, 2, 5, 1, 2,
	2, 2, 0, 2, 2, 3, 0, 10, 5, 2,
	3, 3, 6, 2, 4, 4, 2, 1, 1, 5,
	3, 3,
}
var yyChk = [...]int{

	-1000, -35, -1, -9, -10, -8, -21, -20, -6, -7,
	-23, -22, -24, -25, -26, -27, -28, 21, 20, 22,
	23, 24, 6, 7, 15, 14, 25, 26, 48, 16,
	17, -36, 51, -12, -13, -14, 28, 33, 39, -11,
	-12, -13, -14, 35, -3, 33, -2, 28, -3, 33,
	-3, -4, -29, 8, 13, 12, 11, 36, 9, 37,
	38, 41, -5, -29, -30, 4, -3, -3, -3, 18,
	18, 42, 34, -32, 4, -32, 34, 42, 34, 45,
	46, 47, -32, -32, 49, -32, -32, 47, 49, -32,
	49, -32, -32, -32, -32, 50, -17, 30, 45, 40,
	44, -33, 4, -15, -16, -18, 29, 43, -34, 4,
	-31, 4, -30, -31, -16, -19, 31, -34, -34, 46,
	32, -31, -34, 19, 46, -32,
}
var yyDef = [...]int{

	0, -2, 2, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 0, 0, 0,
	0, 0, 0, 0, 0, 67, 0, 0, 0, 0,
	0, 1, 3, 33, 34, 35, 0, 0, 0, 36,
	37, 38, 39, 0, 29, 0, 31, 0, 45, 0,
	59, 40, 27, 18, 19, 20, 21, 22, 23, 24,
	25, 26, 63, 28, 66, 44, 0, 0, 0, 0,
	0, 0, 0, 0, 43, 46, 0, 0, 0, 0,
	0, 0, 0, 70, 71, 52, 0, 0, 30, 61,
	32, 60, 0, 64, 65, 0, 0, 0, 0, 0,
	0, 69, 68, 47, 48, 56, 0, 0, 51, 41,
	0, 42, 58, 62, 49, 50, 0, 53, 54, 0,
	0, 0, 55, 0, 0, 57,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 51,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50,
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

	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line yacc/console/sql.y:129
		{
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:130
		{
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:135
		{
			setParseTree(yylex, yyDollar[1].create)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:139
		{
			setParseTree(yylex, yyDollar[1].create)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:143
		{
			setParseTree(yylex, yyDollar[1].drop)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:147
		{
			setParseTree(yylex, yyDollar[1].lock)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:151
		{
			setParseTree(yylex, yyDollar[1].unlock)
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:155
		{
			setParseTree(yylex, yyDollar[1].show)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:159
		{
			setParseTree(yylex, yyDollar[1].kill)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:163
		{
			setParseTree(yylex, yyDollar[1].listen)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:167
		{
			setParseTree(yylex, yyDollar[1].shutdown)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:171
		{
			setParseTree(yylex, yyDollar[1].split)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:175
		{
			setParseTree(yylex, yyDollar[1].move)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:179
		{
			setParseTree(yylex, yyDollar[1].unite)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:183
		{
			setParseTree(yylex, yyDollar[1].register_router)
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:187
		{
			setParseTree(yylex, yyDollar[1].unregister_router)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:204
		{
			switch v := string(yyDollar[1].str); v {
			case ShowDatabasesStr, ShowRoutersStr, ShowPoolsStr, ShowShardsStr, ShowKeyRangesStr, ShowShardingRules:
				yyVAL.str = v
			default:
				yyVAL.str = ShowUnsupportedStr
			}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:215
		{
			switch v := string(yyDollar[1].str); v {
			case KillClientsStr:
				yyVAL.str = v
			default:
				yyVAL.str = "unsupp"
			}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:226
		{
			yyVAL.drop = &Drop{Element: yyDollar[2].key_range_selector}
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
//line yacc/console/sql.y:231
		{
			yyVAL.drop = &Drop{Element: &KeyRangeSelector{KeyRangeID: `*`}}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:235
		{
			yyVAL.drop = &Drop{Element: yyDollar[2].sharding_rule_selector}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
//line yacc/console/sql.y:240
		{
			yyVAL.drop = &Drop{Element: &ShardingRuleSelector{ID: `*`}}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:246
		{
			yyVAL.create = &Create{Element: yyDollar[2].sharding_rule}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:251
		{
			yyVAL.create = &Create{Element: yyDollar[2].kr}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:255
		{
			yyVAL.create = &Create{Element: yyDollar[2].shard}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:262
		{
			yyVAL.create = &Create{Element: yyDollar[2].ds}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:267
		{
			yyVAL.create = &Create{Element: yyDollar[2].sharding_rule}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:272
		{
			yyVAL.create = &Create{Element: yyDollar[2].kr}
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:276
		{
			yyVAL.create = &Create{Element: yyDollar[2].shard}
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:283
		{
			yyVAL.show = &Show{Cmd: yyDollar[2].str}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:289
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:296
		{
			yyVAL.bytes = []byte(yyDollar[1].str)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:302
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:308
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:314
		{
			yyVAL.lock = &Lock{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID}
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:322
		{
			yyVAL.ds = &DataspaceDefinition{ID: yyDollar[2].str}
		}
	case 47:
		yyDollar = yyS[yypt-5 : yypt+1]
//line yacc/console/sql.y:334
		{
			yyVAL.sharding_rule = &ShardingRuleDefinition{ID: yyDollar[3].str, TableName: yyDollar[4].str, Entries: yyDollar[5].entrieslist}
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:339
		{
			yyVAL.entrieslist = make([]ShardingRuleEntry, 0)
			yyVAL.entrieslist = append(yyVAL.entrieslist, yyDollar[1].shruleEntry)
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:345
		{
			yyVAL.entrieslist = append(yyDollar[1].entrieslist, yyDollar[2].shruleEntry)
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:351
		{
			yyVAL.shruleEntry = ShardingRuleEntry{
				Column:       yyDollar[1].str,
				HashFunction: yyDollar[2].str,
			}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:360
		{
			yyVAL.str = yyDollar[2].str
		}
	case 52:
		yyDollar = yyS[yypt-0 : yypt+1]
//line yacc/console/sql.y:363
		{
			yyVAL.str = ""
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:367
		{
			yyVAL.str = yyDollar[2].str
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:372
		{
			yyVAL.str = yyDollar[2].str
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line yacc/console/sql.y:378
		{
			yyVAL.str = yyDollar[3].str
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
//line yacc/console/sql.y:381
		{
			yyVAL.str = ""
		}
	case 57:
		yyDollar = yyS[yypt-10 : yypt+1]
//line yacc/console/sql.y:386
		{
			yyVAL.kr = &KeyRangeDefinition{LowerBound: yyDollar[5].bytes, UpperBound: yyDollar[7].bytes, ShardID: yyDollar[10].str, KeyRangeID: yyDollar[3].str}
		}
	case 58:
		yyDollar = yyS[yypt-5 : yypt+1]
//line yacc/console/sql.y:393
		{
			yyVAL.shard = &ShardDefinition{Id: yyDollar[2].str, Hosts: []string{yyDollar[5].str}}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:400
		{
			yyVAL.unlock = &Unlock{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line yacc/console/sql.y:406
		{
			yyVAL.sharding_rule_selector = &ShardingRuleSelector{ID: yyDollar[3].str}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line yacc/console/sql.y:412
		{
			yyVAL.key_range_selector = &KeyRangeSelector{KeyRangeID: yyDollar[3].str}
		}
	case 62:
		yyDollar = yyS[yypt-6 : yypt+1]
//line yacc/console/sql.y:418
		{
			yyVAL.split = &SplitKeyRange{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID, KeyRangeFromID: yyDollar[4].str, Border: yyDollar[6].bytes}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:424
		{
			yyVAL.kill = &Kill{Cmd: yyDollar[2].str}
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
//line yacc/console/sql.y:430
		{
			yyVAL.move = &MoveKeyRange{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID, DestShardID: yyDollar[4].str}
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
//line yacc/console/sql.y:436
		{
			yyVAL.unite = &UniteKeyRange{KeyRangeIDL: yyDollar[2].key_range_selector.KeyRangeID, KeyRangeIDR: yyDollar[4].str}
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
//line yacc/console/sql.y:442
		{
			yyVAL.listen = &Listen{addr: yyDollar[2].str}
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:448
		{
			yyVAL.shutdown = &Shutdown{}
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
//line yacc/console/sql.y:456
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 69:
		yyDollar = yyS[yypt-5 : yypt+1]
//line yacc/console/sql.y:463
		{
			yyVAL.register_router = &RegisterRouter{ID: yyDollar[3].str, Addr: yyDollar[5].str}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line yacc/console/sql.y:469
		{
			yyVAL.unregister_router = &UnregisterRouter{ID: yyDollar[3].str}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line yacc/console/sql.y:474
		{
			yyVAL.unregister_router = &UnregisterRouter{ID: `*`}
		}
	}
	goto yystack /* stack new state and value */
}
