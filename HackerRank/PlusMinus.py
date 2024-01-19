def plusMinus(arr):
    # Write your code here
    nPositive = 0
    nNegative = 0
    nZero = 0
    
    for i in arr:
        if i > 0:
            nPositive += 1
        elif i == 0:
            nZero += 1
        else:
            nNegative += 1
            
    positiveRatios = (nPositive / len(arr))
    zeroRatios = (nZero / len(arr))
    negativeRatios = (nNegative / len(arr))
    
    print(positiveRatios)
    print(negativeRatios)
    print(zeroRatios)