package main

import "html/template"

var html = template.Must(template.New("chat_room").Parse(`
<html>
<head>
	<title>{{ .roomid }}</title>
	<script src="http://ajax.googleapis.com/ajax/libs/jquery/1.7/jquery.js"></script> 
	<script src="http://malsup.github.com/jquery.form.js"></script> 
	<script>
		$('#message_form').focus();
		$(document).ready(function() {
			$('#myForm').ajaxForm(function() {
				$('#message_form').val('');
				$('#message_form').focus();
			});

			if (!!window.EventSource) {
				var source = new EventSource('/stream/{{.roomid}}');
				source.addEventListener('message', function(e) {
					$('#messages').append(e.data + "</br>");
					$('html, body').animate({scrollTop:$(document).height()}, 'slow');
				}, false);
			} else {
				alert("NOT SUPPORTED")
			}
		});
	</script>
</head>
<body>
	<h1>Welcome to {{ .roomid }} room</h1>
	<div id="messages"></id>
	<form id="myForm" action="/room/{{.roomid}}" method="POST">
			User: <input id="user_form" name="user" value="{{.userid}}">
			Message: <input id="message_form" name="message">
			<input type="submit" value="Submit">
	</form>
</body>
</html>
`))
