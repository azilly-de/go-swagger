// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generator

import (
	"encoding/json"
	"regexp"
	"text/template"
)

//go:generate go-bindata -pkg=generator -ignore=.*\.sw? ./templates/...

// fwiw, don't get attached to this, still requires a better abstraction

var (
	modelTemplate *template.Template
	// modelValidatorTemplate *template.Template
	operationTemplate      *template.Template
	parameterTemplate      *template.Template
	responsesTemplate      *template.Template
	builderTemplate        *template.Template
	mainTemplate           *template.Template
	mainDocTemplate        *template.Template
	embeddedSpecTemplate   *template.Template
	configureAPITemplate   *template.Template
	clientTemplate         *template.Template
	clientParamTemplate    *template.Template
	clientResponseTemplate *template.Template
	clientFacadeTemplate   *template.Template
)

var assets = map[string][]byte{
	"validation/primitive.gotmpl":           MustAsset("templates/validation/primitive.gotmpl"),
	"validation/customformat.gotmpl":        MustAsset("templates/validation/customformat.gotmpl"),
	"docstring.gotmpl":                      MustAsset("templates/docstring.gotmpl"),
	"validation/structfield.gotmpl":         MustAsset("templates/validation/structfield.gotmpl"),
	"modelvalidator.gotmpl":                 MustAsset("templates/modelvalidator.gotmpl"),
	"structfield.gotmpl":                    MustAsset("templates/structfield.gotmpl"),
	"tupleserializer.gotmpl":                MustAsset("templates/tupleserializer.gotmpl"),
	"additionalpropertiesserializer.gotmpl": MustAsset("templates/additionalpropertiesserializer.gotmpl"),
	"schematype.gotmpl":                     MustAsset("templates/schematype.gotmpl"),
	"schemabody.gotmpl":                     MustAsset("templates/schemabody.gotmpl"),
	"schema.gotmpl":                         MustAsset("templates/schema.gotmpl"),
	"schemavalidator.gotmpl":                MustAsset("templates/schemavalidator.gotmpl"),
	"model.gotmpl":                          MustAsset("templates/model.gotmpl"),
	"header.gotmpl":                         MustAsset("templates/header.gotmpl"),
	"swagger_json_embed.gotmpl":             MustAsset("templates/swagger_json_embed.gotmpl"),

	"server/parameter.gotmpl":    MustAsset("templates/server/parameter.gotmpl"),
	"server/responses.gotmpl":    MustAsset("templates/server/responses.gotmpl"),
	"server/operation.gotmpl":    MustAsset("templates/server/operation.gotmpl"),
	"server/builder.gotmpl":      MustAsset("templates/server/builder.gotmpl"),
	"server/configureapi.gotmpl": MustAsset("templates/server/configureapi.gotmpl"),
	"server/main.gotmpl":         MustAsset("templates/server/main.gotmpl"),
	"server/doc.gotmpl":          MustAsset("templates/server/doc.gotmpl"),

	"client/parameter.gotmpl": MustAsset("templates/client/parameter.gotmpl"),
	"client/response.gotmpl":  MustAsset("templates/client/response.gotmpl"),
	"client/client.gotmpl":    MustAsset("templates/client/client.gotmpl"),
	"client/facade.gotmpl":    MustAsset("templates/client/facade.gotmpl"),
}

var builtinTemplates = map[string]TemplateDefinition{

	"validatorTempl": {
		Dependencies: []string{
			"primitivevalidator",
			"customformatvalidator",
		},
	},

	"primitivevalidator": {
		Files: []string{"validation/primitive.gotmpl"},
	},
	"customformatvalidator": {
		Files: []string{"validation/customformat.gotmpl"},
	},

	"modelValidatorTemplate": {
		Dependencies: []string{"validatorTempl"},
	},

	"docstring": {
		Files: []string{"docstring.gotmpl"},
	},

	"propertyValidationDocString": {
		Files: []string{"validation/docstring.gotmpl"},
	},
	"schematype": {
		Files: []string{"schematype.gotmpl"},
	},
	"body": {
		Files: []string{"schemabody.gotmpl"},
	},
	"schema": {
		Files: []string{"schema.gotmpl"},
	},
	"schemavalidations": {
		Files: []string{"schemavalidator.gotmpl"},
	},
	"header": {
		Files: []string{"header.gotmpl"},
	},
	"fields": {
		Files: []string{"structfield.gotmpl"},
	},
	"tupleserializer": {
		Files: []string{"tupleserializer.gotmpl"},
	},
	"additionalpropertiesserializer": {
		Files: []string{"additionalpropertiesserializer.gotmpl"},
	},
	"model": {
		Dependencies: []string{
			"docstring",
			"primitivevalidator",
			"customformatvalidator",
			"propertyValidationDocString",
			"schematype",
			"body",
			"schema",
			"schemavalidations",
			"header",
			"fields",
			"tupleserializer",
			"additionalpropertiesserializer",
		},
		Files: []string{
			"model.gotmpl",
		},
	},

	"parameterTemplate": {
		Dependencies: []string{"model"},
		Files:        []string{"server/parameter.gotmpl"},
	},

	"responsesTemplate": {
		Dependencies: []string{"model"},
		Files:        []string{"server/responses.gotmpl"},
	},

	"operationTemplate": {
		Dependencies: []string{"model"},
		Files:        []string{"server/operation.gotmpl"},
	},

	"builderTemplate": {
		Files: []string{"server/builder.gotmpl"},
	},

	"configureAPITemplate": {
		Files: []string{"server/configureapi.gotmpl"},
	},

	"mainTemplate": {
		Files: []string{"server/main.gotmpl"},
	},

	"mainDocTemplate": {
		Files: []string{"server/doc.gotmpl"},
	},

	"embeddedSpecTemplate": {
		Files: []string{"swagger_json_embed.gotmpl"},
	},

	// Client templates
	"clientParamTemplate": {
		Dependencies: []string{"model"},
		Files:        []string{"client/parameter.gotmpl"},
	},

	"clientResponseTemplate": {
		Dependencies: []string{"model"},
		Files:        []string{"client/response.gotmpl"},
	},

	"clientTemplate": {
		Dependencies: []string{

			"docstring",
			"propertyValidationDocString",
			"schematype",
			"body",
		},
		Files: []string{
			"client/client.gotmpl",
		},
	},

	"clientFacadeTemplate": {
		Dependencies: []string{

			"docstring",
			"propertyValidationDocString",
			"schematype",
			"body",
		},
		Files: []string{
			"client/facade.gotmpl",
		},
	},
}

