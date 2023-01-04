const url = document.querySelector("#url-input");
const paste = document.querySelector("#paste");
const shorten = document.querySelector("#shorten");
const shWrap = document.querySelector("#shorty-wrapper");
const shorty = document.querySelector("#shorty");
const copy = document.querySelector("#copy");
const API = "http://localhost:42069/api";

const makeItShort = async (url) => {
    shWrap.style.display = "none";
    if (url === "") {
        return;
    }
    const shortURL = (await (await fetch(API, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        }, 
        mode: "cors", 
        body: JSON.stringify({
            redirectTo: url
        }),
    })).json()).shortURL;

    shorty.innerText = shortURL;
    shorty.href = shortURL;
    shorty.target = "_blank";
    shWrap.style.display = "flex";
}

const copyToClipboard = async () => {
    navigator.clipboard.writeText(shorty.innerText);
}

const pasteFromClipboard = async () => {
    const text = await navigator.clipboard.readText();
    url.value = text;
}

shorten.addEventListener("click", (e) => {
    e.preventDefault();
    makeItShort(url.value);
}) 

copy.addEventListener("click", (e) => {
    e.preventDefault();
    copyToClipboard();
}) 

paste.addEventListener("click", (e) => {
    e.preventDefault();
    pasteFromClipboard();
}) 
