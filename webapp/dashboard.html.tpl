<!DOCTYPE html>
<html>
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta http-equiv="refresh" content="300"> <!-- Refresh every 5 minutes -->
    <title>rlb-stats</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/css/bootstrap.min.css" integrity="sha384-GJzZqFGwb1QTTN6wy59ffF1BuGJpLSa9DkKMp0DgiMDm4iYMj70gZWKYbI706tWS" crossorigin="anonymous">
</head>
<body>
<table class="table">
    <thead>
    <tr>
        <th scope="col" style="width: 50%">Top Files</th>
        <th scope="col" style="width: 50%">Top Nodes</th>
    </tr>
    </thead>
    <tbody>
    <tr>
        <td><img src="/chart?type=by_file{{if .From}}&from={{.From}}{{end}}{{if .To}}&to={{.To}}{{end}}" class="img-fluid"></td>
        <td><img src="/chart?type=by_node{{if .From}}&from={{.From}}{{end}}{{if .To}}&to={{.To}}{{end}}" class="img-fluid"></td>
    </tr>
    <tr>
        <td>
            <table class="table table-hover">
                <thead>
                <tr>
                    <th scope="col">File Name</th>
                    <th scope="col">Volume</th>
                </tr>
                </thead>
                <tbody>
                {{ range .TopFiles }}
                    <tr>
                        <th scope="row"><a href="/file_stats?filename={{.Name}}{{if $.From}}&from={{$.From}}{{end}}{{if $.To}}&to={{$.To}}{{end}}">{{.Name}}</a></th>
                        <td>{{.Volume}}</td>
                    </tr>
                {{ end }}
                </tbody>
            </table>
        </td>
        <td>
            <table class="table table-hover">
                <thead>
                <tr>
                    <th scope="col">Node Name</th>
                    <th scope="col">Volume</th>
                </tr>
                </thead>
                <tbody>
                {{ range .TopNodes }}
                    <tr>
                        <th scope="row">{{.Name}}</th>
                        <td>{{.Volume}}</td>
                    </tr>
                {{ end }}
                </tbody>
            </table>
        </td>
    </tr>
    </tbody>
</table>
</body>
</html>