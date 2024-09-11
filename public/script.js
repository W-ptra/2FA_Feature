async function login(){
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    if (email === "" || password === "")
        return document.getElementById("message").textContent = "Email or Passworld can't empty";

    // fetch to api /login
    document.getElementById("message").textContent = "";
}

async function register(){
    const name = document.getElementById("name").value;
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    const confirmPassword = document.getElementById("confirmPassword").value;

    if (name === "" || email === "" || password === "" || confirmPassword === "")
        return document.getElementById("message").textContent = "Email or Passworld can't empty";

    if (password !== confirmPassword)
        return document.getElementById("message").textContent = "Passworld and Confirm Passwird doesn't match";

    //fetch to api /register
    document.getElementById("message").textContent = "";
}