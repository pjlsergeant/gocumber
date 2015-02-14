
//line ragel/lexer_go.rl:1
package gocumber

import "fmt"
import "strings"
import "regexp"
import "strconv"

// USEFUL URLs:
//    https://github.com/bnoordhuis/ragel/blob/master/examples/go/url.rl
//    http://thingsaaronmade.com/blog/a-simple-intro-to-writing-a-lexer-with-ragel.html
//    https://raw.githubusercontent.com/cucumber/gherkin/master/ragel/lexer.java.rl.erb

  
//line ragel/lexer_go.rl:255


  // START: write data noerror;
  
//line gocumber/parser.go:22
var _lexer_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 1, 9, 1, 10, 1, 11, 
	1, 12, 1, 13, 1, 16, 1, 17, 
	1, 18, 1, 19, 1, 20, 1, 21, 
	1, 22, 1, 23, 2, 2, 18, 2, 
	3, 4, 2, 13, 0, 2, 14, 15, 
	2, 17, 0, 2, 17, 1, 2, 17, 
	16, 2, 17, 19, 2, 18, 6, 2, 
	18, 7, 2, 18, 8, 2, 18, 9, 
	2, 18, 10, 2, 18, 16, 2, 20, 
	21, 2, 22, 0, 2, 22, 1, 2, 
	22, 16, 2, 22, 19, 3, 4, 14, 
	15, 3, 5, 14, 15, 3, 11, 14, 
	15, 3, 12, 14, 15, 3, 13, 14, 
	15, 3, 14, 15, 18, 3, 17, 0, 
	11, 3, 17, 14, 15, 4, 2, 14, 
	15, 18, 4, 3, 4, 14, 15, 4, 
	17, 0, 14, 15, 5, 17, 0, 11, 
	14, 15, 
}

var _lexer_key_offsets []int16 = []int16{
	0, 0, 18, 35, 36, 37, 39, 41, 
	46, 51, 56, 61, 65, 69, 71, 72, 
	73, 74, 75, 76, 77, 78, 79, 80, 
	81, 82, 83, 84, 85, 86, 91, 98, 
	103, 104, 105, 107, 109, 110, 111, 112, 
	113, 114, 115, 116, 117, 118, 119, 120, 
	121, 134, 136, 138, 140, 142, 144, 146, 
	148, 150, 152, 154, 156, 158, 160, 162, 
	164, 166, 183, 184, 185, 186, 187, 188, 
	189, 190, 191, 192, 193, 200, 202, 204, 
	206, 208, 210, 212, 214, 215, 216, 217, 
	218, 219, 220, 221, 222, 223, 234, 236, 
	238, 240, 242, 244, 246, 248, 250, 252, 
	254, 256, 258, 260, 262, 264, 266, 268, 
	270, 272, 274, 276, 278, 280, 282, 284, 
	286, 288, 290, 292, 294, 296, 298, 300, 
	302, 304, 306, 308, 310, 312, 314, 316, 
	318, 320, 322, 325, 327, 329, 331, 333, 
	335, 337, 338, 339, 340, 341, 342, 343, 
	344, 345, 346, 347, 348, 350, 351, 352, 
	353, 354, 355, 356, 357, 358, 359, 360, 
	373, 375, 377, 379, 381, 383, 385, 387, 
	389, 391, 393, 395, 397, 399, 401, 403, 
	405, 407, 409, 411, 413, 415, 417, 419, 
	421, 423, 425, 427, 429, 431, 433, 435, 
	437, 439, 441, 443, 444, 445, 459, 461, 
	463, 465, 467, 469, 471, 473, 475, 477, 
	479, 481, 483, 485, 487, 489, 491, 493, 
	495, 497, 499, 501, 503, 505, 507, 509, 
	511, 513, 515, 517, 519, 521, 523, 525, 
	527, 529, 531, 533, 535, 537, 539, 541, 
	543, 545, 548, 550, 552, 554, 556, 558, 
	560, 562, 563, 567, 573, 576, 578, 584, 
	601, 603, 605, 607, 609, 611, 613, 615, 
	617, 619, 621, 623, 625, 627, 629, 631, 
	633, 635, 637, 640, 642, 644, 646, 648, 
	650, 652, 654, 655, 656, 
}

