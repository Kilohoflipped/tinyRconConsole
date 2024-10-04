package math_utils

import "math"

const NewtonMinSlope = 0.001
const kSplineTableSize = 11
const kSampleStepSize float32 = 1.0 / (kSplineTableSize - 1.0)

const NewtonIterations = 4

const SubdivisionPrecision = 0.0000001
const SubdivisionMaxIterations = 10

func CubicBezier(x, X1, Y1, X2, Y2 float32) float32 {
	return BezierEasing(x, X1, Y1, X2, Y2)
}

func BezierEasing(x, X1, Y1, X2, Y2 float32) float32 {
	if (x == 0) || (x == 1) {
		return x
	}
	sampleValues := make([]float32, kSplineTableSize)
	for i := 0; i < kSplineTableSize; i = i + 1 {
		// i从0到10，sampleValues长度为11
		sampleValues[i] = CalcBezier(float32(i)*kSampleStepSize, X1, X2)
	}
	return CalcBezier(getTFromX(x, X1, X2, sampleValues), Y1, Y2)
}

func BezierPartA(X1, X2 float32) float32 {
	return 1.0 - 3.0*X1 + 3.0*X2
}

func BezierPartB(X1, X2 float32) float32 {
	return 3.0*X1 - 6.0*X2
}
func BezierPartC(X1 float32) float32 {
	return 3.0 * X1
}
func CalcBezier(t, X1, X2 float32) float32 {
	return ((BezierPartA(X1, X2)*t+BezierPartB(X1, X2))*t + BezierPartC(X1)) * t
}

// Returns dx/dt given t, x1, and x2, or dy/dt given t, y1, and y2.
// 求曲线方程的一阶导函数
func getSlope(aT, aA1, aA2 float32) float32 {
	return 3.0*BezierPartA(aA1, aA2)*aT*aT + 2.0*BezierPartB(aA1, aA2)*aT + BezierPartC(aA1)
}
func getTFromX(x, X1, X2 float32, sampleValues []float32) float32 {
	var intervalStart float32 = 0.0
	var currentSample = 1
	// lastSample为10
	var lastSample = kSplineTableSize - 1
	// sampleValues[i]表示i从0以0.1为step，每一步对应的曲线的X坐标值，直到X坐标值小于等于aX
	// 假如aX=0.4，则sampleValues[currentSample]<=aX为止

	for currentSample != lastSample && sampleValues[currentSample] <= x {
		// intervalStart为到aX经过的step步骤
		intervalStart = intervalStart + kSampleStepSize // kSampleStepSize为0.1
		currentSample = currentSample + 1
	}

	//currentSample为什么要减1？sampleValues[currentSample]大于了ax，所以要--，使得sampleValues[currentSample]<=ax
	currentSample = currentSample - 1

	// Interpolate to provide an initial guess for t
	// ax-sampleValues[currentSample]为两者之间的差值，而(sampleValues[currentSample + 1] - sampleValues[currentSample])一个步骤之间的总差值。
	var dist = (x - sampleValues[currentSample]) / (sampleValues[currentSample+1] - sampleValues[currentSample])
	// guessForT为预计的初始T值，很粗糙的一个值，接下来会基于该值求根(t值)。
	var guessForT = intervalStart + dist*kSampleStepSize
	// 预测的T值对应位置的斜率
	var initialSlope = getSlope(guessForT, X1, X2)

	// 当斜率大于0.05729°时，使用newtonRaphsonIterate算法预测T值。0.05729是一个很小的斜率
	if initialSlope >= NewtonMinSlope {
		return newtonRaphsonIterate(x, guessForT, X1, X2)
	} else if initialSlope == 0.0 { // 当斜率为0，则直接返回
		return guessForT
	} else { // 当斜率小于0.05729并且不等于0时，使用binarySubdivide
		// 求得的根t，位于intervalStart和intervalStart + kSampleStepSize之间, mX1、mX2分别对应p1、p2的X坐标
		return binarySubdivide(x, intervalStart, intervalStart+kSampleStepSize, X1, X2)
	}
}

func newtonRaphsonIterate(x, aGuessT, X1, X2 float32) float32 {
	// NEWTON_ITERATIONS为4， 只进行了4次迭代， 根据精度和性能之间做了平衡。
	for i := 0; i < NewtonIterations; i = i + 1 {
		// 计算t值对应位置的斜率
		var currentSlope = getSlope(aGuessT, X1, X2)
		if currentSlope == 0.0 {
			return aGuessT
		}
		// 假设f(t) = 0，求解方程的根。其f(t)=calcBezier(t) - ax
		// 牛顿-拉佛森方法: Xn-1 = Xn - f(t) / f'(t)，应用到求贝塞尔曲线的根：Tn = Tn+1 - (calcBezier(t) - ax) / getSlope(t)
		var currentX = CalcBezier(aGuessT, X1, X2) - x
		aGuessT -= currentX / currentSlope
	}
	// 这里只迭代了4次，求得近似值
	return aGuessT
}

// 二分球根法：
// 求得的根t，位于aA和aB之间, X1、X2分别对应p1、p2的X坐标
func binarySubdivide(x, aA, aB, X1, X2 float32) float32 {
	var currentX, currentT, i float32 = 0, 0, 0
	for {
		currentT = aA + (aB-aA)/2.0
		// 假设f(t) = 0，求解方程的根。其f(t)=calcBezier(t) - ax
		currentX = CalcBezier(currentT, X1, X2) - x
		if currentX > 0.0 {
			aB = currentT
		} else {
			aA = currentT
		}
		// 如果currentX小于等于最小精度(SubdivisionPrecision)SubdivisionMaxIterations，则终止
		i = i + 1
		if math.Abs(float64(currentX)) > SubdivisionPrecision && i < SubdivisionMaxIterations {
			break
		}
	}
	return currentT
}
