package argon2

/*
Copyright © 2015 Andrew Ekstedt
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions
are met:

1. Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright
notice, this list of conditions and the following disclaimer in the
documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED
AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY
 WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 POSSIBILITY OF SUCH DAMAGE.
*/

func round(z, a, b *block) {
	z[0] = a[0] ^ b[0]
	z[1] = a[1] ^ b[1]
	z[2] = a[2] ^ b[2]
	z[3] = a[3] ^ b[3]
	z[4] = a[4] ^ b[4]
	z[5] = a[5] ^ b[5]
	z[6] = a[6] ^ b[6]
	z[7] = a[7] ^ b[7]
	z[8] = a[8] ^ b[8]
	z[9] = a[9] ^ b[9]
	z[10] = a[10] ^ b[10]
	z[11] = a[11] ^ b[11]
	z[12] = a[12] ^ b[12]
	z[13] = a[13] ^ b[13]
	z[14] = a[14] ^ b[14]
	z[15] = a[15] ^ b[15]
	z[16] = a[16] ^ b[16]
	z[17] = a[17] ^ b[17]
	z[18] = a[18] ^ b[18]
	z[19] = a[19] ^ b[19]
	z[20] = a[20] ^ b[20]
	z[21] = a[21] ^ b[21]
	z[22] = a[22] ^ b[22]
	z[23] = a[23] ^ b[23]
	z[24] = a[24] ^ b[24]
	z[25] = a[25] ^ b[25]
	z[26] = a[26] ^ b[26]
	z[27] = a[27] ^ b[27]
	z[28] = a[28] ^ b[28]
	z[29] = a[29] ^ b[29]
	z[30] = a[30] ^ b[30]
	z[31] = a[31] ^ b[31]
	z[32] = a[32] ^ b[32]
	z[33] = a[33] ^ b[33]
	z[34] = a[34] ^ b[34]
	z[35] = a[35] ^ b[35]
	z[36] = a[36] ^ b[36]
	z[37] = a[37] ^ b[37]
	z[38] = a[38] ^ b[38]
	z[39] = a[39] ^ b[39]
	z[40] = a[40] ^ b[40]
	z[41] = a[41] ^ b[41]
	z[42] = a[42] ^ b[42]
	z[43] = a[43] ^ b[43]
	z[44] = a[44] ^ b[44]
	z[45] = a[45] ^ b[45]
	z[46] = a[46] ^ b[46]
	z[47] = a[47] ^ b[47]
	z[48] = a[48] ^ b[48]
	z[49] = a[49] ^ b[49]
	z[50] = a[50] ^ b[50]
	z[51] = a[51] ^ b[51]
	z[52] = a[52] ^ b[52]
	z[53] = a[53] ^ b[53]
	z[54] = a[54] ^ b[54]
	z[55] = a[55] ^ b[55]
	z[56] = a[56] ^ b[56]
	z[57] = a[57] ^ b[57]
	z[58] = a[58] ^ b[58]
	z[59] = a[59] ^ b[59]
	z[60] = a[60] ^ b[60]
	z[61] = a[61] ^ b[61]
	z[62] = a[62] ^ b[62]
	z[63] = a[63] ^ b[63]
	z[64] = a[64] ^ b[64]
	z[65] = a[65] ^ b[65]
	z[66] = a[66] ^ b[66]
	z[67] = a[67] ^ b[67]
	z[68] = a[68] ^ b[68]
	z[69] = a[69] ^ b[69]
	z[70] = a[70] ^ b[70]
	z[71] = a[71] ^ b[71]
	z[72] = a[72] ^ b[72]
	z[73] = a[73] ^ b[73]
	z[74] = a[74] ^ b[74]
	z[75] = a[75] ^ b[75]
	z[76] = a[76] ^ b[76]
	z[77] = a[77] ^ b[77]
	z[78] = a[78] ^ b[78]
	z[79] = a[79] ^ b[79]
	z[80] = a[80] ^ b[80]
	z[81] = a[81] ^ b[81]
	z[82] = a[82] ^ b[82]
	z[83] = a[83] ^ b[83]
	z[84] = a[84] ^ b[84]
	z[85] = a[85] ^ b[85]
	z[86] = a[86] ^ b[86]
	z[87] = a[87] ^ b[87]
	z[88] = a[88] ^ b[88]
	z[89] = a[89] ^ b[89]
	z[90] = a[90] ^ b[90]
	z[91] = a[91] ^ b[91]
	z[92] = a[92] ^ b[92]
	z[93] = a[93] ^ b[93]
	z[94] = a[94] ^ b[94]
	z[95] = a[95] ^ b[95]
	z[96] = a[96] ^ b[96]
	z[97] = a[97] ^ b[97]
	z[98] = a[98] ^ b[98]
	z[99] = a[99] ^ b[99]
	z[100] = a[100] ^ b[100]
	z[101] = a[101] ^ b[101]
	z[102] = a[102] ^ b[102]
	z[103] = a[103] ^ b[103]
	z[104] = a[104] ^ b[104]
	z[105] = a[105] ^ b[105]
	z[106] = a[106] ^ b[106]
	z[107] = a[107] ^ b[107]
	z[108] = a[108] ^ b[108]
	z[109] = a[109] ^ b[109]
	z[110] = a[110] ^ b[110]
	z[111] = a[111] ^ b[111]
	z[112] = a[112] ^ b[112]
	z[113] = a[113] ^ b[113]
	z[114] = a[114] ^ b[114]
	z[115] = a[115] ^ b[115]
	z[116] = a[116] ^ b[116]
	z[117] = a[117] ^ b[117]
	z[118] = a[118] ^ b[118]
	z[119] = a[119] ^ b[119]
	z[120] = a[120] ^ b[120]
	z[121] = a[121] ^ b[121]
	z[122] = a[122] ^ b[122]
	z[123] = a[123] ^ b[123]
	z[124] = a[124] ^ b[124]
	z[125] = a[125] ^ b[125]
	z[126] = a[126] ^ b[126]
	z[127] = a[127] ^ b[127]
	_P(&z[0], &z[1], &z[2], &z[3], &z[4], &z[5], &z[6], &z[7], &z[8], &z[9], &z[10], &z[11], &z[12], &z[13], &z[14], &z[15])
	_P(&z[16], &z[17], &z[18], &z[19], &z[20], &z[21], &z[22], &z[23], &z[24], &z[25], &z[26], &z[27], &z[28], &z[29], &z[30], &z[31])
	_P(&z[32], &z[33], &z[34], &z[35], &z[36], &z[37], &z[38], &z[39], &z[40], &z[41], &z[42], &z[43], &z[44], &z[45], &z[46], &z[47])
	_P(&z[48], &z[49], &z[50], &z[51], &z[52], &z[53], &z[54], &z[55], &z[56], &z[57], &z[58], &z[59], &z[60], &z[61], &z[62], &z[63])
	_P(&z[64], &z[65], &z[66], &z[67], &z[68], &z[69], &z[70], &z[71], &z[72], &z[73], &z[74], &z[75], &z[76], &z[77], &z[78], &z[79])
	_P(&z[80], &z[81], &z[82], &z[83], &z[84], &z[85], &z[86], &z[87], &z[88], &z[89], &z[90], &z[91], &z[92], &z[93], &z[94], &z[95])
	_P(&z[96], &z[97], &z[98], &z[99], &z[100], &z[101], &z[102], &z[103], &z[104], &z[105], &z[106], &z[107], &z[108], &z[109], &z[110], &z[111])
	_P(&z[112], &z[113], &z[114], &z[115], &z[116], &z[117], &z[118], &z[119], &z[120], &z[121], &z[122], &z[123], &z[124], &z[125], &z[126], &z[127])
	_P(&z[0], &z[1], &z[16], &z[17], &z[32], &z[33], &z[48], &z[49], &z[64], &z[65], &z[80], &z[81], &z[96], &z[97], &z[112], &z[113])
	_P(&z[2], &z[3], &z[18], &z[19], &z[34], &z[35], &z[50], &z[51], &z[66], &z[67], &z[82], &z[83], &z[98], &z[99], &z[114], &z[115])
	_P(&z[4], &z[5], &z[20], &z[21], &z[36], &z[37], &z[52], &z[53], &z[68], &z[69], &z[84], &z[85], &z[100], &z[101], &z[116], &z[117])
	_P(&z[6], &z[7], &z[22], &z[23], &z[38], &z[39], &z[54], &z[55], &z[70], &z[71], &z[86], &z[87], &z[102], &z[103], &z[118], &z[119])
	_P(&z[8], &z[9], &z[24], &z[25], &z[40], &z[41], &z[56], &z[57], &z[72], &z[73], &z[88], &z[89], &z[104], &z[105], &z[120], &z[121])
	_P(&z[10], &z[11], &z[26], &z[27], &z[42], &z[43], &z[58], &z[59], &z[74], &z[75], &z[90], &z[91], &z[106], &z[107], &z[122], &z[123])
	_P(&z[12], &z[13], &z[28], &z[29], &z[44], &z[45], &z[60], &z[61], &z[76], &z[77], &z[92], &z[93], &z[108], &z[109], &z[124], &z[125])
	_P(&z[14], &z[15], &z[30], &z[31], &z[46], &z[47], &z[62], &z[63], &z[78], &z[79], &z[94], &z[95], &z[110], &z[111], &z[126], &z[127])
	z[0] ^= a[0] ^ b[0]
	z[1] ^= a[1] ^ b[1]
	z[2] ^= a[2] ^ b[2]
	z[3] ^= a[3] ^ b[3]
	z[4] ^= a[4] ^ b[4]
	z[5] ^= a[5] ^ b[5]
	z[6] ^= a[6] ^ b[6]
	z[7] ^= a[7] ^ b[7]
	z[8] ^= a[8] ^ b[8]
	z[9] ^= a[9] ^ b[9]
	z[10] ^= a[10] ^ b[10]
	z[11] ^= a[11] ^ b[11]
	z[12] ^= a[12] ^ b[12]
	z[13] ^= a[13] ^ b[13]
	z[14] ^= a[14] ^ b[14]
	z[15] ^= a[15] ^ b[15]
	z[16] ^= a[16] ^ b[16]
	z[17] ^= a[17] ^ b[17]
	z[18] ^= a[18] ^ b[18]
	z[19] ^= a[19] ^ b[19]
	z[20] ^= a[20] ^ b[20]
	z[21] ^= a[21] ^ b[21]
	z[22] ^= a[22] ^ b[22]
	z[23] ^= a[23] ^ b[23]
	z[24] ^= a[24] ^ b[24]
	z[25] ^= a[25] ^ b[25]
	z[26] ^= a[26] ^ b[26]
	z[27] ^= a[27] ^ b[27]
	z[28] ^= a[28] ^ b[28]
	z[29] ^= a[29] ^ b[29]
	z[30] ^= a[30] ^ b[30]
	z[31] ^= a[31] ^ b[31]
	z[32] ^= a[32] ^ b[32]
	z[33] ^= a[33] ^ b[33]
	z[34] ^= a[34] ^ b[34]
	z[35] ^= a[35] ^ b[35]
	z[36] ^= a[36] ^ b[36]
	z[37] ^= a[37] ^ b[37]
	z[38] ^= a[38] ^ b[38]
	z[39] ^= a[39] ^ b[39]
	z[40] ^= a[40] ^ b[40]
	z[41] ^= a[41] ^ b[41]
	z[42] ^= a[42] ^ b[42]
	z[43] ^= a[43] ^ b[43]
	z[44] ^= a[44] ^ b[44]
	z[45] ^= a[45] ^ b[45]
	z[46] ^= a[46] ^ b[46]
	z[47] ^= a[47] ^ b[47]
	z[48] ^= a[48] ^ b[48]
	z[49] ^= a[49] ^ b[49]
	z[50] ^= a[50] ^ b[50]
	z[51] ^= a[51] ^ b[51]
	z[52] ^= a[52] ^ b[52]
	z[53] ^= a[53] ^ b[53]
	z[54] ^= a[54] ^ b[54]
	z[55] ^= a[55] ^ b[55]
	z[56] ^= a[56] ^ b[56]
	z[57] ^= a[57] ^ b[57]
	z[58] ^= a[58] ^ b[58]
	z[59] ^= a[59] ^ b[59]
	z[60] ^= a[60] ^ b[60]
	z[61] ^= a[61] ^ b[61]
	z[62] ^= a[62] ^ b[62]
	z[63] ^= a[63] ^ b[63]
	z[64] ^= a[64] ^ b[64]
	z[65] ^= a[65] ^ b[65]
	z[66] ^= a[66] ^ b[66]
	z[67] ^= a[67] ^ b[67]
	z[68] ^= a[68] ^ b[68]
	z[69] ^= a[69] ^ b[69]
	z[70] ^= a[70] ^ b[70]
	z[71] ^= a[71] ^ b[71]
	z[72] ^= a[72] ^ b[72]
	z[73] ^= a[73] ^ b[73]
	z[74] ^= a[74] ^ b[74]
	z[75] ^= a[75] ^ b[75]
	z[76] ^= a[76] ^ b[76]
	z[77] ^= a[77] ^ b[77]
	z[78] ^= a[78] ^ b[78]
	z[79] ^= a[79] ^ b[79]
	z[80] ^= a[80] ^ b[80]
	z[81] ^= a[81] ^ b[81]
	z[82] ^= a[82] ^ b[82]
	z[83] ^= a[83] ^ b[83]
	z[84] ^= a[84] ^ b[84]
	z[85] ^= a[85] ^ b[85]
	z[86] ^= a[86] ^ b[86]
	z[87] ^= a[87] ^ b[87]
	z[88] ^= a[88] ^ b[88]
	z[89] ^= a[89] ^ b[89]
	z[90] ^= a[90] ^ b[90]
	z[91] ^= a[91] ^ b[91]
	z[92] ^= a[92] ^ b[92]
	z[93] ^= a[93] ^ b[93]
	z[94] ^= a[94] ^ b[94]
	z[95] ^= a[95] ^ b[95]
	z[96] ^= a[96] ^ b[96]
	z[97] ^= a[97] ^ b[97]
	z[98] ^= a[98] ^ b[98]
	z[99] ^= a[99] ^ b[99]
	z[100] ^= a[100] ^ b[100]
	z[101] ^= a[101] ^ b[101]
	z[102] ^= a[102] ^ b[102]
	z[103] ^= a[103] ^ b[103]
	z[104] ^= a[104] ^ b[104]
	z[105] ^= a[105] ^ b[105]
	z[106] ^= a[106] ^ b[106]
	z[107] ^= a[107] ^ b[107]
	z[108] ^= a[108] ^ b[108]
	z[109] ^= a[109] ^ b[109]
	z[110] ^= a[110] ^ b[110]
	z[111] ^= a[111] ^ b[111]
	z[112] ^= a[112] ^ b[112]
	z[113] ^= a[113] ^ b[113]
	z[114] ^= a[114] ^ b[114]
	z[115] ^= a[115] ^ b[115]
	z[116] ^= a[116] ^ b[116]
	z[117] ^= a[117] ^ b[117]
	z[118] ^= a[118] ^ b[118]
	z[119] ^= a[119] ^ b[119]
	z[120] ^= a[120] ^ b[120]
	z[121] ^= a[121] ^ b[121]
	z[122] ^= a[122] ^ b[122]
	z[123] ^= a[123] ^ b[123]
	z[124] ^= a[124] ^ b[124]
	z[125] ^= a[125] ^ b[125]
	z[126] ^= a[126] ^ b[126]
	z[127] ^= a[127] ^ b[127]
}

