package blocklogic_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	blocklogic "github.com/navigacontentlab/block-logic"
	"github.com/navigacontentlab/navigadoc/doc"
)

var testdataDir = "testdata"

func MustGetConditionFromFile(conditionFile string) blocklogic.Condition {
	jsonFile, err := os.Open(fmt.Sprintf("%s/%s", testdataDir, conditionFile))
	if err != nil {
		panic(`failed to open file: ` + conditionFile + ` ` + err.Error())
	}

	byteData, err := io.ReadAll(jsonFile)

	if err != nil {
		panic(`failed read file: ` + conditionFile + ` ` + err.Error())
	}

	condition := blocklogic.Condition{}
	err = json.Unmarshal(byteData, &condition)

	if err != nil {
		panic(`failed to unmmarshal ` + conditionFile + ` ` + err.Error())
	}

	return condition
}

func MustGetDocumentFromFile(documentFile string) doc.Document {
	jsonFile, err := os.Open(fmt.Sprintf("%s/%s", testdataDir, documentFile))
	if err != nil {
		panic(`failed to open file: ` + documentFile + ` ` + err.Error())
	}

	byteData, err := io.ReadAll(jsonFile)

	if err != nil {
		panic(`failed read file: ` + documentFile + ` ` + err.Error())
	}

	document := doc.Document{}
	err = json.Unmarshal(byteData, &document)

	if err != nil {
		panic(`failed to unmmarshal ` + documentFile + ` ` + err.Error())
	}

	return document
}

func Test(t *testing.T) {
	tests := []struct {
		name      string
		condition blocklogic.Condition
		want      bool
		document  doc.Document
	}{
		{
			name:      "Test select unit A shared with unit B",
			condition: MustGetConditionFromFile("unit-A-or-B-test-condition.json"),
			document:  MustGetDocumentFromFile("unit-A-shared-B-document.json"),
			want:      true,
		},
		{
			name:      "test documents with links using in links",
			condition: MustGetConditionFromFile("condition-1-links.json"),
			document:  MustGetDocumentFromFile("document-1-links.json"),
			want:      true,
		},
		{
			name:      "test documents with links using in meta",
			condition: MustGetConditionFromFile("condition-1-meta.json"),
			document:  MustGetDocumentFromFile("document-1-meta.json"),
			want:      true,
		},
		{
			name:      "test documents with links using in content",
			condition: MustGetConditionFromFile("condition-1-content.json"),
			document:  MustGetDocumentFromFile("document-1-content.json"),
			want:      true,
		},
		{
			name:      "test documents with links using in *",
			condition: MustGetConditionFromFile("condition-1-all.json"),
			document:  MustGetDocumentFromFile("document-1-links.json"),
			want:      true,
		},
		{
			name:      "test documents with meta using in *",
			condition: MustGetConditionFromFile("condition-1-all.json"),
			document:  MustGetDocumentFromFile("document-1-meta.json"),
			want:      true,
		},
		{
			name:      "test documents with content using in *",
			condition: MustGetConditionFromFile("condition-1-all.json"),
			document:  MustGetDocumentFromFile("document-1-content.json"),
			want:      true,
		},
		{
			name:      "from section y in channel x",
			condition: MustGetConditionFromFile("condition-2.json"),
			document:  MustGetDocumentFromFile("document-2.json"),
			want:      true,
		},
		{
			name:      "from (section y in channel x) or channel x2",
			condition: MustGetConditionFromFile("condition-3.json"),
			document:  MustGetDocumentFromFile("document-3.json"),
			want:      true,
		},
		{
			name:      "test empty document, no links",
			condition: MustGetConditionFromFile("condition-1-links.json"),
			document:  doc.Document{},
			want:      false,
		},
		{
			name:      "test has channel x",
			condition: MustGetConditionFromFile("condition-4.json"),
			document:  MustGetDocumentFromFile("document-1-links.json"),
			want:      true,
		},
		{
			name:      "negative test or",
			condition: MustGetConditionFromFile("condition-negative-or.json"),
			document:  MustGetDocumentFromFile("document-negative.json"),
			want:      false,
		},
		{
			name:      "negative test and",
			condition: MustGetConditionFromFile("condition-negative-and.json"),
			document:  MustGetDocumentFromFile("document-negative.json"),
			want:      false,
		},
		{
			name:      "readme example 1",
			condition: MustGetConditionFromFile("readme-condition-1.json"),
			document:  MustGetDocumentFromFile("readme-document-1.json"),
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.condition
			if got := c.TestDocument(tt.document); got != tt.want {
				t.Errorf("Test() = %v, want %v", got, tt.want)
			}
		})
	}
}
