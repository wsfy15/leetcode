/*
 * @lc app=leetcode.cn id=68 lang=golang
 *
 * [68] 文本左右对齐
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	n, count := len(words), 0
	res := []string{}
	start := 0
	for i := 0; i < n; i++ {
		count += len(words[i]) + 1 // +1表示每个单词后的空格

		// 最后一个单词，或者加上下一个单词长度后超过最大长度
		if i + 1 == n || count + len(words[i + 1]) > maxWidth {
			res = append(res, fillWords(words, start, i, maxWidth, i + 1 == n))
			start = i + 1
			count = 0
		}
	}

	return res
}

func fillWords(words []string, start, end, maxWidth int, lastLine bool) string {
	sb := strings.Builder{}
	wordCount, spaceCount := end - start + 1, maxWidth + 1
	for i := start; i <= end; i++ {
		spaceCount -= len(words[i]) 
		spaceCount-- // 每个单词末尾的空格，最后一个单词不考虑，因此上面已加一
	}
	
	t := 0 // 前t个单词需要多1个空格
	if wordCount > 1 {
		t = spaceCount % (wordCount - 1)
		spaceCount /= (wordCount - 1) 
	}

	for i := start; i < end; i++ {
		sb.WriteString(words[i])
		if lastLine {
			// 最后一行单词间不插入额外空格
			sb.WriteByte(' ')
			continue
		}

		for j := 0; j <= spaceCount; j++ {
			sb.WriteByte(' ')
		}
		if i - start < t {
			sb.WriteByte(' ')
		}
	}

	sb.WriteString(words[end])
	left := maxWidth - len(sb.String())
	for i := 0; i < left; i++ {
		sb.WriteByte(' ')
	}

	return sb.String()
}
// @lc code=end