var _lexer_trans_keys []byte = []byte{
	10, 32, 34, 35, 37, 64, 65, 66, 
	69, 70, 71, 83, 84, 87, 124, 239, 
	9, 13, 10, 32, 34, 35, 37, 64, 
	65, 66, 69, 70, 71, 83, 84, 87, 
	124, 9, 13, 34, 34, 10, 13, 10, 
	13, 10, 32, 34, 9, 13, 10, 32, 
	34, 9, 13, 10, 32, 34, 9, 13, 
	10, 32, 34, 9, 13, 10, 32, 9, 
	13, 10, 32, 9, 13, 10, 13, 10, 
	95, 70, 69, 65, 84, 85, 82, 69, 
	95, 69, 78, 68, 95, 37, 13, 32, 
	64, 9, 10, 9, 10, 13, 32, 64, 
	11, 12, 10, 32, 64, 9, 13, 110, 
	100, 10, 13, 10, 13, 97, 99, 107, 
	103, 114, 111, 117, 110, 100, 58, 10, 
	10, 10, 32, 35, 37, 64, 65, 70, 
	71, 83, 84, 87, 9, 13, 10, 95, 
	10, 70, 10, 69, 10, 65, 10, 84, 
	10, 85, 10, 82, 10, 69, 10, 95, 
	10, 69, 10, 78, 10, 68, 10, 95, 
	10, 37, 10, 110, 10, 100, 10, 32, 
	34, 35, 37, 64, 65, 66, 69, 70, 
	71, 83, 84, 87, 124, 9, 13, 120, 
	97, 109, 112, 108, 101, 115, 58, 10, 
	10, 10, 32, 35, 70, 124, 9, 13, 
	10, 101, 10, 97, 10, 116, 10, 117, 
	10, 114, 10, 101, 10, 58, 101, 97, 
	116, 117, 114, 101, 58, 10, 10, 10, 
	32, 35, 37, 64, 66, 69, 70, 83, 
	9, 13, 10, 95, 10, 70, 10, 69, 
	10, 65, 10, 84, 10, 85, 10, 82, 
	10, 69, 10, 95, 10, 69, 10, 78, 
	10, 68, 10, 95, 10, 37, 10, 97, 
	10, 99, 10, 107, 10, 103, 10, 114, 
	10, 111, 10, 117, 10, 110, 10, 100, 
	10, 58, 10, 120, 10, 97, 10, 109, 
	10, 112, 10, 108, 10, 101, 10, 115, 
	10, 101, 10, 97, 10, 116, 10, 117, 
	10, 114, 10, 101, 10, 99, 10, 101, 
	10, 110, 10, 97, 10, 114, 10, 105, 
	10, 111, 10, 32, 58, 10, 79, 10, 
	117, 10, 116, 10, 108, 10, 105, 10, 
	110, 105, 118, 101, 110, 99, 101, 110, 
	97, 114, 105, 111, 32, 58, 79, 117, 
	116, 108, 105, 110, 101, 58, 10, 10, 
	10, 32, 35, 37, 64, 65, 70, 71, 
	83, 84, 87, 9, 13, 10, 95, 10, 
	70, 10, 69, 10, 65, 10, 84, 10, 
	85, 10, 82, 10, 69, 10, 95, 10, 
	69, 10, 78, 10, 68, 10, 95, 10, 
	37, 10, 110, 10, 100, 10, 101, 10, 
	97, 10, 116, 10, 117, 10, 114, 10, 
	101, 10, 58, 10, 105, 10, 118, 10, 
	101, 10, 110, 10, 99, 10, 101, 10, 
	110, 10, 97, 10, 114, 10, 105, 10, 
	111, 10, 104, 10, 10, 10, 32, 35, 
	37, 64, 65, 66, 70, 71, 83, 84, 
	87, 9, 13, 10, 95, 10, 70, 10, 
	69, 10, 65, 10, 84, 10, 85, 10, 
	82, 10, 69, 10, 95, 10, 69, 10, 
	78, 10, 68, 10, 95, 10, 37, 10, 
	110, 10, 100, 10, 97, 10, 99, 10, 
	107, 10, 103, 10, 114, 10, 111, 10, 
	117, 10, 110, 10, 100, 10, 58, 10, 
	101, 10, 97, 10, 116, 10, 117, 10, 
	114, 10, 101, 10, 105, 10, 118, 10, 
	101, 10, 110, 10, 99, 10, 101, 10, 
	110, 10, 97, 10, 114, 10, 105, 10, 
	111, 10, 32, 58, 10, 79, 10, 117, 
	10, 116, 10, 108, 10, 105, 10, 110, 
	10, 104, 104, 32, 124, 9, 13, 10, 
	32, 92, 124, 9, 13, 10, 92, 124, 
	10, 92, 10, 32, 92, 124, 9, 13, 
	10, 32, 34, 35, 37, 64, 65, 66, 
	69, 70, 71, 83, 84, 87, 124, 9, 
	13, 10, 101, 10, 97, 10, 116, 10, 
	117, 10, 114, 10, 101, 10, 58, 10, 
	105, 10, 118, 10, 101, 10, 110, 10, 
	99, 10, 101, 10, 110, 10, 97, 10, 
	114, 10, 105, 10, 111, 10, 32, 58, 
	10, 79, 10, 117, 10, 116, 10, 108, 
	10, 105, 10, 110, 10, 104, 187, 191, 
	
}

