package main

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow" // импортируем дополнительный анализатор
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
)

func main() {

	// определяем map подключаемых правил
	checks := map[string]bool{
		"SA5000": true,
		"SA6000": true,
		"SA9004": true,
	}

	mychecks := []*analysis.Analyzer{
		ErrCheckAnalyzer,
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
	}
	for _, v := range staticcheck.Analyzers {
		// добавляем в массив нужные проверки
		if checks[v.Analyzer.Name] {
			mychecks = append(mychecks, v.Analyzer)
		}
	}

	multichecker.Main(
		mychecks...,
	)

}
