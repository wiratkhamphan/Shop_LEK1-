function registerUser() {
    var iso_p_code = document.getElementById("iso_p_code").value;
    var p_code = document.getElementById("p_code").value;
    var email = document.getElementById("email").value;
    var password = document.getElementById("password").value;
    var p_name = document.getElementById("p_name").value;

    var userData = {
        iso_p_code: iso_p_code,
        p_code: p_code,
        email: email,
        password: password,
        p_name: p_name
    };

    fetch("http://localhost:8080/Register_user", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(userData)
    })
    .then(response => {
        if (response.ok) {
            alert("User registered successfully!");
        } else {
            alert("Error registering user. Please try again.");
        }
    })
    .catch(error => {
        alert("Request failed. Please check your network connection.");
    });
}