var _lexer_single_lengths []byte = []byte{
	0, 16, 15, 1, 1, 2, 2, 3, 
	3, 3, 3, 2, 2, 2, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 3, 5, 3, 
	1, 1, 2, 2, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	11, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 15, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 5, 2, 2, 2, 
	2, 2, 2, 2, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 9, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 3, 2, 2, 2, 2, 2, 
	2, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 2, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 11, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 1, 1, 12, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 3, 2, 2, 2, 2, 2, 2, 
	2, 1, 2, 4, 3, 2, 4, 15, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 3, 2, 2, 2, 2, 2, 
	2, 2, 1, 1, 0, 
}

var _lexer_range_lengths []byte = []byte{
	0, 1, 1, 0, 0, 0, 0, 1, 
	1, 1, 1, 1, 1, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 1, 1, 1, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	1, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 1, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 1, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 1, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 1, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 1, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 1, 1, 0, 0, 1, 1, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 
}

var _lexer_index_offsets []int16 = []int16{
	0, 0, 18, 35, 37, 39, 42, 45, 
	50, 55, 60, 65, 69, 73, 76, 78, 
	80, 82, 84, 86, 88, 90, 92, 94, 
	96, 98, 100, 102, 104, 106, 111, 118, 
	123, 125, 127, 130, 133, 135, 137, 139, 
	141, 143, 145, 147, 149, 151, 153, 155, 
	157, 170, 173, 176, 179, 182, 185, 188, 
	191, 194, 197, 200, 203, 206, 209, 212, 
	215, 218, 235, 237, 239, 241, 243, 245, 
	247, 249, 251, 253, 255, 262, 265, 268, 
	271, 274, 277, 280, 283, 285, 287, 289, 
	291, 293, 295, 297, 299, 301, 312, 315, 
	318, 321, 324, 327, 330, 333, 336, 339, 
	342, 345, 348, 351, 354, 357, 360, 363, 
	366, 369, 372, 375, 378, 381, 384, 387, 
	390, 393, 396, 399, 402, 405, 408, 411, 
	414, 417, 420, 423, 426, 429, 432, 435, 
	438, 441, 444, 448, 451, 454, 457, 460, 
	463, 466, 468, 470, 472, 474, 476, 478, 
	480, 482, 484, 486, 488, 491, 493, 495, 
	497, 499, 501, 503, 505, 507, 509, 511, 
	524, 527, 530, 533, 536, 539, 542, 545, 
	548, 551, 554, 557, 560, 563, 566, 569, 
	572, 575, 578, 581, 584, 587, 590, 593, 
	596, 599, 602, 605, 608, 611, 614, 617, 
	620, 623, 626, 629, 631, 633, 647, 650, 
	653, 656, 659, 662, 665, 668, 671, 674, 
	677, 680, 683, 686, 689, 692, 695, 698, 
	701, 704, 707, 710, 713, 716, 719, 722, 
	725, 728, 731, 734, 737, 740, 743, 746, 
	749, 752, 755, 758, 761, 764, 767, 770, 
	773, 776, 780, 783, 786, 789, 792, 795, 
	798, 801, 803, 807, 813, 817, 820, 826, 
	843, 846, 849, 852, 855, 858, 861, 864, 
	867, 870, 873, 876, 879, 882, 885, 888, 
	891, 894, 897, 901, 904, 907, 910, 913, 
	916, 919, 922, 924, 926, 
}

