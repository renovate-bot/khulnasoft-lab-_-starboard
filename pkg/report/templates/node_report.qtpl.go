// Code generated by qtc from "node_report.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line pkg/report/templates/node_report.qtpl:1
package templates

//line pkg/report/templates/node_report.qtpl:1
import "github.com/khulnasoft-lab/starboard/pkg/apis/khulnasoft/v1alpha1"

//line pkg/report/templates/node_report.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line pkg/report/templates/node_report.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line pkg/report/templates/node_report.qtpl:3
func (p *NodeReport) StreamTitle(qw422016 *qt422016.Writer) {
//line pkg/report/templates/node_report.qtpl:3
	qw422016.N().S(`
Khulnasoft Starboard Node Security Report - Node: `)
//line pkg/report/templates/node_report.qtpl:4
	qw422016.E().S(p.Node.Name)
//line pkg/report/templates/node_report.qtpl:4
	qw422016.N().S(`
`)
//line pkg/report/templates/node_report.qtpl:5
}

//line pkg/report/templates/node_report.qtpl:5
func (p *NodeReport) WriteTitle(qq422016 qtio422016.Writer) {
//line pkg/report/templates/node_report.qtpl:5
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/report/templates/node_report.qtpl:5
	p.StreamTitle(qw422016)
//line pkg/report/templates/node_report.qtpl:5
	qt422016.ReleaseWriter(qw422016)
//line pkg/report/templates/node_report.qtpl:5
}

//line pkg/report/templates/node_report.qtpl:5
func (p *NodeReport) Title() string {
//line pkg/report/templates/node_report.qtpl:5
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/report/templates/node_report.qtpl:5
	p.WriteTitle(qb422016)
//line pkg/report/templates/node_report.qtpl:5
	qs422016 := string(qb422016.B)
//line pkg/report/templates/node_report.qtpl:5
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/report/templates/node_report.qtpl:5
	return qs422016
//line pkg/report/templates/node_report.qtpl:5
}