func _P(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15 *uint64) {
	var v0 = *p0
	var v1 = *p1
	var v2 = *p2
	var v3 = *p3
	var v4 = *p4
	var v5 = *p5
	var v6 = *p6
	var v7 = *p7
	var v8 = *p8
	var v9 = *p9
	var v10 = *p10
	var v11 = *p11
	var v12 = *p12
	var v13 = *p13
	var v14 = *p14
	var v15 = *p15
	var t uint64
	t = uint64(uint32(v0)) * uint64(uint32(v4))
	v0 = v0 + v4 + t*2
	v12 = v12 ^ v0
	v12 = v12>>32 | v12<<32
	t = uint64(uint32(v8)) * uint64(uint32(v12))
	v8 = v8 + v12 + t*2
	v4 = v4 ^ v8
	v4 = v4>>24 | v4<<40
	t = uint64(uint32(v0)) * uint64(uint32(v4))
	v0 = v0 + v4 + t*2
	v12 = v12 ^ v0
	v12 = v12>>16 | v12<<48
	t = uint64(uint32(v8)) * uint64(uint32(v12))
	v8 = v8 + v12 + t*2
	v4 = v4 ^ v8
	v4 = v4>>63 | v4<<1
	t = uint64(uint32(v1)) * uint64(uint32(v5))
	v1 = v1 + v5 + t*2
	v13 = v13 ^ v1
	v13 = v13>>32 | v13<<32
	t = uint64(uint32(v9)) * uint64(uint32(v13))
	v9 = v9 + v13 + t*2
	v5 = v5 ^ v9
	v5 = v5>>24 | v5<<40
	t = uint64(uint32(v1)) * uint64(uint32(v5))
	v1 = v1 + v5 + t*2
	v13 = v13 ^ v1
	v13 = v13>>16 | v13<<48
	t = uint64(uint32(v9)) * uint64(uint32(v13))
	v9 = v9 + v13 + t*2
	v5 = v5 ^ v9
	v5 = v5>>63 | v5<<1
	t = uint64(uint32(v2)) * uint64(uint32(v6))
	v2 = v2 + v6 + t*2
	v14 = v14 ^ v2
	v14 = v14>>32 | v14<<32
	t = uint64(uint32(v10)) * uint64(uint32(v14))
	v10 = v10 + v14 + t*2
	v6 = v6 ^ v10
	v6 = v6>>24 | v6<<40
	t = uint64(uint32(v2)) * uint64(uint32(v6))
	v2 = v2 + v6 + t*2
	v14 = v14 ^ v2
	v14 = v14>>16 | v14<<48
	t = uint64(uint32(v10)) * uint64(uint32(v14))
	v10 = v10 + v14 + t*2
	v6 = v6 ^ v10
	v6 = v6>>63 | v6<<1
	t = uint64(uint32(v3)) * uint64(uint32(v7))
	v3 = v3 + v7 + t*2
	v15 = v15 ^ v3
	v15 = v15>>32 | v15<<32
	t = uint64(uint32(v11)) * uint64(uint32(v15))
	v11 = v11 + v15 + t*2
	v7 = v7 ^ v11
	v7 = v7>>24 | v7<<40
	t = uint64(uint32(v3)) * uint64(uint32(v7))
	v3 = v3 + v7 + t*2
	v15 = v15 ^ v3
	v15 = v15>>16 | v15<<48
	t = uint64(uint32(v11)) * uint64(uint32(v15))
	v11 = v11 + v15 + t*2
	v7 = v7 ^ v11
	v7 = v7>>63 | v7<<1
	t = uint64(uint32(v0)) * uint64(uint32(v5))
	v0 = v0 + v5 + t*2
	v15 = v15 ^ v0
	v15 = v15>>32 | v15<<32
	t = uint64(uint32(v10)) * uint64(uint32(v15))
	v10 = v10 + v15 + t*2
	v5 = v5 ^ v10
	v5 = v5>>24 | v5<<40
	t = uint64(uint32(v0)) * uint64(uint32(v5))
	v0 = v0 + v5 + t*2
	v15 = v15 ^ v0
	v15 = v15>>16 | v15<<48
	t = uint64(uint32(v10)) * uint64(uint32(v15))
	v10 = v10 + v15 + t*2
	v5 = v5 ^ v10
	v5 = v5>>63 | v5<<1
	t = uint64(uint32(v1)) * uint64(uint32(v6))
	v1 = v1 + v6 + t*2
	v12 = v12 ^ v1
	v12 = v12>>32 | v12<<32
	t = uint64(uint32(v11)) * uint64(uint32(v12))
	v11 = v11 + v12 + t*2
	v6 = v6 ^ v11
	v6 = v6>>24 | v6<<40
	t = uint64(uint32(v1)) * uint64(uint32(v6))
	v1 = v1 + v6 + t*2
	v12 = v12 ^ v1
	v12 = v12>>16 | v12<<48
	t = uint64(uint32(v11)) * uint64(uint32(v12))
	v11 = v11 + v12 + t*2
	v6 = v6 ^ v11
	v6 = v6>>63 | v6<<1
	t = uint64(uint32(v2)) * uint64(uint32(v7))
	v2 = v2 + v7 + t*2
	v13 = v13 ^ v2
	v13 = v13>>32 | v13<<32
	t = uint64(uint32(v8)) * uint64(uint32(v13))
	v8 = v8 + v13 + t*2
	v7 = v7 ^ v8
	v7 = v7>>24 | v7<<40
	t = uint64(uint32(v2)) * uint64(uint32(v7))
	v2 = v2 + v7 + t*2
	v13 = v13 ^ v2
	v13 = v13>>16 | v13<<48
	t = uint64(uint32(v8)) * uint64(uint32(v13))
	v8 = v8 + v13 + t*2
	v7 = v7 ^ v8
	v7 = v7>>63 | v7<<1
	t = uint64(uint32(v3)) * uint64(uint32(v4))
	v3 = v3 + v4 + t*2
	v14 = v14 ^ v3
	v14 = v14>>32 | v14<<32
	t = uint64(uint32(v9)) * uint64(uint32(v14))
	v9 = v9 + v14 + t*2
	v4 = v4 ^ v9
	v4 = v4>>24 | v4<<40
	t = uint64(uint32(v3)) * uint64(uint32(v4))
	v3 = v3 + v4 + t*2
	v14 = v14 ^ v3
	v14 = v14>>16 | v14<<48
	t = uint64(uint32(v9)) * uint64(uint32(v14))
	v9 = v9 + v14 + t*2
	v4 = v4 ^ v9
	v4 = v4>>63 | v4<<1
	*p0 = v0
	*p1 = v1
	*p2 = v2
	*p3 = v3
	*p4 = v4
	*p5 = v5
	*p6 = v6
	*p7 = v7
	*p8 = v8
	*p9 = v9
	*p10 = v10
	*p11 = v11
	*p12 = v12
	*p13 = v13
	*p14 = v14
	*p15 = v15
}