var _lexer_trans_targs []int16 = []int16{
	2, 2, 3, 13, 15, 29, 32, 36, 
	66, 84, 145, 149, 257, 257, 258, 290, 
	2, 0, 2, 2, 3, 13, 15, 29, 
	32, 36, 66, 84, 145, 149, 257, 257, 
	258, 2, 0, 4, 0, 5, 0, 7, 
	6, 6, 7, 6, 6, 8, 8, 9, 
	8, 8, 8, 8, 9, 8, 8, 8, 
	8, 10, 8, 8, 8, 8, 11, 8, 
	8, 2, 12, 12, 0, 2, 12, 12, 
	0, 2, 14, 13, 2, 0, 16, 0, 
	17, 0, 18, 0, 19, 0, 20, 0, 
	21, 0, 22, 0, 23, 0, 24, 0, 
	25, 0, 26, 0, 27, 0, 28, 0, 
	292, 0, 0, 0, 0, 0, 30, 31, 
	2, 31, 31, 29, 30, 30, 2, 31, 
	29, 31, 0, 33, 0, 34, 0, 2, 
	14, 35, 2, 14, 35, 37, 0, 38, 
	0, 39, 0, 40, 0, 41, 0, 42, 
	0, 43, 0, 44, 0, 45, 0, 46, 
	0, 48, 47, 48, 47, 48, 48, 2, 
	49, 2, 63, 264, 271, 275, 289, 289, 
	48, 47, 48, 50, 47, 48, 51, 47, 
	48, 52, 47, 48, 53, 47, 48, 54, 
	47, 48, 55, 47, 48, 56, 47, 48, 
	57, 47, 48, 58, 47, 48, 59, 47, 
	48, 60, 47, 48, 61, 47, 48, 62, 
	47, 48, 2, 47, 48, 64, 47, 48, 
	65, 47, 2, 2, 3, 13, 15, 29, 
	32, 36, 66, 84, 145, 149, 257, 257, 
	258, 2, 0, 67, 0, 68, 0, 69, 
	0, 70, 0, 71, 0, 72, 0, 73, 
	0, 74, 0, 76, 75, 76, 75, 76, 
	76, 2, 77, 2, 76, 75, 76, 78, 
	75, 76, 79, 75, 76, 80, 75, 76, 
	81, 75, 76, 82, 75, 76, 83, 75, 
	76, 65, 75, 85, 0, 86, 0, 87, 
	0, 88, 0, 89, 0, 90, 0, 91, 
	0, 93, 92, 93, 92, 93, 93, 2, 
	94, 2, 108, 118, 125, 131, 93, 92, 
	93, 95, 92, 93, 96, 92, 93, 97, 
	92, 93, 98, 92, 93, 99, 92, 93, 
	100, 92, 93, 101, 92, 93, 102, 92, 
	93, 103, 92, 93, 104, 92, 93, 105, 
	92, 93, 106, 92, 93, 107, 92, 93, 
	2, 92, 93, 109, 92, 93, 110, 92, 
	93, 111, 92, 93, 112, 92, 93, 113, 
	92, 93, 114, 92, 93, 115, 92, 93, 
	116, 92, 93, 117, 92, 93, 65, 92, 
	93, 119, 92, 93, 120, 92, 93, 121, 
	92, 93, 122, 92, 93, 123, 92, 93, 
	124, 92, 93, 117, 92, 93, 126, 92, 
	93, 127, 92, 93, 128, 92, 93, 129, 
	92, 93, 130, 92, 93, 117, 92, 93, 
	132, 92, 93, 133, 92, 93, 134, 92, 
	93, 135, 92, 93, 136, 92, 93, 137, 
	92, 93, 138, 92, 93, 139, 65, 92, 
	93, 140, 92, 93, 141, 92, 93, 142, 
	92, 93, 143, 92, 93, 144, 92, 93, 
	130, 92, 146, 0, 147, 0, 148, 0, 
	34, 0, 150, 0, 151, 0, 152, 0, 
	153, 0, 154, 0, 155, 0, 156, 0, 
	157, 203, 0, 158, 0, 159, 0, 160, 
	0, 161, 0, 162, 0, 163, 0, 164, 
	0, 165, 0, 167, 166, 167, 166, 167, 
	167, 2, 168, 2, 182, 184, 191, 195, 
	202, 202, 167, 166, 167, 169, 166, 167, 
	170, 166, 167, 171, 166, 167, 172, 166, 
	167, 173, 166, 167, 174, 166, 167, 175, 
	166, 167, 176, 166, 167, 177, 166, 167, 
	178, 166, 167, 179, 166, 167, 180, 166, 
	167, 181, 166, 167, 2, 166, 167, 183, 
	166, 167, 65, 166, 167, 185, 166, 167, 
	186, 166, 167, 187, 166, 167, 188, 166, 
	167, 189, 166, 167, 190, 166, 167, 65, 
	166, 167, 192, 166, 167, 193, 166, 167, 
	194, 166, 167, 65, 166, 167, 196, 166, 
	167, 197, 166, 167, 198, 166, 167, 199, 
	166, 167, 200, 166, 167, 201, 166, 167, 
	190, 166, 167, 193, 166, 205, 204, 205, 
	204, 205, 205, 2, 206, 2, 220, 222, 
	232, 238, 242, 256, 256, 205, 204, 205, 
	207, 204, 205, 208, 204, 205, 209, 204, 
	205, 210, 204, 205, 211, 204, 205, 212, 
	204, 205, 213, 204, 205, 214, 204, 205, 
	215, 204, 205, 216, 204, 205, 217, 204, 
	205, 218, 204, 205, 219, 204, 205, 2, 
	204, 205, 221, 204, 205, 65, 204, 205, 
	223, 204, 205, 224, 204, 205, 225, 204, 
	205, 226, 204, 205, 227, 204, 205, 228, 
	204, 205, 229, 204, 205, 230, 204, 205, 
	231, 204, 205, 65, 204, 205, 233, 204, 
	205, 234, 204, 205, 235, 204, 205, 236, 
	204, 205, 237, 204, 205, 231, 204, 205, 
	239, 204, 205, 240, 204, 205, 241, 204, 
	205, 65, 204, 205, 243, 204, 205, 244, 
	204, 205, 245, 204, 205, 246, 204, 205, 
	247, 204, 205, 248, 204, 205, 249, 204, 
	205, 250, 65, 204, 205, 251, 204, 205, 
	252, 204, 205, 253, 204, 205, 254, 204, 
	205, 255, 204, 205, 237, 204, 205, 240, 
	204, 147, 0, 258, 259, 258, 0, 263, 
	262, 261, 259, 262, 260, 0, 261, 259, 
	260, 0, 261, 260, 263, 262, 261, 259, 
	262, 260, 263, 263, 3, 13, 15, 29, 
	32, 36, 66, 84, 145, 149, 257, 257, 
	258, 263, 0, 48, 265, 47, 48, 266, 
	47, 48, 267, 47, 48, 268, 47, 48, 
	269, 47, 48, 270, 47, 48, 65, 47, 
	48, 272, 47, 48, 273, 47, 48, 274, 
	47, 48, 65, 47, 48, 276, 47, 48, 
	277, 47, 48, 278, 47, 48, 279, 47, 
	48, 280, 47, 48, 281, 47, 48, 282, 
	47, 48, 283, 65, 47, 48, 284, 47, 
	48, 285, 47, 48, 286, 47, 48, 287, 
	47, 48, 288, 47, 48, 269, 47, 48, 
	273, 47, 291, 0, 2, 0, 0, 
}

