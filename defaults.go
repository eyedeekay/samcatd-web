// +build webface

package samcatweb

func defaultJS() string {
	r := "" +
		"" +
		"\n"
	return r
}

func defaultCSS() string {
	r := "\n" +
		".btn {\n" +
		"  text-align: center;\n" +
		"  color: #9789a7;\n" +
		"  border: dotted;\n" +
		"  width: 70px;\n" +
		"}\n" +
		"\n" +
		".toolbar {\n" +
		"  margin-left: auto;\n" +
		"  margin-right: auto;\n" +
		"  background-color: #0e1111;\n" +
		"  width: 420px;\n" +
		"}\n" +
		".parent {\n" +
		"  margin-left: 60px;\n" +
		"  background-color: #0e1111;\n" +
		"  color: #9789a7;\n" +
		"}\n" +
		"\n"
	return r
}
