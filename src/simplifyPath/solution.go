package simplifyPath

import p "path"

// https://leetcode.com/problems/simplify-path/
func simplifyPath(path string) string {
	return p.Clean(path)
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Simplify Path.
//Memory Usage: 2.6 MB, less than 97.78% of Go online submissions for Simplify Path.