var _lexer_trans_actions []byte = []byte{
	54, 0, 3, 1, 0, 1, 29, 29, 
	29, 29, 29, 29, 29, 29, 35, 0, 
	0, 43, 54, 0, 3, 1, 0, 1, 
	29, 29, 29, 29, 29, 29, 29, 29, 
	35, 0, 43, 0, 43, 0, 43, 139, 
	48, 7, 102, 9, 0, 134, 45, 45, 
	45, 5, 122, 33, 33, 33, 0, 122, 
	33, 33, 33, 0, 122, 33, 0, 33, 
	0, 106, 11, 11, 43, 54, 0, 0, 
	43, 114, 25, 0, 54, 43, 0, 43, 
	0, 43, 0, 43, 0, 43, 0, 43, 
	0, 43, 0, 43, 0, 43, 0, 43, 
	0, 43, 0, 43, 0, 43, 0, 43, 
	0, 43, 43, 43, 43, 43, 0, 27, 
	118, 27, 27, 51, 27, 0, 54, 0, 
	1, 0, 43, 0, 43, 0, 43, 149, 
	126, 57, 110, 23, 0, 0, 43, 0, 
	43, 0, 43, 0, 43, 0, 43, 0, 
	43, 0, 43, 0, 43, 0, 43, 0, 
	43, 144, 57, 54, 0, 54, 0, 72, 
	33, 72, 84, 84, 84, 84, 84, 84, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 15, 0, 54, 0, 0, 54, 
	15, 0, 130, 31, 60, 57, 31, 57, 
	63, 63, 63, 63, 63, 63, 63, 63, 
	66, 31, 43, 0, 43, 0, 43, 0, 
	43, 0, 43, 0, 43, 0, 43, 0, 
	43, 0, 43, 144, 57, 54, 0, 54, 
	0, 81, 84, 81, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 21, 0, 0, 43, 0, 43, 0, 
	43, 0, 43, 0, 43, 0, 43, 0, 
	43, 144, 57, 54, 0, 54, 0, 69, 
	33, 69, 84, 84, 84, 84, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	13, 0, 54, 0, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 13, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 13, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 0, 43, 0, 43, 0, 43, 
	0, 43, 0, 43, 0, 43, 0, 43, 
	0, 43, 0, 43, 0, 43, 0, 43, 
	0, 0, 43, 0, 43, 0, 43, 0, 
	43, 0, 43, 0, 43, 0, 43, 0, 
	43, 0, 43, 144, 57, 54, 0, 54, 
	0, 78, 33, 78, 84, 84, 84, 84, 
	84, 84, 0, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 0, 0, 54, 19, 0, 54, 0, 
	0, 54, 19, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 19, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 54, 19, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 144, 57, 54, 
	0, 54, 0, 75, 33, 75, 84, 84, 
	84, 84, 84, 84, 84, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 17, 
	0, 54, 0, 0, 54, 17, 0, 54, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 54, 17, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 17, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 0, 17, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 0, 43, 0, 0, 0, 43, 54, 
	37, 37, 87, 37, 37, 43, 0, 39, 
	0, 43, 0, 0, 54, 0, 0, 39, 
	0, 0, 54, 0, 93, 90, 41, 90, 
	96, 96, 96, 96, 96, 96, 96, 96, 
	99, 0, 43, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 15, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 15, 0, 54, 0, 0, 54, 
	0, 0, 54, 0, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 15, 0, 54, 0, 0, 
	54, 0, 0, 54, 0, 0, 54, 0, 
	0, 54, 0, 0, 54, 0, 0, 54, 
	0, 0, 0, 43, 0, 43, 0, 
}

