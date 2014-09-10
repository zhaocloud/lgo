package lgo

import (
	"strings"
	"litgh/lgo/util"
)

const pathSeparator = "/"

type requestMapping struct {
	urlMap map[string][]*controllerInfo
}

func (rm *requestMapping) Add(method, path string, controller RestController) {
	info := &controllerInfo{method: method, pattern: path, handler: controller}
	rm.urlMap[method] = append(rm.urlMap[method], info)
	rm.urlMap[path] = []*controllerInfo{info}
}

func (rm *requestMapping) GetHandler(r *Request) *controllerInfo {
	lookupPath := r.URL.Path
	directPathMatches := rm.urlMap[lookupPath]
	if directPathMatches != nil {
		return directPathMatches[0]
	}
	for _, mapping := range rm.urlMap[r.Method] {
		if match, uriTemplateVariables := matchPattern(lookupPath, mapping.pattern); match {
			r.PathParams = uriTemplateVariables
			return mapping
		}
	}
	return nil
}

func matchPattern(lookupPath, pattern string) (bool, map[string]string) {
	if strings.HasPrefix(pattern, pathSeparator) != strings.HasPrefix(lookupPath, pathSeparator) {
		return false, nil
	}
	patternDirs := strings.Split(pattern, pathSeparator)
	pathDirs := strings.Split(lookupPath, pathSeparator)
	if len(patternDirs) != len(pathDirs) {
		return false, nil
	}
	uriTemplateVariables := map[string]string{}
	patternIdxStart, patternIdxEnd := 0, len(patternDirs)-1
	pathIdxStart, pathIdxEnd := 0, len(pathDirs)-1

	for patternIdxStart <= patternIdxEnd && pathIdxStart <= pathIdxEnd {
		patternDir := patternDirs[patternIdxStart]
		match, key, val := util.GetStringMatcher(patternDir).MatchString(pathDirs[pathIdxStart])
		if !match {
			return false, nil
		}
		if key != "" {
			uriTemplateVariables[key] = val
		}
		patternIdxStart++
		pathIdxStart++
	}
	if pathIdxStart > pathIdxEnd {
		if patternIdxStart > patternIdxEnd {
			if strings.HasSuffix(pattern, pathSeparator) {
				return strings.HasSuffix(lookupPath, pathSeparator), uriTemplateVariables
			}
			return !strings.HasSuffix(lookupPath, pathSeparator), uriTemplateVariables
		}
		return true, uriTemplateVariables
	} else if patternIdxStart > patternIdxEnd {
		return false, nil
	}
	return false, nil
}
