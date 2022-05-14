"".percobaan STEXT size=439 args=0x10 locals=0x98 funcid=0x0
	0x0000 00000 (main.go:9)	TEXT	"".percobaan(SB), ABIInternal, $152-16
	0x0000 00000 (main.go:9)	LEAQ	-24(SP), R12
	0x0005 00005 (main.go:9)	CMPQ	R12, 16(R14)
	0x0009 00009 (main.go:9)	PCDATA	$0, $-2
	0x0009 00009 (main.go:9)	JLS	405
	0x000f 00015 (main.go:9)	PCDATA	$0, $-1
	0x000f 00015 (main.go:9)	SUBQ	$152, SP
	0x0016 00022 (main.go:9)	MOVQ	BP, 144(SP)
	0x001e 00030 (main.go:9)	LEAQ	144(SP), BP
	0x0026 00038 (main.go:9)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0026 00038 (main.go:9)	FUNCDATA	$1, gclocals·87471ff47f1e0f13c6fee0c186afcd29(SB)
	0x0026 00038 (main.go:9)	FUNCDATA	$2, "".percobaan.stkobj(SB)
	0x0026 00038 (main.go:9)	FUNCDATA	$5, "".percobaan.arginfo1(SB)
	0x0026 00038 (main.go:11)	MOVSD	X1, ""..autotmp_62+40(SP)
	0x002c 00044 (main.go:10)	PCDATA	$1, $0
	0x002c 00044 (main.go:10)	CALL	"".Akar(SB)
	0x0031 00049 (main.go:10)	MOVQ	X0, AX
	0x0036 00054 (main.go:10)	CALL	runtime.convT64(SB)
	0x003b 00059 (main.go:10)	LEAQ	""..autotmp_30+112(SP), CX
	0x0040 00064 (main.go:10)	MOVUPS	X15, (CX)
	0x0044 00068 (main.go:10)	LEAQ	""..autotmp_30+128(SP), DX
	0x004c 00076 (main.go:10)	MOVUPS	X15, (DX)
	0x0050 00080 (main.go:10)	LEAQ	type.string(SB), DX
	0x0057 00087 (main.go:10)	MOVQ	DX, ""..autotmp_30+112(SP)
	0x005c 00092 (main.go:10)	LEAQ	""..stmp_0(SB), BX
	0x0063 00099 (main.go:10)	MOVQ	BX, ""..autotmp_30+120(SP)
	0x0068 00104 (main.go:10)	LEAQ	type.float64(SB), BX
	0x006f 00111 (main.go:10)	MOVQ	BX, ""..autotmp_30+128(SP)
	0x0077 00119 (main.go:10)	MOVQ	AX, ""..autotmp_30+136(SP)
	0x007f 00127 (<unknown line number>)	NOP
	0x007f 00127 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), AX
	0x0086 00134 ($GOROOT/src/fmt/print.go:274)	MOVL	$2, DI
	0x008b 00139 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x008e 00142 ($GOROOT/src/fmt/print.go:274)	MOVQ	AX, R8
	0x0091 00145 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0098 00152 ($GOROOT/src/fmt/print.go:274)	MOVQ	R8, BX
	0x009b 00155 ($GOROOT/src/fmt/print.go:274)	NOP
	0x00a0 00160 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x00a5 00165 (main.go:11)	MOVSD	""..autotmp_62+40(SP), X0
	0x00ab 00171 (main.go:11)	CALL	"".bulatBawah(SB)
	0x00b0 00176 (main.go:11)	MOVQ	X0, AX
	0x00b5 00181 (main.go:11)	CALL	runtime.convT64(SB)
	0x00ba 00186 (main.go:11)	LEAQ	""..autotmp_35+80(SP), CX
	0x00bf 00191 (main.go:11)	MOVUPS	X15, (CX)
	0x00c3 00195 (main.go:11)	LEAQ	""..autotmp_35+96(SP), DX
	0x00c8 00200 (main.go:11)	MOVUPS	X15, (DX)
	0x00cc 00204 (main.go:11)	LEAQ	type.string(SB), DX
	0x00d3 00211 (main.go:11)	MOVQ	DX, ""..autotmp_35+80(SP)
	0x00d8 00216 (main.go:11)	LEAQ	""..stmp_1(SB), BX
	0x00df 00223 (main.go:11)	MOVQ	BX, ""..autotmp_35+88(SP)
	0x00e4 00228 (main.go:11)	LEAQ	type.float64(SB), BX
	0x00eb 00235 (main.go:11)	MOVQ	BX, ""..autotmp_35+96(SP)
	0x00f0 00240 (main.go:11)	MOVQ	AX, ""..autotmp_35+104(SP)
	0x00f5 00245 (<unknown line number>)	NOP
	0x00f5 00245 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), AX
	0x00fc 00252 ($GOROOT/src/fmt/print.go:274)	MOVL	$2, DI
	0x0101 00257 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x0104 00260 ($GOROOT/src/fmt/print.go:274)	MOVQ	AX, R8
	0x0107 00263 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x010e 00270 ($GOROOT/src/fmt/print.go:274)	MOVQ	R8, BX
	0x0111 00273 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x0116 00278 (main.go:12)	MOVSD	""..autotmp_62+40(SP), X0
	0x011c 00284 (main.go:12)	NOP
	0x0120 00288 (main.go:12)	CALL	"".bulatAtas(SB)
	0x0125 00293 (main.go:12)	MOVQ	X0, AX
	0x012a 00298 (main.go:12)	CALL	runtime.convT64(SB)
	0x012f 00303 (main.go:12)	LEAQ	""..autotmp_40+48(SP), CX
	0x0134 00308 (main.go:12)	MOVUPS	X15, (CX)
	0x0138 00312 (main.go:12)	LEAQ	""..autotmp_40+64(SP), DX
	0x013d 00317 (main.go:12)	MOVUPS	X15, (DX)
	0x0141 00321 (main.go:12)	LEAQ	type.string(SB), DX
	0x0148 00328 (main.go:12)	MOVQ	DX, ""..autotmp_40+48(SP)
	0x014d 00333 (main.go:12)	LEAQ	""..stmp_2(SB), DX
	0x0154 00340 (main.go:12)	MOVQ	DX, ""..autotmp_40+56(SP)
	0x0159 00345 (main.go:12)	LEAQ	type.float64(SB), DX
	0x0160 00352 (main.go:12)	MOVQ	DX, ""..autotmp_40+64(SP)
	0x0165 00357 (main.go:12)	MOVQ	AX, ""..autotmp_40+72(SP)
	0x016a 00362 (<unknown line number>)	NOP
	0x016a 00362 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), BX
	0x0171 00369 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0178 00376 ($GOROOT/src/fmt/print.go:274)	MOVL	$2, DI
	0x017d 00381 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x0180 00384 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x0185 00389 (main.go:13)	MOVQ	144(SP), BP
	0x018d 00397 (main.go:13)	ADDQ	$152, SP
	0x0194 00404 (main.go:13)	RET
	0x0195 00405 (main.go:13)	NOP
	0x0195 00405 (main.go:9)	PCDATA	$1, $-1
	0x0195 00405 (main.go:9)	PCDATA	$0, $-2
	0x0195 00405 (main.go:9)	MOVSD	X0, 8(SP)
	0x019b 00411 (main.go:9)	MOVSD	X1, 16(SP)
	0x01a1 00417 (main.go:9)	CALL	runtime.morestack_noctxt(SB)
	0x01a6 00422 (main.go:9)	MOVSD	8(SP), X0
	0x01ac 00428 (main.go:9)	MOVSD	16(SP), X1
	0x01b2 00434 (main.go:9)	PCDATA	$0, $-1
	0x01b2 00434 (main.go:9)	JMP	0
	0x0000 4c 8d 64 24 e8 4d 3b 66 10 0f 86 86 01 00 00 48  L.d$.M;f.......H
	0x0010 81 ec 98 00 00 00 48 89 ac 24 90 00 00 00 48 8d  ......H..$....H.
	0x0020 ac 24 90 00 00 00 f2 0f 11 4c 24 28 e8 00 00 00  .$.......L$(....
	0x0030 00 66 48 0f 7e c0 e8 00 00 00 00 48 8d 4c 24 70  .fH.~......H.L$p
	0x0040 44 0f 11 39 48 8d 94 24 80 00 00 00 44 0f 11 3a  D..9H..$....D..:
	0x0050 48 8d 15 00 00 00 00 48 89 54 24 70 48 8d 1d 00  H......H.T$pH...
	0x0060 00 00 00 48 89 5c 24 78 48 8d 1d 00 00 00 00 48  ...H.\$xH......H
	0x0070 89 9c 24 80 00 00 00 48 89 84 24 88 00 00 00 48  ..$....H..$....H
	0x0080 8b 05 00 00 00 00 bf 02 00 00 00 48 89 fe 49 89  ...........H..I.
	0x0090 c0 48 8d 05 00 00 00 00 4c 89 c3 0f 1f 44 00 00  .H......L....D..
	0x00a0 e8 00 00 00 00 f2 0f 10 44 24 28 e8 00 00 00 00  ........D$(.....
	0x00b0 66 48 0f 7e c0 e8 00 00 00 00 48 8d 4c 24 50 44  fH.~......H.L$PD
	0x00c0 0f 11 39 48 8d 54 24 60 44 0f 11 3a 48 8d 15 00  ..9H.T$`D..:H...
	0x00d0 00 00 00 48 89 54 24 50 48 8d 1d 00 00 00 00 48  ...H.T$PH......H
	0x00e0 89 5c 24 58 48 8d 1d 00 00 00 00 48 89 5c 24 60  .\$XH......H.\$`
	0x00f0 48 89 44 24 68 48 8b 05 00 00 00 00 bf 02 00 00  H.D$hH..........
	0x0100 00 48 89 fe 49 89 c0 48 8d 05 00 00 00 00 4c 89  .H..I..H......L.
	0x0110 c3 e8 00 00 00 00 f2 0f 10 44 24 28 0f 1f 40 00  .........D$(..@.
	0x0120 e8 00 00 00 00 66 48 0f 7e c0 e8 00 00 00 00 48  .....fH.~......H
	0x0130 8d 4c 24 30 44 0f 11 39 48 8d 54 24 40 44 0f 11  .L$0D..9H.T$@D..
	0x0140 3a 48 8d 15 00 00 00 00 48 89 54 24 30 48 8d 15  :H......H.T$0H..
	0x0150 00 00 00 00 48 89 54 24 38 48 8d 15 00 00 00 00  ....H.T$8H......
	0x0160 48 89 54 24 40 48 89 44 24 48 48 8b 1d 00 00 00  H.T$@H.D$HH.....
	0x0170 00 48 8d 05 00 00 00 00 bf 02 00 00 00 48 89 fe  .H...........H..
	0x0180 e8 00 00 00 00 48 8b ac 24 90 00 00 00 48 81 c4  .....H..$....H..
	0x0190 98 00 00 00 c3 f2 0f 11 44 24 08 f2 0f 11 4c 24  ........D$....L$
	0x01a0 10 e8 00 00 00 00 f2 0f 10 44 24 08 f2 0f 10 4c  .........D$....L
	0x01b0 24 10 e9 49 fe ff ff                             $..I...
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.float64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.string+0
	rel 3+0 t=24 type.float64+0
	rel 3+0 t=24 type.*os.File+0
	rel 3+0 t=24 type.string+0
	rel 3+0 t=24 type.float64+0
	rel 3+0 t=24 type.string+0
	rel 45+4 t=7 "".Akar+0
	rel 55+4 t=7 runtime.convT64+0
	rel 83+4 t=15 type.string+0
	rel 95+4 t=15 ""..stmp_0+0
	rel 107+4 t=15 type.float64+0
	rel 130+4 t=15 os.Stdout+0
	rel 148+4 t=15 go.itab.*os.File,io.Writer+0
	rel 161+4 t=7 fmt.Fprintln+0
	rel 172+4 t=7 "".bulatBawah+0
	rel 182+4 t=7 runtime.convT64+0
	rel 207+4 t=15 type.string+0
	rel 219+4 t=15 ""..stmp_1+0
	rel 231+4 t=15 type.float64+0
	rel 248+4 t=15 os.Stdout+0
	rel 266+4 t=15 go.itab.*os.File,io.Writer+0
	rel 274+4 t=7 fmt.Fprintln+0
	rel 289+4 t=7 "".bulatAtas+0
	rel 299+4 t=7 runtime.convT64+0
	rel 324+4 t=15 type.string+0
	rel 336+4 t=15 ""..stmp_2+0
	rel 348+4 t=15 type.float64+0
	rel 365+4 t=15 os.Stdout+0
	rel 372+4 t=15 go.itab.*os.File,io.Writer+0
	rel 385+4 t=7 fmt.Fprintln+0
	rel 418+4 t=7 runtime.morestack_noctxt+0
"".main STEXT size=58 args=0x0 locals=0x18 funcid=0x0
	0x0000 00000 (main.go:17)	TEXT	"".main(SB), ABIInternal, $24-0
	0x0000 00000 (main.go:17)	CMPQ	SP, 16(R14)
	0x0004 00004 (main.go:17)	PCDATA	$0, $-2
	0x0004 00004 (main.go:17)	JLS	51
	0x0006 00006 (main.go:17)	PCDATA	$0, $-1
	0x0006 00006 (main.go:17)	SUBQ	$24, SP
	0x000a 00010 (main.go:17)	MOVQ	BP, 16(SP)
	0x000f 00015 (main.go:17)	LEAQ	16(SP), BP
	0x0014 00020 (main.go:17)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0014 00020 (main.go:17)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0014 00020 (main.go:18)	MOVSD	$f64.4030000000000000(SB), X0
	0x001c 00028 (main.go:18)	MOVSD	$f64.402f1eb851eb851f(SB), X1
	0x0024 00036 (main.go:18)	PCDATA	$1, $0
	0x0024 00036 (main.go:18)	CALL	"".printPercobaan(SB)
	0x0029 00041 (main.go:19)	MOVQ	16(SP), BP
	0x002e 00046 (main.go:19)	ADDQ	$24, SP
	0x0032 00050 (main.go:19)	RET
	0x0033 00051 (main.go:19)	NOP
	0x0033 00051 (main.go:17)	PCDATA	$1, $-1
	0x0033 00051 (main.go:17)	PCDATA	$0, $-2
	0x0033 00051 (main.go:17)	CALL	runtime.morestack_noctxt(SB)
	0x0038 00056 (main.go:17)	PCDATA	$0, $-1
	0x0038 00056 (main.go:17)	JMP	0
	0x0000 49 3b 66 10 76 2d 48 83 ec 18 48 89 6c 24 10 48  I;f.v-H...H.l$.H
	0x0010 8d 6c 24 10 f2 0f 10 05 00 00 00 00 f2 0f 10 0d  .l$.............
	0x0020 00 00 00 00 e8 00 00 00 00 48 8b 6c 24 10 48 83  .........H.l$.H.
	0x0030 c4 18 c3 e8 00 00 00 00 eb c6                    ..........
	rel 24+4 t=15 $f64.4030000000000000+0
	rel 32+4 t=15 $f64.402f1eb851eb851f+0
	rel 37+4 t=7 "".printPercobaan+0
	rel 52+4 t=7 runtime.morestack_noctxt+0
os.(*File).close STEXT dupok size=86 args=0x8 locals=0x10 funcid=0x16
	0x0000 00000 (<autogenerated>:1)	TEXT	os.(*File).close(SB), DUPOK|WRAPPER|ABIInternal, $16-8
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	52
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$16, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, 8(SP)
	0x000f 00015 (<autogenerated>:1)	LEAQ	8(SP), BP
	0x0014 00020 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0018 00024 (<autogenerated>:1)	TESTQ	R12, R12
	0x001b 00027 (<autogenerated>:1)	JNE	69
	0x001d 00029 (<autogenerated>:1)	NOP
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$5, os.(*File).close.arginfo1(SB)
	0x001d 00029 (<autogenerated>:1)	MOVQ	AX, ""..this+24(SP)
	0x0022 00034 (<autogenerated>:1)	MOVQ	(AX), AX
	0x0025 00037 (<autogenerated>:1)	PCDATA	$1, $1
	0x0025 00037 (<autogenerated>:1)	CALL	os.(*file).close(SB)
	0x002a 00042 (<autogenerated>:1)	MOVQ	8(SP), BP
	0x002f 00047 (<autogenerated>:1)	ADDQ	$16, SP
	0x0033 00051 (<autogenerated>:1)	RET
	0x0034 00052 (<autogenerated>:1)	NOP
	0x0034 00052 (<autogenerated>:1)	PCDATA	$1, $-1
	0x0034 00052 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0034 00052 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0039 00057 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x003e 00062 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x0043 00067 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0043 00067 (<autogenerated>:1)	JMP	0
	0x0045 00069 (<autogenerated>:1)	LEAQ	24(SP), R13
	0x004a 00074 (<autogenerated>:1)	CMPQ	(R12), R13
	0x004e 00078 (<autogenerated>:1)	JNE	29
	0x0050 00080 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x0054 00084 (<autogenerated>:1)	JMP	29
	0x0000 49 3b 66 10 76 2e 48 83 ec 10 48 89 6c 24 08 48  I;f.v.H...H.l$.H
	0x0010 8d 6c 24 08 4d 8b 66 20 4d 85 e4 75 28 48 89 44  .l$.M.f M..u(H.D
	0x0020 24 18 48 8b 00 e8 00 00 00 00 48 8b 6c 24 08 48  $.H.......H.l$.H
	0x0030 83 c4 10 c3 48 89 44 24 08 e8 00 00 00 00 48 8b  ....H.D$......H.
	0x0040 44 24 08 eb bb 4c 8d 6c 24 18 4d 39 2c 24 75 cd  D$...L.l$.M9,$u.
	0x0050 49 89 24 24 eb c7                                I.$$..
	rel 38+4 t=7 os.(*file).close+0
	rel 58+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
go.info.fmt.Println$abstract SDWARFABSFCN dupok size=42
	0x0000 04 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 11  .fmt.Println....
	0x0010 61 00 00 00 00 00 00 11 6e 00 01 00 00 00 00 11  a.......n.......
	0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
	rel 0+0 t=23 type.[]interface {}+0
	rel 0+0 t=23 type.error+0
	rel 0+0 t=23 type.int+0
	rel 19+4 t=31 go.info.[]interface {}+0
	rel 27+4 t=31 go.info.int+0
	rel 37+4 t=31 go.info.error+0
"".Akar.args_stackmap SRODATA size=10
	0x0000 02 00 00 00 04 00 00 00 00 00                    ..........
"".bulatBawah.args_stackmap SRODATA size=10
	0x0000 02 00 00 00 04 00 00 00 00 00                    ..........
"".bulatAtas.args_stackmap SRODATA size=10
	0x0000 02 00 00 00 04 00 00 00 00 00                    ..........
go.string."Akar dari bilangan 16 adalah ... " SRODATA dupok size=33
	0x0000 41 6b 61 72 20 64 61 72 69 20 62 69 6c 61 6e 67  Akar dari bilang
	0x0010 61 6e 20 31 36 20 61 64 61 6c 61 68 20 2e 2e 2e  an 16 adalah ...
	0x0020 20                                                
go.string."Pembulatan ke bawah dari 15.56 adalah .. " SRODATA dupok size=41
	0x0000 50 65 6d 62 75 6c 61 74 61 6e 20 6b 65 20 62 61  Pembulatan ke ba
	0x0010 77 61 68 20 64 61 72 69 20 31 35 2e 35 36 20 61  wah dari 15.56 a
	0x0020 64 61 6c 61 68 20 2e 2e 20                       dalah .. 
go.string."Pembulatan ke atas dari 15.56 adalah .. " SRODATA dupok size=40
	0x0000 50 65 6d 62 75 6c 61 74 61 6e 20 6b 65 20 61 74  Pembulatan ke at
	0x0010 61 73 20 64 61 72 69 20 31 35 2e 35 36 20 61 64  as dari 15.56 ad
	0x0020 61 6c 61 68 20 2e 2e 20                          alah .. 
"".printPercobaan.args_stackmap SRODATA size=9
	0x0000 01 00 00 00 04 00 00 00 00                       .........
""..stmp_0 SRODATA static size=16
	0x0000 00 00 00 00 00 00 00 00 21 00 00 00 00 00 00 00  ........!.......
	rel 0+8 t=1 go.string."Akar dari bilangan 16 adalah ... "+0
""..stmp_1 SRODATA static size=16
	0x0000 00 00 00 00 00 00 00 00 29 00 00 00 00 00 00 00  ........).......
	rel 0+8 t=1 go.string."Pembulatan ke bawah dari 15.56 adalah .. "+0
""..stmp_2 SRODATA static size=16
	0x0000 00 00 00 00 00 00 00 00 28 00 00 00 00 00 00 00  ........(.......
	rel 0+8 t=1 go.string."Pembulatan ke atas dari 15.56 adalah .. "+0
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=15
	0x0000 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d     ..*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=-32763 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=17
	0x0000 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20 7b  ..*[]interface {
	0x0010 7d                                               }
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=-32763 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.0a SRODATA dupok size=1
	0x0000 0a                                               .
go.itab.*os.File,io.Writer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
	rel 0+8 t=1 type.io.Writer+0
	rel 8+8 t=1 type.*os.File+0
	rel 24+8 t=-32767 os.(*File).Write+0
type..importpath.fmt. SRODATA dupok size=5
	0x0000 00 03 66 6d 74                                   ..fmt
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·87471ff47f1e0f13c6fee0c186afcd29 SRODATA dupok size=10
	0x0000 01 00 00 00 0c 00 00 00 00 00                    ..........
"".percobaan.stkobj SRODATA static size=80
	0x0000 03 00 00 00 00 00 00 00 a0 ff ff ff 20 00 00 00  ............ ...
	0x0010 20 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00   ...............
	0x0020 c0 ff ff ff 20 00 00 00 20 00 00 00 00 00 00 00  .... ... .......
	0x0030 00 00 00 00 00 00 00 00 e0 ff ff ff 20 00 00 00  ............ ...
	0x0040 20 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00   ...............
	rel 24+8 t=1 runtime.gcbits.0a+0
	rel 48+8 t=1 runtime.gcbits.0a+0
	rel 72+8 t=1 runtime.gcbits.0a+0
"".percobaan.arginfo1 SRODATA static dupok size=5
	0x0000 00 08 08 08 ff                                   .....
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
os.(*File).close.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
$f64.402f1eb851eb851f SRODATA size=8
	0x0000 1f 85 eb 51 b8 1e 2f 40                          ...Q../@
$f64.4030000000000000 SRODATA size=8
	0x0000 00 00 00 00 00 00 30 40                          ......0@