var _lexer_eof_actions []byte = []byte{
	0, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 43, 43, 43, 
	43, 43, 43, 43, 43, 
}

const lexer_start int = 1
const lexer_first_final int = 292

const lexer_en_main int = 1


//line ragel/lexer_go.rl:259
  // END: write data noerror;

func currentLineContent(data []byte, lastNewline int) (string) {
    current := string(data[lastNewline:])
    return strings.TrimSpace( current )
}

func unindent(startCol int, text []byte) (string) {
    regex, err := regexp.Compile("(?m)^[\t ]{0," + strconv.Itoa(startCol) + "}")

    if ( err != nil ) {
      panic(err)
    }
    result := regex.ReplaceAll( text, text[:0] )
    return string(result)
}

func keywordContent( data []byte, p int, eof int, nextKeywordStart int, contentStart int ) ([]byte) {
    endPoint := nextKeywordStart
    if ( (nextKeywordStart == -1 || (p == eof)) ) {
        endPoint = p
    }
    con := data[contentStart:endPoint]
    return con
}


func nameAndUnindentedDescription(startCol int, textBytes []byte) (string, []string) {
    text := unindent( startCol, textBytes )
    text = strings.TrimSpace( text )
    lines := strings.Split(text, "\n")

      for index,element := range lines {
        lines[index] = strings.TrimSpace(element)
      }

    if ( len(lines) == 0 ) {
      return "", []string{}
    } else if ( len(lines) == 1 ) {
      return lines[0], []string{}
      } else {
        return lines[0], lines[1:]
      }
}

func makeTable(rows [][]string) (*TableData) {
    return &TableData{}
}

