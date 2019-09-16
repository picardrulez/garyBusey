package main

import ()

type Page struct {
	Username string
	Content  string
}

func pageBuilder() string {
	pageReturn := `
	<html>
	<head>
		<title>The Gary Busey Super Secret Data Vault</title>
		<link REL="StyleSheet" TYPE="text/css" HREF="resources/busey.css">
	</head>
	<body>
	<div class="header">
		<a href="/admin">The Gary Busey Super Secret Data Vault ` + VERSION + ` </a>
		<div class="usermenu">{{ .Username }} | <a href="/logout">Logout</a>
		</div>
	</div>
	<div class="menu">
		<div class="nav">
			<ul>
				<li class="Keys"<a href="/admin">Keys</a>
					<ul>
						<li><a href="/listKeys">List Keys</a></li>
						<li><a href="/addKey">Add Key</a></li>
						<li><a href="/removeKey">Remove Key</a></li>
					</ul>
				</li>
				<li class="Users"<a href="/admin">Users</a>
					<ul>
						<li><a href="/listUsers">List Users</a></li>
						<li><a href="/createUser">Add User</a></li>
						<li><a href="/deleteUser">DeleteUser</a></li>
						<li><a href="/changePassword">Change Password</a></li>
						<li><a href="/logout">Logout</a></li>
					</ul>
				</li>
			</ul>
		</div>
	</div>
	<div class="main">
		{{ .Content }}
	</div>
	</body>
	</html>
	`
	return pageReturn
}