//line pkg/report/templates/node_report.qtpl:7
func (p *NodeReport) StreamBody(qw422016 *qt422016.Writer) {
//line pkg/report/templates/node_report.qtpl:7
	qw422016.N().S(`
<div class="container">

  <div class="col mt-5">
    <div class="row text-center">`)
//line pkg/report/templates/node_report.qtpl:11
	streamimgKhulnasoftLogo(qw422016)
//line pkg/report/templates/node_report.qtpl:11
	qw422016.N().S(`</div>
    <div class="row mt-4 text-center">
      <h2 class="text-muted mx-auto">Khulnasoft Starboard Node Security Report</h2>
    </div>
    <div class="row text-center">
      <h3 class="text-muted mx-auto">`)
//line pkg/report/templates/node_report.qtpl:16
	qw422016.E().S(string(p.Node.Kind))
//line pkg/report/templates/node_report.qtpl:16
	qw422016.N().S(`: `)
//line pkg/report/templates/node_report.qtpl:16
	qw422016.E().S(p.Node.Name)
//line pkg/report/templates/node_report.qtpl:16
	qw422016.N().S(`</h3>
    </div>
    <div class="row text-center">
      <h3 class="text-muted mx-auto">Generated on `)
//line pkg/report/templates/node_report.qtpl:19
	qw422016.E().S(p.GeneratedAt.Format("2 Jan 2006 15:04:01"))
//line pkg/report/templates/node_report.qtpl:19
	qw422016.N().S(`</h3>
    </div>
  </div>

<!-- Resume START -->
  `)
//line pkg/report/templates/node_report.qtpl:24
	if p.CisKubeBenchReport != nil {
//line pkg/report/templates/node_report.qtpl:24
		qw422016.N().S(`

  <div class="row text-center border-bottom mt-4">
      <h3 class="mx-auto " id="vuln_header" style="color: rgb(0, 160, 170);">CIS Benchmarks for Kubernetes </h3>
  </div>
  <!-- Cards -->
  <div class="">
      <div class="row my-5" style="font-size:small;">
          <!-- Scanner -->
          <div class="col-3 border rounded shadow px-3 py-2 ml-4 ">
              <div class="row text-center">
                  <div class="col">
                      <p class="mb-2 pb-1 border-bottom">Scanner</p>
                  </div>
               </div>
               <div class="row">
                  <div class="col">
                  `)
//line pkg/report/templates/node_report.qtpl:42
		report := p.CisKubeBenchReport.Report
		scanner_name := report.Scanner.Name
		scanner_vendor := report.Scanner.Vendor
		scanner_version := report.Scanner.Version
		creation_timestamp := report.UpdateTimestamp.Format("2 Jan 2006 15:04:01")

//line pkg/report/templates/node_report.qtpl:47
		qw422016.N().S(`
                      <p class="my-0">Name:  `)
//line pkg/report/templates/node_report.qtpl:48
		qw422016.E().S(scanner_name)
//line pkg/report/templates/node_report.qtpl:48
		qw422016.N().S(`</p>
                      <p class="my-0">Vendor:  `)
//line pkg/report/templates/node_report.qtpl:49
		qw422016.E().S(scanner_vendor)
//line pkg/report/templates/node_report.qtpl:49
		qw422016.N().S(`</p>
                      <p class="my-0">Version:  `)
//line pkg/report/templates/node_report.qtpl:50
		qw422016.E().S(scanner_version)
//line pkg/report/templates/node_report.qtpl:50
		qw422016.N().S(`</p>
                  </div>
               </div>
          </div>
          <!-- summary -->
          <div class="col-5 border rounded shadow py-2 mx-auto ">
              <div class="row text-center">
                 <div class="col">
                     <p class="mb-2 pb-1 border-bottom">Summary</p>
                 </div>
              </div>
              <div class="row">
                  `)
//line pkg/report/templates/node_report.qtpl:63
		summary := report.Summary

//line pkg/report/templates/node_report.qtpl:64
		qw422016.N().S(`
                  `)
//line pkg/report/templates/node_report.qtpl:65
		if summary.FailCount > 0 {
//line pkg/report/templates/node_report.qtpl:65
			qw422016.N().S(`
                  <div class="col text-center p-0 text-danger font-weight-bold">
                  `)
//line pkg/report/templates/node_report.qtpl:67
		} else {
//line pkg/report/templates/node_report.qtpl:67
			qw422016.N().S(`
                  <div class="col text-center p-0">
                  `)
//line pkg/report/templates/node_report.qtpl:69
		}
//line pkg/report/templates/node_report.qtpl:69
		qw422016.N().S(`
                      <p class="mx-auto mb-1">`)
//line pkg/report/templates/node_report.qtpl:70
		qw422016.N().D(summary.FailCount)
//line pkg/report/templates/node_report.qtpl:70
		qw422016.N().S(`</p>
                      <p class="mx-auto ">FAIL</p>
                  </div>
                  `)
//line pkg/report/templates/node_report.qtpl:73
		if summary.WarnCount > 0 {
//line pkg/report/templates/node_report.qtpl:73
			qw422016.N().S(`
                  <div class="col text-center p-0 text-warning font-weight-bold">
                  `)
//line pkg/report/templates/node_report.qtpl:75
		} else {
//line pkg/report/templates/node_report.qtpl:75
			qw422016.N().S(`
                  <div class="col text-center p-0">
                  `)
//line pkg/report/templates/node_report.qtpl:77
		}
//line pkg/report/templates/node_report.qtpl:77
		qw422016.N().S(`
                      <p class="mx-auto mb-1">`)
//line pkg/report/templates/node_report.qtpl:78
		qw422016.N().D(summary.WarnCount)
//line pkg/report/templates/node_report.qtpl:78
		qw422016.N().S(`</p>
                      <p class="mx-auto ">WARN</p>
                  </div>
                  <div class="col text-center p-0">
                      <p class="mx-auto mb-1">`)
//line pkg/report/templates/node_report.qtpl:82
		qw422016.N().D(summary.InfoCount)
//line pkg/report/templates/node_report.qtpl:82
		qw422016.N().S(`</p>
                      <p class="mx-auto ">INFO</p>
                  </div>
                  <div class="col text-center p-0">
                      <p class="mx-auto mb-1">`)
//line pkg/report/templates/node_report.qtpl:86
		qw422016.N().D(summary.PassCount)
//line pkg/report/templates/node_report.qtpl:86
		qw422016.N().S(`</p>
                      <p class="mx-auto ">PASS</p>
                  </div>
              </div>
          </div>
          <!-- Metadata -->
          <div class="col-3 border rounded shadow px-3 py-2 mr-4">
              <div class="row text-center">
                  <div class="col">
                      <p class="mb-2 pb-1 border-bottom">Metadata</p>
                  </div>
               </div>
               <div class="row">
                  <div class="col">
                      <p class="my-0">
                          Generated at:  `)
//line pkg/report/templates/node_report.qtpl:101
		qw422016.E().S(creation_timestamp)
//line pkg/report/templates/node_report.qtpl:101
		qw422016.N().S(`
                      </p>
                  </div>
               </div>
          </div>
      </div>
  </div>
  `)
//line pkg/report/templates/node_report.qtpl:108
	}
//line pkg/report/templates/node_report.qtpl:108
	qw422016.N().S(`
<!-- Resume END -->

<!-- Sections START -->
  `)
//line pkg/report/templates/node_report.qtpl:113
	report := p.CisKubeBenchReport.Report

//line pkg/report/templates/node_report.qtpl:114
	qw422016.N().S(`
  <div class="row">
  `)
//line pkg/report/templates/node_report.qtpl:116
	for _, section := range report.Sections {
//line pkg/report/templates/node_report.qtpl:116
		qw422016.N().S(`
    <table class="table table-sm table-bordered">
      <thead>
        <tr>
          <th scope="col">Test No.</th>
          <th scope="col">Status</th>
          <th scope="col">Test Description</th>
          <th scope="col">Remediation</th>
        </tr>
      </thead>
      <tbody>
   <h3> `)
//line pkg/report/templates/node_report.qtpl:127
		qw422016.E().S(section.Text)
//line pkg/report/templates/node_report.qtpl:127
		qw422016.N().S(` </h3>
   `)
//line pkg/report/templates/node_report.qtpl:128
		for _, test := range section.Tests {
//line pkg/report/templates/node_report.qtpl:128
			qw422016.N().S(`
    `)
//line pkg/report/templates/node_report.qtpl:129
			for _, result := range test.Results {
//line pkg/report/templates/node_report.qtpl:129
				qw422016.N().S(`
      <tr>
        <td>`)
//line pkg/report/templates/node_report.qtpl:131
				qw422016.E().S(result.TestNumber)
//line pkg/report/templates/node_report.qtpl:131
				qw422016.N().S(`</td>
        <td>`)
//line pkg/report/templates/node_report.qtpl:132
				qw422016.E().S(result.Status)
//line pkg/report/templates/node_report.qtpl:132
				qw422016.N().S(`</td>
        <td>`)
//line pkg/report/templates/node_report.qtpl:133
				qw422016.E().S(result.TestDesc)
//line pkg/report/templates/node_report.qtpl:133
				qw422016.N().S(`</td>
        <td>`)
//line pkg/report/templates/node_report.qtpl:134
				qw422016.E().S(result.Remediation)
//line pkg/report/templates/node_report.qtpl:134
				qw422016.N().S(`</td>
      </tr>
`)
//line pkg/report/templates/node_report.qtpl:136
			}
//line pkg/report/templates/node_report.qtpl:136
			qw422016.N().S(`
      </tbody>
    `)
//line pkg/report/templates/node_report.qtpl:138
		}
//line pkg/report/templates/node_report.qtpl:138
		qw422016.N().S(`
    </table>
`)
//line pkg/report/templates/node_report.qtpl:140
	}
//line pkg/report/templates/node_report.qtpl:140
	qw422016.N().S(`
  </div>
<!-- Sections END -->

</div>
`)
//line pkg/report/templates/node_report.qtpl:145
}