var (
	notNumberExp = regexp.MustCompile("[^0-9]")
)

var templates = NewTemplateRegistry()

func init() {

	for name, asset := range assets {
		templates.AddFile(name, asset)
	}

	for name, template := range builtinTemplates {
		templates.AddTemplate(name, template)
	}

	compileTemplates()
}

func compileTemplates() {

	modelTemplate = templates.MustGet("model")

	// common templates

	// modelValidatorTemplate = templates.MustGet("modelValidatorTemplate")

	// server templates
	parameterTemplate = templates.MustGet("parameterTemplate")

	responsesTemplate = templates.MustGet("responsesTemplate")

	operationTemplate = templates.MustGet("operationTemplate")
	builderTemplate = templates.MustGet("builderTemplate")           //template.Must(template.New("builder").Funcs(FuncMap).Parse(string(assets["server/builder.gotmpl"])))
	configureAPITemplate = templates.MustGet("configureAPITemplate") //template.Must(template.New("configureapi").Funcs(FuncMap).Parse(string(assets["server/configureapi.gotmpl"])))
	mainTemplate = templates.MustGet("mainTemplate")                 //template.Must(template.New("main").Funcs(FuncMap).Parse(string(assets["server/main.gotmpl"])))
	mainDocTemplate = templates.MustGet("mainDocTemplate")           //template.Must(template.New("meta").Funcs(FuncMap).Parse(string(assets["server/doc.gotmpl"])))

	embeddedSpecTemplate = templates.MustGet("embeddedSpecTemplate") //template.Must(template.New("embedded_spec").Funcs(FuncMap).Parse(string(assets["swagger_json_embed.gotmpl"])))

	// Client templates
	clientParamTemplate = templates.MustGet("clientParamTemplate")

	clientResponseTemplate = templates.MustGet("clientResponseTemplate")

	clientTemplate = templates.MustGet("clientTemplate")

	clientFacadeTemplate = templates.MustGet("clientFacadeTemplate")

}

func makeModelTemplate() *template.Template {
	templ := template.Must(template.New("docstring").Funcs(FuncMap).Parse(string(assets["docstring.gotmpl"])))
	templ = template.Must(templ.New("primitivevalidator").Parse(string(assets["validation/primitive.gotmpl"])))
	templ = template.Must(templ.New("customformatvalidator").Parse(string(assets["validation/customformat.gotmpl"])))
	templ = template.Must(templ.New("validationDocString").Parse(string(assets["validation/structfield.gotmpl"])))
	templ = template.Must(templ.New("schematype").Parse(string(assets["schematype.gotmpl"])))
	templ = template.Must(templ.New("body").Parse(string(assets["schemabody.gotmpl"])))
	templ = template.Must(templ.New("schema").Parse(string(assets["schema.gotmpl"])))
	templ = template.Must(templ.New("schemavalidations").Parse(string(assets["schemavalidator.gotmpl"])))
	templ = template.Must(templ.New("header").Parse(string(assets["header.gotmpl"])))
	templ = template.Must(templ.New("fields").Parse(string(assets["structfield.gotmpl"])))
	templ = template.Must(templ.New("tupleSerializer").Parse(string(assets["tupleserializer.gotmpl"])))
	templ = template.Must(templ.New("additionalpropertiesserializer.gotmpl").Parse(string(assets["additionalpropertiesserializer.gotmpl"])))
	templ = template.Must(templ.New("model").Parse(string(assets["model.gotmpl"])))
	return templ
}

func asJSON(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
