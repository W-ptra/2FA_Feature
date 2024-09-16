async function login(){
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    if (email === "" || password === ""){
        document.getElementById("email").value = "";
        document.getElementById("password").value = "";
        return document.getElementById("message").textContent = "Email or Passworld can't empty";
    }
    const respond = await fetch("/login",{
        method:"POST",
        headers:{
            "Content-Type":"application/json"
        },
        body: JSON.stringify({
            email,
            password,
        })
    })
    const result = await respond.json()
    console.log(result)
    if (result.Message === "Login Sucessfully" && respond.ok){
        document.getElementById("otp_email_label").textContent = email;
        document.getElementById("otp_form").style = "display:inherite;"
        localStorage.setItem("email",email);
        return
    }
    document.getElementById("email").value = "";
    document.getElementById("password").value = "";
    document.getElementById("message").textContent = "Login Failed";
}

async function register(){
    const name = document.getElementById("name").value;
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    const confirmPassword = document.getElementById("confirmPassword").value;

    if (name === "" || email === "" || password === "" || confirmPassword === "")
        return document.getElementById("message").textContent = "Email or Passworld can't empty";

    if (password !== confirmPassword)
        return document.getElementById("message").textContent = "Password and Confirm Password doesn't match";

    
    const respond = await fetch("http://localhost:8001/register",{
        method:"POST",
        headers:{
            "Content-Type":"application/json"
        },
        body: JSON.stringify({
            name,
            email,
            password,
            confirmPassword
        })
    })
    const result = await respond.json()
    console.log(result)
    if (result.Message === "Register Sucessfully" && respond.ok){
        return document.getElementById("message").textContent = "Register Successfull";
    }
    document.getElementById("name").value = "";
    document.getElementById("email").value = "";
    document.getElementById("password").value = "";
    document.getElementById("confirmPassword").value = "";
    document.getElementById("message").textContent = "Register failed";
}

async function sendOTP(){
    const otp = document.getElementById("OTP").value;
    if (otp === ""){
        return document.getElementById("otp_message").textContent = "Input can't empty";
    }

    const respond = await fetch("/otp",{
        method:"POST",
        headers:{
            "Content-Type":"application/json"
        },
        body:JSON.stringify({
            code:otp,
            email:localStorage.getItem("email")
        })
    })

    const result = await respond.json()
    console.log(result)
    if (respond.ok){
        document.getElementById("OTP").value = "";
        return document.getElementById("otp_message").textContent = "Login Successfull";
    }
    document.getElementById("otp_message").textContent = result.Message;
}