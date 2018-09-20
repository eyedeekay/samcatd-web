// +build webface

package samcatweb

func defaultJS() string {
	r := "" +
		"" +
		"" +
		""
	return r
}

func defaultCSS() string {
	r := ".btn {" +
		"  text-align: center;" +
		"  color: red;" +
		"  border: dotted;" +
		"}" +
		"" +
		".toolbar {" +
		"  margin-left: auto;" +
		"  margin-right: auto;" +
		"  color: red;" +
		"}" +
		"\n"
	return r
}
