package evalexpr

import (
	"strings"
)

type clause interface {
}

type andOperator struct {
	left  clause
	right clause
}

type orOperator struct {
	left  clause
	right clause
}

func parse(requirements string) clause {
	cleanedString := strings.Replace(requirements, " ", "", -1)

	return extractCondition(cleanedString)
}

func extractCondition(expression string) clause {
	substrings := extractComponents(expression)
	if len(substrings) < 3 {
		return expression
	}

	// Reversing the substring as the conditions construction is from right to left
	for i, j := 0, len(substrings)-1; i < j; i, j = i+1, j-1 {
		substrings[i], substrings[j] = substrings[j], substrings[i]
	}

	var condition = extractCondition(substrings[0])
	for index, substring := range substrings {
		if substring == "&" {
			condition = andOperator{
				left:  extractCondition(substrings[index+1]),
				right: condition,
			}
		}
		if substring == "|" {
			condition = orOperator{
				left:  extractCondition(substrings[index+1]),
				right: condition,
			}
		}
	}
	return condition
}

func extractComponents(requirements string) []string {
	var substrings []string

	componentStartIndex := 0
	parenthesisCount := 0

	for i := 0; i < len(requirements); i++ {
		if (requirements[i] == '&' || requirements[i] == '|') && parenthesisCount == 0 {
			substrings = append(substrings, removeParenthesisIfNeeded(requirements[componentStartIndex:i]))
			substrings = append(substrings, string(requirements[i]))
			componentStartIndex = i + 1
		}
		if requirements[i] == '(' {
			parenthesisCount += 1
		}
		if requirements[i] == ')' {
			parenthesisCount -= 1
		}
	}
	substrings = append(substrings, removeParenthesisIfNeeded(requirements[componentStartIndex:]))
	return substrings
}

func removeParenthesisIfNeeded(expression string) string {
	if len(expression) < 2 {
		return expression
	}
	if expression[0] == '(' && expression[len(expression)-1] == ')' {
		return expression[1 : len(expression)-1]
	}
	return expression
}

func evaluate(condition clause, array []string) bool {
	switch condition.(type) {
	case string:
		for _, value := range array {
			if value == condition {
				return true
			}
		}
	case andOperator:
		condition := condition.(andOperator)
		return evaluate(condition.left, array) && evaluate(condition.right, array)
	case orOperator:
		condition := condition.(orOperator)
		return evaluate(condition.left, array) || evaluate(condition.right, array)
	}
	return false
}

func IsFulfillingCondition(expression string, array []string) bool {
	return evaluate(parse(expression), array)
}
