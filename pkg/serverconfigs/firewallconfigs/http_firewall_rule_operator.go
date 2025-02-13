package firewallconfigs

type HTTPFirewallRuleOperator = string
type HTTPFirewallRuleCaseInsensitive = string

const (
	HTTPFirewallRuleOperatorGt           HTTPFirewallRuleOperator = "gt"
	HTTPFirewallRuleOperatorGte          HTTPFirewallRuleOperator = "gte"
	HTTPFirewallRuleOperatorLt           HTTPFirewallRuleOperator = "lt"
	HTTPFirewallRuleOperatorLte          HTTPFirewallRuleOperator = "lte"
	HTTPFirewallRuleOperatorEq           HTTPFirewallRuleOperator = "eq"
	HTTPFirewallRuleOperatorNeq          HTTPFirewallRuleOperator = "neq"
	HTTPFirewallRuleOperatorEqString     HTTPFirewallRuleOperator = "eq string"
	HTTPFirewallRuleOperatorNeqString    HTTPFirewallRuleOperator = "neq string"
	HTTPFirewallRuleOperatorMatch        HTTPFirewallRuleOperator = "match"
	HTTPFirewallRuleOperatorNotMatch     HTTPFirewallRuleOperator = "not match"
	HTTPFirewallRuleOperatorContains     HTTPFirewallRuleOperator = "contains"
	HTTPFirewallRuleOperatorNotContains  HTTPFirewallRuleOperator = "not contains"
	HTTPFirewallRuleOperatorPrefix       HTTPFirewallRuleOperator = "prefix"
	HTTPFirewallRuleOperatorSuffix       HTTPFirewallRuleOperator = "suffix"
	HTTPFirewallRuleOperatorHasKey       HTTPFirewallRuleOperator = "has key" // has key in slice or map
	HTTPFirewallRuleOperatorVersionGt    HTTPFirewallRuleOperator = "version gt"
	HTTPFirewallRuleOperatorVersionLt    HTTPFirewallRuleOperator = "version lt"
	HTTPFirewallRuleOperatorVersionRange HTTPFirewallRuleOperator = "version range"

	HTTPFirewallRuleOperatorContainsBinary    HTTPFirewallRuleOperator = "contains binary"     // contains binary
	HTTPFirewallRuleOperatorNotContainsBinary HTTPFirewallRuleOperator = "not contains binary" // not contains binary

	// ip
	HTTPFirewallRuleOperatorEqIP       HTTPFirewallRuleOperator = "eq ip"
	HTTPFirewallRuleOperatorGtIP       HTTPFirewallRuleOperator = "gt ip"
	HTTPFirewallRuleOperatorGteIP      HTTPFirewallRuleOperator = "gte ip"
	HTTPFirewallRuleOperatorLtIP       HTTPFirewallRuleOperator = "lt ip"
	HTTPFirewallRuleOperatorLteIP      HTTPFirewallRuleOperator = "lte ip"
	HTTPFirewallRuleOperatorIPRange    HTTPFirewallRuleOperator = "ip range"
	HTTPFirewallRuleOperatorNotIPRange HTTPFirewallRuleOperator = "not ip range"
	HTTPFirewallRuleOperatorIPMod10    HTTPFirewallRuleOperator = "ip mod 10"
	HTTPFirewallRuleOperatorIPMod100   HTTPFirewallRuleOperator = "ip mod 100"
	HTTPFirewallRuleOperatorIPMod      HTTPFirewallRuleOperator = "ip mod"

	HTTPFirewallRuleCaseInsensitiveNone = "none"
	HTTPFirewallRuleCaseInsensitiveYes  = "yes"
	HTTPFirewallRuleCaseInsensitiveNo   = "no"
)

type RuleOperatorDefinition struct {
	Name            string
	Code            string
	Description     string
	CaseInsensitive HTTPFirewallRuleCaseInsensitive // default caseInsensitive setting
}

