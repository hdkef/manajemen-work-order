document.getElementById("btn-logout").onclick = ()=>{
    logout()
}

logout = ()=>{
    document.cookie = "Authorization" + '=; Max-Age=0'
    window.location.href = "/login"
}