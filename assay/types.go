/*

Golang type declaration.

It is possible to declare any types as part of the suite implementation.
The type declartion can be offloaded either to shared library or common
declaration and re-used between suites.

Please see https://assay.it/doc/core for details.
*/

package assay

import (
	"fmt"

	"github.com/assay-it/tk"
)

// News a type used by the example application. This type models a core data of
// the application and used by suites to validates correctness of outputs.
type News struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// List is a sequence of news, a core type of example application.
type List []News

// Value and other functions implements sort.Interface and gurl.Ord interfaces
// for List data type. The implementation of these interfaces is mandatory if
// suite asserts and validates content of the sequence with ƒ.Seq.
//
// Please see https://assay.it/doc/core/style#focus-on-the-sequence
//
func (seq List) Value(i int) interface{} { return seq[i] }
func (seq List) Len() int                { return len(seq) }
func (seq List) Swap(i, j int)           { seq[i], seq[j] = seq[j], seq[i] }
func (seq List) Less(i, j int) bool      { return seq[i].ID < seq[j].ID }
func (seq List) String(i int) string     { return seq[i].ID }

// Settings of assay.it allows developers to customize suite via environment
// variables (See settings of repository). These variables are injected at runtime.
// Here, the example application requires a CONFIG_DOMAIN variable, which declares
// domain name of SUT. The assay toolkit api is used to fetch this variable form environment.
// The subdomain name is deducted from auto variable BUILD_ID. It corresponds to Pull
// Request Number for any assessment originated by WebHook.
var host = fmt.Sprintf("v%s.%s", tk.Env("BUILD_ID", ""), tk.Env("CONFIG_DOMAIN", ""))