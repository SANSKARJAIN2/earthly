// Code generated from earthfile2llb/parser/EarthLexer.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 33, 625,
	8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5,
	9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4,
	11, 9, 11, 4, 12, 9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16,
	9, 16, 4, 17, 9, 17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9,
	21, 4, 22, 9, 22, 4, 23, 9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26,
	4, 27, 9, 27, 4, 28, 9, 28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4,
	32, 9, 32, 4, 33, 9, 33, 4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37,
	9, 37, 4, 38, 9, 38, 4, 39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9,
	42, 4, 43, 9, 43, 4, 44, 9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47,
	4, 48, 9, 48, 4, 49, 9, 49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4,
	53, 9, 53, 4, 54, 9, 54, 4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58,
	9, 58, 4, 59, 9, 59, 4, 60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9,
	63, 4, 64, 9, 64, 4, 65, 9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68,
	4, 69, 9, 69, 4, 70, 9, 70, 4, 71, 9, 71, 3, 2, 6, 2, 149, 10, 2, 13, 2,
	14, 2, 150, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 6,
	26, 380, 10, 26, 13, 26, 14, 26, 381, 3, 26, 3, 26, 3, 27, 5, 27, 387,
	10, 27, 3, 27, 5, 27, 390, 10, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 7,
	28, 397, 10, 28, 12, 28, 14, 28, 400, 11, 28, 3, 28, 6, 28, 403, 10, 28,
	13, 28, 14, 28, 404, 3, 29, 3, 29, 3, 29, 5, 29, 410, 10, 29, 3, 30, 3,
	30, 7, 30, 414, 10, 30, 12, 30, 14, 30, 417, 11, 30, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39,
	3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3,
	42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3,
	46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48,
	3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3,
	50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52,
	3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3,
	54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56,
	3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 6, 58, 554, 10, 58, 13, 58, 14,
	58, 555, 3, 59, 3, 59, 3, 59, 3, 59, 7, 59, 562, 10, 59, 12, 59, 14, 59,
	565, 11, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 61, 5, 61, 572, 10, 61, 3,
	61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3, 62, 3, 63, 3, 63,
	3, 63, 3, 63, 3, 64, 3, 64, 6, 64, 589, 10, 64, 13, 64, 14, 64, 590, 3,
	64, 3, 64, 3, 65, 3, 65, 3, 66, 5, 66, 598, 10, 66, 3, 66, 3, 66, 3, 66,
	3, 66, 3, 66, 3, 67, 3, 67, 3, 67, 3, 67, 3, 68, 3, 68, 3, 68, 3, 68, 3,
	69, 3, 69, 3, 69, 3, 69, 3, 70, 3, 70, 3, 70, 3, 70, 3, 70, 3, 71, 3, 71,
	3, 71, 3, 71, 2, 2, 72, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20,
	39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29,
	57, 30, 59, 31, 61, 2, 63, 2, 65, 2, 67, 2, 69, 2, 71, 2, 73, 2, 75, 2,
	77, 2, 79, 2, 81, 2, 83, 2, 85, 2, 87, 2, 89, 2, 91, 2, 93, 2, 95, 2, 97,
	2, 99, 2, 101, 2, 103, 2, 105, 2, 107, 2, 109, 2, 111, 2, 113, 2, 115,
	2, 117, 2, 119, 32, 121, 2, 123, 2, 125, 2, 127, 2, 129, 33, 131, 2, 133,
	2, 135, 2, 137, 2, 139, 2, 141, 2, 143, 2, 145, 2, 7, 2, 3, 4, 5, 6, 9,
	6, 2, 47, 48, 50, 59, 67, 92, 99, 124, 3, 2, 67, 92, 4, 2, 11, 11, 34,
	34, 4, 2, 12, 12, 15, 15, 3, 2, 36, 36, 6, 2, 11, 12, 15, 15, 34, 34, 36,
	36, 7, 2, 11, 12, 15, 15, 34, 34, 36, 36, 63, 63, 2, 632, 2, 7, 3, 2, 2,
	2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2,
	2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3,
	2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31,
	3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2,
	39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2,
	2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2,
	2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 3, 65, 3, 2,
	2, 2, 3, 67, 3, 2, 2, 2, 3, 69, 3, 2, 2, 2, 3, 71, 3, 2, 2, 2, 3, 73, 3,
	2, 2, 2, 3, 75, 3, 2, 2, 2, 3, 77, 3, 2, 2, 2, 3, 79, 3, 2, 2, 2, 3, 81,
	3, 2, 2, 2, 3, 83, 3, 2, 2, 2, 3, 85, 3, 2, 2, 2, 3, 87, 3, 2, 2, 2, 3,
	89, 3, 2, 2, 2, 3, 91, 3, 2, 2, 2, 3, 93, 3, 2, 2, 2, 3, 95, 3, 2, 2, 2,
	3, 97, 3, 2, 2, 2, 3, 99, 3, 2, 2, 2, 3, 101, 3, 2, 2, 2, 3, 103, 3, 2,
	2, 2, 3, 105, 3, 2, 2, 2, 3, 107, 3, 2, 2, 2, 3, 109, 3, 2, 2, 2, 3, 111,
	3, 2, 2, 2, 3, 113, 3, 2, 2, 2, 3, 115, 3, 2, 2, 2, 3, 117, 3, 2, 2, 2,
	4, 119, 3, 2, 2, 2, 4, 125, 3, 2, 2, 2, 4, 127, 3, 2, 2, 2, 5, 129, 3,
	2, 2, 2, 5, 131, 3, 2, 2, 2, 5, 135, 3, 2, 2, 2, 5, 137, 3, 2, 2, 2, 6,
	139, 3, 2, 2, 2, 6, 141, 3, 2, 2, 2, 6, 143, 3, 2, 2, 2, 6, 145, 3, 2,
	2, 2, 7, 148, 3, 2, 2, 2, 9, 156, 3, 2, 2, 2, 11, 163, 3, 2, 2, 2, 13,
	170, 3, 2, 2, 2, 15, 186, 3, 2, 2, 2, 17, 199, 3, 2, 2, 2, 19, 205, 3,
	2, 2, 2, 21, 214, 3, 2, 2, 2, 23, 223, 3, 2, 2, 2, 25, 229, 3, 2, 2, 2,
	27, 235, 3, 2, 2, 2, 29, 243, 3, 2, 2, 2, 31, 251, 3, 2, 2, 2, 33, 261,
	3, 2, 2, 2, 35, 268, 3, 2, 2, 2, 37, 274, 3, 2, 2, 2, 39, 287, 3, 2, 2,
	2, 41, 299, 3, 2, 2, 2, 43, 313, 3, 2, 2, 2, 45, 327, 3, 2, 2, 2, 47, 333,
	3, 2, 2, 2, 49, 346, 3, 2, 2, 2, 51, 356, 3, 2, 2, 2, 53, 370, 3, 2, 2,
	2, 55, 379, 3, 2, 2, 2, 57, 386, 3, 2, 2, 2, 59, 402, 3, 2, 2, 2, 61, 409,
	3, 2, 2, 2, 63, 411, 3, 2, 2, 2, 65, 418, 3, 2, 2, 2, 67, 423, 3, 2, 2,
	2, 69, 428, 3, 2, 2, 2, 71, 433, 3, 2, 2, 2, 73, 438, 3, 2, 2, 2, 75, 443,
	3, 2, 2, 2, 77, 448, 3, 2, 2, 2, 79, 453, 3, 2, 2, 2, 81, 458, 3, 2, 2,
	2, 83, 463, 3, 2, 2, 2, 85, 468, 3, 2, 2, 2, 87, 473, 3, 2, 2, 2, 89, 478,
	3, 2, 2, 2, 91, 483, 3, 2, 2, 2, 93, 488, 3, 2, 2, 2, 95, 493, 3, 2, 2,
	2, 97, 498, 3, 2, 2, 2, 99, 503, 3, 2, 2, 2, 101, 508, 3, 2, 2, 2, 103,
	513, 3, 2, 2, 2, 105, 518, 3, 2, 2, 2, 107, 523, 3, 2, 2, 2, 109, 528,
	3, 2, 2, 2, 111, 533, 3, 2, 2, 2, 113, 538, 3, 2, 2, 2, 115, 543, 3, 2,
	2, 2, 117, 547, 3, 2, 2, 2, 119, 553, 3, 2, 2, 2, 121, 557, 3, 2, 2, 2,
	123, 568, 3, 2, 2, 2, 125, 571, 3, 2, 2, 2, 127, 578, 3, 2, 2, 2, 129,
	582, 3, 2, 2, 2, 131, 588, 3, 2, 2, 2, 133, 594, 3, 2, 2, 2, 135, 597,
	3, 2, 2, 2, 137, 604, 3, 2, 2, 2, 139, 608, 3, 2, 2, 2, 141, 612, 3, 2,
	2, 2, 143, 616, 3, 2, 2, 2, 145, 621, 3, 2, 2, 2, 147, 149, 9, 2, 2, 2,
	148, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 150,
	151, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 153, 7, 60, 2, 2, 153, 154,
	3, 2, 2, 2, 154, 155, 8, 2, 2, 2, 155, 8, 3, 2, 2, 2, 156, 157, 7, 72,
	2, 2, 157, 158, 7, 84, 2, 2, 158, 159, 7, 81, 2, 2, 159, 160, 7, 79, 2,
	2, 160, 161, 3, 2, 2, 2, 161, 162, 8, 3, 3, 2, 162, 10, 3, 2, 2, 2, 163,
	164, 7, 69, 2, 2, 164, 165, 7, 81, 2, 2, 165, 166, 7, 82, 2, 2, 166, 167,
	7, 91, 2, 2, 167, 168, 3, 2, 2, 2, 168, 169, 8, 4, 3, 2, 169, 12, 3, 2,
	2, 2, 170, 171, 7, 85, 2, 2, 171, 172, 7, 67, 2, 2, 172, 173, 7, 88, 2,
	2, 173, 174, 7, 71, 2, 2, 174, 175, 7, 34, 2, 2, 175, 176, 7, 67, 2, 2,
	176, 177, 7, 84, 2, 2, 177, 178, 7, 86, 2, 2, 178, 179, 7, 75, 2, 2, 179,
	180, 7, 72, 2, 2, 180, 181, 7, 67, 2, 2, 181, 182, 7, 69, 2, 2, 182, 183,
	7, 86, 2, 2, 183, 184, 3, 2, 2, 2, 184, 185, 8, 5, 3, 2, 185, 14, 3, 2,
	2, 2, 186, 187, 7, 85, 2, 2, 187, 188, 7, 67, 2, 2, 188, 189, 7, 88, 2,
	2, 189, 190, 7, 71, 2, 2, 190, 191, 7, 34, 2, 2, 191, 192, 7, 75, 2, 2,
	192, 193, 7, 79, 2, 2, 193, 194, 7, 67, 2, 2, 194, 195, 7, 73, 2, 2, 195,
	196, 7, 71, 2, 2, 196, 197, 3, 2, 2, 2, 197, 198, 8, 6, 3, 2, 198, 16,
	3, 2, 2, 2, 199, 200, 7, 84, 2, 2, 200, 201, 7, 87, 2, 2, 201, 202, 7,
	80, 2, 2, 202, 203, 3, 2, 2, 2, 203, 204, 8, 7, 3, 2, 204, 18, 3, 2, 2,
	2, 205, 206, 7, 71, 2, 2, 206, 207, 7, 90, 2, 2, 207, 208, 7, 82, 2, 2,
	208, 209, 7, 81, 2, 2, 209, 210, 7, 85, 2, 2, 210, 211, 7, 71, 2, 2, 211,
	212, 3, 2, 2, 2, 212, 213, 8, 8, 3, 2, 213, 20, 3, 2, 2, 2, 214, 215, 7,
	88, 2, 2, 215, 216, 7, 81, 2, 2, 216, 217, 7, 78, 2, 2, 217, 218, 7, 87,
	2, 2, 218, 219, 7, 79, 2, 2, 219, 220, 7, 71, 2, 2, 220, 221, 3, 2, 2,
	2, 221, 222, 8, 9, 3, 2, 222, 22, 3, 2, 2, 2, 223, 224, 7, 71, 2, 2, 224,
	225, 7, 80, 2, 2, 225, 226, 7, 88, 2, 2, 226, 227, 3, 2, 2, 2, 227, 228,
	8, 10, 4, 2, 228, 24, 3, 2, 2, 2, 229, 230, 7, 67, 2, 2, 230, 231, 7, 84,
	2, 2, 231, 232, 7, 73, 2, 2, 232, 233, 3, 2, 2, 2, 233, 234, 8, 11, 4,
	2, 234, 26, 3, 2, 2, 2, 235, 236, 7, 78, 2, 2, 236, 237, 7, 67, 2, 2, 237,
	238, 7, 68, 2, 2, 238, 239, 7, 71, 2, 2, 239, 240, 7, 78, 2, 2, 240, 241,
	3, 2, 2, 2, 241, 242, 8, 12, 5, 2, 242, 28, 3, 2, 2, 2, 243, 244, 7, 68,
	2, 2, 244, 245, 7, 87, 2, 2, 245, 246, 7, 75, 2, 2, 246, 247, 7, 78, 2,
	2, 247, 248, 7, 70, 2, 2, 248, 249, 3, 2, 2, 2, 249, 250, 8, 13, 3, 2,
	250, 30, 3, 2, 2, 2, 251, 252, 7, 89, 2, 2, 252, 253, 7, 81, 2, 2, 253,
	254, 7, 84, 2, 2, 254, 255, 7, 77, 2, 2, 255, 256, 7, 70, 2, 2, 256, 257,
	7, 75, 2, 2, 257, 258, 7, 84, 2, 2, 258, 259, 3, 2, 2, 2, 259, 260, 8,
	14, 3, 2, 260, 32, 3, 2, 2, 2, 261, 262, 7, 87, 2, 2, 262, 263, 7, 85,
	2, 2, 263, 264, 7, 71, 2, 2, 264, 265, 7, 84, 2, 2, 265, 266, 3, 2, 2,
	2, 266, 267, 8, 15, 3, 2, 267, 34, 3, 2, 2, 2, 268, 269, 7, 69, 2, 2, 269,
	270, 7, 79, 2, 2, 270, 271, 7, 70, 2, 2, 271, 272, 3, 2, 2, 2, 272, 273,
	8, 16, 3, 2, 273, 36, 3, 2, 2, 2, 274, 275, 7, 71, 2, 2, 275, 276, 7, 80,
	2, 2, 276, 277, 7, 86, 2, 2, 277, 278, 7, 84, 2, 2, 278, 279, 7, 91, 2,
	2, 279, 280, 7, 82, 2, 2, 280, 281, 7, 81, 2, 2, 281, 282, 7, 75, 2, 2,
	282, 283, 7, 80, 2, 2, 283, 284, 7, 86, 2, 2, 284, 285, 3, 2, 2, 2, 285,
	286, 8, 17, 3, 2, 286, 38, 3, 2, 2, 2, 287, 288, 7, 73, 2, 2, 288, 289,
	7, 75, 2, 2, 289, 290, 7, 86, 2, 2, 290, 291, 7, 34, 2, 2, 291, 292, 7,
	69, 2, 2, 292, 293, 7, 78, 2, 2, 293, 294, 7, 81, 2, 2, 294, 295, 7, 80,
	2, 2, 295, 296, 7, 71, 2, 2, 296, 297, 3, 2, 2, 2, 297, 298, 8, 18, 3,
	2, 298, 40, 3, 2, 2, 2, 299, 300, 7, 70, 2, 2, 300, 301, 7, 81, 2, 2, 301,
	302, 7, 69, 2, 2, 302, 303, 7, 77, 2, 2, 303, 304, 7, 71, 2, 2, 304, 305,
	7, 84, 2, 2, 305, 306, 7, 34, 2, 2, 306, 307, 7, 78, 2, 2, 307, 308, 7,
	81, 2, 2, 308, 309, 7, 67, 2, 2, 309, 310, 7, 70, 2, 2, 310, 311, 3, 2,
	2, 2, 311, 312, 8, 19, 3, 2, 312, 42, 3, 2, 2, 2, 313, 314, 7, 70, 2, 2,
	314, 315, 7, 81, 2, 2, 315, 316, 7, 69, 2, 2, 316, 317, 7, 77, 2, 2, 317,
	318, 7, 71, 2, 2, 318, 319, 7, 84, 2, 2, 319, 320, 7, 34, 2, 2, 320, 321,
	7, 82, 2, 2, 321, 322, 7, 87, 2, 2, 322, 323, 7, 78, 2, 2, 323, 324, 7,
	78, 2, 2, 324, 325, 3, 2, 2, 2, 325, 326, 8, 20, 3, 2, 326, 44, 3, 2, 2,
	2, 327, 328, 7, 67, 2, 2, 328, 329, 7, 70, 2, 2, 329, 330, 7, 70, 2, 2,
	330, 331, 3, 2, 2, 2, 331, 332, 8, 21, 3, 2, 332, 46, 3, 2, 2, 2, 333,
	334, 7, 85, 2, 2, 334, 335, 7, 86, 2, 2, 335, 336, 7, 81, 2, 2, 336, 337,
	7, 82, 2, 2, 337, 338, 7, 85, 2, 2, 338, 339, 7, 75, 2, 2, 339, 340, 7,
	73, 2, 2, 340, 341, 7, 80, 2, 2, 341, 342, 7, 67, 2, 2, 342, 343, 7, 78,
	2, 2, 343, 344, 3, 2, 2, 2, 344, 345, 8, 22, 3, 2, 345, 48, 3, 2, 2, 2,
	346, 347, 7, 81, 2, 2, 347, 348, 7, 80, 2, 2, 348, 349, 7, 68, 2, 2, 349,
	350, 7, 87, 2, 2, 350, 351, 7, 75, 2, 2, 351, 352, 7, 78, 2, 2, 352, 353,
	7, 70, 2, 2, 353, 354, 3, 2, 2, 2, 354, 355, 8, 23, 3, 2, 355, 50, 3, 2,
	2, 2, 356, 357, 7, 74, 2, 2, 357, 358, 7, 71, 2, 2, 358, 359, 7, 67, 2,
	2, 359, 360, 7, 78, 2, 2, 360, 361, 7, 86, 2, 2, 361, 362, 7, 74, 2, 2,
	362, 363, 7, 69, 2, 2, 363, 364, 7, 74, 2, 2, 364, 365, 7, 71, 2, 2, 365,
	366, 7, 69, 2, 2, 366, 367, 7, 77, 2, 2, 367, 368, 3, 2, 2, 2, 368, 369,
	8, 24, 3, 2, 369, 52, 3, 2, 2, 2, 370, 371, 7, 85, 2, 2, 371, 372, 7, 74,
	2, 2, 372, 373, 7, 71, 2, 2, 373, 374, 7, 78, 2, 2, 374, 375, 7, 78, 2,
	2, 375, 376, 3, 2, 2, 2, 376, 377, 8, 25, 3, 2, 377, 54, 3, 2, 2, 2, 378,
	380, 9, 3, 2, 2, 379, 378, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 379,
	3, 2, 2, 2, 381, 382, 3, 2, 2, 2, 382, 383, 3, 2, 2, 2, 383, 384, 8, 26,
	3, 2, 384, 56, 3, 2, 2, 2, 385, 387, 5, 59, 28, 2, 386, 385, 3, 2, 2, 2,
	386, 387, 3, 2, 2, 2, 387, 389, 3, 2, 2, 2, 388, 390, 5, 63, 30, 2, 389,
	388, 3, 2, 2, 2, 389, 390, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391, 392,
	5, 61, 29, 2, 392, 58, 3, 2, 2, 2, 393, 403, 9, 4, 2, 2, 394, 398, 7, 94,
	2, 2, 395, 397, 9, 4, 2, 2, 396, 395, 3, 2, 2, 2, 397, 400, 3, 2, 2, 2,
	398, 396, 3, 2, 2, 2, 398, 399, 3, 2, 2, 2, 399, 401, 3, 2, 2, 2, 400,
	398, 3, 2, 2, 2, 401, 403, 5, 61, 29, 2, 402, 393, 3, 2, 2, 2, 402, 394,
	3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404, 402, 3, 2, 2, 2, 404, 405, 3, 2,
	2, 2, 405, 60, 3, 2, 2, 2, 406, 410, 9, 5, 2, 2, 407, 408, 7, 15, 2, 2,
	408, 410, 7, 12, 2, 2, 409, 406, 3, 2, 2, 2, 409, 407, 3, 2, 2, 2, 410,
	62, 3, 2, 2, 2, 411, 415, 7, 37, 2, 2, 412, 414, 10, 5, 2, 2, 413, 412,
	3, 2, 2, 2, 414, 417, 3, 2, 2, 2, 415, 413, 3, 2, 2, 2, 415, 416, 3, 2,
	2, 2, 416, 64, 3, 2, 2, 2, 417, 415, 3, 2, 2, 2, 418, 419, 5, 7, 2, 2,
	419, 420, 3, 2, 2, 2, 420, 421, 8, 31, 6, 2, 421, 422, 8, 31, 2, 2, 422,
	66, 3, 2, 2, 2, 423, 424, 5, 9, 3, 2, 424, 425, 3, 2, 2, 2, 425, 426, 8,
	32, 7, 2, 426, 427, 8, 32, 3, 2, 427, 68, 3, 2, 2, 2, 428, 429, 5, 11,
	4, 2, 429, 430, 3, 2, 2, 2, 430, 431, 8, 33, 8, 2, 431, 432, 8, 33, 3,
	2, 432, 70, 3, 2, 2, 2, 433, 434, 5, 13, 5, 2, 434, 435, 3, 2, 2, 2, 435,
	436, 8, 34, 9, 2, 436, 437, 8, 34, 3, 2, 437, 72, 3, 2, 2, 2, 438, 439,
	5, 15, 6, 2, 439, 440, 3, 2, 2, 2, 440, 441, 8, 35, 10, 2, 441, 442, 8,
	35, 3, 2, 442, 74, 3, 2, 2, 2, 443, 444, 5, 17, 7, 2, 444, 445, 3, 2, 2,
	2, 445, 446, 8, 36, 11, 2, 446, 447, 8, 36, 3, 2, 447, 76, 3, 2, 2, 2,
	448, 449, 5, 19, 8, 2, 449, 450, 3, 2, 2, 2, 450, 451, 8, 37, 12, 2, 451,
	452, 8, 37, 3, 2, 452, 78, 3, 2, 2, 2, 453, 454, 5, 21, 9, 2, 454, 455,
	3, 2, 2, 2, 455, 456, 8, 38, 13, 2, 456, 457, 8, 38, 3, 2, 457, 80, 3,
	2, 2, 2, 458, 459, 5, 23, 10, 2, 459, 460, 3, 2, 2, 2, 460, 461, 8, 39,
	14, 2, 461, 462, 8, 39, 4, 2, 462, 82, 3, 2, 2, 2, 463, 464, 5, 25, 11,
	2, 464, 465, 3, 2, 2, 2, 465, 466, 8, 40, 15, 2, 466, 467, 8, 40, 4, 2,
	467, 84, 3, 2, 2, 2, 468, 469, 5, 27, 12, 2, 469, 470, 3, 2, 2, 2, 470,
	471, 8, 41, 16, 2, 471, 472, 8, 41, 5, 2, 472, 86, 3, 2, 2, 2, 473, 474,
	5, 29, 13, 2, 474, 475, 3, 2, 2, 2, 475, 476, 8, 42, 17, 2, 476, 477, 8,
	42, 3, 2, 477, 88, 3, 2, 2, 2, 478, 479, 5, 31, 14, 2, 479, 480, 3, 2,
	2, 2, 480, 481, 8, 43, 18, 2, 481, 482, 8, 43, 3, 2, 482, 90, 3, 2, 2,
	2, 483, 484, 5, 33, 15, 2, 484, 485, 3, 2, 2, 2, 485, 486, 8, 44, 19, 2,
	486, 487, 8, 44, 3, 2, 487, 92, 3, 2, 2, 2, 488, 489, 5, 35, 16, 2, 489,
	490, 3, 2, 2, 2, 490, 491, 8, 45, 20, 2, 491, 492, 8, 45, 3, 2, 492, 94,
	3, 2, 2, 2, 493, 494, 5, 37, 17, 2, 494, 495, 3, 2, 2, 2, 495, 496, 8,
	46, 21, 2, 496, 497, 8, 46, 3, 2, 497, 96, 3, 2, 2, 2, 498, 499, 5, 39,
	18, 2, 499, 500, 3, 2, 2, 2, 500, 501, 8, 47, 22, 2, 501, 502, 8, 47, 3,
	2, 502, 98, 3, 2, 2, 2, 503, 504, 5, 41, 19, 2, 504, 505, 3, 2, 2, 2, 505,
	506, 8, 48, 23, 2, 506, 507, 8, 48, 3, 2, 507, 100, 3, 2, 2, 2, 508, 509,
	5, 43, 20, 2, 509, 510, 3, 2, 2, 2, 510, 511, 8, 49, 24, 2, 511, 512, 8,
	49, 3, 2, 512, 102, 3, 2, 2, 2, 513, 514, 5, 45, 21, 2, 514, 515, 3, 2,
	2, 2, 515, 516, 8, 50, 25, 2, 516, 517, 8, 50, 3, 2, 517, 104, 3, 2, 2,
	2, 518, 519, 5, 47, 22, 2, 519, 520, 3, 2, 2, 2, 520, 521, 8, 51, 26, 2,
	521, 522, 8, 51, 3, 2, 522, 106, 3, 2, 2, 2, 523, 524, 5, 49, 23, 2, 524,
	525, 3, 2, 2, 2, 525, 526, 8, 52, 27, 2, 526, 527, 8, 52, 3, 2, 527, 108,
	3, 2, 2, 2, 528, 529, 5, 51, 24, 2, 529, 530, 3, 2, 2, 2, 530, 531, 8,
	53, 28, 2, 531, 532, 8, 53, 3, 2, 532, 110, 3, 2, 2, 2, 533, 534, 5, 53,
	25, 2, 534, 535, 3, 2, 2, 2, 535, 536, 8, 54, 29, 2, 536, 537, 8, 54, 3,
	2, 537, 112, 3, 2, 2, 2, 538, 539, 5, 55, 26, 2, 539, 540, 3, 2, 2, 2,
	540, 541, 8, 55, 30, 2, 541, 542, 8, 55, 3, 2, 542, 114, 3, 2, 2, 2, 543,
	544, 5, 57, 27, 2, 544, 545, 3, 2, 2, 2, 545, 546, 8, 56, 31, 2, 546, 116,
	3, 2, 2, 2, 547, 548, 5, 59, 28, 2, 548, 549, 3, 2, 2, 2, 549, 550, 8,
	57, 32, 2, 550, 118, 3, 2, 2, 2, 551, 554, 5, 123, 60, 2, 552, 554, 5,
	121, 59, 2, 553, 551, 3, 2, 2, 2, 553, 552, 3, 2, 2, 2, 554, 555, 3, 2,
	2, 2, 555, 553, 3, 2, 2, 2, 555, 556, 3, 2, 2, 2, 556, 120, 3, 2, 2, 2,
	557, 563, 7, 36, 2, 2, 558, 562, 10, 6, 2, 2, 559, 560, 7, 94, 2, 2, 560,
	562, 7, 36, 2, 2, 561, 558, 3, 2, 2, 2, 561, 559, 3, 2, 2, 2, 562, 565,
	3, 2, 2, 2, 563, 561, 3, 2, 2, 2, 563, 564, 3, 2, 2, 2, 564, 566, 3, 2,
	2, 2, 565, 563, 3, 2, 2, 2, 566, 567, 7, 36, 2, 2, 567, 122, 3, 2, 2, 2,
	568, 569, 10, 7, 2, 2, 569, 124, 3, 2, 2, 2, 570, 572, 5, 59, 28, 2, 571,
	570, 3, 2, 2, 2, 571, 572, 3, 2, 2, 2, 572, 573, 3, 2, 2, 2, 573, 574,
	5, 61, 29, 2, 574, 575, 3, 2, 2, 2, 575, 576, 8, 61, 31, 2, 576, 577, 8,
	61, 33, 2, 577, 126, 3, 2, 2, 2, 578, 579, 5, 59, 28, 2, 579, 580, 3, 2,
	2, 2, 580, 581, 8, 62, 32, 2, 581, 128, 3, 2, 2, 2, 582, 583, 7, 63, 2,
	2, 583, 584, 3, 2, 2, 2, 584, 585, 8, 63, 34, 2, 585, 130, 3, 2, 2, 2,
	586, 589, 5, 133, 65, 2, 587, 589, 5, 121, 59, 2, 588, 586, 3, 2, 2, 2,
	588, 587, 3, 2, 2, 2, 589, 590, 3, 2, 2, 2, 590, 588, 3, 2, 2, 2, 590,
	591, 3, 2, 2, 2, 591, 592, 3, 2, 2, 2, 592, 593, 8, 64, 35, 2, 593, 132,
	3, 2, 2, 2, 594, 595, 10, 8, 2, 2, 595, 134, 3, 2, 2, 2, 596, 598, 5, 59,
	28, 2, 597, 596, 3, 2, 2, 2, 597, 598, 3, 2, 2, 2, 598, 599, 3, 2, 2, 2,
	599, 600, 5, 61, 29, 2, 600, 601, 3, 2, 2, 2, 601, 602, 8, 66, 31, 2, 602,
	603, 8, 66, 33, 2, 603, 136, 3, 2, 2, 2, 604, 605, 5, 59, 28, 2, 605, 606,
	3, 2, 2, 2, 606, 607, 8, 67, 32, 2, 607, 138, 3, 2, 2, 2, 608, 609, 7,
	63, 2, 2, 609, 610, 3, 2, 2, 2, 610, 611, 8, 68, 36, 2, 611, 140, 3, 2,
	2, 2, 612, 613, 5, 131, 64, 2, 613, 614, 3, 2, 2, 2, 614, 615, 8, 69, 35,
	2, 615, 142, 3, 2, 2, 2, 616, 617, 5, 135, 66, 2, 617, 618, 3, 2, 2, 2,
	618, 619, 8, 70, 31, 2, 619, 620, 8, 70, 33, 2, 620, 144, 3, 2, 2, 2, 621,
	622, 5, 137, 67, 2, 622, 623, 3, 2, 2, 2, 623, 624, 8, 71, 32, 2, 624,
	146, 3, 2, 2, 2, 25, 2, 3, 4, 5, 6, 148, 150, 381, 386, 389, 398, 402,
	404, 409, 415, 553, 555, 561, 563, 571, 588, 590, 597, 37, 7, 3, 2, 7,
	4, 2, 7, 5, 2, 7, 6, 2, 9, 5, 2, 9, 6, 2, 9, 7, 2, 9, 8, 2, 9, 9, 2, 9,
	10, 2, 9, 11, 2, 9, 12, 2, 9, 13, 2, 9, 14, 2, 9, 15, 2, 9, 16, 2, 9, 17,
	2, 9, 18, 2, 9, 19, 2, 9, 20, 2, 9, 21, 2, 9, 22, 2, 9, 23, 2, 9, 24, 2,
	9, 25, 2, 9, 26, 2, 9, 27, 2, 9, 28, 2, 9, 29, 2, 9, 30, 2, 9, 31, 2, 6,
	2, 2, 4, 4, 2, 9, 32, 2, 9, 33, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "RECIPE", "COMMAND_ARGS", "COMMAND_ARGS_KEY_VALUE", "COMMAND_ARGS_KEY_VALUE_LABEL",
}