var AllRuleOperators = []*RuleOperatorDefinition{
	{
		Name:            "数值大于",
		Code:            HTTPFirewallRuleOperatorGt,
		Description:     "使用数值对比大于",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNone,
	},
	{
		Name:            "数值大于等于",
		Code:            HTTPFirewallRuleOperatorGte,
		Description:     "使用数值对比大于等于",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNone,
	},
	{
		Name:            "数值小于",
		Code:            HTTPFirewallRuleOperatorLt,
		Description:     "使用数值对比小于",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNone,
	},
	{
		Name:            "数值小于等于",
		Code:            HTTPFirewallRuleOperatorLte,
		Description:     "使用数值对比小于等于",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNone,
	},
	{
		Name:            "数值等于",
		Code:            HTTPFirewallRuleOperatorEq,
		Description:     "使用数值对比等于",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNone,
	},
	{
		Name:            "数值不等于",
		Code:            HTTPFirewallRuleOperatorNeq,
		Description:     "使用数值对比不等于",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNone,
	},
	{
		Name:            "字符串等于",
		Code:            HTTPFirewallRuleOperatorEqString,
		Description:     "使用字符串对比等于",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "字符串不等于",
		Code:            HTTPFirewallRuleOperatorNeqString,
		Description:     "使用字符串对比不等于",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "正则匹配",
		Code:            HTTPFirewallRuleOperatorMatch,
		Description:     "使用正则表达式匹配，在头部使用(?i)表示不区分大小写，<a href=\"http://teaos.cn/doc/regexp/Regexp.md\" target=\"_blank\">正则表达式语法 &raquo;</a>",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveYes,
	},
	{
		Name:            "正则不匹配",
		Code:            HTTPFirewallRuleOperatorNotMatch,
		Description:     "使用正则表达式不匹配，在头部使用(?i)表示不区分大小写，<a href=\"http://teaos.cn/doc/regexp/Regexp.md\" target=\"_blank\">正则表达式语法 &raquo;</a>",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveYes,
	},
	{
		Name:            "包含字符串",
		Code:            HTTPFirewallRuleOperatorContains,
		Description:     "包含某个字符串",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "不包含字符串",
		Code:            HTTPFirewallRuleOperatorNotContains,
		Description:     "不包含某个字符串",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "包含前缀",
		Code:            HTTPFirewallRuleOperatorPrefix,
		Description:     "包含某个前缀",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "包含后缀",
		Code:            HTTPFirewallRuleOperatorSuffix,
		Description:     "包含某个后缀",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "包含二进制数据",
		Code:            HTTPFirewallRuleOperatorContainsBinary,
		Description:     "包含一组二进制数据",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "不包含二进制数据",
		Code:            HTTPFirewallRuleOperatorNotContainsBinary,
		Description:     "不包含一组二进制数据",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "包含索引",
		Code:            HTTPFirewallRuleOperatorHasKey,
		Description:     "对于一组数据拥有某个键值或者索引",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "版本号大于",
		Code:            HTTPFirewallRuleOperatorVersionGt,
		Description:     "对比版本号大于",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "版本号小于",
		Code:            HTTPFirewallRuleOperatorVersionLt,
		Description:     "对比版本号小于",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "版本号范围",
		Code:            HTTPFirewallRuleOperatorVersionRange,
		Description:     "判断版本号在某个范围内，格式为version1,version2",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP等于",
		Code:            HTTPFirewallRuleOperatorEqIP,
		Description:     "将参数转换为IP进行对比",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP大于",
		Code:            HTTPFirewallRuleOperatorGtIP,
		Description:     "将参数转换为IP进行对比",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP大于等于",
		Code:            HTTPFirewallRuleOperatorGteIP,
		Description:     "将参数转换为IP进行对比",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP小于",
		Code:            HTTPFirewallRuleOperatorLtIP,
		Description:     "将参数转换为IP进行对比",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP小于等于",
		Code:            HTTPFirewallRuleOperatorLteIP,
		Description:     "将参数转换为IP进行对比",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP范围",
		Code:            HTTPFirewallRuleOperatorIPRange,
		Description:     "IP在某个范围之内，范围格式可以是英文逗号分隔的ip1,ip2，或者CIDR格式的ip/bits",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "不在IP范围",
		Code:            HTTPFirewallRuleOperatorNotIPRange,
		Description:     "IP不在某个范围之内，范围格式可以是英文逗号分隔的ip1,ip2，或者CIDR格式的ip/bits",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP取模10",
		Code:            HTTPFirewallRuleOperatorIPMod10,
		Description:     "对IP参数值取模，除数为10，对比值为余数",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP取模100",
		Code:            HTTPFirewallRuleOperatorIPMod100,
		Description:     "对IP参数值取模，除数为100，对比值为余数",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
	{
		Name:            "IP取模",
		Code:            HTTPFirewallRuleOperatorIPMod,
		Description:     "对IP参数值取模，对比值格式为：除数,余数，比如10,1",
		CaseInsensitive: HTTPFirewallRuleCaseInsensitiveNo,
	},
}

func FindRuleOperatorName(code string) string {
	for _, operator := range AllRuleOperators {
		if operator.Code == code {
			return operator.Name
		}
	}
	return ""
}
