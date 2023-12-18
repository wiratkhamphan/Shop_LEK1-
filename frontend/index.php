<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>User Registration</title>
</head>
<body>
    <h1>User Registration</h1>
    <form id="registrationForm"method="post">
        <label for="iso_p_code">ISO P Code:</label>
        <input type="text" id="iso_p_code" name="iso_p_code" required>
        <br>
        <label for="p_code">P Code:</label>
        <input type="text" id="p_code" name="p_code" required>
        <br>
        <label for="email">Email:</label>
        <input type="email" id="email" name="email" required>
        <br>
        <label for="password">Password:</label>
        <input type="password" id="password" name="password" required>
        <br>
        <label for="p_name">P Name:</label>
        <input type="text" id="p_name" name="p_name" required>
        <br>
        <button type="button" onclick="registerUser()">Register</button>
    </form>

    <script src="script.js"></script>
</body>
</html>