//line pkg/report/templates/node_report.qtpl:145
func (p *NodeReport) WriteBody(qq422016 qtio422016.Writer) {
//line pkg/report/templates/node_report.qtpl:145
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/report/templates/node_report.qtpl:145
	p.StreamBody(qw422016)
//line pkg/report/templates/node_report.qtpl:145
	qt422016.ReleaseWriter(qw422016)
//line pkg/report/templates/node_report.qtpl:145
}

//line pkg/report/templates/node_report.qtpl:145
func (p *NodeReport) Body() string {
//line pkg/report/templates/node_report.qtpl:145
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/report/templates/node_report.qtpl:145
	p.WriteBody(qb422016)
//line pkg/report/templates/node_report.qtpl:145
	qs422016 := string(qb422016.B)
//line pkg/report/templates/node_report.qtpl:145
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/report/templates/node_report.qtpl:145
	return qs422016
//line pkg/report/templates/node_report.qtpl:145
}

//line pkg/report/templates/node_report.qtpl:147
func streamnodeReference(qw422016 *qt422016.Writer, section []v1alpha1.CISKubeBenchSection) {
//line pkg/report/templates/node_report.qtpl:147
	qw422016.N().S(`
  `)
//line pkg/report/templates/node_report.qtpl:148
	qw422016.E().S(section[0].ID)
//line pkg/report/templates/node_report.qtpl:148
	qw422016.N().S(`/`)
//line pkg/report/templates/node_report.qtpl:148
	qw422016.E().S(section[0].Text)
//line pkg/report/templates/node_report.qtpl:148
	qw422016.N().S(`:`)
//line pkg/report/templates/node_report.qtpl:148
	qw422016.E().S(section[0].NodeType)
//line pkg/report/templates/node_report.qtpl:148
	qw422016.N().S(`/
`)
//line pkg/report/templates/node_report.qtpl:149
}

//line pkg/report/templates/node_report.qtpl:149
func writenodeReference(qq422016 qtio422016.Writer, section []v1alpha1.CISKubeBenchSection) {
//line pkg/report/templates/node_report.qtpl:149
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/report/templates/node_report.qtpl:149
	streamnodeReference(qw422016, section)
//line pkg/report/templates/node_report.qtpl:149
	qt422016.ReleaseWriter(qw422016)
//line pkg/report/templates/node_report.qtpl:149
}

//line pkg/report/templates/node_report.qtpl:149
func nodeReference(section []v1alpha1.CISKubeBenchSection) string {
//line pkg/report/templates/node_report.qtpl:149
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/report/templates/node_report.qtpl:149
	writenodeReference(qb422016, section)
//line pkg/report/templates/node_report.qtpl:149
	qs422016 := string(qb422016.B)
//line pkg/report/templates/node_report.qtpl:149
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/report/templates/node_report.qtpl:149
	return qs422016
//line pkg/report/templates/node_report.qtpl:149
}
