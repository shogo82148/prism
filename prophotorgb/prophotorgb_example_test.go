package prophotorgb_test

import (
	"fmt"
	"github.com/mandykoh/prism/prophotorgb"
)

func ExampleConvert8BitToLinear_generateLUT() {
	lut := [256]float32{}
	for i := range lut {
		lut[i] = float32(prophotorgb.Convert8BitToLinear(uint8(i)))
	}

	fmt.Printf("[256]float32{")
	for i, v := range lut {
		if i%8 == 0 {
			fmt.Printf("\n\t")
		}
		fmt.Printf(" %v,", v)
	}
	fmt.Printf("\n}\n")

	// Output: [256]float32{
	//	 0, 0.00024509805, 0.0004901961, 0.00073529413, 0.0009803922, 0.0012254902, 0.0014705883, 0.0017156863,
	//	 0.0019669335, 0.0024314437, 0.00293919, 0.0034892694, 0.004080881, 0.0047133067, 0.005385898, 0.0060980637,
	//	 0.00684926, 0.007638986, 0.008466778, 0.009332202, 0.010234854, 0.011174354, 0.012150342, 0.013162482,
	//	 0.014210452, 0.015293951, 0.016412688, 0.017566387, 0.018754786, 0.019977635, 0.021234691, 0.022525722,
	//	 0.023850508, 0.025208835, 0.026600495, 0.028025292, 0.029483033, 0.030973535, 0.032496616, 0.034052104,
	//	 0.035639834, 0.03725964, 0.038911358, 0.040594846, 0.042309947, 0.04405652, 0.045834422, 0.047643516,
	//	 0.04948367, 0.05135475, 0.05325663, 0.055189185, 0.057152297, 0.059145845, 0.061169714, 0.06322379,
	//	 0.06530796, 0.06742212, 0.06956617, 0.07173999, 0.07394349, 0.07617656, 0.078439124, 0.080731064,
	//	 0.08305229, 0.08540272, 0.087782264, 0.09019082, 0.09262831, 0.09509464, 0.09758974, 0.10011351,
	//	 0.10266589, 0.10524678, 0.10785612, 0.11049381, 0.11315979, 0.11585399, 0.11857632, 0.121326715,
	//	 0.1241051, 0.12691142, 0.12974559, 0.13260755, 0.13549723, 0.13841455, 0.14135946, 0.1443319,
	//	 0.1473318, 0.1503591, 0.15341371, 0.15649562, 0.15960473, 0.162741, 0.16590436, 0.16909477,
	//	 0.17231214, 0.17555645, 0.17882763, 0.18212561, 0.18545036, 0.18880181, 0.1921799, 0.19558461,
	//	 0.19901586, 0.20247361, 0.2059578, 0.20946838, 0.21300532, 0.21656854, 0.22015803, 0.2237737,
	//	 0.22741553, 0.23108347, 0.23477747, 0.23849748, 0.24224345, 0.24601535, 0.24981314, 0.25363678,
	//	 0.25748616, 0.26136133, 0.26526222, 0.26918873, 0.27314088, 0.27711862, 0.28112188, 0.28515065,
	//	 0.2892049, 0.29328454, 0.29738957, 0.30151993, 0.3056756, 0.3098565, 0.31406265, 0.31829402,
	//	 0.3225505, 0.3268321, 0.33113876, 0.3354705, 0.3398272, 0.3442089, 0.3486155, 0.353047,
	//	 0.35750338, 0.36198458, 0.36649057, 0.37102133, 0.3755768, 0.38015696, 0.38476178, 0.38939124,
	//	 0.3940453, 0.39872387, 0.403427, 0.40815464, 0.4129067, 0.41768324, 0.42248416, 0.42730945,
	//	 0.4321591, 0.43703303, 0.44193125, 0.4468537, 0.45180038, 0.45677125, 0.4617663, 0.46678546,
	//	 0.47182873, 0.47689608, 0.48198745, 0.48710287, 0.49224225, 0.4974056, 0.50259286, 0.50780404,
	//	 0.5130391, 0.51829803, 0.5235808, 0.5288873, 0.5342176, 0.53957164, 0.5449494, 0.55035084,
	//	 0.55577594, 0.5612247, 0.56669706, 0.572193, 0.57771254, 0.5832556, 0.5888222, 0.5944122,
	//	 0.6000258, 0.6056627, 0.61132306, 0.61700684, 0.622714, 0.62844443, 0.6341982, 0.6399753,
	//	 0.6457757, 0.6515992, 0.6574461, 0.66331613, 0.6692093, 0.67512566, 0.68106514, 0.68702775,
	//	 0.6930135, 0.69902223, 0.70505404, 0.71110886, 0.7171867, 0.7232875, 0.72941124, 0.735558,
	//	 0.7417276, 0.7479201, 0.7541355, 0.7603737, 0.76663476, 0.77291864, 0.7792253, 0.7855547,
	//	 0.7919069, 0.7982818, 0.8046794, 0.81109965, 0.8175426, 0.8240082, 0.83049643, 0.8370073,
	//	 0.8435407, 0.8500967, 0.85667527, 0.8632763, 0.8698999, 0.87654597, 0.88321453, 0.8899055,
	//	 0.89661896, 0.90335476, 0.91011304, 0.91689366, 0.92369664, 0.93052197, 0.9373696, 0.94423956,
	//	 0.9511318, 0.9580463, 0.96498305, 0.97194207, 0.97892326, 0.9859267, 0.9929522, 1,
	//}
}

