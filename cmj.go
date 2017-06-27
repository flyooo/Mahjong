type stPAI struct{
	mType int
	mValue int
}

type stCHI{
	mType int
	mValue1 int
	mValue2 int
	mValue3 int
}

type stGoodInfo struct{
	mGoodName string
	mGoodValue int
}

type stPAIEx struct{
	mNewPai stPAI
	mPaiNum int
	mIsHz bool
}

type Integer int

type Lesser interface{
	Less (b Integer) bool
}