func ParseFeature(data []byte, filename string) (feature Feature, err error) {

  // Original ragel parser assumes this will be there, who am I to argue?
  data = append(data, []byte("%_FEATURE_END_%")...)

  cs := 0 // No idea what this is
  p := 0 // Position?
  pe := len(data) // No idea
  eof := len(data) // Location of EOF

  lineNumber := 1
  lastNewline := 0

  contentStart := -1
  currentLine := -1

  docstringContentTypeStart := -1
  docstringContentTypeEnd := -1
  startCol := -1;
  nextKeywordStart := -1
  keywordStart := -1

  var keyword string
  var currentTable [][]string
  var currentRow []string

  var tags []string
//  var lastScenario *Scenario = nil;
//  var lastStep *Step = nil;

  var examplesTable bool = false;

  // START: write init
    
//line gocumber/parser.go:660
	{
	cs = lexer_start
	}

//line ragel/lexer_go.rl:342
  // END: write init

  // START: write exec
    
//line gocumber/parser.go:670
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_keys = int(_lexer_key_offsets[cs])
	_trans = int(_lexer_index_offsets[cs])

	_klen = int(_lexer_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _lexer_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _lexer_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_lexer_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _lexer_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _lexer_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	cs = int(_lexer_trans_targs[_trans])

	if _lexer_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_lexer_trans_actions[_trans])
	_nacts = uint(_lexer_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _lexer_actions[_acts-1] {
		case 0:
//line ragel/lexer_go.rl:17

      contentStart = p
      currentLine = lineNumber

      if(len(keyword) != 0) {
        startCol = p - lastNewline - (len(keyword) + 1);
      }
    
		case 1:
//line ragel/lexer_go.rl:26

      currentLine = lineNumber
      startCol = p - lastNewline
    
		case 2:
//line ragel/lexer_go.rl:31

      contentStart = p
    
		case 3:
//line ragel/lexer_go.rl:35

      docstringContentTypeStart = p
    
		case 4:
//line ragel/lexer_go.rl:39

      docstringContentTypeEnd = p
    
		case 5:
//line ragel/lexer_go.rl:43

      rawcon := data[contentStart:nextKeywordStart-1]

      con := unindent( startCol, rawcon )
      con = strings.Replace( con, "\\\"\\\"\\\"", "\"\"\"", -1 )
      con = strings.TrimLeft( con, " \t\r\n" )

      conType := string(data[docstringContentTypeStart:docstringContentTypeEnd])
      conType = strings.TrimSpace( conType )

      fmt.Printf("DocString Type:[%s]\nDocString Cont:[%s]\n", conType, con)
      // listener.docString(conType, con, currentLine);
    
		case 6:
//line ragel/lexer_go.rl:57

        kcon := keywordContent(data, p, eof, nextKeywordStart, contentStart)
        name, description := nameAndUnindentedDescription( startCol, kcon );

        feature = Feature{
            Name: name,
            StartsAt: DocumentLocation{ Filename: filename, Line: currentLine },
            ConditionsOfSatisfaction: description,
            Tags: tags,
            Language: "en",
            Background: nil,
            Scenario: []Scenario{},
        }

        tags = tags[:0]

        if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
        nextKeywordStart = -1
    
		case 7:
//line ragel/lexer_go.rl:77

        kcon := keywordContent(data, p, eof, nextKeywordStart, contentStart)
        name, _ := nameAndUnindentedDescription( startCol, kcon );

        background := Scenario{
            Feature:    &feature,
            Name:       name,
            StartsAt:   DocumentLocation{ Filename: filename, Line: currentLine },
            Tags:       tags,
            Background: true,
            Steps:      []Step{},
            TableData:  nil,
        }

        //lastScenario = &background
        tags = tags[:0]

        if ( examplesTable ) {
            background.TableData = makeTable(currentTable)
            currentTable = currentTable[:0]
            examplesTable = false
        }

        feature.Background = &background

        if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
        nextKeywordStart = -1
    
		case 8:
//line ragel/lexer_go.rl:106

        kcon := keywordContent(data, p, eof, nextKeywordStart, contentStart)
        name, _ := nameAndUnindentedDescription( startCol, kcon );

        scenario := Scenario{
            Feature:    &feature,
            Name:       name,
            StartsAt:   DocumentLocation{ Filename: filename, Line: currentLine },
            Tags:       tags,
            Background: false,
            Steps:      []Step{},
            TableData:  nil,
        }

        //lastScenario = &scenario
        tags = tags[:0]

        if ( examplesTable ) {
            scenario.TableData = makeTable(currentTable)
            currentTable = currentTable[:0]
            examplesTable = false
        }

        feature.Scenario = append( feature.Scenario, scenario )

        if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
        nextKeywordStart = -1
    
		case 9:
//line ragel/lexer_go.rl:135

        kcon := keywordContent(data, p, eof, nextKeywordStart, contentStart)
        name, _ := nameAndUnindentedDescription( startCol, kcon );

        scenario := Scenario{
            Feature:    &feature,
            Name:       name,
            StartsAt:   DocumentLocation{ Filename: filename, Line: currentLine },
            Tags:       tags,
            Background: false,
            Steps:      []Step{},
            TableData:  nil,
        }

        //lastScenario = &scenario
        tags = tags[:0]

        if ( examplesTable ) {
            scenario.TableData = makeTable(currentTable)
            currentTable = currentTable[:0]
            examplesTable = false
        }

        feature.Scenario = append( feature.Scenario, scenario )

        if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
        nextKeywordStart = -1
    
		case 10:
//line ragel/lexer_go.rl:164

        //kcon := keywordContent(data, p, eof, nextKeywordStart, contentStart)
        //name, description := nameAndUnindentedDescription( startCol, kcon );

        examplesTable = true;

        if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
        nextKeywordStart = -1
    
		case 11:
//line ragel/lexer_go.rl:174

        con := string( data[contentStart:p] )
        con = strings.TrimSpace( con )
        fmt.Printf("Store Thing: [Step]\nStore Keyword: [%s]\nStore Content: [%s]\n", keyword, con )
      // listener.step(keyword, substring(data, contentStart, p).trim(), currentLine);

    
		case 12:
//line ragel/lexer_go.rl:182

        con := string( data[contentStart:p] )
        con = strings.TrimSpace( con )
        // Discard comments
        // fmt.Printf("Store Thing: [Comment]\nStore Keyword: [%s]\nStore Content: [%s]\n", keyword, con )
        // listener.comment(substring(data, contentStart, p).trim(), lineNumber);
        keywordStart = -1
    
		case 13:
//line ragel/lexer_go.rl:191

        con := string( data[contentStart:p] )
        con = strings.TrimSpace( con )
        con = strings.TrimLeft( con, "@" ) // Don't need the @

        tags = append( tags, con )

        keywordStart = -1
    
		case 14:
//line ragel/lexer_go.rl:201

      lineNumber++;
    
		case 15:
//line ragel/lexer_go.rl:205

      lastNewline = p + 1
    
		case 16:
//line ragel/lexer_go.rl:209

      if (keywordStart == -1) { keywordStart = p }
    
		case 17:
//line ragel/lexer_go.rl:213

      rawKeyword := string(data[keywordStart:p])
      rawKeyword = strings.Replace(rawKeyword, ":", "", 1)
      keyword = rawKeyword[0:]
      keywordStart = -1
    
		case 18:
//line ragel/lexer_go.rl:220

      nextKeywordStart = p
    
		case 19:
//line ragel/lexer_go.rl:224

      p = p - 1
      currentRow = currentRow[:0]
      currentLine = lineNumber
    
		case 20:
//line ragel/lexer_go.rl:230

      contentStart = p;
    
		case 21:
//line ragel/lexer_go.rl:234

        con := string(data[contentStart:p])
        con = strings.TrimSpace(con)
        con = strings.Replace(con, "\\|", "|", -1)
        con = strings.Replace(con, "\\n", "\n", -1)
        con = strings.Replace(con, "\\\\", "\\", -1)
        currentRow = append( currentRow, con )
    
		case 22:
//line ragel/lexer_go.rl:243

        currentTable = append( currentTable, currentRow )
    
		case 23:
//line ragel/lexer_go.rl:247

      if(cs < lexer_first_final) {
          content := currentLineContent( data, lastNewline )
          panic(fmt.Sprintf("Lexing error on line %d: '%s'. See http://wiki.github.com/cucumber/gherkin/lexingerror for more information.", lineNumber, content))
      }
    
//line gocumber/parser.go:1009
		}
	}

_again:
	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		__acts := _lexer_eof_actions[cs]
		__nacts := uint(_lexer_actions[__acts]); __acts++
		for ; __nacts > 0; __nacts-- {
			__acts++
			switch _lexer_actions[__acts-1] {
			case 23:
//line ragel/lexer_go.rl:247

      if(cs < lexer_first_final) {
          content := currentLineContent( data, lastNewline )
          panic(fmt.Sprintf("Lexing error on line %d: '%s'. See http://wiki.github.com/cucumber/gherkin/lexingerror for more information.", lineNumber, content))
      }
    
//line gocumber/parser.go:1036
			}
		}
	}

	_out: {}
	}

//line ragel/lexer_go.rl:346
  // END: write init


    return
}
