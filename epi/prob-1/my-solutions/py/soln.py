def isPalindrome(s):
	from collections import Counter
	cntr = Counter(s)
	numOddCount = 0
	for v in cntr.values():
		if v % 2 == 1:
			numOddCount += 1
		if numOddCount > 1:
			return 0
	if numOddCount > 1:
		return 0
	else:
		return 1

def main():
	x = raw_input()
	print(isPalindrome(x))

main()