var lexerLiteralNames = []string{
	"", "", "", "", "'FROM'", "'COPY'", "'SAVE ARTIFACT'", "'SAVE IMAGE'",
	"'RUN'", "'EXPOSE'", "'VOLUME'", "'ENV'", "'ARG'", "'LABEL'", "'BUILD'",
	"'WORKDIR'", "'USER'", "'CMD'", "'ENTRYPOINT'", "'GIT CLONE'", "'DOCKER LOAD'",
	"'DOCKER PULL'", "'ADD'", "'STOPSIGNAL'", "'ONBUILD'", "'HEALTHCHECK'",
	"'SHELL'",
}

var lexerSymbolicNames = []string{
	"", "INDENT", "DEDENT", "Target", "FROM", "COPY", "SAVE_ARTIFACT", "SAVE_IMAGE",
	"RUN", "EXPOSE", "VOLUME", "ENV", "ARG", "LABEL", "BUILD", "WORKDIR", "USER",
	"CMD", "ENTRYPOINT", "GIT_CLONE", "DOCKER_LOAD", "DOCKER_PULL", "ADD",
	"STOPSIGNAL", "ONBUILD", "HEALTHCHECK", "SHELL", "Command", "NL", "WS",
	"Atom", "EQUALS",
}

var lexerRuleNames = []string{
	"Target", "FROM", "COPY", "SAVE_ARTIFACT", "SAVE_IMAGE", "RUN", "EXPOSE",
	"VOLUME", "ENV", "ARG", "LABEL", "BUILD", "WORKDIR", "USER", "CMD", "ENTRYPOINT",
	"GIT_CLONE", "DOCKER_LOAD", "DOCKER_PULL", "ADD", "STOPSIGNAL", "ONBUILD",
	"HEALTHCHECK", "SHELL", "Command", "NL", "WS", "CRLF", "COMMENT", "Target_R",
	"FROM_R", "COPY_R", "SAVE_ARTIFACT_R", "SAVE_IMAGE_R", "RUN_R", "EXPOSE_R",
	"VOLUME_R", "ENV_R", "ARG_R", "LABEL_R", "BUILD_R", "WORKDIR_R", "USER_R",
	"CMD_R", "ENTRYPOINT_R", "GIT_CLONE_R", "DOCKER_LOAD_R", "DOCKER_PULL_R",
	"ADD_R", "STOPSIGNAL_R", "ONBUILD_R", "HEALTHCHECK_R", "SHELL_R", "Command_R",
	"NL_R", "WS_R", "Atom", "QuotedAtom", "NonWSNLQuote", "NL_C", "WS_C", "EQUALS",
	"Atom_CAKV", "NonWSNLQuote_CAKV", "NL_CAKV", "WS_CAKV", "EQUALS_L", "Atom_CAKVL",
	"NL_CAKVL", "WS_CAKVL",
}

type EarthLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewEarthLexer(input antlr.CharStream) *EarthLexer {

	l := new(EarthLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "EarthLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EarthLexer tokens.
const (
	EarthLexerINDENT        = 1
	EarthLexerDEDENT        = 2
	EarthLexerTarget        = 3
	EarthLexerFROM          = 4
	EarthLexerCOPY          = 5
	EarthLexerSAVE_ARTIFACT = 6
	EarthLexerSAVE_IMAGE    = 7
	EarthLexerRUN           = 8
	EarthLexerEXPOSE        = 9
	EarthLexerVOLUME        = 10
	EarthLexerENV           = 11
	EarthLexerARG           = 12
	EarthLexerLABEL         = 13
	EarthLexerBUILD         = 14
	EarthLexerWORKDIR       = 15
	EarthLexerUSER          = 16
	EarthLexerCMD           = 17
	EarthLexerENTRYPOINT    = 18
	EarthLexerGIT_CLONE     = 19
	EarthLexerDOCKER_LOAD   = 20
	EarthLexerDOCKER_PULL   = 21
	EarthLexerADD           = 22
	EarthLexerSTOPSIGNAL    = 23
	EarthLexerONBUILD       = 24
	EarthLexerHEALTHCHECK   = 25
	EarthLexerSHELL         = 26
	EarthLexerCommand       = 27
	EarthLexerNL            = 28
	EarthLexerWS            = 29
	EarthLexerAtom          = 30
	EarthLexerEQUALS        = 31
)

// EarthLexer modes.
const (
	EarthLexerRECIPE = iota + 1
	EarthLexerCOMMAND_ARGS
	EarthLexerCOMMAND_ARGS_KEY_VALUE
	EarthLexerCOMMAND_ARGS_KEY_VALUE_LABEL
)
