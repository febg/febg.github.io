<html>
<head>
<title>A BASIC HTML FORM</title>
<?PHP
	$username = $_GET['username'];
	print $username;
?>
</head>
<body>

<FORM NAME ="form1" METHOD ="GET" ACTION = "htmlbd.php">

<INPUT TYPE = "TEXT" NAME= "username" VALUE ="username">
<INPUT TYPE = "Submit" Name = "Submit1" VALUE = "Login">

</FORM>

<FORM NAME ="form2" METHOD ="GET" ACTION = "htmlbd.php">

<INPUT TYPE = "TEXT" NAME= "password" VALUE ="password">
<INPUT TYPE = "Submit" Name = "Submit2" VALUE = "password">

</FORM>

</body>
</html>
