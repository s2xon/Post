const login = document.getElementById("login");
const logout = document.getElementById("logout");

const post = document.getElementById("post");

const fb_loginURL = "https://www.facebook.com/v19.0/dialog/oauth?client_id=939552021128756&redirect_uri=https://localhost:3000/facebook&state=s";  

login.addEventListener("click", () => {
    window.location.href = fb_loginURL;
});
