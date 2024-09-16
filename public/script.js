async function login(){
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    if (email === "" || password === "")
        return document.getElementById("message").textContent = "Email or Passworld can't empty";
    
    const respond = await fetch("http://localhost:8001/login",{
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
        document.getElementById("message").textContent = "Login Successfull";
        document.getElementById("otp_email_label").textContent = email;
        document.getElementById("otp_form").style = "position: absolute; left: 0; right: 0; top: 0; bottom: 0; background-color: rgb(179, 179, 179); z-index: 10; display:inherite;"
        return
    }
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
            Code:otp
        })
    })

    const result = await respond.json()
    console.log(result)

    document.getElementById("otp_message").textContent = "Login Successfull";
}