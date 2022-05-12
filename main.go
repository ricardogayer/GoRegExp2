package main

import (
	"fmt"
	"strings"
)

func main() {

	var texto string = `EXEC sp_updateextendedproperty 
	@name = N'MS_Description', @value = 'Número da Apólice 
	  teste ',
	@level0type = N'Schema', @level0name = 'dbo', 
	@level1type = N'Table',  @level1name = 'tb_apolice', 
	@level2type = N'Column', @level2name = 'nr_apolice';
	
	EXEC sp_updateextendedproperty 
	@name = N'MS_Description', @value = 'Número da Sinistro',
	@level0type = N'Schema', @level0name = 'dbo', 
	@level1type = N'Table',  @level1name = 'tb_sinistro', 
	@level2type = N'Column', @level2name = 'nr_sinistro';`

	var padrao string = `EXEC sp_updateextendedproperty 
	@name = N'MS_Description', @value = '(?P<comment>(.*))',
	@level0type = N'Schema', @level0name = '(?P<schema>(.*))', 
	@level1type = N'Table',  @level1name = '(?P<table>(.*))', 
	@level2type = N'Column', @level2name = '(?P<column>(.*))';`

	padrao = strings.Replace(padrao, "\n", "", -1)
	padrao = strings.Replace(padrao, "\r", "", -1)
	padrao = strings.Replace(padrao, "EXEC sp_updateextendedproperty", "", -1)
	padrao = strings.Replace(padrao, "@name = N'MS_Description', ", "", -1)
	padrao = strings.Replace(padrao, "@level0type = N'Schema', ", "", -1)
	padrao = strings.Replace(padrao, "@level1type = N'Table',  ", "", -1)
	padrao = strings.Replace(padrao, "@level2type = N'Column', ", "", -1)

	texto = strings.Replace(texto, "\n", "", -1)
	texto = strings.Replace(texto, "\r", "", -1)
	texto = strings.Replace(texto, "EXEC sp_updateextendedproperty", "", -1)
	texto = strings.Replace(texto, "@name = N'MS_Description', ", "", -1)
	texto = strings.Replace(texto, "EXEC sp_updateextendedproperty", "", -1)
	texto = strings.Replace(texto, "@level0type = N'Schema', ", "", -1)
	texto = strings.Replace(texto, "@level1type = N'Table',  ", "", -1)
	texto = strings.Replace(texto, "@level2type = N'Column', ", "", -1)
	texto = strings.Replace(texto, "@value = ", "\n", -1)
	texto = strings.Replace(texto, ";", "", -1)
	// @level0name =  @level1name = @level2name =
	texto = strings.Replace(texto, "@level0name = ", "", -1)
	texto = strings.Replace(texto, "@level1name = ", "", -1)
	texto = strings.Replace(texto, "@level2name = ", "", -1)
	texto = strings.Replace(texto, "'", "\"", -1)

	fmt.Println(texto)

	// fmt.Println(texto)

	// datePattern := regexp.MustCompile(padrao)
	// dataResults := datePattern.FindAllStringSubmatch(texto, -1)

	// for _, dataResult := range dataResults {

	// 	var comment, schema, table, column string

	// 	schema = dataResult[datePattern.SubexpIndex("schema")]
	// 	table = dataResult[datePattern.SubexpIndex("table")]
	// 	column = dataResult[datePattern.SubexpIndex("column")]
	// 	comment = dataResult[datePattern.SubexpIndex("comment")]

	// 	fmt.Println("Schema:", schema)
	// 	fmt.Println("Table:", table)
	// 	fmt.Println("Column:", column)
	// 	fmt.Println("Comment:", comment)
	// 	fmt.Println("")
	// }

}
