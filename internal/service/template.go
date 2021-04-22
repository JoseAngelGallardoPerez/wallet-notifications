package service

import (
	"text/template"
)

var (
	mailTemplate = template.Must(template.New("mail").Parse(`
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN""http://www.w3.org/TR/html4/strict.dtd">
<html lang="en">
<head>
  <meta http-equiv="Content-Type" content="text/html charset=UTF-8" />
  <meta name="viewport"
        content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title></title>
</head>
<body>
<table align="center" style="width: 100%;bg-color:#E5E5E5;">
  <tr style="bg-color:#E5E5E5;">
    <td style="padding-top:0;
    padding-bottom:0;
    padding-right:0;
    padding-left:0;
    margin:0px;
bg-color:#E5E5E5;">
      <table border="0" cellpadding="0" cellspacing="0" style="background-color: #FFFFFF;font-family: 'SF Pro Display', 'Graphik', sans-serif; margin:0; padding:0" width="100%">
        <tr width="100%">
          <td width="100%" height="100%" style="text-align: center;">
            <table border="0" cellpadding="0" cellspacing="0" width="100%" style="margin: 0 auto; padding:0; max-width: 788px">
              {{if .Logo}}
              <tr class="logo" style="margin: 0 auto;">
                <td style="padding: 78px 0 60px; margin: 0 auto; " >
                  <center>
                    <img src="{{ .Logo }}" alt="dominicacapital" style="
                    height: 30px;
                    width: 154.34px;">
                  </center>
                </td>
              </tr>
              {{end}}
              
              <tr width="100%">
                <td width="100%" style="
                border-top-right-radius: 5px;
                border-top-left-radius: 5px;
                overflow: hidden;
                text-align: left;
                color:#000000">
                  <table border="0" cellpadding="0" cellspacing="0" width="100%" style="border-spacing: 0;
                  border-collapse: collapse;
                  overflow: hidden;
                  border-top-left-radius: 5px;
                  border-top-right-radius: 5px;
                  width: 100%;">
                    <tr>
                      <td style="width: 100%;
                      border-top: 5px solid #229932;
                      border-top-left-radius: 5px;
                      border-top-right-radius: 5px;
                      padding: 0;
                      text-align: center;">
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <p style="margin: 32px 68px 24px;
                        font-family: 'Graphik', sans-serif;
                        text-align: left;
                        font-style: normal;
                        font-weight: normal;
                        font-size: 18px;
                        line-height: 200%;
                        color: #000000;">
                          
						{{ .Body }}

						</p>
                      </td>
                    </tr>
                    <tr style="
                      width: 100%;">
                      <td style="
                      width: 100%;
                      padding: 0 68px 44px;">
                        <table style="">
                          <tr style="width: 100%;">
                            <td style="width: 100%;">
                                <a href="#" target="_blank" style="
                                font-family: 'SF Pro Display', 'Arial', sans-serif;
                                text-decoration: underline;
                                color: #229932;
                                text-align: center;
                                font-size: 16px;
                                font-style: normal;
                                font-weight: bold;
                                ">
                                  <span style="color: #229932; text-decoration: underline">Visit {{.SiteName}}</span></a>
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
​
					{{if .Signature}}
                    <tr style="width: 100%;">
                      <td style="width: 100%;">
                        <table border="0" cellpadding="0" cellspacing="0" style="
                         border-spacing: 0;
                         width: 100%;
                         padding: 24px 68px;
                         border-spacing: 0;
                         border-bottom-right-radius: 5px ;
                         border-bottom-left-radius: 5px;
                         background-color: #F8F8F8;
                         background: linear-gradient(180deg, #F8F8F8 0%, rgba(248, 248, 248, 0) 91.22%);">
                          <tr style="width: 100%;">
                            <td style="width: 100%;
                            text-align-last: justify;
                            text-align: justify;">
                              <table style="
                              display: inline-table;
                              text-align: left;
                              text-align-last: left;
                              max-width: 330px;
                              vertical-align:top;">
                                <tbody style="width: 100%;">
                                  <tr style="width: 100%;">
                                    <td style="width: 100%;">
                                      <p style="
                                      margin:0;
                                      font-family: 'SF Pro Display', 'Arial', sans-serif;
                                      font-style: normal;
                                      font-weight: normal;
                                      font-size: 13px;
                                      line-height: 150%;
                                      letter-spacing: 0.5px;
                                      color: #999999;">
                                        {{ .Signature }}
                                      </p>
                                    </td>
                                  </tr>
                                </tbody>
                              </table>
                            </td>
                          </tr>
                        </table>
                      </td>
                    </tr>
                    {{end}}
                  </table>
                </td>
              </tr>
              <tr>
                <td style="">
                  <p style="
                  margin: 32px auto;
                  text-align: center;
                  font-family: 'Graphik', 'Arial', sans-serif;
                  font-style: normal;
                  font-weight: normal;
                  font-size: 13px;
                  line-height: 100%;
                  letter-spacing: 0.5px;
                  text-decoration: none;
                  color: #999999;">
                    This message was automatically generated by {{.SiteName}}.
                  </p>
                </td>
              </tr>
            </table>
          </td>
        </tr>
      </table>
    </td>
  </tr>
</table>
​
</body>
</html>
`))

	logoTemplate = template.Must(template.New("mail").Parse(`
<img src="{{ .Logo }}" alt="mail-logo" border="0" width="200"  style="display:block; margin: auto;"/>		      
`))
)