func ExampleConvertLinearTo8Bit_generateLUT() {
	lut := [512]uint8{}
	for i := range lut {
		lut[i] = prophotorgb.ConvertLinearTo8Bit(float32(i) / 511)
	}

	fmt.Printf("[512]uint8{")
	for i, v := range lut {
		if i%16 == 0 {
			fmt.Printf("\n\t")
		}
		fmt.Printf(" %d,", v)
	}
	fmt.Printf("\n}\n")

	// Output: [512]uint8{
	//	 0, 8, 12, 15, 17, 20, 22, 24, 25, 27, 29, 30, 32, 33, 35, 36,
	//	 37, 38, 40, 41, 42, 43, 44, 46, 47, 48, 49, 50, 51, 52, 53, 54,
	//	 55, 56, 57, 58, 58, 59, 60, 61, 62, 63, 64, 64, 65, 66, 67, 68,
	//	 69, 69, 70, 71, 72, 72, 73, 74, 75, 75, 76, 77, 78, 78, 79, 80,
	//	 80, 81, 82, 82, 83, 84, 85, 85, 86, 87, 87, 88, 88, 89, 90, 90,
	//	 91, 92, 92, 93, 94, 94, 95, 95, 96, 97, 97, 98, 98, 99, 100, 100,
	//	 101, 101, 102, 102, 103, 104, 104, 105, 105, 106, 106, 107, 108, 108, 109, 109,
	//	 110, 110, 111, 111, 112, 112, 113, 113, 114, 115, 115, 116, 116, 117, 117, 118,
	//	 118, 119, 119, 120, 120, 121, 121, 122, 122, 123, 123, 124, 124, 125, 125, 126,
	//	 126, 127, 127, 128, 128, 129, 129, 130, 130, 130, 131, 131, 132, 132, 133, 133,
	//	 134, 134, 135, 135, 136, 136, 137, 137, 137, 138, 138, 139, 139, 140, 140, 141,
	//	 141, 141, 142, 142, 143, 143, 144, 144, 145, 145, 145, 146, 146, 147, 147, 148,
	//	 148, 148, 149, 149, 150, 150, 151, 151, 151, 152, 152, 153, 153, 154, 154, 154,
	//	 155, 155, 156, 156, 156, 157, 157, 158, 158, 158, 159, 159, 160, 160, 160, 161,
	//	 161, 162, 162, 162, 163, 163, 164, 164, 164, 165, 165, 166, 166, 166, 167, 167,
	//	 168, 168, 168, 169, 169, 170, 170, 170, 171, 171, 171, 172, 172, 173, 173, 173,
	//	 174, 174, 174, 175, 175, 176, 176, 176, 177, 177, 177, 178, 178, 179, 179, 179,
	//	 180, 180, 180, 181, 181, 181, 182, 182, 183, 183, 183, 184, 184, 184, 185, 185,
	//	 185, 186, 186, 187, 187, 187, 188, 188, 188, 189, 189, 189, 190, 190, 190, 191,
	//	 191, 191, 192, 192, 192, 193, 193, 194, 194, 194, 195, 195, 195, 196, 196, 196,
	//	 197, 197, 197, 198, 198, 198, 199, 199, 199, 200, 200, 200, 201, 201, 201, 202,
	//	 202, 202, 203, 203, 203, 204, 204, 204, 205, 205, 205, 206, 206, 206, 207, 207,
	//	 207, 208, 208, 208, 209, 209, 209, 210, 210, 210, 211, 211, 211, 212, 212, 212,
	//	 212, 213, 213, 213, 214, 214, 214, 215, 215, 215, 216, 216, 216, 217, 217, 217,
	//	 218, 218, 218, 219, 219, 219, 219, 220, 220, 220, 221, 221, 221, 222, 222, 222,
	//	 223, 223, 223, 223, 224, 224, 224, 225, 225, 225, 226, 226, 226, 227, 227, 227,
	//	 227, 228, 228, 228, 229, 229, 229, 230, 230, 230, 230, 231, 231, 231, 232, 232,
	//	 232, 233, 233, 233, 233, 234, 234, 234, 235, 235, 235, 236, 236, 236, 236, 237,
	//	 237, 237, 238, 238, 238, 238, 239, 239, 239, 240, 240, 240, 241, 241, 241, 241,
	//	 242, 242, 242, 243, 243, 243, 243, 244, 244, 244, 245, 245, 245, 245, 246, 246,
	//	 246, 247, 247, 247, 247, 248, 248, 248, 249, 249, 249, 249, 250, 250, 250, 251,
	//	 251, 251, 251, 252, 252, 252, 252, 253, 253, 253, 254, 254, 254, 254, 255, 255,
	//}
}
