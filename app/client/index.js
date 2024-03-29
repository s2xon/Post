const login = document.getElementById("login");
const logout = document.getElementById("logout");

const post = document.getElementById("post");

const fb_loginURL = "https://www.facebook.com/v19.0/dialog/oauth?client_id=399965006007248&redirect_uri=https://localhost:3000/facebook&state=st=state123gitygiabc,ds=123456789";

login.addEventListener("click", () => {
    window.location.href = fb_loginURL;
});
