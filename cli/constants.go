package cli

var ARGUMENTS Arguments

type Flag struct {
	Name        string
	Default     bool
	Description string
}

type Arguments struct {
	flags     map[string]*bool
	InputFile string
}

var OPTIONS []Flag = []Flag{
	{
		"minify",
		false,
		"minifies the html output",
	},
	{
		"toc-full",
		false,
		"generates a full toc, includes headings 1,2,3,4,5,6",
	},
	{
		"silent",
		false,
		"disables info logs, reduces output significantly",
	},
	{
		"toc",
		false,
		"generates a table of contents at the top of the output file, includes headings 1,2,3",
	},
	{
		"keep-temp",
		false,
		"keeps fleck from removing temporary files, used for debug purposes",
	},
	{
		"preprocessor-enabled",
		false,
		"enables the preprocessor to replace macros, decreases performance",
	},
	{
		"shell-macro-enabled",
		false,
		"enables the dangerous '@shell{command}' macro, which allows the preprocessor to run any command on your system",
	},
}
